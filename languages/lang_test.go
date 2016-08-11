package languages

import "testing"

func TestGetLangObject(t *testing.T) {
	tests := map[string]bool{
		"english": true,
		"en":      true,
		"german":  false,
	}

	for lang, exists := range tests {

		_, err := GetLangObject(lang)

		if err != nil && exists {
			t.Fatalf("For existent language %s, got error %s.", lang, err.Error())
		} else if err == nil && !exists {
			t.Fatalf("For not existent language %s, got it exists.")
		}

	}
}
