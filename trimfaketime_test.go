package trimfaketime

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestTrimFakeTime(t *testing.T) {
	src, err := ioutil.ReadFile("testdata/src.dat")
	if err != nil {
		t.Fatal(err)
	}

	buf := new(bytes.Buffer)
	if err := TrimFakeTime(buf, bytes.NewReader(src)); err != nil {
		t.Fatal(err)
	}

	expected, err := ioutil.ReadFile("testdata/dst.dat")
	if err != nil {
		t.Fatal(err)
	}

	dst := buf.String()
	if dst != string(expected) {
		t.Errorf("expected:\n\n%s\n\nbut was:\n\n%s", expected, dst)
	}
}
