package main

import (
    "github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
)

var rootCmd = &cobra.Command{
	Run: rootCmdHandler,
}

var credentialFile string
var credentialProfile string

func init() {
	rootCmd.PersistentFlags().StringVarP(&credentialFile, "credentials", "c", "~/.aws/credentials", "path of the aws credentials file")
	rootCmd.PersistentFlags().StringVarP(&credentialProfile, "profile", "p", "default", "credential profile given inside the credentials file")
    rootCmd.MarkFlagRequired("credentials")
    rootCmd.MarkFlagRequired("profile")
}

func rootCmdHandler(cmd *cobra.Command, args []string) {
    log.WithField("credential file", credentialFile).WithField("profile", credentialProfile).Info("AWS credentials")
	creds := credentials.NewSharedCredentials(credentialFile, credentialProfile)
	_, err := creds.Get()
	if err != nil {
		log.Panic("Error getting credentials information ", err)
	}
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		log.Panic("Error during execution: ", err)
	}
}

