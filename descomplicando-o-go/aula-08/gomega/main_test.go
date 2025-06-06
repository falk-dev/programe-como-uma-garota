package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestSoma(t *testing.T) {
	// Ele precisa do "t" da biblioteca "testing" para funcionar
	g := NewGomegaWithT(t)

	resultado := Soma(4, 5)

	g.Expect(resultado).To(Equal(9))
}
