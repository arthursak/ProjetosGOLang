package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	//"reflect" para verificar o tipo da variavel
)

const monitoramentos = 3
const delay = 5

//não precisa declarar var, podendo usar := no lugar

func main() {
	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs")
		case 0:
			fmt.Println("Saindo")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	nome := "Arthur"
	versao := 1.1

	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("")
}

func exibeMenu() {
	fmt.Println("-------------------------")
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
	fmt.Println("-------------------------")

}

func leComando() int {
	var comandoLido int

	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Iniciando monitoramento...")
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")

	}
	fmt.Println("") //Espaçamento

}

func testaSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, " - Online")
	} else {
		fmt.Println("Site: ", site, " com erro. Status Code: ", resp.StatusCode)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("teste.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)

	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha) //tira espaços
		fmt.Println(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}
