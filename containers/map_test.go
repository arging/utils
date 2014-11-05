// Copyright 2014 li. All rights reserved.
// Use of this source code is governed by a MIT/X11
// license that can be found in the LICENSE file.

package containers

import (
	"reflect"
	"testing"
)

func TestConcurrentMap(t *testing.T) {
	cmap := NewCMap()

	if cmap.Size() != 0 {
		t.Fatal()
	}

	cmap.Put("China", "BeiJing")
	cmap.Put("Japan", "Tokyo")

	if cmap.Size() != 2 {
		t.Fatal()
	}

	if !cmap.ContainsKey("China") || cmap.ContainsKey("CHINA") {
		t.Fatal()
	}

	cmap.PutIfAbsent("China", "ShangHai")
	cmap.PutIfAbsent("America", "NewYork")
	if cmap.Size() != 3 {
		t.Fatal()
	}
	v, _ := cmap.Get("China")
	if v != "BeiJing" {
		t.Fatal()
	}

	if !reflect.DeepEqual([]interface{}{"China", "Japan", "America"},
		cmap.Keys()) {
		t.Fatal()
	}

	cmap.Remove("China")
	if cmap.Size() != 2 {
		t.Fatal()
	}

	cmap.Clear()
	if cmap.Size() != 0 {
		t.Fatal()
	}
}
