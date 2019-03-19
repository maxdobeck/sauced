package main

// http://ghodss.com/2014/the-right-way-to-handle-yaml-in-golang/

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

// T is an example struct
type T struct {
	Name string `yaml:"name"`
	Arg1 string `yaml:"arg1"`
	Arg2 int    `yaml:"arg2"`
	Arg3 []int  `yaml:"arg3"`
}

func main() {
	data := []byte(`
name: My first test
arg1: bumpitybump
arg2: 2
arg3: [3, 4]
`)

	t := T{}
	err := yaml.Unmarshal(data, &t)
	if err != nil {
		fmt.Print("uh oh", err)
	}
	fmt.Println(t)
	fmt.Println(t.Name)

	d, err := yaml.Marshal(&t)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(string(d))
	// Cannot do a d.Name call here.  Its not a struct but a YAML object

	fmt.Printf("hello, world\n")
	for {
	}
}
