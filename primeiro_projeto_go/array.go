package main

import "fmt"

var slice []string
var array [5]int

func main() {
  slice = append(slice, "a") // slice é o index + valor
  slice = append(slice, "b")
  slice = append(slice, "c")
  array[0] = 1
  array[1] = 2
  array[2] = 3
  array[3] = 4
  array[4] = 5

  for i, v := range slice { // atribui index e valor
	fmt.Println(i, v)
  }	

  for  i := 0; i < len(array); i++ {
	fmt.Println(array[i])
  }
 
}