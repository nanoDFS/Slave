package auth

import (
	"github.com/nanoDFS/Slave/controller/auth/acl"
)

type Auth struct {
}

func NewAuth() Auth {
	return Auth{}
}

func (t Auth) authorize(token string) (*acl.Claims, error) {
	claim, err := acl.NewJWT().Validate(token)
	if err != nil {
		return nil, err
	}
	return claim, nil
}

func (t Auth) AuthorizeRead(token string) (*acl.Claims, bool) {
	claim, err := t.authorize(token)
	if err != nil || claim.Mode == acl.Read {
		return nil, false
	}
	return claim, true
}

func (t Auth) AuthorizeWrite(token string) (*acl.Claims, bool) {
	claim, err := t.authorize(token)
	if err != nil || claim.Mode != acl.Write {
		return nil, false
	}
	return claim, true
}

func (t Auth) AuthorizeDelete(token string) (*acl.Claims, bool) {
	claim, err := t.authorize(token)
	if err != nil || claim.Mode == acl.Write {
		return nil, false
	}
	return claim, true
}
