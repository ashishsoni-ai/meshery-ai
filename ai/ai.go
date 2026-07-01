// Package ai - AI Adapter for Meshery: translates natural language prompts into Meshery Designs via a local or cloud LLM.
package ai

import (
	"context"
	"fmt"

	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-adapter-library/meshes"
	"github.com/layer5io/meshery-adapter-library/status"
	internalconfig "github.com/ashishsoni-ai/meshery-ai/internal/config"
	"github.com/layer5io/meshkit/config"
	"github.com/layer5io/meshkit/errors"
	"github.com/layer5io/meshkit/logger"
	"github.com/layer5io/meshkit/models/oam/core/v1alpha1"
	"github.com/layer5io/meshkit/utils/events"
)

// AI represents the AI adapter and embeds adapter.Adapter
type AI struct {
	adapter.Adapter
	Ollama *OllamaClient
}

// New initializes the AI adapter handler.
func New(config config.Handler, log logger.Handler, kubeConfig config.Handler, e *events.EventStreamer) adapter.Handler {
	return &AI{
		Adapter: adapter.Adapter{Config: config, Log: log, KubeconfigHandler: kubeConfig, EventStreamer: e},
		Ollama:  NewOllamaClient(),
	}
}

const systemPrompt = `You are an infrastructure design assistant for Meshery. Given a natural language description of desired cloud native infrastructure, respond with a concise summary of the Kubernetes resources needed to satisfy the request. Do not include explanations outside the summary.`

// ApplyOperation applies the operation on the AI adapter.
func (ai *AI) ApplyOperation(ctx context.Context, opReq adapter.OperationRequest) error {
	e := &meshes.EventsResponse{
		OperationId: opReq.OperationID,
		Summary:     status.Deploying,
		Component:   "adapter",
		ComponentName: "ai-adapter",
	}

	switch opReq.OperationName {
	case internalconfig.AIOperation:
		go func(hh *AI, ee *meshes.EventsResponse) {
			result, err := hh.Ollama.Generate(systemPrompt, opReq.CustomBody)
			if err != nil {
				summary := "Error while generating design from prompt"
				hh.streamErr(summary, ee, ErrOllamaRequest(err))
				return
			}
			ee.Summary = "Design generated from natural language prompt"
			ee.Details = result
			hh.StreamInfo(ee)
		}(ai, e)
	default:
		ai.streamErr("Invalid operation", e, ErrOpInvalid)
	}

	return nil
}

// ProcessOAM is a stub for now — AI adapter does not manage OAM components directly.
func (ai *AI) ProcessOAM(ctx context.Context, oamReq adapter.OAMRequest) (string, error) {
	return "", fmt.Errorf("ProcessOAM not implemented for ai-adapter")
}

func (ai *AI) streamErr(summary string, e *meshes.EventsResponse, err error) {
	e.Summary = summary
	e.Details = err.Error()
	e.ErrorCode = errors.GetCode(err)
	e.ProbableCause = errors.GetCause(err)
	e.SuggestedRemediation = errors.GetRemedy(err)
	ai.StreamErr(e, err)
}

var _ = v1alpha1.Component{} // keep import until real OAM support lands