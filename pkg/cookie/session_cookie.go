package cookie

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/flowista2/models"
	"github.com/flowista2/pkg"
	"github.com/flowista2/pkg/caching"
	"github.com/spf13/viper"
)


func SessionCookie(user models.User) (*http.Cookie, error) {
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

	caching.RedisClient.Set(encodedData, encryptedData, 0)
	return &http.Cookie{
		Name:     "session",
		Value:    encodedData,
		HttpOnly: false,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}, nil
}

func SessionUser(cookie *http.Cookie) (*models.User, error) {
	var key = viper.GetString("session.key")
	encryptedData, err := caching.RedisClient.Get(cookie.Value).Bytes()
	if err != nil {
		return nil, err
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
