// Copyright (c) 2026 Redacto KMS
// SPDX-License-Identifier: MPL-2.0

package command

import (
	"os"
	"testing"
)

func TestBridgeThirdPartyEnv(t *testing.T) {
	// Legacy library names set directly must be ignored.
	t.Setenv("BAO_KMIP_ENDPOINT", "legacy:5696")
	t.Setenv("VAULT_TRANSIT_SEAL_ADDR", "http://legacy:8200")
	t.Setenv("REDACTO_KMS_KMIP_ENDPOINT", "redacto:5696")

	bridgeThirdPartyEnv()

	if got := os.Getenv("BAO_KMIP_ENDPOINT"); got != "redacto:5696" {
		t.Fatalf("expected bridged value, got %q", got)
	}
	if _, ok := os.LookupEnv("VAULT_TRANSIT_SEAL_ADDR"); ok {
		t.Fatal("legacy VAULT_TRANSIT_SEAL_ADDR must not act as a fallback")
	}
}
