package main

import "fmt"

const x int = 3
const y int = 4
const z int = 3

func multiply_matrix(A [][]int, B [][]int) [x][z]int {
  var C [x][z]int

  for i := 0; i < x; i++ {
    for j := 0; j < z; j++ {
      C[i][j] = 0
      for k := 0; k < y; k++ {
        C[i][j] += A[i][k] * B[k][j]
      }
    }
  }

  return C
}

func main() {
  A := [][]int{
    []int{2,4,1,4},
    []int{1,3,5,2},
    []int{4,5,6,7},
  }

  B := [][]int{
    []int{4,1,4},
    []int{3,5,2},
    []int{4,6,7},
    []int{3,6,3},
  }

  C := multiply_matrix(A,B)

  for i := 0; i < x; i++ {
    for j := 0; j < z; j++ {
      fmt.Printf(" %d ", C[i][j])
    }

    fmt.Printf("\n")
  }
}
