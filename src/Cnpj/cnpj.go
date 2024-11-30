package CNPJ

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Generate(withMask bool) string {
	var withMaskString string;
	if withMask {
		withMaskString = "S"
	} else {
		withMaskString = "N"
	}

	formData := url.Values{}
	formData.Set("acao", "gerar_cnpj")
	formData.Set("pontuacao", withMaskString)

	resp, err := http.PostForm("https://www.4devs.com.br/ferramentas_online.php", formData)
	if err != nil {
		panic("Erro ao gerar o CNPJ: " + err.Error())
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

func Validator(cnpj string) string {
	data := url.Values{}
	data.Set("acao", "validar_cnpj")
	data.Set("txt_cnpj", cnpj)

	response, err := http.PostForm("https://www.4devs.com.br/ferramentas_online.php", data)
	if err != nil {
		panic("Erro ao validar o CNPJ: " + err.Error())
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	cnpjStatus := strings.Split(bodyString, " - ")[1]
	return strings.TrimSpace(cnpjStatus)
}