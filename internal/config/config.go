package config

import (
	"path"

	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-adapter-library/config"
	"github.com/layer5io/meshery-adapter-library/status"
	configprovider "github.com/layer5io/meshkit/config/provider"
	"github.com/layer5io/meshkit/utils"
)

const (
	Development  = "development"
	Production   = "production"
	AIOperation  = "ai-generate"
)

var (
	configRootPath = path.Join(utils.GetHome(), ".meshery")

	ServerConfig = map[string]string{
		"name":     "ai-adapter",
		"type":     "adapter",
		"port":     "10015",
		"traceurl": status.None,
	}

	MeshSpec = map[string]string{
		"name":    "ai-adapter",
		"status":  status.NotInstalled,
		"version": status.None,
	}

	Operations = getOperations(map[string]*adapter.Operation{})
)

func New(provider string) (h config.Handler, err error) {
	opts := configprovider.Options{
		FilePath: configRootPath,
		FileName: "ai-adapter",
		FileType: "yaml",
	}
	switch provider {
	case configprovider.ViperKey:
		h, err = configprovider.NewViper(opts)
		if err != nil {
			return nil, err
		}
	case configprovider.InMemKey:
		h, err = configprovider.NewInMem(opts)
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrEmptyConfig
	}

	if err := h.SetObject(adapter.ServerKey, ServerConfig); err != nil {
		return nil, err
	}
	if err := h.SetObject(adapter.MeshSpecKey, MeshSpec); err != nil {
		return nil, err
	}
	if err := h.SetObject(adapter.OperationsKey, Operations); err != nil {
		return nil, err
	}
	return h, nil
}

func NewKubeconfigBuilder(provider string) (config.Handler, error) {
	opts := configprovider.Options{
		FilePath: configRootPath,
		FileType: "yaml",
		FileName: "kubeconfig",
	}
	switch provider {
	case configprovider.ViperKey:
		return configprovider.NewViper(opts)
	case configprovider.InMemKey:
		return configprovider.NewInMem(opts)
	}
	return nil, config.ErrEmptyConfig
}

func RootPath() string {
	return configRootPath
}