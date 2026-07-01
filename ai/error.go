package ai

import (
	"github.com/layer5io/meshkit/errors"
)

const (
	ErrOpInvalidCode = "ai_adapter_1001"
	ErrOllamaCode    = "ai_adapter_1002"
)

var (
	ErrOpInvalid = errors.New(ErrOpInvalidCode, errors.Alert, []string{"invalid operation"}, []string{}, []string{}, []string{})
)

func ErrOllamaRequest(err error) error {
	return errors.New(ErrOllamaCode, errors.Alert, []string{"error calling Ollama"}, []string{err.Error()}, []string{"Ollama server unreachable or model not pulled"}, []string{"Ensure `ollama serve` is running and the model is pulled with `ollama pull <model>`"})
}