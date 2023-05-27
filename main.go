package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sanyogpatel-tecblic/ecom-without-ORM/controller/config"
	"github.com/sanyogpatel-tecblic/ecom-without-ORM/controller/routes"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {
	fmt.Println("Server is getting started...")
	fmt.Println("Listening at port 8080 ...")

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	routes.Routes()
}
