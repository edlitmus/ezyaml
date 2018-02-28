/*
  Copyright (c) 2012 José Carlos Nieto, http://xiam.menteslibres.org/

  Permission is hereby granted, free of charge, to any person obtaining
  a copy of this software and associated documentation files (the
  "Software"), to deal in the Software without restriction, including
  without limitation the rights to use, copy, modify, merge, publish,
  distribute, sublicense, and/or sell copies of the Software, and to
  permit persons to whom the Software is furnished to do so, subject to
  the following conditions:

  The above copyright notice and this permission notice shall be
  included in all copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
  OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
  WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package yaml

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/esilva-everbridge/dig"
	yaml "gopkg.in/yaml.v2"
)

// Yaml data type
type Yaml struct {
	File   string
	Values map[string]interface{}
}

// New creates and returns a YAML struct.
func New() *Yaml {
	self := &Yaml{}
	self.Values = map[string]interface{}{}
	return self
}

// Open creates and returns a YAML struct, from a file.
func Open(file string) (*Yaml, error) {
	var err error

	self := New()

	_, err = os.Stat(file)

	if err != nil {
		return nil, err
	}

	self.File = file

	err = self.Read(self.File)

	if err != nil {
		return nil, err
	}

	return self, nil
}

// Set sets a YAML setting
func (y *Yaml) Set(params ...interface{}) error {

	l := len(params)

	if l < 2 {
		return fmt.Errorf("missing value")
	}

	route := params[0 : l-1]
	value := params[l-1]

	err := dig.Dig(&y.Values, route...)
	if err != nil {
		return err
	}
	return dig.Set(&y.Values, value, route...)
}

// Get returns a YAML setting
func (y *Yaml) Get(route ...interface{}) interface{} {
	var i interface{}

	err := dig.Get(&y.Values, &i, route...)
	if err != nil {
		if debug := os.Getenv("YAML_DEBUG"); debug != "" {
			log.Print(err)
		}
	}

	return i
}

// Save writes changes to the currently opened YAML file.
func (y *Yaml) Save() error {
	if y.File != "" {
		return y.Write(y.File)
	}

	return fmt.Errorf("no file specified")
}

// Write writes the current YAML struct to disk.
func (y *Yaml) Write(filename string) error {

	out, err := yaml.Marshal(y.Values)

	if err != nil {
		return err
	}

	fp, err := os.Create(filename)

	if err != nil {
		return err
	}

	_, err = fp.Write(out)
	if err != nil {
		return err
	}

	return fp.Close()
}

// Read loads a YAML file from disk.
func (y *Yaml) Read(filename string) error {
	_, err := os.Stat(filename)
	if err != nil {
		return err
	}

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(buf, &y.Values)
}
