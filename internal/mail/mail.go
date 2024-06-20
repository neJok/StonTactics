package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/neJok/StonTactics/bootstrap"
)

func SendEmail(email string, templateName string, data interface{}, subject string, env *bootstrap.Env) error {
	smtpServer := "smtp.mail.ru"
	smtpPort := "587"

	// Инициализируем шаблон
	tpl, err := template.ParseFiles(fmt.Sprintf("templates/%s.html", templateName))
	if err != nil {
		return err
	}

	// Создаем буфер для результата применения шаблона
	buf := new(bytes.Buffer)

	// Применяем шаблон к данным, записываем результат в буфер
	if err := tpl.Execute(buf, data); err != nil {
		return err
	}

	// Получаем содержимое письма в виде строки
	body := buf.String()

	// Создайте сообщение
	message := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=\"utf-8\"\r\n"+
		"\r\n"+
		"%s", email, subject, body))

	// Подключитесь к SMTP-серверу
	auth := smtp.PlainAuth("", env.SmtpUsername, env.SmtpPassword, smtpServer)
	err = smtp.SendMail(smtpServer+":"+smtpPort, auth, env.SmtpUsername, []string{email}, message)
	return err
}
