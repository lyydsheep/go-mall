package demo

import (
	"fmt"
	"os"
	"testing"
)

func Test(t *testing.T) {
	fileName, ok := os.LookupEnv("fileName")
	if !ok {
		fmt.Println("???")
	} else {
		fmt.Println(fileName)
	}
}
