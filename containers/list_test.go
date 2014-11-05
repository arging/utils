// Copyright 2014 li. All rights reserved.
// Use of this source code is governed by a MIT/X11
// license that can be found in the LICENSE file.

package containers

import (
	"container/list"
	"testing"
)

func TestIsEmptyList(t *testing.T) {
	l1 := list.New()
	l2 := list.New()
	l2.PushBack(3)

	if !IsEmptyList(nil) || !IsEmptyList(l1) || IsEmptyList(l2) {
		t.Error("utils/collection: list should be empty.")
	}
}

func TestIsNotEmptyList(t *testing.T) {
	l1 := list.New()
	l2 := list.New()
	l2.PushBack(3)

	if IsNotEmptyList(nil) || IsNotEmptyList(l1) || !IsNotEmptyList(l2) {
		t.Error("utils/collection: list should not be empty.")
	}
}
