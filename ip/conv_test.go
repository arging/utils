// Copyright 2014 li. All rights reserved.
// Use of this source code is governed by a MIT/X11
// license that can be found in the LICENSE file.

package ip

import (
	"math"
	"testing"
)

func assertTrue(rs bool, msg string, t *testing.T) {
	if !rs {
		t.Log(msg)
	}
}

func assertFalse(rs bool, t *testing.T) {
	if rs {
		t.Log("")
	}
}

func TestIsIpv4(t *testing.T) {
	assertFalse(IsIpv4("192.168.1"), t)
	assertFalse(IsIpv4("192.168.1.1.1"), t)
	assertFalse(IsIpv4("192.168..1"), t)
	assertFalse(IsIpv4("192.168,1.1"), t)
	assertFalse(IsIpv4("256.168.1.1"), t)
	assertFalse(IsIpv4("-1.168,1.1"), t)
	assertFalse(IsIpv4("-1.168.1.1"), t)
	assertFalse(IsIpv4("2E.168,1.1"), t)
	assertTrue(IsIpv4("2.168.1.1"), "", t)
	assertTrue(IsIpv4("192.168.1.1"), "", t)
}

func TestIpv4ToInt(t *testing.T) {
	//bad case
	_, err1 := Ipv4ToInt("192.168.1")
	_, err2 := Ipv4ToInt("192.168.1.1.1")
	_, err3 := Ipv4ToInt("192.168..1")
	_, err4 := Ipv4ToInt("192.168,1.1")
	_, err5 := Ipv4ToInt("256.168,1.1")
	_, err6 := Ipv4ToInt("-1.168,1.1")

	assertTrue(err1 != nil, "192.168.1 bad ip", t)
	assertTrue(err2 != nil, "192.168.1.1.1 bad ip", t)
	assertTrue(err3 != nil, "192.168..1 bad ip", t)
	assertTrue(err4 != nil, "192.168,1.1 bad ip", t)
	assertTrue(err5 != nil, "256.168,1.1 bad ip", t)
	assertTrue(err6 != nil, "-1.168,1.1 bad ip", t)

	//normal case
	intV7, err7 := Ipv4ToInt("0.0.0.0")
	assertTrue(err7 == nil && intV7 == 0, "0.0.0.0 good ip", t)

	intV8, err8 := Ipv4ToInt("255.255.255.255")
	ip1 := int64(math.Pow(2, 24))
	ip2 := int64(math.Pow(2, 16))
	ip3 := int64(math.Pow(2, 8))
	ip4 := int64(1)
	assertTrue(err8 == nil && intV8 == (ip1+ip2+ip3+ip4)*255, "255.255.255.255 good ip", t)
}

func TestIntToIpv4(t *testing.T) {
	ip4Int1, _ := Ipv4ToInt("0.0.0.0")
	ip4Int2, _ := Ipv4ToInt("163.87.54.33")
	ip4Int3, _ := Ipv4ToInt("255.255.255.255")

	ip1, _ := IntToIpv4(ip4Int1)
	ip2, _ := IntToIpv4(ip4Int2)
	ip3, _ := IntToIpv4(ip4Int3)

	assertTrue(ip1 == "0.0.0.0", "ip1", t)
	assertTrue(ip2 == "163.87.54.33", "ip2", t)
	assertTrue(ip3 == "255.255.255.255", "ip3", t)
}

func TestIsIpv4In(t *testing.T) {
	assertTrue(ipIn("163.87.54.33", "0.0.0.0", "255.255.255.255"), "", t)
	assertTrue(ipIn("0.0.0.0", "0.0.0.0", "255.255.255.255"), "", t)
	assertTrue(ipIn("255.255.255.255", "0.0.0.0", "255.255.255.255"), "", t)
	assertTrue(ipIn("163.87.54.33", "163.87.54.33", "163.87.54.33"), "", t)
	assertTrue(ipIn("163.87.54.33", "163.87.54.32", "163.87.54.33"), "", t)
	assertTrue(ipIn("163.87.54.33", "163.87.54.32", "163.87.54.34"), "", t)
	assertTrue(!ipIn("163.87.54.33", "163.87.54.34", "163.87.54.33"), "", t)
	assertTrue(!ipIn("163.87.54.33", "163.87.54.33", "163.87.54.32"), "", t)
}

// help function
func ipIn(ip string, left string, right string) bool {
	rs, _ := IsIpv4In(ip, left, right)
	return rs
}
