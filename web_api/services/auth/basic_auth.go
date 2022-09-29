package auth

import "fmt"

type Basic struct {
	Username string
	Password string
}

func (b *Basic) Exec() error {
	if b.Username != "usename" && b.Password != "password" {
		return fmt.Errorf("invalid credentials")
	}

	return nil
}
