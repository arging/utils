// Copyright 2014 li. All rights reserved.
// Use of this source code is governed by a MIT/X11
// license that can be found in the LICENSE file.

package hash

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"sync/atomic"
	"time"
)

// objectIdCounter is atomically incremented when generating a new ObjectId
// using NewObjectId() function. It's used as a counter part of an id.
var objectIdCounter uint32 = 0

var machineId = readMachineId()

// initMachineId generates machine id and puts it into the machineId global
// variable. If this function fails to get the hostname, it will cause
// a runtime error.
func readMachineId() []byte {
	var sum [3]byte
	id := sum[:]
	hostname, err1 := os.Hostname()
	if err1 != nil {
		_, err2 := io.ReadFull(rand.Reader, id)
		if err2 != nil {
			panic(fmt.Errorf("cannot get hostname: %v; %v", err1, err2))
		}
		return id
	}
	hw := md5.New()
	hw.Write([]byte(hostname))
	copy(id, hw.Sum(nil))
	return id
}

func UniqueId() string {
	var b [12]byte

	// Timestamp, 4 bytes, big endian
	binary.BigEndian.PutUint32(b[:], uint32(time.Now().Unix()))

	// Machine, first 3 bytes of md5(hostname)
	b[4] = machineId[0]
	b[5] = machineId[1]
	b[6] = machineId[2]

	// Pid, 2 bytes, specs don't specify endianness, but we use big endian.
	pid := os.Getpid()
	b[7] = byte(pid >> 8)
	b[8] = byte(pid)

	// Increment, 3 bytes, big endian
	i := atomic.AddUint32(&objectIdCounter, 1)

	b[9] = byte(i >> 16)
	b[10] = byte(i >> 8)
	b[11] = byte(i)
	return string(b[:])
}
