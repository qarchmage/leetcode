package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2}
	b := []int{3, 4}
	fmt.Println(findMedianSortedArrays(a, b))

}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	ln1 := len(nums1) - 1
	if ln1%2 == 0 {
		return float64(nums1[ln1/2])
	} else {
		// add two middle elems
		return float64(nums1[ln1/2]+nums1[(ln1/2)+1]) / 2.0
	}
}
