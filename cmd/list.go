/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "list config tunnels",
	Run: func(cmd *cobra.Command, args []string) {
		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()
		tbl := table.New("Name", "LocalPort", "Remote", "JumpServer")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		for _, widget := range conf.Tunnel {
			tbl.AddRow(widget.Name, widget.LocalPort, fmt.Sprintf("%s:%d", widget.RemoteHost, widget.RemotePort), widget.JumpServer)
		}

		tbl.Print()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
