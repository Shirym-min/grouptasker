/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Shirym-min/grouptasker/internal/config"
	"github.com/spf13/cobra"
)
import "sort"

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Displays a list of registered commands",
	Long:  `Displays a list of registered commands`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			fmt.Println(err)
			return
		}

		names := make([]string, 0, len(cfg.Tasks))
		for name := range cfg.Tasks {
			names = append(names, name)
		}
		sort.Strings(names)

		fmt.Println("--- Available Tasks ---")
		for _, name := range names {
			fmt.Println(name + ":")
			for _, command := range cfg.Tasks[name] {
				fmt.Println("  - " + command)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
