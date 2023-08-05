package main

import "testing"

func assertNil(t *testing.T, got any) {
	if got != nil {
		t.Errorf("got is %+v but should be nil", got)
	}
}
