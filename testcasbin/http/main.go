/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
"fmt"
"github.com/casbin/casbin"
	"github.com/kazekim/golang-test/testcasbin/http/authorization"
	"github.com/kazekim/golang-test/testcasbin/http/model"
	"log"
"net/http"
"time"
"github.com/alexedwards/scs"
)

var sessionManager *scs.SessionManager

func main() {
	// setup casbin auth rules
	authEnforcer, err := casbin.NewEnforcer("./auth_model.conf", "./policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	// setup session store
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	users := createUsers()

	// setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginHandler(users))
	mux.HandleFunc("/logout", logoutHandler())
	mux.HandleFunc("/member/current", currentMemberHandler())
	mux.HandleFunc("/member/role", memberRoleHandler())
	mux.HandleFunc("/admin/stuff", adminHandler())

	log.Print("Server started on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", sessionManager.LoadAndSave(authorization.Authorizer(authEnforcer, sessionManager, users)(mux))))

}

func createUsers() model.Users {
	users := model.Users{}

	users = append(users, model.User{ID: 1, Name: "Admin", Role: "admin"})
	users = append(users, model.User{ID: 2, Name: "Sabine", Role: "member"})
	users = append(users, model.User{ID: 3, Name: "Sepp", Role: "member"})
	return users
}

func loginHandler(users model.Users) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		name := r.PostFormValue("name")
		user, err := users.FindByName(name)
		if err != nil {
			writeError(http.StatusBadRequest, "WRONG_CREDENTIALS", w, err)
			return
		}

		// setup session
		if err := sessionManager.RenewToken(ctx); err != nil {
			writeError(http.StatusInternalServerError, "ERROR", w, err)
			return
		}

		sessionManager.Put(ctx, "userID", user.ID)
		sessionManager.Put(ctx, "role", user.Role)
		writeSuccess("SUCCESS", w)
	})
}

func logoutHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if err := sessionManager.Clear(ctx); err != nil {
			writeError(http.StatusInternalServerError, "ERROR", w, err)
			return
		}
		writeSuccess("SUCCESS", w)
	})
}

func currentMemberHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		uid := sessionManager.GetInt(ctx, "userID")
		if uid == 0 {
			writeError(http.StatusInternalServerError, "ERROR", w, fmt.Errorf("uid is not correct"))
			return
		}
		writeSuccess(fmt.Sprintf("User with ID: %d", uid), w)
	})
}

func memberRoleHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		role := sessionManager.GetString(ctx, "role")
		if role == "" {
			writeError(http.StatusInternalServerError, "ERROR", w, fmt.Errorf("role is not correct"))
			return
		}
		writeSuccess(fmt.Sprintf("User with Role: %s", role), w)
	})
}

func adminHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeSuccess("I'm an Admin!", w)
	})
}

func writeError(status int, message string, w http.ResponseWriter, err error) {
	log.Print("ERROR: ", err.Error())
	w.WriteHeader(status)
	w.Write([]byte(message))
}

func writeSuccess(message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
