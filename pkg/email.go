package pkg

import (
	"fmt"
	"net/smtp"

	"github.com/Gurveer1510/task-scheduler/internal/core"
)


func SendMsg(email core.Email) {
	from := "lannister251@gmail.com"
	pass := "sjmfbqvxsqakfrfu"

	host := "smtp.gmail.com"
	port := "587"

	msg := []byte(fmt.Sprintf("Subject:%s\r\n\r%s\r\n",email.Subject,email.Body))

	auth := smtp.PlainAuth("", from, pass, host)
	err := smtp.SendMail(host+":"+port, auth, from , []string{email.To}, msg)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Email Sent !")
}