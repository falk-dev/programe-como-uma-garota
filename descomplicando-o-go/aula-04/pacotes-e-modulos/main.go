package main

// Para poder dividir o código em múltiplos arquivos, é necessário criar um módulo em Go.
// Para criar, basta utilizar: go mod init github.com/falk-dev/exemplo
import goodbye "github.com/falk-dev/programe-como-uma-garota/tree/main/descomplicando-o-go/aula-04/pacotes-e-modulos/despedida"

func main() {
	sayHello()
	goodbye.SayGoodbye()
}
