// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package command

import (
	"strings"

	"github.com/hashicorp/cli"
)

var _ cli.Command = (*OperatorCommand)(nil)

type OperatorCommand struct {
	*BaseCommand
}

func (c *OperatorCommand) Synopsis() string {
	return "Perform operator-specific tasks"
}

func (c *OperatorCommand) Help() string {
	helpText := `
Usage: redacto-kms operator <subcommand> [options] [args]

  This command groups subcommands for operators interacting with Redacto KMS. Most
  users will not need to interact with these commands. Here are a few examples
  of the operator commands:

  Initialize a new Redacto KMS cluster:

      $ redacto-kms operator init

  Force a Redacto KMS node to resign leadership in a cluster:

      $ redacto-kms operator step-down

  Rotate Redacto KMS's underlying encryption key:

      $ redacto-kms operator rotate

  Please see the individual subcommand help for detailed usage information.
`

	return strings.TrimSpace(helpText)
}

func (c *OperatorCommand) Run(args []string) int {
	return cli.RunResultHelp
}
