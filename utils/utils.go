package utils

func Swap(a, b *int) {
  aux := *b
  *b = *a
  *a = aux
}

