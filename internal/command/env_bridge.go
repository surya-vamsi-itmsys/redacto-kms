// Copyright (c) 2026 Redacto KMS
// SPDX-License-Identifier: MPL-2.0

package command

import "os"

// thirdPartyEnvBridge maps public REDACTO_KMS_* environment variables to the
// names read directly by in-process third-party libraries (go-kms-wrapping
// seal wrappers), which cannot be renamed without forking those libraries.
//
// REDACTO_KMS_* is the only supported public interface: the library-specific
// names are cleared first and then populated solely from their REDACTO_KMS_*
// equivalents, so legacy BAO_* / VAULT_* values never act as a fallback.
var thirdPartyEnvBridge = map[string]string{
	// go-kms-wrapping/wrappers/kmip
	"REDACTO_KMS_KMIP_CA_CERT":        "BAO_KMIP_CA_CERT",
	"REDACTO_KMS_KMIP_CLIENT_CERT":    "BAO_KMIP_CLIENT_CERT",
	"REDACTO_KMS_KMIP_CLIENT_KEY":     "BAO_KMIP_CLIENT_KEY",
	"REDACTO_KMS_KMIP_ENCRYPT_ALG":    "BAO_KMIP_ENCRYPT_ALG",
	"REDACTO_KMS_KMIP_ENDPOINT":       "BAO_KMIP_ENDPOINT",
	"REDACTO_KMS_KMIP_SERVER_NAME":    "BAO_KMIP_SERVER_NAME",
	"REDACTO_KMS_KMIP_TIMEOUT":        "BAO_KMIP_TIMEOUT",
	"REDACTO_KMS_KMIP_TLS12_CIPHERS":  "BAO_KMIP_TLS12_CIPHERS",
	"REDACTO_KMS_KMIP_WRAPPER_KEY_ID": "BAO_KMIP_WRAPPER_KEY_ID",
	"REDACTO_KMS_KMIP_SEAL_KEY_ID":    "VAULT_KMIP_SEAL_KEY_ID",

	// go-kms-wrapping/wrappers/static
	"REDACTO_KMS_STATIC_SEAL_CURRENT_KEY":     "BAO_STATIC_SEAL_CURRENT_KEY",
	"REDACTO_KMS_STATIC_SEAL_CURRENT_KEY_ID":  "BAO_STATIC_SEAL_CURRENT_KEY_ID",
	"REDACTO_KMS_STATIC_SEAL_PREVIOUS_KEY":    "BAO_STATIC_SEAL_PREVIOUS_KEY",
	"REDACTO_KMS_STATIC_SEAL_PREVIOUS_KEY_ID": "BAO_STATIC_SEAL_PREVIOUS_KEY_ID",

	// go-kms-wrapping/wrappers/transit
	"REDACTO_KMS_TRANSIT_SEAL_ADDR":            "VAULT_TRANSIT_SEAL_ADDR",
	"REDACTO_KMS_TRANSIT_SEAL_TOKEN":           "VAULT_TRANSIT_SEAL_TOKEN",
	"REDACTO_KMS_TRANSIT_SEAL_KEY_NAME":        "VAULT_TRANSIT_SEAL_KEY_NAME",
	"REDACTO_KMS_TRANSIT_SEAL_MOUNT_PATH":      "VAULT_TRANSIT_SEAL_MOUNT_PATH",
	"REDACTO_KMS_TRANSIT_SEAL_DISABLE_RENEWAL": "VAULT_TRANSIT_SEAL_DISABLE_RENEWAL",
	"REDACTO_KMS_TRANSIT_SEAL_NAMESPACE":       "VAULT_NAMESPACE",
}

// bridgeThirdPartyEnv applies thirdPartyEnvBridge to the process environment.
func bridgeThirdPartyEnv() {
	for public, library := range thirdPartyEnvBridge {
		_ = os.Unsetenv(library)
		if v, ok := os.LookupEnv(public); ok {
			_ = os.Setenv(library, v)
		}
	}
}
