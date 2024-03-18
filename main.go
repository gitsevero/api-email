package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
)
type Config struct {
	Email    string `json:"email"`
	Password string `json:"password"`

}
func loadConfig(filename string) (Config, error) {
	var config Config
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, &config)
	return config, err
}

var authToken = "token_de_autenticacao_secreto"

func sendEmail(emissor, assunto, destinatario, mensagem, usuario string, config Config) {
	auth := smtp.PlainAuth("", config.Email, config.Password, "smtp.gmail.com")
	msg := []byte("Subject: " + assunto + "\r\n" +
		"From: " + emissor + " - " + usuario + " <" + config.Email + ">\r\n" +
		"To: " + destinatario + "\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n\r\n" +
		mensagem)
	serverAddr := "smtp.gmail.com:587"
	if err := smtp.SendMail(serverAddr, auth, config.Email, []string{destinatario}, msg); err != nil {
		fmt.Println("Erro ao enviar email:", err)
		return
	}
	fmt.Println("Email para:", destinatario)
}

func validToken(token string) bool {
	return token == authToken
}
func sendEmailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	token := r.Header.Get("Authorization")
	if !validToken(token) {
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var data map[string]string
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, "Erro ao decodificar dados", http.StatusInternalServerError)
		return
	}
	config, err := loadConfig("config.json")
	if err != nil {
		http.Error(w, "Erro ao carregar configurações", http.StatusInternalServerError)
		return
	}
	sendEmail(data["emissor"], data["assunto"], data["destinatario"], data["mensagem"], data["usuario"], config)
	w.WriteHeader(http.StatusOK)
}
func main() {
	fmt.Println("API rodando...")
	http.HandleFunc("/send-email", sendEmailHandler)
	http.ListenAndServe(":32781", nil)
}
