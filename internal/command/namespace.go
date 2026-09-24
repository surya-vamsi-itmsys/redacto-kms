// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package command

import (
	"strings"

	"github.com/hashicorp/cli"
)

var _ cli.Command = (*NamespaceCommand)(nil)

type NamespaceCommand struct {
	*BaseCommand
}

func (c *NamespaceCommand) Synopsis() string {
	return "Interact with namespaces"
}

func (c *NamespaceCommand) Help() string {
	helpText := `
Usage: redacto-kms namespace <subcommand> [options] [args]

  This command groups subcommands for interacting with Redacto KMS namespaces.

  List enabled child namespaces:

      $ redacto-kms namespace list

  List enabled child namespaces recursively:

      $ redacto-kms namespace scan

  Look up an existing namespace:

      $ redacto-kms namespace lookup

  Create a new namespace:

      $ redacto-kms namespace create

  Patch an existing namespace:

      $ redacto-kms namespace patch

  Delete an existing namespace:

      $ redacto-kms namespace delete

  Lock the API for an existing namespace:

      $ redacto-kms namespace lock

  Unlock the API for an existing namespace:

      $ redacto-kms namespace unlock

  Seal the namespace:  

      $ redacto-kms namespace seal

  Unseal the namespace:

      $ redacto-kms namespace unseal

  Please see the individual subcommand help for detailed usage information.
`

	return strings.TrimSpace(helpText)
}

func (c *NamespaceCommand) Run(args []string) int {
	return cli.RunResultHelp
}
