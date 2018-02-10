package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-check-database-tables",
	Short: "A CLI Tool for checking database tables",
}

var host string
var port int
var user string
var password string
var databaseType string
var database string

func init() {
	rootCmd.PersistentFlags().StringVarP(&host, "Host", "H", "localhost", "Host Name（localhost）")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 3306, "Port（3306）")
	rootCmd.PersistentFlags().StringVarP(&user, "user", "u", "root", "User Name（root）")
	rootCmd.PersistentFlags().StringVarP(&password, "Password", "P", "", "Password")
	rootCmd.PersistentFlags().StringVarP(&databaseType, "type", "t", "mysql", "Database Type（mysql）")
	rootCmd.PersistentFlags().StringVarP(&database, "database", "d", "", "Database Name")
}

//Execute run command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
