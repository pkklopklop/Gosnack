package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func (u User) Notify() {
	fmt.Println("Normal Notify ", u.Name)
}

func (u User) SMSNoti() {
	fmt.Println("SMS Notify ", u.Name)
}

func (u User) SendEmail(cc string) (result string) {
	fmt.Println("Email send to", u.Email, "and CC to", cc, ".")
	result = "Success"
	return result
}

type Admin struct {
	User
	Level string
}

func (a Admin) Notify() {
	fmt.Println("Secret Notify ", a.Name)
}

func (a Admin) SMSNoti() {
	fmt.Println("SMS Notify ", a.Name)
}

func (a Admin) SendEmail(cc string) (result string) {
	fmt.Println("Email send to", a.Email, "and BCC to", cc, "secretly.")
	result = "Success"
	return
}

type Notifier interface {
	Emailer
	Notify()
	SMSNoti()
}

type Emailer interface {
	SendEmail(cc string) (result string)
}

func SendNotify(n Notifier, cc string) {
	n.Notify()
	n.SMSNoti()
	fmt.Println(n.SendEmail(cc))
}

func SendEmailToCC(e Emailer, cc string) {
	fmt.Println(e.SendEmail(cc))
}

func main() {

	admin := Admin{
		User: User{
			Name:  "Test Admin",
			Email: "tAdmin@hotmail.com",
		},
		Level: "Super user",
	}

	var admin2 Admin
	admin2.User.Name = "Admin 4 test"
	admin2.User.Email = "admin4@hotmail.com"
	admin2.Level = "Normal user 2"

	// fmt.Println("Admin: ", admin)
	// fmt.Println("Admin2: ", admin2)
	// fmt.Println("Admin2 level: ", admin2.Level)

	// fmt.Printf("%#v\n", admin)
	// fmt.Printf("%#v\n", admin2)

	//Admin notify.
	// admin.Notify()
	SendNotify(admin, "บอย")
	SendEmailToCC(admin, "ก้อย")

	//User under admin to be notified.
	// admin.User.Notify()
	// User.Notify(admin2.User)
	SendNotify(admin.User, "บอย")
	SendEmailToCC(admin.User, "ก้อย")
}
