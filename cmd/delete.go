/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/Shirym-min/grouptasker/internal/config"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"os"
	"sort"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <task>",
	Short: "Delete a registered command",
	Long:  `Delete a registered command from the configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		const (
			Reset  = "\033[0m"
			Red    = "\033[31m"
			Green  = "\033[32m"
			Yellow = "\033[33m"
			Blue   = "\033[34m"
		)
		cfg, err := config.Load()
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(args) == 0 {
			fmt.Printf(Red + "Error: Task name required. Please specify a task to delete.\n" + Reset)
			names := make([]string, 0, len(cfg.Tasks))
			for name := range cfg.Tasks {
				names = append(names, name)
			}
			sort.Strings(names)

			fmt.Printf(Green + "--- Available Tasks ---" + Reset + "\n")
			for _, name := range names {
				fmt.Printf(Green+"%s:"+Reset+"\n", name)
				for _, command := range cfg.Tasks[name] {
					fmt.Printf(Blue+"  - %s"+Reset+"\n", command)
				}
			}
			return
		}
		taskName := args[0]

		if _, ok := cfg.Tasks[taskName]; !ok {
			fmt.Printf("Task \"%s\" not found\n", taskName)
			return
		}
		if !Confirmdelete(fmt.Sprintf("Are you sure you want to delete the task \"%s\"? [y/N]: ", taskName)) {
			fmt.Println("Task deletion cancelled.")
			return
		}
		if err := cfg.DeleteTask(taskName); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println()
		fmt.Printf("Task \"%s\" deleted successfully\n", taskName)
	},
}

func init() {
	configCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func Confirmdelete(pronpt string) bool {
	const (
		Reset  = "\033[0m"
		Red    = "\033[31m"
		Green  = "\033[32m"
		Yellow = "\033[33m"
		Blue   = "\033[34m"
	)
	fmt.Print(Red + pronpt + Reset)
	oldstate, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return false
	}
	defer term.Restore(int(os.Stdin.Fd()), oldstate)

	var b [1]byte

	for {
		_, err := os.Stdin.Read(b[:])
		if err != nil {
			return false
		}
		switch b[0] {
		case 'y', 'Y':
			fmt.Println("y")
			return true
		case 'n', 'N':
			fmt.Println("n")
			return false
		case '\r', '\n':
			fmt.Println()
			return false
		}
	}
}
