package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/NucleonGodX/Distributed-Task-Queue-System/api"
	"github.com/NucleonGodX/Distributed-Task-Queue-System/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError= errors.New("Invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func ( w http.ResponseWriter, r *http.Request){
		var username string= r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error
		
		if username== "" || token== "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		} 

		var database *tools.DatabaseInterface
		database, err= tools.NewDatabse()
	})
}