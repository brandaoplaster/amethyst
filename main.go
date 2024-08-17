package main

import (
	"fmt"
	"os"

	"github.com/brandaoplaster/amethyst/commands"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "mycli"}

	// Adiciona o comando 'generateCmd' da pasta commands
	rootCmd.AddCommand(commands.GenerateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
