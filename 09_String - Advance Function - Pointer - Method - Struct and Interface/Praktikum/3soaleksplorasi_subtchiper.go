package main

import (

	"fmt"
)

type student struct {
	name string
	nameEncode string
	score int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
  var en string

  for _, b := range []byte(s.name) {

  	n := int(b) - 97 // start at 0
  	k := func() int {

  		if n <= 12 {

  			return n
  		}

		// 13 - 25, 12 - 0
  		return 12 - (n - 13)
  	}()
  	z := 25 - (k * 2)

	if n <= 12 { n += z } else { n += -z }

  	en += string(byte(n + 97))
  }
	
  return en
}

func (s *student) Decode() string {
  var de string

  for _, b := range []byte(s.name) {

  	n := int(b) - 97 // start at 0
  	k := func() int {

  		if n <= 12 {

  			return n
  		}

		// 13 - 25, 12 - 0
  		return 12 - (n - 13)
  	}()
  	z := 25 - (k * 2)

	if n <= 12 { n += z } else { n += -z }

  	de += string(byte(n + 97))
  }
	
  return de
}

func main() {

  // rizky is irapb, flip half

  var menu int
  var a student = student{}
  var c Chiper = &a

  fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
  fmt.Scan(&menu)
	
  if menu == 1 {
    fmt.Print("\nInput Student Name: ")
    fmt.Scan(&a.name)
    fmt.Print("\nEncode of student’s name " + a.name + " is : " + c.Encode())
  } else if menu == 2 {
    fmt.Print("\nInput Student Name: ")
    fmt.Scan(&a.name)
    fmt.Print("\nDecode of student’s name " + a.name + " is : " + c.Decode())
  }

  fmt.Println()
}
