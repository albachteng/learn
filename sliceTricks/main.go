package main

func main() {

}

func makeCopy[T any](a []T) []T {
  return append([]T(nil), a...)
}

func appendVector[T any](a, b []T) []T {
  return append(a, b...)
}

func deleteAt[T any](a []T, i int) []T {
  return append(a[:i], a[i+1:]...)
}

func cut[T any](a []T, i, j int) []T {
  return append(a[:i], a[j:]...)
}

// presumably faster?
func deleteWithoutOrder[T any](a []T, i int) []T{
  a[i] = a[len(a) - 1] // move last item to i
  a = a[:len(a) -1 ] // remove extra item at the end
  return a
}

func :
