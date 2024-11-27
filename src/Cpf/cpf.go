package CPF

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

func Generate(withMask bool, ufOrigin string) string {
	var withMaskString string
	if withMask {
		withMaskString = "S"
	} else {
		withMaskString = "N"
	}

	var requestBuffer bytes.Buffer
	multipartRequest := multipart.NewWriter(&requestBuffer)
	multipartRequest.WriteField("acao", "gerar_cpf")
	multipartRequest.WriteField("pontuacao", withMaskString)
	multipartRequest.WriteField("cpf_estado", ufOrigin)

	req, _ := http.NewRequest("POST", "https://www.4devs.com.br/ferramentas_online.php", &requestBuffer)
	req.Header.Set("Content-Type", multipartRequest.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao gerar o CPF: ", err)
	}
	defer req.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	return string(body)
}