package model

import (
	"dto"
	"time"

	"github.com/mjibson/goon"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

type VerifyResult int

const (
	VERIFY_NOT_FOUND VerifyResult = iota
	VERIFY_EXPIRED
	VERIFY_OK
)

// Model
type AuthModel struct {
	Ctx  context.Context
	kind string
}

func NewAuthModel(ctx *context.Context) *AuthModel {
	return &AuthModel{
		Ctx:  *ctx,
		kind: "Auth",
	}
}

func (u *AuthModel) CreateToken(userId string) (*dto.Token, error) {
	g := goon.FromContext(u.Ctx)

	t := u.createToken()
	t.UserId = userId

	log.Debugf(u.Ctx, ("token uuid = " + t.Token))
	_, err := g.Put(t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (u *AuthModel) Verify(token string) (VerifyResult, *dto.Token) {
	g := goon.FromContext(u.Ctx)

	t := &dto.Token{
		Token: token,
	}

	if err := g.Get(t); err != nil {
		return VERIFY_NOT_FOUND, nil
	}

	if t.ExpiredAt < time.Now().UnixNano() {
		return VERIFY_EXPIRED, nil
	} else {
		return VERIFY_OK, t
	}

	return VERIFY_NOT_FOUND, nil
}

func (u *AuthModel) createToken() *dto.Token {
	return &dto.Token{
		Token:     getUUID().String(),
		ExpiredAt: time.Now().Add(time.Duration(24) * time.Hour).UnixNano(),
	}
}
