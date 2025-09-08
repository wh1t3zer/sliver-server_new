package operators

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
	operatorsCmd := &cobra.Command{
		Use:   consts.OperatorsStr,
		Short: "Manage operators",
		Long:  help.GetHelpFor([]string{consts.OperatorsStr}),
		Run: func(cmd *cobra.Command, args []string) {
			OperatorsCmd(cmd, con, args)
		},
		GroupID: consts.GenericHelpGroup,
	}
	flags.Bind("operators", false, operatorsCmd, func(f *pflag.FlagSet) {
		f.IntP("timeout", "t", flags.DefaultTimeout, "grpc timeout in seconds")
	})

	return []*cobra.Command{operatorsCmd}
}
