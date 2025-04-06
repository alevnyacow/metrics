// Package config obtains configuration data for agent and server.
//
// It provides configuration based on env variables, cli arguments
// and default values. Priority list: env variables, cli arguments,
// default preset values.
package config

// ParseAgentConfigs returns configuration data for agent application.
func ParseAgentConfigs() AgentConfigs {
	envConfigs := parseAgentEnvData()
	argsConfigs := parseAgentArgsConfigs()
	agentConfigs := mergeClientConfigs(envConfigs, argsConfigs)

	if !isLinkCorrect(agentConfigs.APIHost) {
		agentConfigs.APIHost = defaultAgentConfigs.APIHost
		return agentConfigs
	}

	if isLocalhostWithoutPrefix(agentConfigs.APIHost) {
		agentConfigs.APIHost = withHTTPPrefix(agentConfigs.APIHost)
	}

	return agentConfigs
}

// ParseServerConfigs returns configuration data for server application.
func ParseServerConfigs() ServerConfigs {
	envConfigs := parseServerEnvData()
	argsConfigs := parseServerArgsConfigs()
	serverConfigs := mergeServerConfigs(envConfigs, argsConfigs)

	if !isLinkCorrect(serverConfigs.APIHost) {
		serverConfigs.APIHost = defaultServerConfigs.APIHost
	}

	return serverConfigs
}
