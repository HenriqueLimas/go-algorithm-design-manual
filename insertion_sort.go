package main

import "fmt"

func insertion_sort(v []int, n int) {
  var i, j int;

  for i = 1; i < n; i++ {
    j = i;

    for j > 0 && v[j] < v[j - 1] {
      swap(&v[j], &v[j - 1])
      j--
    }
  }
}

func swap(a *int, b *int) {
  aux := *a
  *a = *b
  *b = aux
}

func main() {
  v := []int{3,2,4,1}

  insertion_sort(v, len(v))

  for _, n := range v {
    fmt.Println(n)
  }
}
