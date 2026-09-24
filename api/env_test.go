package api

import (
	"testing"
)

func TestReadBaoVariable_RedactoKMS(t *testing.T) {
	actual := "example_value"
	t.Setenv("REDACTO_KMS_TEST", actual)
	expected := ReadBaoVariable("REDACTO_KMS_TEST")
	if actual != expected {
		t.Fatalf("bad: Failed to Read Environment Variable actual: %s expected: %s", actual, expected)
	}
}

// Legacy BAO_ and VAULT_ names must not act as fallbacks for REDACTO_KMS_.
func TestReadBaoVariable_NoLegacyFallback(t *testing.T) {
	t.Setenv("BAO_TEST", "bao_value")
	t.Setenv("VAULT_TEST", "vault_value")
	if v := ReadBaoVariable("REDACTO_KMS_TEST"); v != "" {
		t.Fatalf("bad: legacy variable used as fallback: %q", v)
	}
	if v, ok := LookupBaoVariable("REDACTO_KMS_TEST"); ok || v != "" {
		t.Fatalf("bad: legacy variable used as fallback: %q, %v", v, ok)
	}
}

func TestDefaultConfig_NoLegacyAddrFallback(t *testing.T) {
	t.Setenv("BAO_ADDR", "http://bao.example:1234")
	t.Setenv("VAULT_ADDR", "http://vault.example:1234")
	t.Setenv(EnvVaultAddress, "")
	cfg := DefaultConfig()
	if cfg.Error != nil {
		t.Fatal(cfg.Error)
	}
	if cfg.Address != "https://127.0.0.1:8200" {
		t.Fatalf("bad: legacy address used: %q", cfg.Address)
	}

	t.Setenv(EnvVaultAddress, "http://redacto.example:8200")
	cfg = DefaultConfig()
	if cfg.Address != "http://redacto.example:8200" {
		t.Fatalf("bad: REDACTO_KMS_ADDR not used: %q", cfg.Address)
	}
}
