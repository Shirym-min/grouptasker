/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/Shirym-min/grouptasker/internal/config"
	"github.com/Shirym-min/grouptasker/internal/runner"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "grouptasker",
	Short: "Combine multiple commands into a single command and execute them in one go",
	Long: `GroupTasker v1.0.0
Combine multiple commands into a single command and execute them in one go`,

	DisableSuggestions: true,
	CompletionOptions: cobra.CompletionOptions{

		DisableDefaultCmd: true,
	},
	// ★これを追加
	Args: cobra.ArbitraryArgs,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("task name required")
			return
		}

		taskName := args[0]
		taskArgs := args[1:]
		cfg, err := config.Load()
		if err != nil {
			fmt.Println("config error:", err)
			return
		}

		commands, ok := cfg.Tasks[taskName]
		if !ok {
			fmt.Println("task not found:", taskName)
			return
		}

		requiredPlaceholders := placeholderIndices(commands)
		required := maxPlaceholderIndex(commands)

		// If no placeholders are used in any command, do not ask for interactive input.
		if required == 0 {
			processedCommands := make([]string, len(commands))
			for i, command := range commands {
				processedCommands[i] = applyArgs(command, taskArgs)
			}

			if err := runner.RunCommands(processedCommands); err != nil {
				fmt.Println("error:", err)
			}
			return
		}

		if len(taskArgs) == 0 && len(requiredPlaceholders) > 0 {
			reader := bufio.NewReader(os.Stdin)
			for _, index := range requiredPlaceholders {
				promptCommand := placeholderCommandForIndex(commands, index)
				fmt.Printf("Input task %d %s : ", index, promptCommand)
				line, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("error reading input:", err)
					return
				}
				taskArgs = append(taskArgs, strings.TrimSpace(line))
			}
		}
		if len(taskArgs) < required {

			fmt.Printf("error: missing argument %d\n", len(taskArgs)+1)
			fmt.Printf("Please provide all of the required arguments or run the task again with no arguments.\n")
			return
		}
		processedCommands := make([]string, len(commands))
		for i, command := range commands {
			processedCommands[i] = applyArgs(command, taskArgs)
		}

		if err := runner.RunCommands(processedCommands); err != nil {
			fmt.Println("error:", err)
			return
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.grouptasker.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func applyArgs(command string, args []string) string {
	for i, v := range args {
		placeholder := fmt.Sprintf("{{%d}}", i+1)
		command = strings.ReplaceAll(command, placeholder, v)
	}
	return command
}

func maxPlaceholderIndex(commands []string) int {
	re := regexp.MustCompile(`\{\{(\d+)\}\}`)

	maxIndex := 0

	for _, cmd := range commands {
		matches := re.FindAllStringSubmatch(cmd, -1)
		for _, m := range matches {
			if len(m) < 2 {
				continue
			}
			n, err := strconv.Atoi(m[1])
			if err != nil {
				continue
			}
			if n > maxIndex {
				maxIndex = n
			}
		}
	}

	return maxIndex
}

func placeholderIndices(commands []string) []int {
	re := regexp.MustCompile(`\{\{(\d+)\}\}`)
	seen := map[int]struct{}{}

	for _, cmd := range commands {
		matches := re.FindAllStringSubmatch(cmd, -1)
		for _, m := range matches {
			if len(m) < 2 {
				continue
			}
			n, err := strconv.Atoi(m[1])
			if err != nil {
				continue
			}
			seen[n] = struct{}{}
		}
	}

	indices := make([]int, 0, len(seen))
	for n := range seen {
		indices = append(indices, n)
	}
	sort.Ints(indices)
	return indices
}

func placeholderCommandForIndex(commands []string, index int) string {
	placeholder := fmt.Sprintf("{{%d}}", index)
	for _, command := range commands {
		if strings.Contains(command, placeholder) {
			return command
		}
	}
	return ""
}
