package config

import (
	"github.com/layer5io/meshkit/errors"
)

const (
	ErrEmptyConfigCode = "ai_adapter_1000"
)

var (
	ErrEmptyConfig = errors.New(ErrEmptyConfigCode, errors.Alert, []string{"config cannot be empty"}, []string{}, []string{}, []string{})
)