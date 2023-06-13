package healper

import (
	"TheaterCinemaAPISystem/initializers"
	"fmt"
	"os"

	"net/smtp"
)

func init() {
	initializers.LoadEnvVarables()
}



type Email struct {
	To      string
	Message string
}

func (email *Email) SendEmail() error{

		host :=  os.Getenv("SMTPHost")
		intity:=os.Getenv("SMTPINTITY")
		username:=os.Getenv("SMTPUsername")
		password:=os.Getenv("SMTPPassword")
		port:=os.Getenv("SMTPPORT")

// اطلاعات اتصال به سرور SMTP را وارد کنید
auth := smtp.PlainAuth(intity, username, password, host)


// تنظیمات ایمیل را وارد کنید
to := []string{email.To}
msg := []byte(fmt.Sprintf("To: %s\r\n"+
	"Subject: Your Subject\r\n"+
	"\r\n"+
	"%s\r\n", email.To, email.Message))

// ارسال ایمیل
		err := smtp.SendMail(host+":"+port, auth, host, to, msg)
		if err != nil {
			fmt.Println("err : ", err)
			return err
		}
		return nil
}