package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const HMAC_SAMPLE_SECRET = "hmacSampleSecret"
const ACCESS_TOKEN_DURATION = time.Hour

type AccessTokenClaims struct {
	CustomerId string `json:"customer_id"`
	Username   string `json:"username"`
	Role       string `json:"role"`
	jwt.StandardClaims
}

func (c AccessTokenClaims) IsUserRole() bool {
	return c.Role == "user"
}

func (c AccessTokenClaims) IsValidCustomerId(customerId string) bool {
	return c.CustomerId == customerId
}

func (c AccessTokenClaims) IsRequestVerifiedWithTokenClaims(urlParams map[string]string) bool {
	return c.CustomerId == urlParams["customer_id"]
}
