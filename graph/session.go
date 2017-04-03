package graph

import (
	"encoding/hex"
	"fmt"
)

type Session struct {
	DB        *Database					`json:"-"`
	Uid       int64
	User      int64
	TokenHash string
	Timestamp int64
}

func (db *Database) FindSession(token string) (bool, *Session) {

	_, err := hex.DecodeString(token)

	if db.Log.Error(err) {
		return false, nil
	}

	q := fmt.Sprintf("SELECT * FROM %v WHERE tokenHash = ?;", db.internalClasses.sessions.Table())

	ok, session := db.Client.QuerySession(q, sessionKeyHash(token))

	return ok, session
}
