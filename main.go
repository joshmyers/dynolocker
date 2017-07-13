package main

import (
	"github.com/urfave/cli"
	"github.com/zencoder/ddbsync"
	"os"
	"sort"
	"time"
)

const (
	DB_LOCK_RETRY time.Duration = 10 * time.Second
)

func lock(c *cli.Context) {
	s := ddbsync.NewLockService(c.GlobalString("table_name"), c.GlobalString("region"), "", c.GlobalBool("disable_ssl"))
	m := s.NewLock(c.GlobalString("lock_name"), c.GlobalInt64("lock_ttl"), DB_LOCK_RETRY)
	m.Lock()
}

func unlock(c *cli.Context) {
	s := ddbsync.NewLockService(c.GlobalString("table_name"), c.GlobalString("region"), "", c.GlobalBool("disable_ssl"))
	m := s.NewLock(c.GlobalString("lock_name"), c.GlobalInt64("lock_ttl"), DB_LOCK_RETRY)
	m.Unlock()
}

func main() {
	app := cli.NewApp()
	app.Name = "dynolocker"
	app.Usage = "distributed locking using DynamoDB"
	app.Version = "0.0.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "GOV.UK Pay Operations",
			Email: "gds-team-payments-web-ops@digital.cabinet-office.gov.uk",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "table_name",
			Value:  "dynolocker",
			Usage:  "DynamoDB table for locks",
			EnvVar: "DB_TABLE_NAME",
		},
		cli.StringFlag{
			Name:   "lock_name",
			Value:  "lock",
			Usage:  "DynamoDB lock name",
			EnvVar: "DB_LOCK_NAME",
		},
		cli.StringFlag{
			Name:   "region",
			Value:  "eu-west-1",
			Usage:  "AWS region",
			EnvVar: "AWS_DEFAULT_REGION",
		},
		cli.Int64Flag{
			Name:   "lock_ttl",
			Value:  60,
			Usage:  "Lock duration",
			EnvVar: "DB_TTL",
		},
		cli.BoolFlag{
			Name:  "disable_ssl",
			Usage: "Disable SSL on calls to AWS (default: false)",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "lock",
			Usage:  "Create a lock",
			Action: lock,
		},
		{
			Name:   "unlock",
			Usage:  "Force an unlock",
			Action: unlock,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}
