package bit

import (
	"fmt"
	"testing"
)

func TestReverseBits(t *testing.T) {
	var num uint32 = 43261596
	ret := ReverseBits(num)
	fmt.Println(ret)
}
