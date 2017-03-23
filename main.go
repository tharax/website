package websites

import (
	"net/http"
)

func PersonalHandler() http.Handler {
	return http.FileServer(http.Dir("./peterrosser"))
}

func TrustHandler() http.Handler {
	return http.FileServer(http.Dir("./thefirsttrust"))
}

func BusinessHandler() http.Handler {
	return http.FileServer(http.Dir("./rossersoftware"))
}
