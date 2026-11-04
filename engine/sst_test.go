package engine_test

import (
	"fmt"
	"testing"
)

func SSTTest(t *testing.T) {
	key := "hello"
	if key == "" {
		fmt.Println("the key is not valid")
	}
}
