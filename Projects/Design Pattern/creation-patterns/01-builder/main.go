package main

import "fmt"

func main() {
	var builder = newNotificationBuilder()

	builder.SetTitle("XYZ Notification")
	builder.SetIcon("alert-icon.png")
	builder.SetSubTitle("Bu bir subtitle")
	builder.SetImage("alert-image.png")
	builder.SetPriority(4)
	builder.SetMessage("Uygulama çöktü. Buna bi el at!")
	builder.SetType("alert")

	notif, err := builder.Build()
	if err != nil {
		fmt.Println("Bu bildirimi oluştururken bir hata meydana geldi: ", err)
	} else {
		fmt.Printf("Notification: %+v", notif)
	}

	// fmt.Println("Title : ", notif.title)
	// fmt.Println("Message : ", notif.message)
}

/*
	> go run *	// deprecated
	> go run .
	> go run .\builder.go .\notification.go .\main.go
*/
