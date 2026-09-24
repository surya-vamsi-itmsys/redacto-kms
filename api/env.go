package api

import (
	"os"
	"strings"
)

const (
	// RedactoKMSEnvPrefix is the prefix of every public Redacto KMS
	// environment variable. Legacy BAO_* and VAULT_* names are deliberately
	// not consulted as fallbacks.
	RedactoKMSEnvPrefix = "REDACTO_KMS_"

	// OpenBaoEnvPrefix and UpstreamEnvPrefix are retained only for the
	// server-to-plugin process protocol (see PluginAutoMTLSEnv and friends),
	// where external plugins built against the upstream SDK expect them.
	OpenBaoEnvPrefix  = "BAO_"
	UpstreamEnvPrefix = "VAULT_"
)

// UpstreamVariableName maps a BAO_-prefixed plugin protocol variable to its
// VAULT_-prefixed equivalent. It is used only when constructing the
// environment of plugin processes.
func UpstreamVariableName(name string) string {
	if !strings.HasPrefix(name, OpenBaoEnvPrefix) {
		return name
	}

	nonPrefixedName := strings.Replace(name, OpenBaoEnvPrefix, "", 1)
	return UpstreamEnvPrefix + nonPrefixedName
}

// ReadBaoVariable returns the value of the named environment variable. No
// alternate-prefix fallback is performed.
func ReadBaoVariable(name string) string {
	return os.Getenv(name)
}

// LookupBaoVariable looks up the named environment variable. No
// alternate-prefix fallback is performed.
func LookupBaoVariable(name string) (string, bool) {
	return os.LookupEnv(name)
}
