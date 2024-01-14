package helper

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/itsLeonB/posyandu-api/core/config"
	"time"
)

func GenerateJWT(id int, role string) (string, error) {
	var (
		cfg       = config.ProvideConfig()
		exp       = cfg.GetInt("JWT_EXPIRE")
		jwtSecret = cfg.Get("JWT_SECRET")
		jwtExpire = time.Now().Add(time.Hour * time.Duration(exp))
	)

	claims := jwt.MapClaims{
		"id":   id,
		"role": role,
		"exp":  jwtExpire.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecret))
}
