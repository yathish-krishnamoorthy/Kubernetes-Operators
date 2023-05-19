// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package buildid

import (
	"bytes"
	"cmd/internal/codesign"
	"crypto/sha256"
	"debug/macho"
	"fmt"
	"io"
)

// FindAndHash reads all of r and returns the offsets of occurrences of id.
// While reading, findAndHash also computes and returns
// a hash of the content of r, but with occurrences of id replaced by zeros.
// FindAndHash reads bufSize bytes from r at a time.
// If bufSize == 0, FindAndHash uses a reasonable default.
func FindAndHash(r io.Reader, id string, bufSize int) (matches []int64, hash [32]byte, err error) {
	if bufSize == 0 {
		bufSize = 31 * 1024 // bufSize+little will likely fit in 32 kB
	}
	if len(id) == 0 {
		return nil, [32]byte{}, fmt.Errorf("buildid.FindAndHash: no id specified")
	}
	if len(id) > bufSize {
		return nil, [32]byte{}, fmt.Errorf("buildid.FindAndHash: buffer too small")
	}
	zeros := make([]byte, len(id))
	idBytes := []byte(id)

	// For Mach-O files, we want to exclude the code signature.
	// The code signature contains hashes of the whole file (except the signature
	// itself), including the buildid. So the buildid cannot contain the signature.
	r = excludeMachoCodeSignature(r)

	// The strategy is to read the file through buf, looking for id,
	// but we need to worry about what happens if id is broken up
	// and returned in parts by two different reads.
	// We allocate a tiny buffer (at least len(id)) and a big buffer (bufSize bytes)
	// next to each other in memory and then copy the tail of
	// one read into the tiny buffer before reading new data into the big buffer.
	// The search for id is over the entire tiny+big buffer.
	tiny := (len(id) + 127) &^ 127 // round up to 128-aligned
	buf := make([]byte, tiny+bufSize)
	h := sha256.New()
	start := tiny
	for offset := int64(0); ; {
		// The file offset maintained by the loop corresponds to &buf[tiny].
		// buf[start:tiny] is left over from previous iteration.
		// After reading n bytes into buf[tiny:], we process buf[start:tiny+n].
		n, err := io.ReadFull(r, buf[tiny:])
		if err != io.ErrUnexpectedEOF && err != io.EOF && err != nil {
			return nil, [32]byte{}, err
		}

		// Process any matches.
		for {
			i := bytes.Index(buf[start:tiny+n], idBytes)
			if i < 0 {
				break
			}
			matches = append(matches, offset+int64(start+i-tiny))
			h.Write(buf[start : start+i])
			h.Write(zeros)
			start += i + len(id)
		}
		if n < bufSize {
			// Did not fill buffer, must be at end of file.
			h.Write(buf[start : tiny+n])
			break
		}

		// Process all but final tiny bytes of buf (bufSize = len(buf)-tiny).
		// Note that start > len(buf)-tiny is possible, if the search above
		// found an id ending in the final tiny fringe. That's OK.
		if start < len(buf)-tiny {
			h.Write(buf[start : len(buf)-tiny])
			start = len(buf) - tiny
		}

		// Slide ending tiny-sized fringe to beginning of buffer.
		copy(buf[0:], buf[bufSize:])
		start -= bufSize
		offset += int64(bufSize)
	}
	h.Sum(hash[:0])
	return matches, hash, nil
}

func Rewrite(w io.WriterAt, pos []int64, id string) error {
	b := []byte(id)
	for _, p := range pos {
		if _, err := w.WriteAt(b, p); err != nil {
			return err
		}
	}

	// Update Mach-O code signature, if any.
	if f, cmd, ok := findMachoCodeSignature(w); ok {
		if codesign.Size(int64(cmd.Dataoff), "a.out") == int64(cmd.Datasize) {
			// Update the signature if the size matches, so we don't need to
			// fix up headers. Binaries generated by the Go linker should have
			// the expected size. Otherwise skip.
			text := f.Segment("__TEXT")
			cs := make([]byte, cmd.Datasize)
			codesign.Sign(cs, w.(io.Reader), "a.out", int64(cmd.Dataoff), int64(text.Offset), int64(text.Filesz), f.Type == macho.TypeExec)
			if _, err := w.WriteAt(cs, int64(cmd.Dataoff)); err != nil {
				return err
			}
		}
	}

	return nil
}

func excludeMachoCodeSignature(r io.Reader) io.Reader {
	_, cmd, ok := findMachoCodeSignature(r)
	if !ok {
		return r
	}
	return &excludedReader{r, 0, int64(cmd.Dataoff), int64(cmd.Dataoff + cmd.Datasize)}
}

// excludedReader wraps an io.Reader. Reading from it returns the bytes from
// the underlying reader, except that when the byte offset is within the
// range between start and end, it returns zero bytes.
type excludedReader struct {
	r          io.Reader
	off        int64 // current offset
	start, end int64 // the range to be excluded (read as zero)
}

func (r *excludedReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if n > 0 && r.off+int64(n) > r.start && r.off < r.end {
		cstart := r.start - r.off
		if cstart < 0 {
			cstart = 0
		}
		cend := r.end - r.off
		if cend > int64(n) {
			cend = int64(n)
		}
		zeros := make([]byte, cend-cstart)
		copy(p[cstart:cend], zeros)
	}
	r.off += int64(n)
	return n, err
}

func findMachoCodeSignature(r any) (*macho.File, codesign.CodeSigCmd, bool) {
	ra, ok := r.(io.ReaderAt)
	if !ok {
		return nil, codesign.CodeSigCmd{}, false
	}
	f, err := macho.NewFile(ra)
	if err != nil {
		return nil, codesign.CodeSigCmd{}, false
	}
	cmd, ok := codesign.FindCodeSigCmd(f)
	return f, cmd, ok
}
