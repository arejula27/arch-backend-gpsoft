package main
import (
    "testing")

func TestSomething(t *testing.T) {
    value:=1
    
    if value != 1 {
        t.Fatalf("expected 100, got: %d", value)
    }
}
