package silk

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	//
	"github.com/tilinna/z85"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
	//
	"github.com/golangdaddy/tarantula/graph"
)

const (
	SHORT_ID_BUF_LEN = 12
	SESSION_KEY_LEN  = 64
)

func sessionKey() (string, string) {

	buf := make([]byte, graph.SESSION_KEY_LEN)

	rand.Read(buf)

	secret := hex.EncodeToString(buf)[:64]

	return secret, sessionKeyHash(secret)
}

func sessionKeyHash(secret string) string {

	hashedSecret := make([]byte, graph.SESSION_KEY_LEN)

	sha3.ShakeSum128(hashedSecret, []byte(secret))

	return hex.EncodeToString(hashedSecret)[:graph.SESSION_KEY_LEN]
}

func newHashedPassword(username, password string) (string, []byte) {

	sp := []string{
		username,
		password,
	}

	b, _ := json.Marshal(sp)
	c, _ := bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)

	return hex.EncodeToString(c), b
}

func comparePassword(passwordHash string, password []byte) error {

	hashBytes, err := hex.DecodeString(passwordHash)
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword(hashBytes, []byte(password))
}

func random() string {

	buf := make([]byte, graph.SHORT_ID_BUF_LEN)

	rand.Read(buf)

	return encode(buf, graph.SHA3_ID_LENGTH)
}

func encode(b []byte, length int) string {

	out := make([]byte, length)

	z85.Encode(out, b)

	return string(out)
}

func iD(input string) string {

	hash := make([]byte, graph.SHORT_ID_BUF_LEN)

	sha3.ShakeSum128(hash, []byte(input))

	return encode(hash, graph.SHA3_ID_LENGTH)
}

func Hash256(input []byte) string {

	b := make([]byte, 64)

	sha3.ShakeSum256(b, input)

	return hex.EncodeToString(b)
}
