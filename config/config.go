// Copyright (c) li. All rights reserved.
// Use of this source code is governed by a MIT/X11
// license that can be found in the LICENSE file.

package config

import (
	"bufio"
	"github.com/roverli/utils/errors"
	"io"
	"os"
	"strconv"
	"strings"
)

// Read options from the file.
func Load(fname string) (*Config, errors.Error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, errors.Wrapf(err, "cann't open file %s", fname)
	}

	reader := bufio.NewReader(file)
	options := make(map[string]string)
	for {

		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, errors.Wrapf(err, "read file %s error.", fname)
		}

		line = strings.TrimSpace(line)

		if len(line) == 0 || line[0] == '#' {
			continue
		}

		i := strings.Index(line, "=")
		if i < 0 {
			return nil, errors.New("parse error: " + line)
		}

		options[strings.TrimSpace(line[:i])] = strings.TrimSpace(line[i+1:])
	}

	if err = file.Close(); err != nil {
		return nil, errors.Wrapf(err, "close file %s error.", fname)
	}

	return &Config{options: options}, nil
}

// Return a config instance with empty options.
func New() *Config {
	return &Config{options: make(map[string]string)}
}

// Config is not safe for multiply goroutines access.
type Config struct {
	options map[string]string
}

// IsEmpty check whether the configuration is empty.
// Return true if the configuration contains no property, false otherwise.
func (c *Config) IsEmpty() bool {
	return len(c.options) == 0
}

// Remove all options from the configuration.
func (c *Config) Clear() {
	c.options = make(map[string]string)
}

// Remove a option from the configuration.
func (c *Config) ClearOption(key string) {
	delete(c.options, key)
}

// Set a option, this will replace any previously set values.
// If previous value does not exist, eq to add a new option.
func (c *Config) SetOption(key string, value string) {
	c.options[key] = value
}

// Get the list of the keys contained in the configuration.
// The returned slice can be used to obtain all defined keys.
func (c *Config) Keys() []string {
	i := 0
	keys := make([]string, len(c.options))
	for k, _ := range c.options {
		keys[i] = k
		i++
	}
	return keys
}

// String gets the string value for the given key in the configuration.
// It returns default value if the key does not exist.
func (c *Config) String(key string, defaultv string) string {
	if v, ok := c.options[key]; ok {
		return v
	}

	// Support get config from env.
	v := os.Getenv(key)
	if v != "" {
		return v
	}
	return defaultv
}

// Bool gets the bool value for the given key in the configuration.
// It returns default value if the key does not exist or cann't be converted to bool.
func (c *Config) Bool(key string, defaultv bool) bool {
	switch strings.ToLower(c.String(key, "")) {
	case "y", "on", "1":
		return true
	case "n", "off", "0":
		return false
	default:
		return defaultv
	}
}

// Float gets the float value for the given key in the configuration.
// It returns default value if the key does not exist or cann't be converted to float64.
func (c *Config) Float(key string, defaultv float64) float64 {
	v, err := strconv.ParseFloat(c.String(key, ""), 64)
	if err == nil {
		return v
	} else {
		return defaultv
	}
}

// Int gets the int value for the given key in the configuration.
// It returns default value if the key does not exist or cann't be converted to int.
func (c *Config) Int(key string, defaultv int) int {
	v, err := strconv.Atoi(c.String(key, ""))
	if err == nil {
		return v
	} else {
		return defaultv
	}
}

// Merge the target config options in current config.
// If target config has the same key with current config, the value was ignored.
func (c *Config) Merge(target *Config) {
	kvs := target.toKvs()

	for _, kv := range kvs {
		k := kv[0]
		v := kv[1]
		if _, ok := c.options[k]; !ok {
			c.options[k] = v
		}
	}
}

// Inner method, for "Merge" method.
func (c *Config) toKvs() [][2]string {

	i := 0
	kvs := make([][2]string, len(c.options))
	for k, v := range c.options {
		kvs[i] = [2]string{k, v}
		i++
	}
	return kvs
}
