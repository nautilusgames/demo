package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type jwtImpl struct {
	signingKey string
	issuer     string
	audience   string
}

func New(signingKey, issuer, audience string) (Maker, error) {
	return &jwtImpl{
		signingKey: signingKey,
		issuer:     issuer,
		audience:   audience,
	}, nil
}

func (j *jwtImpl) CreateToken(object string, playerID int64, username string, duration time.Duration) (string, *Payload, error) {
	payload, err := newPayload(object, playerID, username, duration)
	if err != nil {
		return "", payload, err
	}

	claims := &MyClaim{
		Object:   object,
		Username: username,
		PlayerID: playerID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        payload.ID.String(),
			ExpiresAt: jwt.NewNumericDate(payload.ExpiredAt),
			IssuedAt:  jwt.NewNumericDate(payload.IssuedAt),
			Subject:   username,
			Issuer:    j.issuer,
			Audience:  make(jwt.ClaimStrings, 0),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(j.signingKey))
	if err != nil {
		return "", payload, err
	}

	return token, payload, err
}

func (j *jwtImpl) VerifyToken(token string) (*Payload, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(*MyClaim)
	if !ok {
		return nil, ErrInvalidToken
	}

	return &Payload{
		ID:        uuid.MustParse(claims.ID),
		Object:    claims.Object,
		PlayerID:  claims.PlayerID,
		Username:  claims.Username,
		IssuedAt:  claims.IssuedAt.Time,
		ExpiredAt: claims.ExpiresAt.Time,
	}, nil
}
