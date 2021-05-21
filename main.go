package main

import (
	"net/http"

	"github.com/OCain/go-user-webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
