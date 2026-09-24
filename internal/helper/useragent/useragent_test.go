// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package useragent

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserAgent(t *testing.T) {
	rt = "go5.0"
	versionFunc = func() string { return "1.2.3" }

	act := String()

	exp := "Redacto KMS/1.2.3 (go5.0)"
	require.Equal(t, exp, act)
}

// TestUserAgent_VaultAgent tests the AgentString() function works
// as expected
func TestUserAgent_VaultAgent(t *testing.T) {
	rt = "go5.0"
	versionFunc = func() string { return "1.2.3" }

	act := AgentString()

	exp := "Redacto KMS Agent/1.2.3 (go5.0)"
	require.Equal(t, exp, act)
}

// TestUserAgent_VaultAgentTemplating tests the AgentTemplatingString() function works
// as expected
func TestUserAgent_VaultAgentTemplating(t *testing.T) {
	rt = "go5.0"
	versionFunc = func() string { return "1.2.3" }

	act := AgentTemplatingString()

	exp := "Redacto KMS Agent Templating/1.2.3 (go5.0)"
	require.Equal(t, exp, act)
}

// TestUserAgent_VaultAgentProxy tests the AgentProxyString() function works
// as expected
func TestUserAgent_VaultAgentProxy(t *testing.T) {
	rt = "go5.0"
	versionFunc = func() string { return "1.2.3" }

	act := AgentProxyString()

	exp := "Redacto KMS Agent API Proxy/1.2.3 (go5.0)"
	require.Equal(t, exp, act)
}

// TestUserAgent_VaultAgentProxyWithProxiedUserAgent tests the AgentProxyStringWithProxiedUserAgent()
// function works as expected
func TestUserAgent_VaultAgentProxyWithProxiedUserAgent(t *testing.T) {
	rt = "go5.0"
	versionFunc = func() string { return "1.2.3" }
	userAgent := "my-user-agent"

	act := AgentProxyStringWithProxiedUserAgent(userAgent)

	exp := "Redacto KMS Agent API Proxy/1.2.3 (go5.0); my-user-agent"
	require.Equal(t, exp, act)
}

// TestUserAgent_VaultAgentAutoAuth tests the AgentAutoAuthString() function works
// as expected
func TestUserAgent_VaultAgentAutoAuth(t *testing.T) {
	rt = "go5.0"
	versionFunc = func() string { return "1.2.3" }

	act := AgentAutoAuthString()

	exp := "Redacto KMS Agent Auto-Auth/1.2.3 (go5.0)"
	require.Equal(t, exp, act)
}

// TestUserAgent_VaultProxy tests the ProxyString() function works
// as expected
func TestUserAgent_VaultProxy(t *testing.T) {
	rt = "go5.0"
	versionFunc = func() string { return "1.2.3" }

	act := ProxyString()

	exp := "Redacto KMS Proxy/1.2.3 (go5.0)"
	require.Equal(t, exp, act)
}

// TestUserAgent_VaultProxyAPIProxy tests the ProxyAPIProxyString() function works
// as expected
func TestUserAgent_VaultProxyAPIProxy(t *testing.T) {
	rt = "go5.0"
	versionFunc = func() string { return "1.2.3" }

	act := ProxyAPIProxyString()

	exp := "Redacto KMS Proxy API Proxy/1.2.3 (go5.0)"
	require.Equal(t, exp, act)
}

// TestUserAgent_VaultProxyWithProxiedUserAgent tests the ProxyStringWithProxiedUserAgent()
// function works as expected
func TestUserAgent_VaultProxyWithProxiedUserAgent(t *testing.T) {
	rt = "go5.0"
	versionFunc = func() string { return "1.2.3" }
	userAgent := "my-user-agent"

	act := ProxyStringWithProxiedUserAgent(userAgent)

	exp := "Redacto KMS Proxy API Proxy/1.2.3 (go5.0); my-user-agent"
	require.Equal(t, exp, act)
}

// TestUserAgent_VaultProxyAutoAuth tests the ProxyAPIProxyString() function works
// as expected
func TestUserAgent_VaultProxyAutoAuth(t *testing.T) {
	rt = "go5.0"
	versionFunc = func() string { return "1.2.3" }

	act := ProxyAutoAuthString()

	exp := "Redacto KMS Proxy Auto-Auth/1.2.3 (go5.0)"
	require.Equal(t, exp, act)
}
