// Copyright 2020 Michael Treanor. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package stringutils implements additional functions to
// support the go standard library strings module.
//
// The algorithms chosen are based on benchmarks from
// the stringbenchmarks module. ymmv...
//
// The current implementation at the start of this project was
// .../go/1.15.3/libexec/src/strings/strings.go
//
// For information about UTF-8 strings in Go,
// see https://blog.golang.org/strings.
package stringutils

import (
	"math/rand"
	"time"

	// Functions exported from stringbenchmarks are the
	// most efficient versions
	. "github.com/skeptycal/stringutils/stringbenchmarks"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func ByteToDigit(c byte) byte {
	if IsDigit(c) {
		return c - 65
	}
	return '0'
}
