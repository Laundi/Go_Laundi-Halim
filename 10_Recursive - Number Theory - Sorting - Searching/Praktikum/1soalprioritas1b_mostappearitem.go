package main

import (
	"fmt"
)

type Pair struct {

	name string
	count int
}

func (this *Pair)GetName() string {

	return this.name
}

func (this *Pair)SetName(name string) {

	this.name = name
}

func (this *Pair)GetCount() int {

	return this.count
}

func (this *Pair)SetCount(count int) {

	this.count = count
}

func Contains[K comparable, V any](data map[K]V, key K) bool {

	for k := range data {

		if k == key { return true }
	}

	return false
}

func MostAppearItem(items []string) []Pair {

	temp := make(map[string]*Pair)

	for _, v := range items {

		if Contains(temp, v) {

			p := temp[v]
			p.SetCount(p.GetCount() + 1)

			continue
		}

		temp[v] = &Pair{v, 1}
	}

	result := make([]Pair, 0)

	for _, pair := range temp {

		if len(result) == 0 {

			// setup - initial
			result = append(result, *pair)
			continue
		}

		result = append(result, *pair)

		n := len(result)
		k := n - 1
		o := result[k] // last sample

		for i := 0; i < k; i++ {

			v := result[i]
			a := o.count
			b := v.count
			
			if a <= b { 

				result[k], result[i] = result[i], result[k]
				o = v
			}
		}
	}

	return result
}

func main() {

	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}
