// Copyright 2014 li. All rights reserved.
// Use of this source code is governed by a MIT/X11
// license that can be found in the LICENSE file.

package slice

import (
	"container/list"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestAsList(t *testing.T) {
	l1 := AsList(1, 2, "3", "Hello, GoLang")

	l2 := list.New()
	l2.PushBack(1)
	l2.PushBack(2)
	l2.PushBack("3")
	l2.PushBack("Hello, GoLang")

	if !reflect.DeepEqual(l1, l2) {
		t.Fatal()
	}
}

func TestToList(t *testing.T) {
	s := []int{1, 2, 3, 4}

	l1 := list.New()
	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)
	l1.PushBack(4)

	l2 := ToList(s)
	l3 := ToList(s)

	if !reflect.DeepEqual(l1, l2) || !reflect.DeepEqual(l2, l3) {
		t.Fatal()
	}
}

func TestToSlice(t *testing.T) {
	l1 := list.New()
	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)
	l1.PushBack(4)

	if !reflect.DeepEqual(ToSlice(l1), []interface{}{1, 2, 3, 4}) {
		t.Fatal()
	}
}

func TestForeach(t *testing.T) {
	sum := 0
	Foreach([]int{1, 2, 3, 4}, func(i int) { sum += i })
	if sum != 10 {
		t.Fatal()
	}
}

func TestMap(t *testing.T) {
	r := Map([]int{1, 2, 3, 4}, func(i int) int { return i * 100 })
	if !reflect.DeepEqual(r, []interface{}{100, 200, 300, 400}) {
		t.Fatal()
	}
}

func TestMapInt(t *testing.T) {
	r := MapInt([]int{1, 2, 3, 4}, func(i int) int { return i * 100 })
	if !reflect.DeepEqual(r, []int64{100, 200, 300, 400}) {
		t.Fatal()
	}
}

func TestMapString(t *testing.T) {
	r := MapString([]string{"hello", "golang"}, func(s string) string { return strings.ToUpper(s) })
	if !reflect.DeepEqual(r, []string{"HELLO", "GOLANG"}) {
		t.Fatal()
	}
}

func TestMapFilter(t *testing.T) {
	r := MapFilter([]int{1, 2, 3, 4}, func(i int) (bool, string) {
		if i%2 == 0 {
			return true, fmt.Sprintf("%v", i*100)
		} else {
			return false, ""
		}
	})
	if !reflect.DeepEqual(r, []interface{}{"200", "400"}) {
		t.Fatal()
	}
}

func TestMapFilterInt(t *testing.T) {
	r := MapFilterInt([]string{"Hello", "Go"}, func(s string) (bool, int) {
		if len(s) < 3 {
			return true, len(s)
		} else {
			return false, 0
		}
	})

	if !reflect.DeepEqual(r, []int64{2}) {
		t.Fatal()
	}
}

func TestMapFilterString(t *testing.T) {
	r := MapFilterString([]int{1, 2, 3, 4}, func(i int) (bool, string) {
		if i%2 == 0 {
			return true, fmt.Sprintf("%v", i*100)
		} else {
			return false, ""
		}
	})
	if !reflect.DeepEqual(r, []string{"200", "400"}) {
		t.Fatal()
	}
}

func TestExist(t *testing.T) {
	r1 := Exist([]int{1, 2, 3, 4}, func(i int) bool { return i%3 == 0 })
	r2 := Exist([]int{1, 2, 3, 4}, func(i int) bool { return i%5 == 0 })
	if r1 == false {
		t.Fatal()
	}
	if r2 == true {
		t.Fatal()
	}
}

func TestFilter(t *testing.T) {
	rs := Filter([]int{1, 2, 3, 4}, func(i int) bool { return i%2 == 0 })
	if !reflect.DeepEqual([]interface{}{2, 4}, rs) {
		t.Fatal()
	}
}

func TestFilterInt(t *testing.T) {
	rs := FilterInt([]int{1, 2, 3, 4}, func(i int) bool { return i%2 == 0 })
	if !reflect.DeepEqual([]int{2, 4}, rs) {
		t.Fatal()
	}
}

func TestFilterString(t *testing.T) {
	rs := FilterString([]string{"ABC", "EFG"}, func(s string) bool { return strings.Contains(s, "A") })
	if !reflect.DeepEqual([]string{"ABC"}, rs) {
		t.Fatal()
	}
}

func TestIndex(t *testing.T) {
	i1 := Index([]int{1, 2, 3, 4, 6}, func(i int) bool { return i%3 == 0 })
	i2 := Index([]int{1, 2, 3, 4}, func(i int) bool { return i%5 == 0 })

	if i1 != 2 && i2 != -1 {
		t.Fatal()
	}
}

func TestIndexLast(t *testing.T) {
	i1 := IndexLast([]int{1, 2, 3, 4, 6}, func(i int) bool { return i%3 == 0 })
	i2 := IndexLast([]int{1, 2, 3, 4}, func(i int) bool { return i%5 == 0 })

	if i1 != 4 && i2 != -1 {
		t.Fatal()
	}
}

func TestFind(t *testing.T) {
	ok1, r1 := Find([]int{1, 2, 3, 4, 6}, func(i int) bool { return i%3 == 0 })
	ok2, _ := Find([]int{1, 2, 3, 4}, func(i int) bool { return i%5 == 0 })

	n1, _ := r1.(int)
	if ok1 != true || n1 != 3 {
		t.Fatal()
	}
	if ok2 != false {
		t.Fatal()
	}
}

func TestFindInt(t *testing.T) {
	ok1, r1 := FindInt([]int{1, 2, 3, 4, 6}, func(i int) bool { return i%3 == 0 })
	ok2, _ := FindInt([]int{1, 2, 3, 4}, func(i int) bool { return i%5 == 0 })

	if ok1 != true || r1 != 3 {
		t.Fatal()
	}
	if ok2 != false {
		t.Fatal()
	}
}

func TestFindString(t *testing.T) {
	ok1, s := FindString([]string{"ABC", "EFG", "AXY"}, func(s string) bool { return strings.Contains(s, "A") })
	if ok1 == false || s != "ABC" {
		t.Fatal()
	}
}

func TestFindLast(t *testing.T) {
	ok1, r1 := FindLast([]int{1, 2, 3, 4, 6}, func(i int) bool { return i%3 == 0 })
	ok2, r2 := FindLast([]int{6, 1, 2, 3, 4}, func(i int) bool { return i%6 == 0 })

	if ok1 != true || r1 != 6 {
		t.Fatal()
	}

	if ok2 == false || r2 != 6 {
		t.Fatal()
	}
}

func TestFindLastInt(t *testing.T) {
	ok1, r1 := FindLastInt([]int{1, 2, 3, 4, 6}, func(i int) bool { return i%3 == 0 })
	ok2, r2 := FindLastInt([]int{6, 1, 2, 3, 4}, func(i int) bool { return i%6 == 0 })

	if ok1 != true || r1 != 6 {
		t.Fatal()
	}

	if ok2 == false || r2 != 6 {
		t.Fatal()
	}
}

func TestFindLastString(t *testing.T) {
	ok1, s := FindLastString([]string{"ABC", "EFG", "AXY"}, func(s string) bool { return strings.Contains(s, "A") })
	if ok1 == false || s != "AXY" {
		t.Fatal()
	}
}

func TestJoin(t *testing.T) {
	type p struct {
		name string
		age  int
	}

	s := Join([]interface{}{1, 2, "3", 4.0, p{"li", 25}}, ",")
	if s != "[1,2,3,4,{li 25}]" {
		t.Fatal()
	}
	if Join([]interface{}{}, ",") != "[]" {
		t.Fatal()
	}
}
