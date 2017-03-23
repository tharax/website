package websites

import (
	"net/http"
)

func PersonalServer() http.Handler {
	return http.FileServer(http.Dir("./website/peterrosser"))
}