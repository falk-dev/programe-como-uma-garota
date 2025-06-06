package tokens

import (
	"testing"

	"github.com/cristalhq/jwt/v5"
)

func TestGenerate(t *testing.T) {
	token := Generate()

	// Uma string, em Go, nunca vai ser nula, mas sim vazia
	if token == "" {
		t.Error("Token vazia")
	}

	if len(token) < 100 {
		t.Error("Token muito curto")
	}

	if len(token) > 1000 {
		t.Error("Token muito grande")
	}
}

func TestVerifyToken(t *testing.T) {
	token := Generate()

	key := []byte(`secret`)

	verifier, err := jwt.NewVerifierHS(jwt.HS256, key)

	if err != nil {
		t.Error("Falha ao criar verifier")
	}

	newToken, err := jwt.Parse([]byte(token), verifier)

	if err != nil {
		t.Error("Falha ao fazer o parse do token")
	}

	if newToken == nil {
		t.Error("Token não pode ser nulo.")
	}
}
