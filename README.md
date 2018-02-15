# go-check-database-tables

go-check-database-tables is command line tool for cheking database table

# install

```
go get github.com/hlts2/go-check-database-tables
```

## CLI Usage

```
$ go-check-database-tables --help

Usage:
  go-check-database-tables [flags]
  go-check-database-tables [command]

Available Commands:
  help        Help about any command
  ls          List for database tables

Flags:
  -H, --Host string       Host Name (default "localhost")
  -P, --Password string   Password
  -d, --dbms string       Database Management System (default "mysql")
  -h, --help              help for go-check-database-tables
  -n, --name string       Database Name
  -p, --port int          Port (default 3306)
  -t, --table string      Database Table
  -u, --user string       User Name (default "root")
```

```
$ go-check-database-tables ls --help

Usage:
  go-check-database-tables ls [flags]

Flags:
  -h, --help   help for ls

Global Flags:
  -H, --Host string       Host Name (default "localhost")
  -P, --Password string   Password
  -d, --dbms string       Database Management System (default "mysql")
  -n, --name string       Database Name
  -p, --port int          Port (default 3306)
  -u, --user string       User Name (default "root")

```

### Usage

You can check the table information by executing this command.
This confirms the user table of mysql.

```
$ go-check-database-tables -H 192.168.33.10 -p 3306 -d mysql -u root -P root -n mysql -t user

+------------------------+-----------------------------------+------+-----+---------+-------+
|         FIELD          |               TYPE                | NULL | KEY | DEFAULT | EXTRA |
+------------------------+-----------------------------------+------+-----+---------+-------+
| Host                   | char(60)                          | NO   | PRI |         |       |
| User                   | char(16)                          | NO   | PRI |         |       |
| Password               | char(41)                          | NO   |     |         |       |
| Insert_priv            | enum('N','Y')                     | NO   |     | N       |       |
+------------------------+-----------------------------------+------+-----+---------+-------+
```

You can check the table list by executing this command.
This confirms the tables of mysql. 

```
$ go-check-database-tables ls -H 192.168.33.10 -p 3306 -d mysql -u root -P root -n mysql 
+---------------------------+
|        TABLE NAME         |
+---------------------------+
| columns_priv              |
| db                        |
| event                     |
+---------------------------+
```
