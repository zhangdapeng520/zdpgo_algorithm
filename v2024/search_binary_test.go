package v2024_test

import (
	v2024 "github.com/zhangdapeng520/zdpgo_algorithm/v2024"
	"testing"
)

func TestSearchBinary_Basic(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	target := 3
	index := v2024.SearchBinary(arr, target)
	t.Log(index, index == 2)
}
