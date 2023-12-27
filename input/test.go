/*
test test Big comment

// fjgjgjg

***
///
*/

/*
 */
package input

import (
	"fmt"
	"math/big"
)

// test function
func test() {
	var x = 10
	y := 0xfa10

	var text string = "123"
	cnt := 0
	var flag bool = true
	for _, r := range text {
		if r == 'a' && flag {
			cnt++
		}
	}

	fmt.Println(x + y + cnt)

	a := new(big.Int).SetUint64(uint64(10))
	fmt.Println(a)

	type A struct {
		x int
		y float64
	}

	aval := A{
		x: 10,
		y: 0.10,
	}

	fmt.Println(aval.y)
}

/*
Big comment

*/
