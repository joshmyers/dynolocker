## Dynolocker

A CLI tool for distributed locks using DynamoDB:

```
NAME:
   dynolocker - distributed locking using DynamoDB

USAGE:
   dynolocker [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR:
   Joshua Myers <joshuajmyers@gmail.com>

COMMANDS:
     lock     Create a lock
     unlock   Force an unlock
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug         Show debug output
   --disable-ssl   Disable SSL on calls to AWS (default: false)
   --name value    DynamoDB lock name (default: "lock") [$DB_LOCK_NAME]
   --region value  AWS region (default: "eu-west-1") [$AWS_DEFAULT_REGION]
   --retry value   Lock reattempt wait duration (default: 3s)
   --table value   DynamoDB table for locks (default: "dynolocker") [$DB_TABLE_NAME]
   --ttl value     Lock duration (default: 60) [$DB_TTL]
   --help, -h      show help
   --version, -v   print the version
```
