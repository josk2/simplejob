package test

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := make(map[string][]string)
	m["A"] = []string{"Anna", "AI", "Andy", "Alpha"}
	m["B"] = []string{"Bottom", "Bi", "Brother"}
	m["C"] = []string{"Chris", "Cherry", "Carrot"}

	vals, ok := m["A"]
	if ok {
		for _, v := range vals {
			fmt.Println(v)
		}
	}

	val2, ok := m["D"]
	fmt.Println(len(val2))

}
