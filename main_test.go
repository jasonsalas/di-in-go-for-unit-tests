package main

import (
	"testing"

	"github.com/jasonsalas/isolate-astros/utils"
)

type testWebRequest struct{}

// FetchBytes satisfies the astros.GetWebRequest interface
func (testWebRequest) FetchBytes(url string) ([]byte, error) {
	return []byte(`{"number":3}`), nil
}

func TestGetAstronauts(t *testing.T) {
	want := 3
	got, err := utils.GetAstronauts(testWebRequest{})
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Errorf("people in space || WANT: %d, GOT: %d", want, got)
	}
}
