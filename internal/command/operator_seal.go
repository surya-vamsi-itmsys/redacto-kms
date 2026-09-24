// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package command

import (
	"fmt"
	"strings"

	"github.com/hashicorp/cli"
	"github.com/posener/complete"
)

var (
	_ cli.Command             = (*OperatorSealCommand)(nil)
	_ cli.CommandAutocomplete = (*OperatorSealCommand)(nil)
)

type OperatorSealCommand struct {
	*BaseCommand
}

func (c *OperatorSealCommand) Synopsis() string {
	return "Seals the Redacto KMS server"
}

func (c *OperatorSealCommand) Help() string {
	helpText := `
Usage: redacto-kms operator seal [options]

  Seals the Redacto KMS server. Sealing tells the Redacto KMS server to stop responding
  to any operations until it is unsealed. When sealed, the Redacto KMS server
  discards its in-memory root key to unlock the data, so it is physically
  blocked from responding to operations unsealed.

  If an unseal is in progress, sealing Redacto KMS will reset the unsealing
  process. Users will have to re-enter their portions of the root key again.

  This command does nothing if the Redacto KMS server is already sealed.

  Seal the Redacto KMS server:

      $ redacto-kms operator seal

` + c.Flags().Help()

	return strings.TrimSpace(helpText)
}

func (c *OperatorSealCommand) Flags() *FlagSets {
	return c.flagSet(FlagSetHTTP)
}

func (c *OperatorSealCommand) AutocompleteArgs() complete.Predictor {
	return nil
}

func (c *OperatorSealCommand) AutocompleteFlags() complete.Flags {
	return c.Flags().Completions()
}

func (c *OperatorSealCommand) Run(args []string) int {
	f := c.Flags()

	if err := f.Parse(args); err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	args = f.Args()
	if len(args) > 0 {
		c.UI.Error(fmt.Sprintf("Too many arguments (expected 0, got %d)", len(args)))
		return 1
	}

	client, err := c.Client()
	if err != nil {
		c.UI.Error(err.Error())
		return 2
	}

	if err := client.Sys().Seal(); err != nil {
		c.UI.Error(fmt.Sprintf("Error sealing: %s", err))
		return 2
	}

	c.UI.Output("Success! Redacto KMS is sealed.")
	return 0
}
