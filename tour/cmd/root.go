package cmd

import "github.com/spf13/cobra"

/**
 *@Author tudou
 *@Date 2020/7/22
 **/

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(sqlStructCmd)
	rootCmd.AddCommand(sqlCmd)
	rootCmd.AddCommand(jsonCmd)
}
