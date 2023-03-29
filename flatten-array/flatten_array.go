package flatten

import "fmt"

func Flatten(nested interface{}) []interface{} {
	for x := range nested {
		fmt.Println(x)
	}
}
