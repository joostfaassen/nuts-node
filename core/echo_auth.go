package core

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lestrrat-go/jwx/jwk"
	"os"
	"time"
)

const authTokenSignAlg = "ES256"

func HTTPAuthenticatorProvider(tokenKeyFile string) func(authType AuthType) (HTTPAuthenticator, error) {
	return func(authType AuthType) (HTTPAuthenticator, error) {
		switch authType {
		case TokenAuthType:
			return NewTokenAuthenticator(tokenKeyFile)
		case NoAuthAuthType:
			fallthrough
		case "":
			return noopAuthenticator{}, nil
		default:
			return nil, fmt.Errorf("invalid auth type: %s", authType)
		}
	}
}

type HTTPAuthenticator interface {
	authenticator() echo.MiddlewareFunc
}

type noopAuthenticator struct {}

func (n noopAuthenticator) authenticator() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			return next(context)
		}
	}
}

func NewTokenAuthenticator(jwkKeyFile string) (*TokenAuthenticator, error) {
	_, err := os.Lstat(jwkKeyFile)
	if os.IsNotExist(err) {
		Logger().Info("No key found for HTTP token authentication, generating.")
		privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		keyAsJWK, _ := jwk.New(privateKey)
		jwkBytes, _ := json.Marshal(keyAsJWK)
		err := os.WriteFile(jwkKeyFile, jwkBytes, 0600)
		if err != nil {
			return nil, fmt.Errorf("unable to write key file (file=%s): %w", jwkKeyFile, err)
		}
		return &TokenAuthenticator{key: privateKey}, nil
	} else if err != nil {
		return nil, fmt.Errorf("unable to stat existing key file (file=%s): %w", jwkKeyFile, err)
	} else {
		jwkBytes, err := os.ReadFile(jwkKeyFile)
		if err != nil && !os.IsNotExist(err) {
			return nil, fmt.Errorf("unable to read existing key file (file=%s): %w", jwkKeyFile, err)
		}
		jwkSet, err := jwk.Parse(jwkBytes)
		if err != nil {
			return nil, fmt.Errorf("unable to read JWK from key file (file=%s): %w", jwkKeyFile, err)
		}
		existingJWK, ok := jwkSet.Get(0)
		if !ok {
			return nil, fmt.Errorf("invalid number of keys in JWK file (file=%s)", jwkKeyFile)
		}
		var key ecdsa.PrivateKey
		err = existingJWK.Raw(&key)
		if err != nil {
			return nil, fmt.Errorf("unable to read ECDSA private key from JWK file (file=%s): %w", jwkKeyFile, err)
		}
		middleware.Logger()
		return &TokenAuthenticator{key: &key}, nil
	}
}

type TokenAuthenticator struct {
	key *ecdsa.PrivateKey
}

func (p TokenAuthenticator) CreateToken(subject string, validity time.Duration) (string, error) {
	Logger().Infof("Creating HTTP authentication token for %s", subject)
	token := jwt.NewWithClaims(jwt.GetSigningMethod(authTokenSignAlg), jwt.StandardClaims{
		ExpiresAt: time.Now().Add(validity).Unix(),
		IssuedAt:  time.Now().Unix(),
		NotBefore: time.Now().Unix(),
		Subject:   subject,
	})
	return token.SignedString(p.key)
}

func (p *TokenAuthenticator) authenticator() echo.MiddlewareFunc {
	cfg := middleware.DefaultJWTConfig
	cfg.SigningMethod = authTokenSignAlg
	cfg.SigningKey = p.key.Public()
	return middleware.JWTWithConfig(cfg)
}
