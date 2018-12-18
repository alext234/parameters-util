package main

import (
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
)

var rootCmd = &cobra.Command{
	Run: rootCmdHandler,
}

var credentialFile string

func init() {
	rootCmd.PersistentFlags().StringVarP(&credentialFile, "credentials", "c", "~/.aws/credentials", "path of the aws credentials file")
    rootCmd.MarkFlagRequired("credentials")
}

func rootCmdHandler(cmd *cobra.Command, args []string) {
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		log.Panic("Error during execution: ", err)
	}
}

