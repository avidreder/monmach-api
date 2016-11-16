package auth

import (
	"os"

	"github.com/gorilla/sessions"
)

var sessionStore *sessions.FilesystemStore

func init() {
	sessionStore = sessions.NewFilesystemStore(os.TempDir(), []byte("monmach-sessions"))
}

// Get returns a postgres instance
func Get() (*sessions.FilesystemStore, error) {
	return sessionStore, nil
}

// Set sets the store (mostly for testing)
func Set(s *sessions.FilesystemStore) {
	sessionStore = s
}
