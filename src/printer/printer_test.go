package printer

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/sam8helloworld/json-go/value"
)

func TestSuccess(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	stdout := os.Stdout
	os.Stdout = w
	input := value.Object{
		"number": value.NumberInt(123),
	}
	sut := NewPrinter(input)
	sut.Execute()

	os.Stdout = stdout
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)
	got := buf.String()
	want := "{\"number\":123}"

	if got != want {
		t.Errorf("want %s, but got %s", want, got)
	}
}
