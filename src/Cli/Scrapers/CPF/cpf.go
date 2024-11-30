package cpf

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Generate(withMask bool, ufOrigin string) string {
	var withMaskString string
	if withMask {
		withMaskString = "S"
	} else {
		withMaskString = "N"
	}

	data := url.Values{}
	data.Set("acao", "gerar_cpf")
	data.Set("pontuacao", withMaskString)
	data.Set("cpf_estado", ufOrigin)

	resp, err := http.PostForm("https://www.4devs.com.br/ferramentas_online.php", data)
	if err != nil {
		fmt.Println("Erro ao gerar o CPF: ", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

func Validator(cpf string) string {
	data := url.Values{}
	data.Set("acao", "validar_cpf")
	data.Set("txt_cpf", cpf)

	response, err := http.PostForm("https://www.4devs.com.br/ferramentas_online.php", data)
	if err != nil {
		panic("Erro ao validar o CPF: " + err.Error())
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	cpfStatus := strings.Split(bodyString, " - ")[1]
	return strings.TrimSpace(cpfStatus)
}