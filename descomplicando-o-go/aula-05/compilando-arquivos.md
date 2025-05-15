# Compilando arquivos em Go

Até o momento, o comando de compilar o arquivo que foi utilizado:
```cmd
go run main.go
```
Esse comando compila o código e depois ele roda de forma "transparente" pois roda o binário gerado sem que seja necessário um comando específico para isso. 



Para fazer esse processo de compilar e rodar **um** arquivo separadamente, faça:
```cmd
go build -o nome_que_eu_quiser main.go 
```
Caso deseje rodar todo o projeto, o comando abaixo irá procurar pelo arquivo `go.mod`:

```cmd
go build
```

Daí, para executar o arquivo após gerar o arquivo, basta:
```cmd
./nome_que_eu_quiser
```

Para verificar a arquitetura da máquina, basta:
```cmd
go env GOOS GOARCH
```

Para gerar um binário específico para determinado OS e determinada arquitetura:
```cmd
GOOS=windows GOARCH=arm64 go build main.go
```