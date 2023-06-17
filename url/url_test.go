package url

import "testing"

func TestParse(t *testing.T) {
	// Better error message
	const rawurl = "broken url"

	// Fail test
	if err := Parse(rawurl); err != nil {
		t.Logf("Parse(%q) err = %q, want nil", rawurl, err)
		t.Fail()
	}
}
