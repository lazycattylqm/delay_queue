package test_temp

import (
	"fmt"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println("test name for a")
}

func TestNameB(t *testing.T) {
	fmt.Println("test name for b")
}

func TestNameC(t *testing.T) {
	fmt.Println("test name for c")
}

func TestMain(m *testing.M) {
	printNow()
	code := m.Run()
	os.Exit(code)
}

func printNow() {
	fmt.Println("time.Now()")
}
