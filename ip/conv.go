// Copyright 2014 li. All rights reserved.
// Use of this source code is governed by a MIT/X11
// license that can be found in the LICENSE file.

package ip

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	ipv4Min = 0
	ipv4Max = 255
)

const (
	ipSep   = "."
	ipV4Len = 4
)

// Value is not a correct ipv4.
var (
	BadIpv4Error = fmt.Errorf("utils/ip: bad ipv4 value.")
)

var (
	ip1x      = int64(math.Pow(2, 24)) * 255
	ip2x      = int64(math.Pow(2, 16)) * 255
	ip3x      = int64(math.Pow(2, 8)) * 255
	ip4x      = int64(math.Pow(2, 0)) * 255
	ipv4Shift = [4]uint8{24, 16, 8, 0}
)
var (
	ipv4MaxInt, _ = Ipv4ToInt("255.255.255.255")
	ipv4MinInt, _ = Ipv4ToInt("0.0.0.0")
)

// Test whether a correct ipv4 string
func IsIpv4(s string) bool {
	ips := strings.Split(s, ipSep)
	if len(ips) != ipV4Len {
		return false
	}
	for _, v := range ips {
		num, e := strconv.Atoi(v)
		if e != nil || num > ipv4Max || num < ipv4Min {
			return false
		}
	}

	return true
}

// Convert a ip string to int64.
// If ip string is not a corret ipv4, error.
func Ipv4ToInt(ip string) (intv int64, err error) {
	ips := strings.Split(ip, ipSep)
	if len(ips) != ipV4Len {
		err = BadIpv4Error
		return
	}

	for i, v := range ips {
		num, e := strconv.Atoi(v)
		if e != nil || num < ipv4Min || num > ipv4Max {
			err = BadIpv4Error
			return
		}
		intv += int64(num) << ipv4Shift[i]
	}
	return
}

// Convert a int64 value to ipv4 string
// If int64 value isn't in the rang of ipv4 value, error.
func IntToIpv4(intv int64) (ip string, err error) {
	if intv < ipv4MinInt || intv > ipv4MaxInt {
		err = BadIpv4Error
		return
	}

	ip = strings.Join([]string{
		strconv.Itoa(int(intv & ip1x >> ipv4Shift[0])),
		strconv.Itoa(int(intv & ip2x >> ipv4Shift[1])),
		strconv.Itoa(int(intv & ip3x >> ipv4Shift[2])),
		strconv.Itoa(int(intv & ip4x >> ipv4Shift[3]))},
		ipSep)

	return
}

// Test is ip in the ip segment.
// Both left and right is included.
func IsIpv4In(ip string, left string, right string) (in bool, err error) {

	ipInt, e := Ipv4ToInt(ip)
	if e != nil {
		err = BadIpv4Error
		return
	}

	leftInt, e := Ipv4ToInt(left)
	if e != nil {
		err = BadIpv4Error
		return
	}

	rightInt, e := Ipv4ToInt(right)
	if e != nil {
		err = BadIpv4Error
		return
	}

	in = ipInt >= leftInt && ipInt <= rightInt
	return
}
