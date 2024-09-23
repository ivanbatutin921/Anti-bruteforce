package tests

import (
	
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	tests.InitDB()
	defer tests.CloseDB()
	os.Exit(m.Run())
}
