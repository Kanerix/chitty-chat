package session

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

func StringToSession(s string) (*Session, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 3 {
		return nil, errors.New("parts of session string not equal to 3")
	}

	username := parts[0]
	anonymous := parts[1] == "true"
	fmt.Println(parts[2])
	token, err := uuid.Parse(parts[2])
	if err != nil {
		return nil, errors.New("error parsing token into UUID")
	}

	return &Session{username, anonymous, token}, nil
}

func GetSessionFileContent() ([]byte, error) {
	if _, err := os.Stat(SessionFile); err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(SessionFile)
			if err != nil {
				return nil, err
			}
		}
		return nil, err
	}

	return os.ReadFile(SessionFile)
}

func SaveSessionToken(token string) error {
	return os.WriteFile(SessionFile, []byte(token), 0644)
}
