package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entity_policies"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/scope"
)

type Policy struct {
	network *net.Network
}

func NewPolicy(network *net.Network) (cmd *Policy) {
	cmd = new(Policy)
	cmd.network = network
	return
}

func (cmd *Policy) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "policy",
		Description: "Show the status of a policy for an application and entity",
		Usage:       "BROOKLYN_NAME [ SCOPE ] policy POLICY",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Policy) Run(scope scope.Scope, c *cli.Context) {
	policy := entity_policies.PolicyStatus(cmd.network, scope.Application, scope.Entity, c.Args().First())
	fmt.Println(policy)
}