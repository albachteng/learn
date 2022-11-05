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
