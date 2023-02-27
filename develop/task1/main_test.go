package main

import (
	"testing"
	"time"
)

func TestTask1(t *testing.T) {
	a := TimeNow()
	b := time.Now()
	if a != b {
		t.Errorf("got %q want %q", a, b)
	}
}
