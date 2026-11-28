package auth

import (
    "errors"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
    secretKey     string
    expirationDur time.Duration
}

func NewTokenService(secret string, expirationHours int) *TokenService {
    return &TokenService{
        secretKey:     secret,
        expirationDur: time.Hour * time.Duration(expirationHours),
    }
}

func (s *TokenService) ValidateConfig() error {
    if len([]byte(s.secretKey)) < 32 {
        return errors.New("JWT secret must be at least 32 bytes long")
    }

    if s.expirationDur <= 0 {
        return errors.New("JWT expiration must be greater than zero")
    }

    return nil
}

func (s *TokenService) GenerateToken(userID string) (string, error) {
    if err := s.ValidateConfig(); err != nil {
        return "", err
    }

    claims := jwt.MapClaims{
        "sub": userID,
        "iat": time.Now().Unix(),
        "exp": time.Now().Add(s.expirationDur).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return token.SignedString([]byte(s.secretKey))
}

func (s *TokenService) ValidateToken(tokenString string) (string, error) {
    if err := s.ValidateConfig(); err != nil {
        return "", err
    }

    token, err := jwt.Parse(
        tokenString,
        func(token *jwt.Token) (interface{}, error) {
            if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
                return nil, errors.New("unexpected signing method")
            }

            return []byte(s.secretKey), nil
        },
        jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
    )

    if err != nil {
        return "", errors.New("invalid token")
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        userID, ok := claims["sub"].(string)
        if !ok {
            return "", errors.New("invalid token claims: subject not found")
        }

        return userID, nil
    }

    return "", errors.New("invalid token")
}

func (s *TokenService) GetExpirationSeconds() int {
    return int(s.expirationDur.Seconds())
}
