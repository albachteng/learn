package arrays

func Sum(nums []int) int {
	sum := 0
	for _, number := range nums {
		sum += number
	}
	return sum
}

func SumAll(arrs ...[]int) []int {
  var sums []int

  for _, arr := range arrs {
    sums = append(sums, Sum(arr))
  }
  return sums
}

func SumAllTails(arrs ...[]int) []int {
  var sums []int

  for _, arr := range arrs {
    tail := arr[1:]
    sums = append(sums, Sum(tail))
  }
  return sums;
}
