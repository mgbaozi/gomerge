package gomerge

import (
	"strings"
	"testing"
)

type People struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Age  int    `json:"age"`
}

var tom = People{
	Name: "Tom",
	Sex:  "male",
	Age:  18,
}

type P map[string]interface{}

var new_toms = map[string]P{
	"-1": {
		"name": "Tom",
		"sex":  "female",
		"age":  "19",
	},
	"2": P{
		"name": "Tom",
		"age":  19,
	},
	"3": P{
		"name": "Tom",
		"sex":  "female",
	},
}

var results = map[string]People{
	"-1": {Name: "Tom", Sex: "female", Age: 18},
	"2":  {Name: "Tom", Sex: "male", Age: 19},
	"3":  {Name: "Tom", Sex: "female", Age: 18},
}

// merge tom and new tom
// dst as pointer
func TestMerge(t *testing.T) {
	for k, v := range new_toms {
		new_tom := tom
		err := Merge(&new_tom, v)
		if strings.HasPrefix(k, "-") == (err == nil) {
			t.Fail()
		}
		if results[k] != new_tom {
			t.Fail()
		}
	}
}
