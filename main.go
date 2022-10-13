package main

import (
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/monk78anthony/apiv6/package/storage/aws"
	"github.com/monk78anthony/apiv6/user"
)

func main() {
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Instantiate HTTPs app
	usr := user.Controller{
		Storage: aws.NewUserStorage(session, time.Second*5),
	}

	// Instantiate HTTPs router
	rtr := http.NewServeMux()
	rtr.HandleFunc("/api/v1/users/create", usr.Create)
	rtr.HandleFunc("/api/v1/users/find", usr.Find)
	rtr.HandleFunc("/api/v1/users/delete", usr.Delete)
	rtr.HandleFunc("/api/v1/users/update", usr.Update)

	// Start HTTP server
	//log.Fatal(http.ListenAndServe(":8080", rtr))
	//Start HTTPs server
	log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", rtr))
}
