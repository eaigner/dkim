package dkim

import (
	"fmt"
	"testing"
)

var utilSampleEML1 [][]byte = [][]byte{
	[]byte("A: X\r\n" + "B : Y\t\r\n" + "\tZ  \r\n" + "\r\n" + " C \r\n" + "D \t E\r\n" + "\r\n" + "\r\n"),
	[]byte("A: X\r\nB : Y\t\r\n\tZ  "),
	[]byte(" C \r\nD \t E\r\n\r\n\r\n"),
}

var utilSampleEML2 [][]byte = [][]byte{
	[]byte("A: X\n" + "B : Y\t\n" + "\tZ  \n" + "\n" + " C \n" + "D \t E\n" + "\n" + "\n"),
	[]byte("A: X\nB : Y\t\n\tZ  "),
	[]byte(" C \nD \t E\n\n\n"),
}

var utilSampleEMLs [][][]byte = [][][]byte{utilSampleEML1, utilSampleEML2}

func TestSplitEML(t *testing.T) {
	for _, utilSampleEMLparts := range utilSampleEMLs {
		utilSampleEML := utilSampleEMLparts[0]
		header, body, err := splitEML(utilSampleEML)
		if err != nil {
			t.Fatal("error not nil", err)
		}
		if len(header) == 0 {
			t.Fatal("wrong header length", len(header))
		}
		if len(body) == 0 {
			t.Fatal("wrong body length", len(body))
		}
		if x := string(header); x != string(utilSampleEMLparts[1]) {
			t.Fatal(fmt.Sprintf("wrong header: %q should be %q", x, string(utilSampleEMLparts[1])))
		}
		if x := string(body); x != string(utilSampleEMLparts[2]) {
			t.Fatal(fmt.Sprintf("wrong body: %q should be %q", x, string(utilSampleEMLparts[2])))
		}
	}
}
