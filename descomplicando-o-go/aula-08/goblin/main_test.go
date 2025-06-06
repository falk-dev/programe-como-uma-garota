package main

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestSoma(t *testing.T) {
	g := Goblin(t)

	g.Describe("Soma", func() {
		g.Describe("Quando realiza a soma", func() {
			resultado := Soma(2, 2)
			g.It("Expect result equal 4", func() {
				g.Assert(resultado).Equal(4)
			})
		})
		g.Describe("Quando verifica o valor", func() {
			resultado := Soma(2, 2)
			g.It("Expect result less than 10", func() {
				g.Assert(resultado < 10).IsTrue()
			})
		})
	})
}
