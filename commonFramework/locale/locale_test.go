package locale

import (
	"testing"
)

func TestLocaleGet(t *testing.T) {
	t.Parallel()

	var brand = "cb"
	var expectedLocale = "cb-en-us"

	result := Get(brand, "")

	if result != expectedLocale {
		t.Errorf("Expected %s, but got %s", expectedLocale, result)
	}

	brand = "cb2"
	expectedLocale = "c2-en-us"

	result = Get(brand, "")

	if result != expectedLocale {
		t.Errorf("Expected %s, but got %s", expectedLocale, result)
	}

	brand = "lon"
	expectedLocale = "nd-en-us"

	result = Get(brand, "")

	if result != expectedLocale {
		t.Errorf("Expected %s, but got %s", expectedLocale, result)
	}

	brand = "CANCB"
	expectedLocale = "cn-en-ca"

	result = Get(brand, "")

	if result != expectedLocale {
		t.Errorf("Expected %s, but got %s", expectedLocale, result)
	}

	brand = "CANCB2"
	expectedLocale = "c2-en-ca"

	result = Get(brand, "")

	if result != expectedLocale {
		t.Errorf("Expected %s, but got %s", expectedLocale, result)
	}

	brand = "CANCB"
	expectedLocale = "cn-fr-ca"

	result = Get(brand, "fr_CA")

	if result != expectedLocale {
		t.Errorf("Expected %s, but got %s", expectedLocale, result)
	}

	brand = "CANCB2"
	expectedLocale = "c2-fr-ca"

	result = Get(brand, "fr_CA")

	if result != expectedLocale {
		t.Errorf("Expected %s, but got %s", expectedLocale, result)
	}

}

func TestMultiLang(t *testing.T) {

	t.Parallel()

	var locale = "cn-fr-ca"
	var expectedLocale = "cn-en-ca"

	result, bLocale := IsMultiLang(locale)

	if !result || bLocale != expectedLocale {
		t.Errorf("Expected %s, but got %s", expectedLocale, bLocale)
	}

	locale = "cn-en-ca"
	expectedLocale = ""

	result, bLocale = IsMultiLang(locale)

	if result || bLocale != expectedLocale {
		t.Errorf("Expected %s, but got %s", expectedLocale, bLocale)
	}

	locale = "cb-en-us"
	expectedLocale = ""

	result, bLocale = IsMultiLang(locale)

	if result || bLocale != expectedLocale {
		t.Errorf("Expected %s, but got %s", expectedLocale, bLocale)
	}
}
