# Copyright 2010 Utkan Güngördü. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file of Go. 

include $(GOROOT)/src/Make.inc

TARG=guess

CGOFILES=guess.go
CGO_LDFLAGS=-lguess

include $(GOROOT)/src/Make.pkg 
