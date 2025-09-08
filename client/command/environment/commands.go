package environment

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/wh1t3zer/sliver-server_new/client/command/flags"
	"github.com/wh1t3zer/sliver-server_new/client/command/help"
	"github.com/wh1t3zer/sliver-server_new/client/console"
	consts "github.com/wh1t3zer/sliver-server_new/client/constants"
)

// Commands returns the “ command and its subcommands.
func Commands(con *console.SliverClient) []*cobra.Command {
	envCmd := &cobra.Command{
		Use:   consts.EnvStr,
		Short: "List environment variables",
		Long:  help.GetHelpFor([]string{consts.EnvStr}),
		Args:  cobra.RangeArgs(0, 1),
		Run: func(cmd *cobra.Command, args []string) {
			EnvGetCmd(cmd, con, args)
		},
		GroupID: consts.InfoHelpGroup,
	}
	flags.Bind("", true, envCmd, func(f *pflag.FlagSet) {
		f.Int64P("timeout", "t", flags.DefaultTimeout, "grpc timeout in seconds")
	})
	carapace.Gen(envCmd).PositionalCompletion(carapace.ActionValues().Usage("environment variable to fetch (optional)"))

	envSetCmd := &cobra.Command{
		Use:   consts.SetStr,
		Short: "Set environment variables",
		Long:  help.GetHelpFor([]string{consts.EnvStr, consts.SetStr}),
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			EnvSetCmd(cmd, con, args)
		},
	}
	envCmd.AddCommand(envSetCmd)
	carapace.Gen(envSetCmd).PositionalCompletion(
		carapace.ActionValues().Usage("environment variable name"),
		carapace.ActionValues().Usage("value to assign"),
	)

	envUnsetCmd := &cobra.Command{
		Use:   consts.UnsetStr,
		Short: "Clear environment variables",
		Long:  help.GetHelpFor([]string{consts.EnvStr, consts.UnsetStr}),
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			EnvUnsetCmd(cmd, con, args)
		},
	}
	envCmd.AddCommand(envUnsetCmd)
	carapace.Gen(envUnsetCmd).PositionalCompletion(carapace.ActionValues().Usage("environment variable name"))

	return []*cobra.Command{envCmd}
}
