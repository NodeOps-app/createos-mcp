package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port                   uint16 `yaml:"port"`
	BaseURL                string `yaml:"base_url"`
	AuthorizationServerUrl string `yaml:"authorization_server_url"`
	APIBaseUrl             string `yaml:"api_base_url"`
	Transport              string `yaml:"transport"`
	LogLevel               string `yaml:"log_level"`
	Toolset                string `yaml:"toolset"`

	// Scopes and claims
	SupportedScopes                   []string `yaml:"supported_scopes"` // OAuth server scopes
	ResponseTypesSupported            []string `yaml:"response_types_supported"`
	GrantTypesSupported               []string `yaml:"grant_types_supported"`
	CodeChallengeMethodsSupported     []string `yaml:"code_challenge_methods_supported"`
	TokenEndpointAuthMethodsSupported []string `yaml:"token_endpoint_auth_methods_supported"`

	// Well-known OAuth Authorization Server Metadata
	Issuer                string `yaml:"issuer"`
	AuthorizationEndpoint string `yaml:"authorization_endpoint"`
	TokenEndpoint         string `yaml:"token_endpoint"`
	RevokeEndpoint        string `yaml:"revoke_endpoint"`
}

var Cfg *Config

func LoadConfig(path string) error {
	cfg, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	return yaml.Unmarshal(cfg, &Cfg)
}
