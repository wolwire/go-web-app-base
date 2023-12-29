package cookie

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	"github.com/flowista2/models"
	"github.com/flowista2/pkg"
	"github.com/flowista2/pkg/caching"
	"github.com/spf13/viper"
)

func SessionCookie(user models.User, expiry_duration time.Duration) (*http.Cookie, error) {
	var key = viper.GetString("session.key")
	userJson, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	encryptedData, err := pkg.Encrypt([]byte(userJson), []byte(key))
	if err != nil {
		return nil, err
	}
	encodedData := base64.StdEncoding.EncodeToString(encryptedData)
	
	if time.Now().Add(expiry_duration).Before(time.Now().Add(time.Second)) {
		caching.RedisClient.Del(encodedData)
	} else {
		caching.RedisClient.Set(encodedData, encryptedData, expiry_duration)
	}

	return &http.Cookie{
		Name:     "session",
		Value:    encodedData,
		HttpOnly: false,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		Expires: time.Now().Add(expiry_duration),
	}, nil
}

type UserNotLoggedIn struct{}

func (user_error *UserNotLoggedIn) Error() string {
	return "User Session Expired"
}

func SessionUser(cookie *http.Cookie) (*models.User, error) {
	var key = viper.GetString("session.key")
	encryptedData, err := caching.RedisClient.Get(cookie.Value).Bytes()
	if err != nil {
		return nil, &UserNotLoggedIn{}
	}
	decryptedData, err := pkg.Decrypt(encryptedData, []byte(key))
	if err != nil {
		return nil, err
	}

	var user models.User
	err = json.Unmarshal(decryptedData, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
