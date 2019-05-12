package main

import (
	"github.com/joshmyers/dynolocker/dynamodb"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/zencoder/ddbsync"
	"os"
	"sort"
	"sync"
	"time"
)

func lock(c *cli.Context, m sync.Locker) {
	log.WithFields(log.Fields{"name": c.GlobalString("name"), "table": c.GlobalString("table")}).Debug("Locking table...")
	m.Lock()
}

func unlock(c *cli.Context, m sync.Locker) {
	log.WithFields(log.Fields{"name": c.GlobalString("name"), "table": c.GlobalString("table")}).Debug("Unlocking table...")
	m.Unlock()
}

func newDynolocker(c *cli.Context) sync.Locker {
	s := ddbsync.NewLockService(c.GlobalString("table"), c.GlobalString("region"), "", c.GlobalBool("disable_ssl"))
	return s.NewLock(c.GlobalString("name"), c.GlobalInt64("ttl"), c.GlobalDuration("retry"))
}

func main() {
	app := cli.NewApp()
	app.Name = "dynolocker"
	app.Usage = "distributed locking using DynamoDB"
	app.Version = "0.0.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Joshua Myers",
			Email: "joshuajmyers@gmail.com",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "table",
			Value:  "dynolocker",
			Usage:  "DynamoDB table for locks",
			EnvVar: "DB_TABLE_NAME",
		},
		cli.StringFlag{
			Name:   "name",
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
		cli.DurationFlag{
			Name:  "retry",
			Value: time.Second * 3,
			Usage: "Lock reattempt wait duration",
		},
		cli.BoolTFlag{
			Name:  "disable-ssl",
			Usage: "Disable SSL on calls to AWS (default: false)",
		},
		cli.Int64Flag{
			Name:   "ttl",
			Value:  60,
			Usage:  "Lock duration",
			EnvVar: "DB_TTL",
		},
		cli.BoolTFlag{
			Name:  "debug",
			Usage: "Show debug output",
		},
	}
	app.Before = func(c *cli.Context) error {
		if c.GlobalBoolT("debug") {
			log.SetLevel(log.DebugLevel)
		}
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:  "lock",
			Usage: "Create a lock",
			Action: func(c *cli.Context) error {
				if c.BoolT("create-table") {
					dynamodb.CreateLockTableIfNecessary(c.GlobalString("table"), c.GlobalString("region"))
				}
				m := newDynolocker(c)
				lock(c, m)
				return nil
			},
			Flags: []cli.Flag{
				cli.BoolTFlag{
					Name:  "create-table",
					Usage: "If we should create the DynamoDB table (default: true)",
				},
			},
		},
		{
			Name:  "unlock",
			Usage: "Force an unlock",
			Action: func(c *cli.Context) error {
				m := newDynolocker(c)
				unlock(c, m)
				return nil
			},
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)
}
