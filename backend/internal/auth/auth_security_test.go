package auth

import (
    "testing"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "github.com/stretchr/testify/require"
)

func TestValidateConfig(t *testing.T) {
    tests := []struct {
        name            string
        secret          string
        expirationHours int
        wantError       bool
    }{
        {
            name:            "valid configuration",
            secret:          "01234567890123456789012345678901",
            expirationHours: 1,
            wantError:       false,
        },
        {
            name:            "secret too short",
            secret:          "short-secret",
            expirationHours: 1,
            wantError:       true,
        },
        {
            name:            "invalid expiration",
            secret:          "01234567890123456789012345678901",
            expirationHours: 0,
            wantError:       true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            service := NewTokenService(tt.secret, tt.expirationHours)
            err := service.ValidateConfig()

            if tt.wantError {
                require.Error(t, err)
            } else {
                require.NoError(t, err)
            }
        })
    }
}

func TestValidateTokenRejectsExpiredToken(t *testing.T) {
    secret := "01234567890123456789012345678901"
    service := NewTokenService(secret, 1)

    claims := jwt.MapClaims{
        "sub": "test-user",
        "iat": time.Now().Add(-2 * time.Hour).Unix(),
        "exp": time.Now().Add(-1 * time.Hour).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString([]byte(secret))
    require.NoError(t, err)

    _, err = service.ValidateToken(signedToken)
    require.Error(t, err)
}

func TestValidateTokenRejectsNonHS256Algorithm(t *testing.T) {
    secret := "01234567890123456789012345678901"
    service := NewTokenService(secret, 1)

    claims := jwt.MapClaims{
        "sub": "test-user",
        "iat": time.Now().Unix(),
        "exp": time.Now().Add(time.Hour).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
    signedToken, err := token.SignedString([]byte(secret))
    require.NoError(t, err)

    _, err = service.ValidateToken(signedToken)
    require.Error(t, err)
}
