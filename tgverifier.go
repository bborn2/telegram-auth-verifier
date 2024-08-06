package tgverifier

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"sort"
)

// ErrInvalidCreds ...
// Represents error in case of having invalid Telegram auth credentials
var ErrInvalidCreds = errors.New("invalid telegram creds")

// Credentials ...
// Telegram Login credentials available for parsing from JSON.
type Credentials struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	PhotoURL  string `json:"photo_url"`
	AuthDate  int64  `json:"auth_date"`
	Hash      string `json:"hash"`
}

// Verify ...
// Checks if the credentials are from Telegram.
// Returns nil error if credentials are from Telegram.
// https://core.telegram.org/widgets/login
func (c *Credentials) Verify(token []byte) error {
	secret := sha256.Sum256(token)

	checkString := c.String()

	authCode := computeHmac256([]byte(checkString), secret[:])
	hexAuthCode := hex.EncodeToString(authCode)

	if hexAuthCode != c.Hash {
		return ErrInvalidCreds
	}

	return nil
}

// String ...
// Builds credentials string, excluding hash field.
func (c *Credentials) String() string {

	val := reflect.ValueOf(c).Elem()
	typ := val.Type()

	// 存储 JSON 标签和对应的值
	var kvPairs []string
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" && jsonTag != "hash" {

			v, err := url.QueryUnescape(fmt.Sprintf("%v", val.Field(i).Interface()))
			if err != nil || v == "" {
				continue
			}

			kvPairs = append(kvPairs, fmt.Sprintf("%s=%v", jsonTag, v))
		}
	}

	// 对 JSON 标签进行排序
	sort.Strings(kvPairs)

	// 拼接成 k1=v1&k2=v2&... 的形式
	result := ""
	for i, kv := range kvPairs {
		if i > 0 {
			result += "\n"
		}
		result += kv
	}

	return result
}

func computeHmac256(msg []byte, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(msg)
	return h.Sum(nil)
}
