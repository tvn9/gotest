package url

import (
	"testing"
)

func TestParse(t *testing.T) {
	var url = "https://twitter.com/anything"
	u, err := Parse(url)
	if err != nil {
		t.Fatalf("Parse(%q) err = %q, want nil", url, err)
	}

	if got, want := u.Scheme, "https"; got != want {
		t.Errorf("Parse(%q), want: %q, got u.Scheme = %q", url, want, got)
	}

	if got, want := u.Host, "twitter.com"; got != want {
		t.Errorf("Parse(%q), Host = %q; want %q, got: %q", url, u.Host, got, want)
	}

	if got, want := u.Path, "/anything"; got != want {
		t.Errorf("Parse(%q), Path: %q, want: %q, got %q", url, u.Path, want, got)

	}

}
