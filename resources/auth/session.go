package auth

import (
	"crypto/rand"
	"os"

	"github.com/gorilla/sessions"
)

var sessionStore *sessions.FilesystemStore

func init() {
	key := make([]byte, 64)
	rand.Read(key)
	sessionStore = sessions.NewFilesystemStore(os.TempDir(), key)
}

// Get returns a postgres instance
func Get() (*sessions.FilesystemStore, error) {
	return sessionStore, nil
}

// Set sets the store (mostly for testing)
func Set(s *sessions.FilesystemStore) {
	sessionStore = s
}
