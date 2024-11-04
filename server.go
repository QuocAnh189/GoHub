package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
    var rootCmd = &cobra.Command{
        Use:   "app",
        Short: "A brief description of your application",
    }

    var serverCmd = &cobra.Command{
        Use:   "server",
        Short: "Starts the server",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Server is running...")
        },
    }

    rootCmd.AddCommand(serverCmd)

	if len(os.Args) == 1 {
        fmt.Println("No command provided, starting server by default...")
        serverCmd.Run(serverCmd, nil) // Chạy lệnh server
    } else {
        if err := rootCmd.Execute(); err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    }
}
