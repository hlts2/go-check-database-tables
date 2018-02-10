package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/olekukonko/tablewriter"

	"github.com/hlts2/go-check-database-tables/dao/databases/config"
	"github.com/hlts2/go-check-database-tables/dao/factories"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-check-database-tables",
	Short: "A CLI Tool for checking database tables",
	Run: func(cmd *cobra.Command, args []string) {
		if err := root(cmd, args); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var host string
var port int
var user string
var password string
var database string

func init() {
	rootCmd.PersistentFlags().StringVarP(&host, "Host", "H", "localhost", "Host Name（localhost）")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 3306, "Port（3306）")
	rootCmd.PersistentFlags().StringVarP(&user, "user", "u", "root", "User Name（root）")
	rootCmd.PersistentFlags().StringVarP(&password, "Password", "P", "", "Password")
	rootCmd.PersistentFlags().StringVarP(&database, "database", "d", "", "Database Name")
}

func root(cmd *cobra.Command, args []string) error {
	c := config.Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Database: database,
	}

	progress := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	progress.Start()

	dao := factories.FactoryTableDao("mysql", c)
	tables, err := dao.GetTables()
	if err != nil {
		progress.Stop()
		return err
	}

	writer := tablewriter.NewWriter(os.Stdout)
	writer.SetHeader([]string{"Table Name"})

	for _, table := range tables {
		writer.Append([]string{table.Name})
	}

	progress.Stop()
	writer.Render()

	return nil
}

//Execute run command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
