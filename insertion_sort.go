package main

import (
  "fmt"
  "./utils"
)

func insertion_sort(v []int, n int) {
  var i, j int;

  for i = 1; i < n; i++ {
    j = i;

    for j > 0 && v[j] < v[j - 1] {
      utils.Swap(&v[j], &v[j - 1])
      j--
    }
  }
}

func main() {
  v := []int{3,2,4,1}

  insertion_sort(v, len(v))

  for _, n := range v {
    fmt.Println(n)
  }
}
