package internal

import "errors"

var (
	ErrBadParam   = errors.New("bad_parameter")
	ErrSvc        = errors.New("service_error")
	ErrRepo       = errors.New("repo_error")
	ErrInvalidEnv = errors.New("invalid_env")
	ErrConfig     = errors.New("config_error")
)
