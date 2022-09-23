package constant

import (
	"os"
	"time"
)

var (
	JWTSecret = os.Getenv("JWT_SECRET")
	JWTExp    = time.Now().Add(time.Hour * 24).Unix()
)
