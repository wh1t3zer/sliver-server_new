package generate

import (
	"context"
	"fmt"

	"github.com/wh1t3zer/sliver-server_new/client/command/settings"
	"github.com/wh1t3zer/sliver-server_new/client/console"
	"github.com/wh1t3zer/sliver-server_new/protobuf/clientpb"
	"github.com/wh1t3zer/sliver-server_new/protobuf/commonpb"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// CanariesCmd - Display canaries from the database and their status.
func CanariesCmd(cmd *cobra.Command, con *console.SliverClient, args []string) {
	canaries, err := con.Rpc.Canaries(context.Background(), &commonpb.Empty{})
	if err != nil {
		con.PrintErrorf("Failed to list canaries %s", err)
		return
	}
	if 0 < len(canaries.Canaries) {
		burnedOnly, _ := cmd.Flags().GetBool("burned")
		PrintCanaries(con, canaries.Canaries, burnedOnly)
	} else {
		con.PrintInfof("No canaries in database\n")
	}
}

// PrintCanaries - Print the canaries tracked by the server.
func PrintCanaries(con *console.SliverClient, canaries []*clientpb.DNSCanary, burnedOnly bool) {
	tw := table.NewWriter()
	tw.SetStyle(settings.GetTableStyle(con))
	tw.AppendHeader(table.Row{
		"Sliver Name",
		"Domain",
		"Triggered",
		"First Trigger",
		"Latest Trigger",
	})
	for _, canary := range canaries {
		if burnedOnly && !canary.Triggered {
			continue
		}
		lineColor := console.Normal
		if canary.Triggered {
			lineColor = console.Bold + console.Red
		}
		firstTrigger := "Never"
		latestTrigger := "Never"
		if canary.Triggered {
			firstTrigger = fmt.Sprintf(lineColor+"%s"+console.Normal, canary.FirstTriggered)
			latestTrigger = fmt.Sprintf(lineColor+"%s"+console.Normal, canary.LatestTrigger)
		}
		row := table.Row{
			fmt.Sprintf(lineColor+"%s"+console.Normal, canary.ImplantName),
			fmt.Sprintf(lineColor+"%s"+console.Normal, canary.Domain),
			fmt.Sprintf(lineColor+"%v"+console.Normal, canary.Triggered),
			firstTrigger,
			latestTrigger,
		}
		tw.AppendRow(row)
	}
	con.Printf("%s\n", tw.Render())
}
