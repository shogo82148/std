package main

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGodoc(t *testing.T) {
	data, err := godoc("testdata/mypackage/mypackage.go")
	if err != nil {
		t.Fatal(err)
	}
	got := string(data)

	data, err = os.ReadFile("testdata/want/mypackage.go")
	if err != nil {
		t.Fatal(err)
	}
	want := string(data)

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("(-want +got):\n%s", diff)
	}
}
