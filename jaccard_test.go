package zdpgo_algorithm_test

import (
	"github.com/zhangdapeng520/zdpgo_algorithm"
	"testing"
)

func TestJaccard_Basic(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{4, 5, 6, 7}
	t.Log(zdpgo_algorithm.Jaccard(arr1, arr2))
}

func TestJaccardSorted_Basic(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{4, 5, 6, 7}
	t.Log(zdpgo_algorithm.JaccardSorted(arr1, arr2))
}
