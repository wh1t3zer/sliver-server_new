package use

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/wh1t3zer/sliver-server_new/client/command/beacons"
	"github.com/wh1t3zer/sliver-server_new/client/command/flags"
	"github.com/wh1t3zer/sliver-server_new/client/command/help"
	"github.com/wh1t3zer/sliver-server_new/client/console"
	consts "github.com/wh1t3zer/sliver-server_new/client/constants"
)

// Commands returns the “ command and its subcommands.
func Commands(con *console.SliverClient) []*cobra.Command {
	useCmd := &cobra.Command{
		Use:   consts.UseStr,
		Short: "Switch the active session or beacon",
		Long:  help.GetHelpFor([]string{consts.UseStr}),
		Run: func(cmd *cobra.Command, args []string) {
			UseCmd(cmd, con, args)
		},
		GroupID: consts.SliverHelpGroup,
	}
	flags.Bind("use", true, useCmd, func(f *pflag.FlagSet) {
		f.Int64P("timeout", "t", flags.DefaultTimeout, "grpc timeout in seconds")
	})
	carapace.Gen(useCmd).PositionalCompletion(BeaconAndSessionIDCompleter(con))

	useSessionCmd := &cobra.Command{
		Use:   consts.SessionsStr,
		Short: "Switch the active session",
		Long:  help.GetHelpFor([]string{consts.UseStr, consts.SessionsStr}),
		Run: func(cmd *cobra.Command, args []string) {
			UseSessionCmd(cmd, con, args)
		},
	}
	carapace.Gen(useSessionCmd).PositionalCompletion(SessionIDCompleter(con))
	useCmd.AddCommand(useSessionCmd)

	useBeaconCmd := &cobra.Command{
		Use:   consts.BeaconsStr,
		Short: "Switch the active beacon",
		Long:  help.GetHelpFor([]string{consts.UseStr, consts.BeaconsStr}),
		Run: func(cmd *cobra.Command, args []string) {
			UseBeaconCmd(cmd, con, args)
		},
	}
	carapace.Gen(useBeaconCmd).PositionalCompletion(beacons.BeaconIDCompleter(con))
	useCmd.AddCommand(useBeaconCmd)

	return []*cobra.Command{useCmd}
}
