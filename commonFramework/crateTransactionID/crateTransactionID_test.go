package crateTransactionID

import (
	"testing"
)

func TestTransactionID(t *testing.T) {
	t.Parallel()

	ID := GetTransactionID()
	if len(ID) != 19 {
		t.Errorf("transaction ID should be 19 digits: " + ID)
	}
}

func TestTransactionIDV2(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		locale   string
		expected string
	}{
		{"cb-en-us", "1"},
		{"c2-en-us", "2"},
		{"nd-en-us", "3"},
		{"cn-en-ca", "4"},
		{"c2-en-ca", "5"},
		{"cn-fr-ca", "4"},
		{"c2-fr-ca", "5"},
	}

	for _, test := range tests {
		actual := GetTransactionIDV2(test.locale)[:1]
		if actual != test.expected {
			t.Errorf("Expected First digit of GetTransactionIDV2(%s) to be %s, got %s", test.locale, test.expected, actual)
		}
	}
}
