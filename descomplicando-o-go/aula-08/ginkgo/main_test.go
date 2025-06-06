package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSoma(t *testing.T) {
	// A função abaixo é o que interliga o Gomega e o Ginkgo. Então, o que ela faz é:
	// quando ocorrer uma falha, reporta ao Ginkgo através da função "Fail",
	// que sabe como parar a execução do teste atual, marcar como falho e seguir
	// para o próximo teste, sem interromper todo o processo.
	RegisterFailHandler(Fail)

	// Principal função do Ginkgo. O que ela faz é vasculhar o código procurando os
	// blocos de "Describe" e "It", executando 1 por 1.
	RunSpecs(t, "Soma Suite")
}

var _ = Describe("Soma", func() {
	Context("Quando realiza a soma", func() {
		resultado := Soma(2, 2)
		It("Expect result equal 4", func() {
			Expect(resultado).To(Equal(4))
		})
	})
})
