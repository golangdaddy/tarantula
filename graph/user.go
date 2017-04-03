package graph

import (
	"fmt"
	"time"
)

type User struct {
	DB           *Database				`json:"-"`
	Uid          int64
	Name         string
	PasswordHash string
	Vertex       *Vertex
}

func (user *User) NewSession() (bool, string) {

	secret, hashedSecret := sessionKey()

	sesh := &Session{
		DB:        user.DB,
		User:      user.Uid,
		TokenHash: hashedSecret,
		Timestamp: time.Now().UTC().Add(9 * time.Hour).UnixNano(),
	}

	q := fmt.Sprintf("INSERT INTO %v (uid, user, tokenHash) VALUES (null, ?, ?);", user.DB.Table(TABLE_SESSIONS))

	if ok, _ := user.DB.Client.Exec(q, sesh.User, sesh.TokenHash); !ok {
		return false, ""
	}

	return true, secret
}
