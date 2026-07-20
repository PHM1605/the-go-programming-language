package storage2

import (
	"fmt"
	"log"
	"net/smtp"
)

func bytesInUse(username string) int64 {
	return 980000000 // just a dummy value to simulate 98% usage of storage
}

const sender = "notifications@example.com"
const password = "correcthorsebatterystaple"
const hostname = "smtp.example.com"

const template = `Warning: you are using %d bytes of storage, %d%% of your quota.`

var notifyUser = func(username, msg string) {
	// setup authentication identity before sending email
	// sender: who is sending the email (e.g. phm1605@gmail.com)
	// password: password of email sender
	// hostname: Google SMTP Server if using @gmail
	auth := smtp.PlainAuth("", sender, password, hostname)
	// send email when Storage > 90%
	// username: Client who is using our Storage service
	err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendEmail(%s) failed: %s", username, err)
	}
}

func CheckQuota(username string) {
	used := bytesInUse(username)
	const quota = 1000000000 // 1GB
	percent := 100 * used / quota
	if percent < 90 {
		return
	} // OK
	msg := fmt.Sprintf(template, used, percent)
	notifyUser(username, msg)
}
