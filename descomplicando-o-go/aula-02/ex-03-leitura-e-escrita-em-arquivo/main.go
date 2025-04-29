package main

import (
	"fmt"
	"os"
)

func main() {
	txt := "MEU_REMEDIO_NOVO;61.190.096/0001-92;EUROFARMA LABORATORIOS S.A.;508011804138416;1004306960107;7891317469610;7891317020118;    -     ;SIMECO PLUS;120 MG/ML + 60 MG/ML + 7 MG/ML SUS OR CT FR VD AMB X 60 ML;A2A4 - ANTIÁCIDOS COM ANTIFLATULENTOS OU CARMINATIVOS;Específico;Liberado;8,13;9,11;10,53;9,24;11,26;9,80;11,34;9,86;11,42;9,92;11,58;10,04;11,66;10,10;11,75;10,17;11,83;11,83;11,92;10,30;12,10;10,43;12,19;10,50;12,28;10,56;11,24;12,22;14,07;12,77;15,01;13,55;15,11;13,63;15,21;13,71;15,42;13,88;15,52;13,96;15,64;14,06;15,74;16,35;15,86;14,24;16,09;14,42;16,20;14,52;16,32;14,60;Não;Não;Não;Não;;Negativa;Não;Tarja Sem Tarja;"

	// Estrutura: nome do arquivo, flags e permissão
	file, err := os.OpenFile("TA_PRECO_MEDICAMENTO.csv", os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// Escreve ao fim do arquivo
	// Usa WriteString pois se trata de uma string, não um slice de bytes
	file.WriteString(txt)
}
