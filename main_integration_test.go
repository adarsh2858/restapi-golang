// +build integration

package main

import (
	"fmt"
	"testing"
)

func TestIntegration(t *testing.T) {
	message := "Integration testing start\n"
	fmt.Printf("%s", message)
}
