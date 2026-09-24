// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package useragent

import (
	"fmt"
	"runtime"

	"github.com/openbao/openbao/v2/internal/version"
)

var (

	// rt is the runtime - variable for tests.
	rt = runtime.Version()

	// versionFunc is the func that returns the current version.
	versionFunc = func() string {
		return version.GetVersion().VersionNumber()
	}
)

// String returns the consistent user-agent string for Redacto KMS.
//
// e.g. Redacto KMS/0.10.4 (go1.10.1)
func String() string {
	return fmt.Sprintf("Redacto KMS/%s (%s)",
		versionFunc(), rt)
}

// AgentString returns the consistent user-agent string for Redacto KMS Agent.
//
// e.g. Redacto KMS Agent/0.10.4 (go1.10.1)
func AgentString() string {
	return fmt.Sprintf("Redacto KMS Agent/%s (%s)",
		versionFunc(), rt)
}

// AgentTemplatingString returns the consistent user-agent string for Redacto KMS Agent Templating.
//
// e.g. Redacto KMS Agent Templating/0.10.4 (go1.10.1)
func AgentTemplatingString() string {
	return fmt.Sprintf("Redacto KMS Agent Templating/%s (%s)",
		versionFunc(), rt)
}

// AgentProxyString returns the consistent user-agent string for Redacto KMS Agent API Proxying.
//
// e.g. Redacto KMS Agent API Proxy/0.10.4 (go1.10.1)
func AgentProxyString() string {
	return fmt.Sprintf("Redacto KMS Agent API Proxy/%s (%s)",
		versionFunc(), rt)
}

// AgentProxyStringWithProxiedUserAgent returns the consistent user-agent
// string for Vault Agent API Proxying, keeping the User-Agent of the proxied
// client as an extension to this UserAgent
//
// e.g. Redacto KMS Agent API Proxy/0.10.4 (go1.10.1); proxiedUserAgent
func AgentProxyStringWithProxiedUserAgent(proxiedUserAgent string) string {
	return fmt.Sprintf("Redacto KMS Agent API Proxy/%s (%s); %s",
		versionFunc(), rt, proxiedUserAgent)
}

// AgentAutoAuthString returns the consistent user-agent string for Redacto KMS Agent Auto-Auth.
//
// e.g. Redacto KMS Agent Auto-Auth/0.10.4 (go1.10.1)
func AgentAutoAuthString() string {
	return fmt.Sprintf("Redacto KMS Agent Auto-Auth/%s (%s)",
		versionFunc(), rt)
}

// ProxyString returns the consistent user-agent string for Redacto KMS Proxy.
//
// e.g. Redacto KMS Proxy/0.10.4 (go1.10.1)
func ProxyString() string {
	return fmt.Sprintf("Redacto KMS Proxy/%s (%s)",
		versionFunc(), rt)
}

// ProxyAPIProxyString returns the consistent user-agent string for Redacto KMS Proxy API Proxying.
//
// e.g. Redacto KMS Proxy API Proxy/0.10.4 (go1.10.1)
func ProxyAPIProxyString() string {
	return fmt.Sprintf("Redacto KMS Proxy API Proxy/%s (%s)",
		versionFunc(), rt)
}

// ProxyStringWithProxiedUserAgent returns the consistent user-agent
// string for Vault Proxy API Proxying, keeping the User-Agent of the proxied
// client as an extension to this UserAgent
//
// e.g. Redacto KMS Proxy API Proxy/0.10.4 (go1.10.1); proxiedUserAgent
func ProxyStringWithProxiedUserAgent(proxiedUserAgent string) string {
	return fmt.Sprintf("Redacto KMS Proxy API Proxy/%s (%s); %s",
		versionFunc(), rt, proxiedUserAgent)
}

// ProxyAutoAuthString returns the consistent user-agent string for Redacto KMS Agent Auto-Auth.
//
// e.g. Redacto KMS Proxy Auto-Auth/0.10.4 (go1.10.1)
func ProxyAutoAuthString() string {
	return fmt.Sprintf("Redacto KMS Proxy Auto-Auth/%s (%s)",
		versionFunc(), rt)
}
