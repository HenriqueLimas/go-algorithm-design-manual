package main

import "fmt"

func selection_sort(v []int, n int) {
  var i,j int;
  var min int;

  for i = 0; i < n; i++ {
    min = i

    for j = i + 1; j < n; j++ {
      if (v[j] < v[min]) {
        min = j
      }
    }

    swap(&v[min], &v[i])
  }
}

func swap(a, b *int) {
  aux := *b
  *b = *a
  *a = aux
}

func main() {
  v := []int{3,2,4,1}

  selection_sort(v, len(v))

  for _, n := range v {
    fmt.Println(n)
  }
}
