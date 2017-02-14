package main

import "fmt"

func findmatch(p string, t string) int {
  var i, j int;
  var p_len, t_len int;

  p_len = len(p)
  t_len = len(t)

  for i = 0 ; i < (t_len - p_len); i++ {
    j = 0

    for j < p_len && t[i + j] == p[j] {
      j++
    }

    if j == p_len {
      return i
    }
  }

  return -1
}

func main() {
  t := "This is a long text"
  p := "long"

  fmt.Println(findmatch(p, t))
}
