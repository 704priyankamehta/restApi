package services

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kataras/go-mailer"
)

//Config ...
type Config struct {
	Host string

	Port int

	Username string

	Password string

	FromAddr string
	// FromAlias is the from part, if empty this is the first part before @ from the Username field.
	FromAlias string
	// UseCommand enable it if you want to send e-mail with the mail command  instead of smtp.
	//
	// Host,Port & Password will be ignored.
	// ONLY FOR UNIX.
	UseCommand bool
}

func Mail(c *fiber.Ctx) error {

	config := mailer.Config{
		Host:     "smtp..com",
		Username: "user",
		Password: "pass",
		FromAddr: "pp",
		Port:     587,

		UseCommand: false,
	}

	sender := mailer.New(config)

	subject := "Hello subject"

	content := `<h1>Hello</h1> <br/><br/> <span style="color:red"> This is the rich message body </span>`

	to := []string{"to"}

	err := sender.Send(subject, content, to...)

	if err != nil {
		println("error while sending the e-mail: " + err.Error())
	}
	return c.SendString("mail send")
}

//Hello ...
func Hello(s string) string {
	if len(s) == 0 {
		return "bye"
	}
	return "hello"

}
