package config

import (
	"github.com/layer5io/meshery-adapter-library/adapter"
)

func getOperations(ops adapter.Operations) adapter.Operations {
	ops[AIOperation] = &adapter.Operation{
		Type:        int32(0),
		Description: "Generate infrastructure design from a natural language prompt",
	}
	return ops
}