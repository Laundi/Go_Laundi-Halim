package main

import (
	"fmt"
)

func contains[T comparable](data []T, v T) bool {

	for _, value := range data {

		if value == v {

			return true
		}
	}

	return false
}

func ArrayMerge(a, b []string) []string {

  for _, v := range b {

  	if !contains(a, v) {

  		a = append(a, v)
  	}
  }

  return a
}



func main() {

	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yohimitsu"}, []string{"devil jin", "yohimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
