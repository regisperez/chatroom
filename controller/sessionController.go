package controller

import (
	"chatroom/model"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Session struct {
	Username string
	Expiry   time.Time
}
var (
	sessions = map[string]Session{}
	SessionTest bool
)

func isInvalidSession(w http.ResponseWriter, r *http.Request) bool {
	if SessionTest{
		return false
	}
	code, message := CheckSession(w, r)
	if code != 0 {
		respondWithError(w, code, message)
		return true
	}
	return false
}

func CheckSession(w http.ResponseWriter, r *http.Request) (int,string){
	// We can obtain the session token from the requests cookies, which come with every request
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			return http.StatusUnauthorized, "not authorized"
		}
		// For any other type of error, return a bad request status
		return http.StatusBadRequest, "bad request"
	}
	sessionToken := c.Value

	// We then get the session from our session map
	userSession, exists := sessions[sessionToken]
	if !exists {
		// If the session token is not present in session map, return an unauthorized error
		return http.StatusUnauthorized, "not authorized"
	}

	// If the session is present, but has expired, we can delete the session, and return
	// an unauthorized status
	if userSession.isExpired() {
		delete(sessions, sessionToken)
		return http.StatusUnauthorized, "session expired"
	}

	return 0,""
}

func (s Session) isExpired() bool {
	return s.Expiry.Before(time.Now())
}

func initSession(w http.ResponseWriter, user model.User) {
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(60 * time.Minute)

	sessions[sessionToken] = Session{
		Username: user.Name,
		Expiry:   expiresAt,
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
		Path: "/",
		SameSite: 4,
	})

	http.SetCookie(w, &http.Cookie{
		Name:    "user_name",
		Value:   user.Name,
		Expires: expiresAt,
		Path: "/",
		SameSite: 4,
	})
}

func closeSession(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	// remove the users session from the session map
	delete(sessions, sessionToken)

	// We need to let the client know that the cookie is expired
	// In the response, we set the session token to an empty
	// value and set its expiry as the current time
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	})

	http.SetCookie(w, &http.Cookie{
		Name:    "user_name",
		Value:   "",
		Expires: time.Now(),
	})
}


