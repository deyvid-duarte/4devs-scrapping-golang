package person

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type Person struct {
	Gender string
	Age uint8
}

// method only to read attributes from struct
func (p Person) Generate(withMask bool, amount uint8) string {
	var withMaskString string
	if withMask {
		withMaskString = "S"
	} else {
		withMaskString = "N"
	}
	
	data := url.Values{}
	data.Set("acao", "gerar_pessoa")
	data.Set("sexo", p.Gender)
	data.Set("pontuacao", withMaskString)
	data.Set("idade", strconv.Itoa(int(p.Age)))
	data.Set("cep_estado", "")
	data.Set("txt_qtde", strconv.Itoa(int(amount)))
	data.Set("cep_cidade", "")

	resp, err := http.PostForm("https://www.4devs.com.br/ferramentas_online.php", data)
	if err != nil {
		panic("Erro ao gerar a pessoa: " + err.Error())
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return string(body)
}