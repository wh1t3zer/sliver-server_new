package shell

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
	shellCmd := &cobra.Command{
		Use:         consts.ShellStr,
		Short:       "Start an interactive shell",
		Long:        help.GetHelpFor([]string{consts.ShellStr}),
		GroupID:     consts.ExecutionHelpGroup,
		Annotations: flags.RestrictTargets(consts.SessionCmdsFilter),
		Run: func(cmd *cobra.Command, args []string) {
			ShellCmd(cmd, con, args)
		},
	}
	flags.Bind("", false, shellCmd, func(f *pflag.FlagSet) {
		f.BoolP("no-pty", "y", false, "disable use of pty on macos/linux")
		f.StringP("shell-path", "s", "", "path to shell interpreter")

		f.Int64P("timeout", "t", flags.DefaultTimeout, "grpc timeout in seconds")
	})

	return []*cobra.Command{shellCmd}
}
