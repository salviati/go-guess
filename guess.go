// Copyright 2011 Utkan Güngördü. All rights reserved.
// Use of this source code is governed by a BSD-style
// See the LICENSE file of the official Go distrubtion.

/*
 * This code is derivative of guess.c of Gauche-0.8.3.
 * The following is the original copyright notice.
 */

/*
 *   Copyright (c) 2000-2003 Shiro Kawai, All rights reserved.
 *
 *   Redistribution and use in source and binary forms, with or without
 *   modification, are permitted provided that the following conditions
 *   are met:
 *
 *   1. Redistributions of source code must retain the above copyright
 *      notice, this list of conditions and the following disclaimer.
 *
 *   2. Redistributions in binary form must reproduce the above copyright
 *      notice, this list of conditions and the following disclaimer in the
 *      documentation and/or other materials provided with the distribution.
 *
 *   3. Neither the name of the authors nor the names of its contributors
 *      may be used to endorse or promote products derived from this
 *      software without specific prior written permission.
 *
 *   THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 *   "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 *   LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 *   A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 *   OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 *   SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED
 *   TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
 *   PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
 *   LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
 *   NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 *   SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */

// cgo wrapper for libguess (http://www.atheme.org/project/libguess)
//
// libguess employs discrete-finite automata to deduce the character set of the
// input buffer. The advantage of this is that all character sets can be
// checked in parallel, and quickly. Right now, libguess passes a byte to each
// DFA on the same pass, meaning that the winning character set can be deduced
//as efficiently as possible.
//
//libguess is fully reentrant, using only local stack memory for DFA operations.
package guess

// #cgo pkg-config: libguess
//#include <libguess/libguess.h>
import "C"

import (
  "unsafe"
)

const (
  AR = C.GUESS_REGION_AR // Arabic
  BL = C.GUESS_REGION_BL // Baltic
  CN = C.GUESS_REGION_CN // Chinese
  GR = C.GUESS_REGION_GR // Greek
  HW = C.GUESS_REGION_HW // Hebrew
  JP = C.GUESS_REGION_JP // Japanese
  KR = C.GUESS_REGION_KR // Korean
  PL = C.GUESS_REGION_PL // Polish
  RU = C.GUESS_REGION_RU // Russian
  TW = C.GUESS_REGION_TW // Taiwanese
  TR = C.GUESS_REGION_TR // Turkish
)

// This employs libguess's DFA-based character set validation rules to ensure
// that a string is pure UTF-8. GLib's UTF-8 validation functions are broken,
// for example.
func ValidateUTF8(buf []byte) bool {
  x := C.libguess_validate_utf8((*C.char)(unsafe.Pointer(&buf[0])), C.int(len(buf)))
  return x != 0
}

// This detects a character set.
// Returns an appropriate charset name that can be passed to iconv_open().
// Region is the name of the language or region that the data is related to,
// e.g. 'Baltic' for the Baltic states, or 'Japanese' for Japan.
func DetermineEncoding(inbuf []byte, region string) string {
	cbuf := (*C.char)(unsafe.Pointer(&inbuf[0]))
	cbuflen := C.int(len(inbuf))
	cregion := C.CString(region)
	defer C.free(unsafe.Pointer(cregion))

	cguess := C.libguess_determine_encoding(cbuf, cbuflen, cregion)
	return C.GoString(cguess)
}
