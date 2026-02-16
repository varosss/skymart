package auth

import (
	"clirzy/auth/internal/domain"
	"clirzy/auth/internal/domain/port"
	"clirzy/auth/internal/domain/valueobject"
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v5"
)

type JWTVerifier struct {
	publicKey *rsa.PublicKey
	issuer    string
}

func NewJWTVerifier(
	publicKey *rsa.PublicKey,
	issuer string,
) *JWTVerifier {
	return &JWTVerifier{
		publicKey: publicKey,
		issuer:    issuer,
	}
}

func (v *JWTVerifier) VerifyAccess(
	tokenString string,
) (*port.AccessTokenClaims, error) {

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, domain.ErrInvalidToken
			}
			return v.publicKey, nil
		},
		jwt.WithIssuer(v.issuer),
		jwt.WithValidMethods([]string{"RS256"}),
	)
	if err != nil || !token.Valid {
		return nil, domain.ErrInvalidToken
	}

	sub, ok := claims["sub"].(string)
	if !ok || sub == "" {
		return nil, domain.ErrInvalidClaims
	}

	rawRoles, ok := claims["roles"].([]interface{})
	if !ok {
		return nil, domain.ErrInvalidClaims
	}

	roles := make([]valueobject.RoleCode, 0, len(rawRoles))
	for _, r := range rawRoles {
		roleStr, ok := r.(string)
		if !ok {
			return nil, domain.ErrInvalidClaims
		}
		roles = append(roles, valueobject.RoleCode(roleStr))
	}

	return &port.AccessTokenClaims{
		UserID: valueobject.UserID(sub),
		Roles:  roles,
	}, nil
}

func (v *JWTVerifier) VerifyRefresh(
	tokenString string,
) (valueobject.TokenID, valueobject.UserID, error) {

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, domain.ErrInvalidToken
			}
			return v.publicKey, nil
		},
		jwt.WithIssuer(v.issuer),
		jwt.WithValidMethods([]string{"RS256"}),
	)
	if err != nil {
		return "", "", domain.ErrInvalidToken
	}

	if !token.Valid {
		return "", "", domain.ErrInvalidToken
	}

	sub, ok := claims["sub"].(string)
	if !ok || sub == "" {
		return "", "", domain.ErrInvalidClaims
	}

	jti, ok := claims["jti"].(string)
	if !ok || jti == "" {
		return "", "", domain.ErrInvalidClaims
	}

	return valueobject.TokenID(jti), valueobject.UserID(sub), nil
}
