package go_vcard

import (
	"bytes"
	"reflect"
	"testing"
)

func TestEncoder(t *testing.T) {
	var b bytes.Buffer
	if err := NewEncoder(&b).Encode(testCard); err != nil {
		t.Fatal("Expected no error when formatting card, got:", err)
	}

	expected := "BEGIN:VCARD\r\nVERSION:4.0\r\nCLIENTPIDMAP:1;urn:uuid:53e374d9-337e-4727-8803-a1e9c14e0556\r\nEMAIL;PID=1.1:jdoe@example.com\r\nFN;PID=1.1:J. Doe\r\nN:Doe;J.;;;\r\nUID:urn:uuid:4fbe8971-0bc3-424c-9c26-36c3e1eff6b1\r\nEND:VCARD\r\n"
	if b.String() != expected {
		t.Errorf("Excpected vcard to be %q,\n but got %q", expected, b.String())
	}

	card, err := NewDecoder(&b).Decode()
	if err != nil {
		t.Fatal("Expected no error when parsing formatted card, got:", err)
	}

	if !reflect.DeepEqual(card, testCard) {
		t.Errorf("Invalid parsed card: expected \n%+v\n but got \n%+v", testCard, card)
	}
}



var testValue = []struct {
	v         string
	formatted string
}{
	{"Hello World!", "Hello World!"},
	{"this is a single value, with a comma encoded", "this is a single value\\, with a comma encoded"},
	{"Mythical Manager\nHyjinx Software Division", "Mythical Manager\\nHyjinx Software Division"},
	{"aa\\\nbb", "aa\\\\\\nbb"},
}


