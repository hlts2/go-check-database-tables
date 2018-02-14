package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/hlts2/go-check-database-tables/dao/databases/config"
	"github.com/hlts2/go-check-database-tables/dao/factories"
	"github.com/olekukonko/tablewriter"
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
var dbms string
var database string
var table string

func init() {
	rootCmd.PersistentFlags().StringVarP(&host, "Host", "H", "localhost", "Host Name")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 3306, "Port")
	rootCmd.PersistentFlags().StringVarP(&user, "user", "u", "root", "User Name")
	rootCmd.PersistentFlags().StringVarP(&password, "Password", "P", "", "Password")
	rootCmd.PersistentFlags().StringVarP(&dbms, "dbms", "d", "mysql", "Database Management System")
	rootCmd.PersistentFlags().StringVarP(&database, "name", "n", "", "Database Name")
	rootCmd.Flags().StringVarP(&table, "table", "t", "", "Database Table")
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

	dao := factories.FactoryTableDao(dbms, c)
	if dao == nil {
		return errors.New("Invalid dabase management system")
	}
	describeTables, err := dao.GetDescribeTable(table)
	if err != nil {
		progress.Stop()
		return err
	}

	writer := tablewriter.NewWriter(os.Stdout)
	writer.SetHeader([]string{"Field", "Type", "Null", "Key", "Default", "Extra"})

	for _, describeTable := range describeTables {
		writer.Append(describeTable.FieldValue())
	}
	progress.Stop()
	writer.Render()

	return nil
}

//Execute run command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
