package main

import (
	"fmt"
	"io/ioutil"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func main() {

	html, err := ioutil.ReadFile("./src/index.html")
	htmlContexto := fmt.Sprintf(string(html), "Leo desde un fmt.string")
	e := email.NewEmail()
	e.From = "Leonardo Antonio Nolasco Leyva <leo2001.nl08@gmail.com>"
	e.To = []string{"mcc2001.nl8@gmail.com"}
	e.Subject = "Se registro correctamente en nuestra pagina web!!"
	e.HTML = []byte(htmlContexto)
	err = e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "leo2001.nl08@gmail.com", "gxyt rxyx vimx yshl", "smtp.gmail.com"))

	if err != nil {
		fmt.Println("Informamos el error", err)
	} else {
		fmt.Println("enviado")
	}
}
