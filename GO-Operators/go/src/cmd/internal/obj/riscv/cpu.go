//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2008 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2008 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors.  All rights reserved.
//	Portions Copyright © 2019 The Go Authors.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package riscv

import "cmd/internal/obj"

//go:generate go run ../stringer.go -i $GOFILE -o anames.go -p riscv

const (
	// Base register numberings.
	REG_X0 = obj.RBaseRISCV + iota
	REG_X1
	REG_X2
	REG_X3
	REG_X4
	REG_X5
	REG_X6
	REG_X7
	REG_X8
	REG_X9
	REG_X10
	REG_X11
	REG_X12
	REG_X13
	REG_X14
	REG_X15
	REG_X16
	REG_X17
	REG_X18
	REG_X19
	REG_X20
	REG_X21
	REG_X22
	REG_X23
	REG_X24
	REG_X25
	REG_X26
	REG_X27
	REG_X28
	REG_X29
	REG_X30
	REG_X31

	// FP register numberings.
	REG_F0
	REG_F1
	REG_F2
	REG_F3
	REG_F4
	REG_F5
	REG_F6
	REG_F7
	REG_F8
	REG_F9
	REG_F10
	REG_F11
	REG_F12
	REG_F13
	REG_F14
	REG_F15
	REG_F16
	REG_F17
	REG_F18
	REG_F19
	REG_F20
	REG_F21
	REG_F22
	REG_F23
	REG_F24
	REG_F25
	REG_F26
	REG_F27
	REG_F28
	REG_F29
	REG_F30
	REG_F31

	// This marks the end of the register numbering.
	REG_END

	// General registers reassigned to ABI names.
	REG_ZERO = REG_X0
	REG_RA   = REG_X1 // aka REG_LR
	REG_SP   = REG_X2
	REG_GP   = REG_X3 // aka REG_SB
	REG_TP   = REG_X4
	REG_T0   = REG_X5
	REG_T1   = REG_X6
	REG_T2   = REG_X7
	REG_S0   = REG_X8
	REG_S1   = REG_X9
	REG_A0   = REG_X10
	REG_A1   = REG_X11
	REG_A2   = REG_X12
	REG_A3   = REG_X13
	REG_A4   = REG_X14
	REG_A5   = REG_X15
	REG_A6   = REG_X16
	REG_A7   = REG_X17
	REG_S2   = REG_X18
	REG_S3   = REG_X19
	REG_S4   = REG_X20
	REG_S5   = REG_X21
	REG_S6   = REG_X22
	REG_S7   = REG_X23
	REG_S8   = REG_X24
	REG_S9   = REG_X25
	REG_S10  = REG_X26 // aka REG_CTXT
	REG_S11  = REG_X27 // aka REG_G
	REG_T3   = REG_X28
	REG_T4   = REG_X29
	REG_T5   = REG_X30
	REG_T6   = REG_X31 // aka REG_TMP

	// Go runtime register names.
	REG_CTXT = REG_S10 // Context for closures.
	REG_G    = REG_S11 // G pointer.
	REG_LR   = REG_RA  // Link register.
	REG_TMP  = REG_T6  // Reserved for assembler use.

	// ABI names for floating point registers.
	REG_FT0  = REG_F0
	REG_FT1  = REG_F1
	REG_FT2  = REG_F2
	REG_FT3  = REG_F3
	REG_FT4  = REG_F4
	REG_FT5  = REG_F5
	REG_FT6  = REG_F6
	REG_FT7  = REG_F7
	REG_FS0  = REG_F8
	REG_FS1  = REG_F9
	REG_FA0  = REG_F10
	REG_FA1  = REG_F11
	REG_FA2  = REG_F12
	REG_FA3  = REG_F13
	REG_FA4  = REG_F14
	REG_FA5  = REG_F15
	REG_FA6  = REG_F16
	REG_FA7  = REG_F17
	REG_FS2  = REG_F18
	REG_FS3  = REG_F19
	REG_FS4  = REG_F20
	REG_FS5  = REG_F21
	REG_FS6  = REG_F22
	REG_FS7  = REG_F23
	REG_FS8  = REG_F24
	REG_FS9  = REG_F25
	REG_FS10 = REG_F26
	REG_FS11 = REG_F27
	REG_FT8  = REG_F28
	REG_FT9  = REG_F29
	REG_FT10 = REG_F30
	REG_FT11 = REG_F31

	// Names generated by the SSA compiler.
	REGSP = REG_SP
	REGG  = REG_G
)

// https://github.com/riscv-non-isa/riscv-elf-psabi-doc/blob/master/riscv-dwarf.adoc#dwarf-register-numbers
var RISCV64DWARFRegisters = map[int16]int16{
	// Integer Registers.
	REG_X0:  0,
	REG_X1:  1,
	REG_X2:  2,
	REG_X3:  3,
	REG_X4:  4,
	REG_X5:  5,
	REG_X6:  6,
	REG_X7:  7,
	REG_X8:  8,
	REG_X9:  9,
	REG_X10: 10,
	REG_X11: 11,
	REG_X12: 12,
	REG_X13: 13,
	REG_X14: 14,
	REG_X15: 15,
	REG_X16: 16,
	REG_X17: 17,
	REG_X18: 18,
	REG_X19: 19,
	REG_X20: 20,
	REG_X21: 21,
	REG_X22: 22,
	REG_X23: 23,
	REG_X24: 24,
	REG_X25: 25,
	REG_X26: 26,
	REG_X27: 27,
	REG_X28: 28,
	REG_X29: 29,
	REG_X30: 30,
	REG_X31: 31,

	// Floating-Point Registers.
	REG_F0:  32,
	REG_F1:  33,
	REG_F2:  34,
	REG_F3:  35,
	REG_F4:  36,
	REG_F5:  37,
	REG_F6:  38,
	REG_F7:  39,
	REG_F8:  40,
	REG_F9:  41,
	REG_F10: 42,
	REG_F11: 43,
	REG_F12: 44,
	REG_F13: 45,
	REG_F14: 46,
	REG_F15: 47,
	REG_F16: 48,
	REG_F17: 49,
	REG_F18: 50,
	REG_F19: 51,
	REG_F20: 52,
	REG_F21: 53,
	REG_F22: 54,
	REG_F23: 55,
	REG_F24: 56,
	REG_F25: 57,
	REG_F26: 58,
	REG_F27: 59,
	REG_F28: 60,
	REG_F29: 61,
	REG_F30: 62,
	REG_F31: 63,
}

// Prog.Mark flags.
const (
	// USES_REG_TMP indicates that a machine instruction generated from the
	// corresponding *obj.Prog uses the temporary register.
	USES_REG_TMP = 1 << iota

	// NEED_CALL_RELOC is set on JAL instructions to indicate that a
	// R_RISCV_CALL relocation is needed.
	NEED_CALL_RELOC

	// NEED_PCREL_ITYPE_RELOC is set on AUIPC instructions to indicate that
	// it is the first instruction in an AUIPC + I-type pair that needs a
	// R_RISCV_PCREL_ITYPE relocation.
	NEED_PCREL_ITYPE_RELOC

	// NEED_PCREL_STYPE_RELOC is set on AUIPC instructions to indicate that
	// it is the first instruction in an AUIPC + S-type pair that needs a
	// R_RISCV_PCREL_STYPE relocation.
	NEED_PCREL_STYPE_RELOC
)

// RISC-V mnemonics, as defined in the "opcodes" and "opcodes-pseudo" files
// at https://github.com/riscv/riscv-opcodes.
//
// As well as some pseudo-mnemonics (e.g. MOV) used only in the assembler.
//
// See also "The RISC-V Instruction Set Manual" at https://riscv.org/specifications/.
//
// If you modify this table, you MUST run 'go generate' to regenerate anames.go!
const (
	// Unprivileged ISA (Document Version 20190608-Base-Ratified)

	// 2.4: Integer Computational Instructions
	AADDI = obj.ABaseRISCV + obj.A_ARCHSPECIFIC + iota
	ASLTI
	ASLTIU
	AANDI
	AORI
	AXORI
	ASLLI
	ASRLI
	ASRAI
	ALUI
	AAUIPC
	AADD
	ASLT
	ASLTU
	AAND
	AOR
	AXOR
	ASLL
	ASRL
	ASUB
	ASRA

	// The SLL/SRL/SRA instructions differ slightly between RV32 and RV64,
	// hence there are pseudo-opcodes for the RV32 specific versions.
	ASLLIRV32
	ASRLIRV32
	ASRAIRV32

	// 2.5: Control Transfer Instructions
	AJAL
	AJALR
	ABEQ
	ABNE
	ABLT
	ABLTU
	ABGE
	ABGEU

	// 2.6: Load and Store Instructions
	ALW
	ALWU
	ALH
	ALHU
	ALB
	ALBU
	ASW
	ASH
	ASB

	// 2.7: Memory Ordering Instructions
	AFENCE
	AFENCEI
	AFENCETSO

	// 5.2: Integer Computational Instructions (RV64I)
	AADDIW
	ASLLIW
	ASRLIW
	ASRAIW
	AADDW
	ASLLW
	ASRLW
	ASUBW
	ASRAW

	// 5.3: Load and Store Instructions (RV64I)
	ALD
	ASD

	// 7.1: Multiplication Operations
	AMUL
	AMULH
	AMULHU
	AMULHSU
	AMULW
	ADIV
	ADIVU
	AREM
	AREMU
	ADIVW
	ADIVUW
	AREMW
	AREMUW

	// 8.2: Load-Reserved/Store-Conditional Instructions
	ALRD
	ASCD
	ALRW
	ASCW

	// 8.3: Atomic Memory Operations
	AAMOSWAPD
	AAMOADDD
	AAMOANDD
	AAMOORD
	AAMOXORD
	AAMOMAXD
	AAMOMAXUD
	AAMOMIND
	AAMOMINUD
	AAMOSWAPW
	AAMOADDW
	AAMOANDW
	AAMOORW
	AAMOXORW
	AAMOMAXW
	AAMOMAXUW
	AAMOMINW
	AAMOMINUW

	// 10.1: Base Counters and Timers
	ARDCYCLE
	ARDCYCLEH
	ARDTIME
	ARDTIMEH
	ARDINSTRET
	ARDINSTRETH

	// 11.2: Floating-Point Control and Status Register
	AFRCSR
	AFSCSR
	AFRRM
	AFSRM
	AFRFLAGS
	AFSFLAGS
	AFSRMI
	AFSFLAGSI

	// 11.5: Single-Precision Load and Store Instructions
	AFLW
	AFSW

	// 11.6: Single-Precision Floating-Point Computational Instructions
	AFADDS
	AFSUBS
	AFMULS
	AFDIVS
	AFMINS
	AFMAXS
	AFSQRTS
	AFMADDS
	AFMSUBS
	AFNMADDS
	AFNMSUBS

	// 11.7: Single-Precision Floating-Point Conversion and Move Instructions
	AFCVTWS
	AFCVTLS
	AFCVTSW
	AFCVTSL
	AFCVTWUS
	AFCVTLUS
	AFCVTSWU
	AFCVTSLU
	AFSGNJS
	AFSGNJNS
	AFSGNJXS
	AFMVXS
	AFMVSX
	AFMVXW
	AFMVWX

	// 11.8: Single-Precision Floating-Point Compare Instructions
	AFEQS
	AFLTS
	AFLES

	// 11.9: Single-Precision Floating-Point Classify Instruction
	AFCLASSS

	// 12.3: Double-Precision Load and Store Instructions
	AFLD
	AFSD

	// 12.4: Double-Precision Floating-Point Computational Instructions
	AFADDD
	AFSUBD
	AFMULD
	AFDIVD
	AFMIND
	AFMAXD
	AFSQRTD
	AFMADDD
	AFMSUBD
	AFNMADDD
	AFNMSUBD

	// 12.5: Double-Precision Floating-Point Conversion and Move Instructions
	AFCVTWD
	AFCVTLD
	AFCVTDW
	AFCVTDL
	AFCVTWUD
	AFCVTLUD
	AFCVTDWU
	AFCVTDLU
	AFCVTSD
	AFCVTDS
	AFSGNJD
	AFSGNJND
	AFSGNJXD
	AFMVXD
	AFMVDX

	// 12.6: Double-Precision Floating-Point Compare Instructions
	AFEQD
	AFLTD
	AFLED

	// 12.7: Double-Precision Floating-Point Classify Instruction
	AFCLASSD

	// 13.1 Quad-Precision Load and Store Instructions
	AFLQ
	AFSQ

	// 13.2: Quad-Precision Computational Instructions
	AFADDQ
	AFSUBQ
	AFMULQ
	AFDIVQ
	AFMINQ
	AFMAXQ
	AFSQRTQ
	AFMADDQ
	AFMSUBQ
	AFNMADDQ
	AFNMSUBQ

	// 13.3 Quad-Precision Convert and Move Instructions
	AFCVTWQ
	AFCVTLQ
	AFCVTSQ
	AFCVTDQ
	AFCVTQW
	AFCVTQL
	AFCVTQS
	AFCVTQD
	AFCVTWUQ
	AFCVTLUQ
	AFCVTQWU
	AFCVTQLU
	AFSGNJQ
	AFSGNJNQ
	AFSGNJXQ
	AFMVXQ
	AFMVQX

	// 13.4 Quad-Precision Floating-Point Compare Instructions
	AFEQQ
	AFLEQ
	AFLTQ

	// 13.5 Quad-Precision Floating-Point Classify Instruction
	AFCLASSQ

	// Privileged ISA (Version 20190608-Priv-MSU-Ratified)

	// 3.1.9: Instructions to Access CSRs
	ACSRRW
	ACSRRS
	ACSRRC
	ACSRRWI
	ACSRRSI
	ACSRRCI

	// 3.2.1: Environment Call and Breakpoint
	AECALL
	ASCALL
	AEBREAK
	ASBREAK

	// 3.2.2: Trap-Return Instructions
	AMRET
	ASRET
	AURET
	ADRET

	// 3.2.3: Wait for Interrupt
	AWFI

	// 4.2.1: Supervisor Memory-Management Fence Instruction
	ASFENCEVMA

	// Hypervisor Memory-Management Instructions
	AHFENCEGVMA
	AHFENCEVVMA

	// The escape hatch. Inserts a single 32-bit word.
	AWORD

	// Pseudo-instructions.  These get translated by the assembler into other
	// instructions, based on their operands.
	ABEQZ
	ABGEZ
	ABGT
	ABGTU
	ABGTZ
	ABLE
	ABLEU
	ABLEZ
	ABLTZ
	ABNEZ
	AFABSD
	AFABSS
	AFNEGD
	AFNEGS
	AFNED
	AFNES
	AMOV
	AMOVB
	AMOVBU
	AMOVF
	AMOVD
	AMOVH
	AMOVHU
	AMOVW
	AMOVWU
	ANEG
	ANEGW
	ANOT
	ASEQZ
	ASNEZ

	// End marker
	ALAST
)

// All unary instructions which write to their arguments (as opposed to reading
// from them) go here. The assembly parser uses this information to populate
// its AST in a semantically reasonable way.
//
// Any instructions not listed here are assumed to either be non-unary or to read
// from its argument.
var unaryDst = map[obj.As]bool{
	ARDCYCLE:    true,
	ARDCYCLEH:   true,
	ARDTIME:     true,
	ARDTIMEH:    true,
	ARDINSTRET:  true,
	ARDINSTRETH: true,
}

// Instruction encoding masks.
const (
	// JTypeImmMask is a mask including only the immediate portion of
	// J-type instructions.
	JTypeImmMask = 0xfffff000

	// ITypeImmMask is a mask including only the immediate portion of
	// I-type instructions.
	ITypeImmMask = 0xfff00000

	// STypeImmMask is a mask including only the immediate portion of
	// S-type instructions.
	STypeImmMask = 0xfe000f80

	// UTypeImmMask is a mask including only the immediate portion of
	// U-type instructions.
	UTypeImmMask = 0xfffff000
)
