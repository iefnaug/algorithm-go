package bit

import (
	"fmt"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	distance := HammingDistance(1, 4)
	//distance := HammingDistance(-4, 4)
	fmt.Println(distance)

}
