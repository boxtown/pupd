package cmd

import "github.com/spf13/cobra"

var apiUser string
var apiPass string

func init() {
	rootCmd.PersistentFlags().StringVar(&apiUser, "api-user", "pupd_api", "Username for API user role")
	rootCmd.PersistentFlags().StringVar(&apiPass, "api-pass", "changeme", "Password for API user role")
}

var rootCmd = &cobra.Command{
	Use:   "pupd",
	Short: "Pick Up Put Down Backend",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute the root command
func Execute() error {
	return rootCmd.Execute()
}
