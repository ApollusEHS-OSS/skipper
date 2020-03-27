package auth

import (
	"testing"

	"github.com/zalando/skipper/secrets"
)

func TestFlowState(t *testing.T) {
	secrets := secrets.NewRegistry()
	defer secrets.Close()

	fs := newFlowState(secrets, "testdata/secret-key")
	const u = "https://www.example.org/foo"
	s, err := fs.createState(u)
	if err != nil {
		t.Fatal(err)
	}

	st, err := fs.extractState(s)
	if err != nil {
		t.Fatal(err)
	}

	if st.RedirectURL != u {
		t.Errorf("invalid redirect url: '%s', expected: '%s'", st.RedirectURL, u)
	}
}
