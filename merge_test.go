package gomerge

import (
	"fmt"
	"testing"
)

type Body struct {
	Height int `json:"height"`
	Weight int `json:"weight"`
}

type Job struct {
	Company string `json:"company"`
}

type People struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Age  int    `json:"age"`
	Body Body   `json:"body"`
	Job
}

var tom = People{
	Name: "Tom",
	Sex:  "male",
	Age:  18,
}

type P map[string]interface{}

var newToms = map[string]P{
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
		"body": map[string]interface{}{
			"height": 170,
			"weight": 80,
		},
		"company": "github",
	},
}

var results = map[string]People{
	"-1": {Name: "Tom", Sex: "female", Age: 18},
	"2":  {Name: "Tom", Sex: "male", Age: 19},
	"3":  {Name: "Tom", Sex: "female", Age: 18, Body: Body{Height: 170, Weight: 80}, Job: Job{Company: "github"}},
}

// merge tom and new tom
// dst as pointer
func TestMerge(t *testing.T) {
	value := newToms["2"]
	newTom := tom
	err := Merge(&newTom, value)
	if err != nil {
		fmt.Println("Merge with error:", err)
		t.Fail()
	}
	if results["2"] != newTom {
		fmt.Printf("Merge result not right, expect %v, but %v", results["2"], newTom)
		t.Fail()
	}
}

func TestMergeEmbedded(t *testing.T) {
	value := newToms["3"]
	newTom := tom
	err := Merge(&newTom, value)
	if err != nil {
		fmt.Println("Merge with error:", err)
		t.Fail()
		return
	}
	if results["3"] != newTom {
		fmt.Printf("Merge result not right, expect %v, but %v", results["3"], newTom)
		t.Fail()
	}
}

func TestMergeFailed(t *testing.T) {
	value := newToms["-1"]
	newTom := tom
	err := Merge(&newTom, value)
	if err == nil {
		t.Fail()
	}
}
