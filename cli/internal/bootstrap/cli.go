package bootstrap

import (
	"github.com/filprpx/conduit-demo/cli/internal/auth"
	"github.com/filprpx/conduit-demo/cli/internal/cli"
	"github.com/filprpx/conduit-demo/cli/internal/config"
	"github.com/filprpx/conduit-demo/cli/internal/platform"
	"github.com/spf13/cobra"
)

func NewCLICommand() *cobra.Command {
	return cli.NewCommand(cli.Dependencies{
		LoadConfig: config.FromEnv,
		NewAuth:    auth.NewManager,
		NewAPI:     platform.NewClient,
	})
}
