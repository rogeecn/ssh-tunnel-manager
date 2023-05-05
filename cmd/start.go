/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/rogeecn/stm/config"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a ssh tunnel",
	Run: func(cmd *cobra.Command, args []string) {
		name := strings.ToLower(args[0])

		var tunnel config.Tunnel
		for _, t := range conf.Tunnel {
			if strings.ToLower(t.Name) == name {
				tunnel = t
			}
		}

		tunnelStr := fmt.Sprintf("%d:%s:%d", tunnel.LocalPort, tunnel.RemoteHost, tunnel.RemotePort)
		cmder := exec.CommandContext(context.Background(), "ssh", "-N", "-L", tunnelStr, "-f", tunnel.JumpServer)
		b, err := cmder.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("OUTPUT: \n\n%s\n\n", b)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
