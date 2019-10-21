/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package authorization

import (
	"errors"
	"fmt"
	"github.com/alexedwards/scs"
	"github.com/casbin/casbin"
	"github.com/kazekim/golang-test/testcasbin/http/model"
	"log"
	"net/http"
)

// Authorizer is a middleware for authorization
func Authorizer(e *casbin.Enforcer, sessionManager *scs.SessionManager, users model.Users) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			role := sessionManager.GetString(ctx, "role")
			if role == "" {
				role = "anonymous"
			}
			// if it's a member, check if the user still exists
			if role == "member" {
				uid := sessionManager.GetInt(ctx, "userID")
				if uid == 0 {
					writeError(http.StatusInternalServerError, "ERROR", w, fmt.Errorf("userID is not correct"))
					return
				}
				exists := users.Exists(uid)
				if !exists {
					writeError(http.StatusForbidden, "FORBIDDEN", w, errors.New("user does not exist"))
					return
				}
			}
			// casbin enforce
			res, err := e.Enforce(role, r.URL.Path, r.Method)
			if err != nil {
				writeError(http.StatusInternalServerError, "ERROR", w, err)
				return
			}
			if res {
				next.ServeHTTP(w, r)
			} else {
				writeError(http.StatusForbidden, "FORBIDDEN", w, errors.New("unauthorized, or already logged in"))
				return
			}
		}

		return http.HandlerFunc(fn)
	}
}

func writeError(status int, message string, w http.ResponseWriter, err error) {
	log.Print("ERROR: ", err.Error())
	w.WriteHeader(status)
	w.Write([]byte(message))
}
