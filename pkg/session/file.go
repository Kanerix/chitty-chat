package session

import (
	"errors"
	"os"
	"path"
	"strings"

	"github.com/google/uuid"
)

func FromString(s string) (*Session, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 3 {
		return nil, errors.New("parts of session string not equal to 3")
	}

	username := parts[0]
	anonymous := parts[1] == "true"
	token, err := uuid.Parse(parts[2])
	if err != nil {
		return nil, errors.New("error parsing token into UUID")
	}

	return &Session{username, anonymous, token}, nil
}

func GetSessionFileContent() ([]byte, error) {
	err := os.MkdirAll(path.Dir(SessionFile), os.ModePerm)
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(SessionFile, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return os.ReadFile(SessionFile)
}

func SaveSessionToken(token string) error {
	return os.WriteFile(SessionFile, []byte(token), 0644)
}
