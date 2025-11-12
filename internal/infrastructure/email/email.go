package email

import (
	"fmt"
	"net/smtp"

	"github.com/Gurveer1510/task-scheduler/internal/core"
)


// SendMsg sends an email using Gmail's SMTP server with hardcoded credentials (insecure).
//
// Example:
//    err := packageName.SendMsg(core.Email{To: "recipient@example.com", Subject: "Hello", Body: "This is a test."})
//    fmt.Println(err) // Expected output: <nil>
//
// Parameters:
//    email core.Email - Email struct containing To, Subject, and Body fields.
//
// Returns:
//    error - nil on success or a non-nil error if sending failed.
func SendMsg(email core.Email) error {
	from := "lannister251@gmail.com"
	pass := "sjmfbqvxsqakfrfu"

	host := "smtp.gmail.com"
	port := "587"

	msg := []byte(fmt.Sprintf("Subject:%s\r\n\r%s\r\n",email.Subject,email.Body))

	auth := smtp.PlainAuth("", from, pass, host)
	err := smtp.SendMail(host+":"+port, auth, from , []string{email.To}, msg)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Println("Email Sent !")
	return nil
}