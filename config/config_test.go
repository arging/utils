// Copyright (c) li. All rights reserved.
// Use of this source code is governed by a MIT/X11
// license that can be found in the LICENSE file.

package config

import (
	"testing"
)

func loadConfig(t *testing.T) *Config {
	conf, err := Load("testdata/read.conf")
	if err != nil {
		t.Fatal("Read config from testdata/read.conf should not error.")
	}
	return conf
}

func TestLoad(t *testing.T) {
	conf := loadConfig(t)

	if len(conf.options) != 4 {
		t.Error("Expected config only has 4 options.")
	}

	name := conf.String("name", "")
	if name != "tom" {
		t.Errorf("Expected name property to be tom, but was %v", name)
	}

	age := conf.Int("age", -1)
	if age != 25 {
		t.Errorf("Expected age property to be 25, but was %v", age)
	}

	isMan := conf.Bool("man", false)
	if isMan != true {
		t.Errorf("Expected man property to be true, but was %v", isMan)
	}

	height := conf.Float("height", -1)
	if height != 1.7 {
		t.Errorf("Expected height property to be 1.7, but was %v", height)
	}
}

func TestIsEmpty(t *testing.T) {
	conf := New()
	if !conf.IsEmpty() {
		t.Error("Expected config to be empty.")
	}

	conf.SetOption("name", "li")
	if conf.IsEmpty() {
		t.Error("Expected config to be not empty.")
	}
}

func TestClearOption(t *testing.T) {
	conf := loadConfig(t)
	conf.ClearOption("name")

	v := conf.String("name", "")
	if v != "" {
		t.Errorf("Expected name to be deleted.")
	}
}

func TestClear(t *testing.T) {
	conf := loadConfig(t)
	conf.Clear()

	if !conf.IsEmpty() {
		t.Errorf("Expected all config options to be deleted.")
	}
}

func TestSetOption(t *testing.T) {
	conf := loadConfig(t)
	conf.Clear()

	conf.SetOption("name", "li")
	name := conf.String("name", "")
	if name != "li" {
		t.Errorf(`Expected name to be "li", but was %v`, name)
	}

	conf.SetOption("aNewKey", "aNewValue")
	value := conf.String("aNewKey", "")
	if value != "aNewValue" {
		t.Errorf(`Expected aNewKey to be "aNewValue", but was %v`, value)
	}
}

func TestKeys(t *testing.T) {
	conf := loadConfig(t)
	keys := conf.Keys()
	if len(keys) != 4 {
		t.Error("Expected keys length is 4.")
	}
	for _, k := range keys {
		switch k {
		case "name", "age", "height", "man":
		default:
			t.Errorf(`Expected keys are in ("name", "age", "height", "man"), but was %v`, k)
		}
	}
}

func TestMerge(t *testing.T) {
	conf1 := loadConfig(t)
	conf2 := New()

	conf2.SetOption("name", "li")
	conf2.SetOption("city", "Tokyo")

	conf1.Merge(conf2)
	if len(conf1.options) != 5 {
		t.Error("Expected conf1 len to be 5.")
	}

	name := conf1.String("name", "")
	if name != "tom" {
		t.Error(`Expected name to be "tom".But was %v`, name)
	}
	city := conf1.String("city", "")
	if city != "Tokyo" {
		t.Error(`Expected city to be "Tokyo".But was %v`, city)
	}
}
