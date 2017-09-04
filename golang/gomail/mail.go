package main

import gomail "gopkg.in/gomail.v2"

func main() {

	m := gomail.NewMessage()
	m.SetHeader("From", "GanEasy@qq.com")
	m.SetHeader("To", "2325379239@qq.com")
	// m.SetAddressHeader("Cc", "zenghuitrue@gmail.com", "易增辉")
	m.SetAddressHeader("Cc", "245561237@qq.com", "易增辉")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.qq.com", 465, "GanEasy@qq.com", "yizeme2you")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	// m := gomail.NewMessage()
	// m.SetHeader("From", "from@example.com")
	// // m.SetHeader("To", "to@example.com")
	// m.SetHeader("To", "2325379239@qq.com")

	// m.SetHeader("Subject", "Hello!")
	// m.SetBody("text/plain", "Hello!")

	// d := gomail.Dialer{Host: "localhost", Port: 587}
	// if err := d.DialAndSend(m); err != nil {
	// 	panic(err)
	// }

	// m := gomail.NewMessage()
	// m.SetHeader("From", "GanEasy@qq.com")
	// // m.SetHeader("To", "lj221108@163.com")
	// m.SetHeader("To", "2325379239@qq.com")
	// m.SetHeader("Subject", "Hello!")
	// m.SetBody("text/plain", "Hello!")

	// s := gomail.SendFunc(func(from string, to []string, msg io.WriterTo) error {
	// 	// Implements you email-sending function, for example by calling
	// 	// an API, or running postfix, etc.
	// 	fmt.Println("From:", from)
	// 	fmt.Println("To:", to)
	// 	return nil
	// })

	// if err := gomail.Send(s, m); err != nil {
	// 	panic(err)
	// }
}
