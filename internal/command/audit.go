// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package command

import (
	"strings"

	"github.com/hashicorp/cli"
)

var _ cli.Command = (*AuditCommand)(nil)

type AuditCommand struct {
	*BaseCommand
}

func (c *AuditCommand) Synopsis() string {
	return "Interact with audit devices"
}

func (c *AuditCommand) Help() string {
	helpText := `
Usage: redacto-kms audit <subcommand> [options] [args]

  This command groups subcommands for interacting with Redacto KMS's audit devices.
  Users can list, enable, and disable audit devices.

  *NOTE*: Once an audit device has been enabled, failure to audit could prevent
  Redacto KMS from servicing future requests. It is highly recommended that you enable
  multiple audit devices.

  List all enabled audit devices:

      $ redacto-kms audit list

  Enable a new audit device "file";

       $ redacto-kms audit enable file file_path=/var/log/audit.log

  Please see the individual subcommand help for detailed usage information.
`

	return strings.TrimSpace(helpText)
}

func (c *AuditCommand) Run(args []string) int {
	return cli.RunResultHelp
}
