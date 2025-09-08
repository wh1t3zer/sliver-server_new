package builders

import (
	"github.com/wh1t3zer/sliver-server_new/client/command/flags"
	"github.com/wh1t3zer/sliver-server_new/client/command/help"
	"github.com/wh1t3zer/sliver-server_new/client/console"
	consts "github.com/wh1t3zer/sliver-server_new/client/constants"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Commands returns the “ command and its subcommands.
func Commands(con *console.SliverClient) []*cobra.Command {
	buildersCmd := &cobra.Command{
		Use:   consts.BuildersStr,
		Short: "List external builders",
		Long:  help.GetHelpFor([]string{consts.BuildersStr}),
		Run: func(cmd *cobra.Command, args []string) {
			BuildersCmd(cmd, con, args)
		},
		GroupID: consts.PayloadsHelpGroup,
	}
	flags.Bind("builders", false, buildersCmd, func(f *pflag.FlagSet) {
		f.Int64P("timeout", "t", flags.DefaultTimeout, "grpc timeout in seconds")
	})

	return []*cobra.Command{buildersCmd}
}
