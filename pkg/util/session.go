package util

import (
	"os"
	"path"
)

type contextKey string

const SessionContextKey = contextKey("session")

var sessionFile = path.Join(ChippyPath, "session.txt")

func GetSessionFileContent() ([]byte, error) {
	if _, err := os.Stat(sessionFile); err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(sessionFile)
			if err != nil {
				return nil, err
			}
		}
		return nil, err
	}

	return os.ReadFile(sessionFile)
}

func SaveSessionToken(token string) error {
	return os.WriteFile(sessionFile, []byte(token), 0644)
}
