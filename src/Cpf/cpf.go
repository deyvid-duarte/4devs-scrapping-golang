package CPF

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
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