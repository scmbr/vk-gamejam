package usecase

import "errors"

var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrUserAlreadyExists =  errors.New("user already exists")