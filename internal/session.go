package internal

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

// Session struct wraps around the Gorilla sessions cookie store.
type Session struct {
	Store *sessions.CookieStore
}

// NewSession initializes a new session store with the given secret key.
func NewSession(secretKey []byte) *Session {
	store := sessions.NewCookieStore(secretKey)
	return &Session{Store: store}
}

// GetUserID retrieves the user ID from the session.
func (s *Session) GetUserID(r *http.Request) (string, error) {
	session, err := s.Store.Get(r, "session-name") // Use your preferred session name here
	if err != nil {
		return "", fmt.Errorf("could not get session: %v", err)
	}

	userID, ok := session.Values["user_id"].(string)
	if !ok {
		return "", fmt.Errorf("user ID not found in session")
	}

	return userID, nil
}

// SetUserID sets the user ID in the session.
func (s *Session) SetUserID(w http.ResponseWriter, r *http.Request, userID string) error {
	session, err := s.Store.Get(r, "session-name") // Use your preferred session name here
	if err != nil {
		return fmt.Errorf("could not get session: %v", err)
	}

	// Set user ID in session
	session.Values["user_id"] = userID

	// Save the session
	err = session.Save(r, w)
	if err != nil {
		return fmt.Errorf("could not save session: %v", err)
	}

	return nil
}
