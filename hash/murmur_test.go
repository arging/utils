// Copyright 2014 li. All rights reserved.
// Use of this source code is governed by a MIT/X11
// license that can be found in the LICENSE file.

package hash

import (
	"testing"
)

func TestMurmur3_32(t *testing.T) {
	if Murmur3_32([]byte{}, 0) != 0 {
		t.Fail()
	}
	if Murmur3_32([]byte("Murmur3_32,Murmur3_32,Murmur3_32"), 0) != 2714439771 {
		t.Fail()
	}
}

func TestMurmur3_128(t *testing.T) {
	h1, h2 := Murmur3_128([]byte{}, 0)
	if h1 != 0 || h2 != 0 {
		t.Fail()
	}
	h1, h2 = Murmur3_128([]byte("Murmur3_128,Murmur3_128,Murmur3_128,Murmur3_128"), 0)
	if h1 != 5324627280710961006 || h2 != 4655615913267170926 {
		t.Fail()
	}
}
