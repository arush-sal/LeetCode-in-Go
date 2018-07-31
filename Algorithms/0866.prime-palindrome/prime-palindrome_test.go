package problem0866

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N   int
	ans int
}{

	{
		6,
		7,
	},

	{
		8,
		11,
	},

	{
		13,
		101,
	},

	// 可以有多个 testcase
}

func Test_primePalindrome(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, primePalindrome(tc.N), "输入:%v", tc)
	}
}

func Benchmark_primePalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			primePalindrome(tc.N)
		}
	}
}
