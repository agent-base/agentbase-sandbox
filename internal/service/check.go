package service

import (
	"errors"

	"github.com/agent-base/agentbase-sandbox/internal/core/runner/types"
	"github.com/agent-base/agentbase-sandbox/internal/static"
)

var (
	ErrNetworkDisabled = errors.New("network is disabled, please enable it in the configuration")
)

func checkOptions(options *types.RunnerOptions) error {
	configuration := static.GetAgentBaseSandboxGlobalConfigurations()

	if options.EnableNetwork && !configuration.EnableNetwork {
		return ErrNetworkDisabled
	}

	return nil
}
