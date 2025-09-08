package licenses

import (
	"github.com/spf13/cobra"

	"github.com/wh1t3zer/sliver-server_new/client/command/help"
	"github.com/wh1t3zer/sliver-server_new/client/console"
	consts "github.com/wh1t3zer/sliver-server_new/client/constants"
	"github.com/wh1t3zer/sliver-server_new/client/licenses"
)

// Commands returns the `licences` command.
func Commands(con *console.SliverClient) []*cobra.Command {
	licensesCmd := &cobra.Command{
		Use:   consts.LicensesStr,
		Short: "Open source licenses",
		Long:  help.GetHelpFor([]string{consts.LicensesStr}),
		Run: func(cmd *cobra.Command, args []string) {
			con.Println(licenses.All)
		},
		GroupID: consts.GenericHelpGroup,
	}

	return []*cobra.Command{licensesCmd}
}
