// Copyright 2014 li. All rights reserved.
// Use of this source code is governed by a MIT/X11
// license that can be found in the LICENSE file.

package containers

import (
	"container/list"
)

// Check if the list is empty.
// Return true if list is nil or length eq 0.
func IsEmptyList(l *list.List) bool {
	return l == nil || l.Len() == 0
}

// Check if the list is not empty.
// Return true only if the list has at least one element.
func IsNotEmptyList(l *list.List) bool {
	return !IsEmptyList(l)
}
