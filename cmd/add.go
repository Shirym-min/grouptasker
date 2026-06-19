/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/Shirym-min/grouptasker/internal/config"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"os"
	"strings"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new tasks",
	Long:  `Add new tasks. You can specify a task name and the commands to be executed for that task.`,
	Run: func(cmd *cobra.Command, args []string) {
		const (
			Reset  = "\033[0m"
			Red    = "\033[31m"
			Green  = "\033[32m"
			Yellow = "\033[33m"
			Blue   = "\033[34m"
		)
		reservedNames := []string{"list", "config", "add", "remove", "help", "version", "root"}
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Task name: ")
		taskName, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		taskName = strings.TrimSpace(taskName)
		for _, name := range reservedNames {
			if taskName == name {
				fmt.Println(Red + "Error: '" + taskName + "' is a reserved name. You cannot use this name. Please try another word." + Reset)
				return
			}
		}
		// gpx.ymlにすでに存在している名前の場合、エラーにする
		cfg, err := config.Load()
		for name := range cfg.Tasks {
			if taskName == name {
				fmt.Println(Red + "Error: '" + taskName + "' already exists. Please try another word." + Reset)
				return
			}
		}
		fmt.Println(Green + "Creating Task: " + taskName + Reset)
		fmt.Println(Yellow + "Enter commands for the task (just press Enter on an empty line to finish):" + Reset)
		fmt.Println(Yellow + "You can use placeholders like {{1}}, {{2}}, etc. for arguments. You can use like : " + Blue + `gpx [your command] "[argument1]" "[argument2]"` + Reset)
		var commands []string
		for {
			fmt.Printf(Blue+"Command #%d :"+Reset, len(commands)+1)
			line, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}
			line = strings.TrimSpace(line)
			if line == "" {
				break
			}
			commands = append(commands, line)
		}
		fmt.Println()
		fmt.Println("Task:", taskName)

		for _, c := range commands {
			fmt.Println(" -", c)
		}

		if !Confirm("Are you sure you want to add this task? [Y/n]: ") {
			fmt.Println("Task creation cancelled.")
			return
		}

		cfg, err = config.Load()
		if err != nil {
			fmt.Println(err)
			return
		}
		if cfg.Tasks == nil {
			cfg.Tasks = make(map[string][]string)
		}
		cfg.Tasks[taskName] = commands
		if err := config.Save(cfg); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println()
		fmt.Println(Green + "Task saved successfully!" + Reset)
	},
}

func init() {
	configCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Confirm(pronpt string) bool {
	const (
		Reset  = "\033[0m"
		Red    = "\033[31m"
		Green  = "\033[32m"
		Yellow = "\033[33m"
		Blue   = "\033[34m"
	)
	fmt.Print(Green + pronpt + Reset)
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
			return true
		}
	}
}
