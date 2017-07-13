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
   Josh Myers <josh@joshmyers.io>

COMMANDS:
     lock     Create a lock
     unlock   Force an unlock
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --create_table      If we should create the DynamoDB table (default: true)
   --disable_ssl       Disable SSL on calls to AWS
   --lock_name value   DynamoDB lock name (default: "lock") [$DB_LOCK_NAME]
   --lock_ttl value    Lock duration (default: 60) [$DB_TTL]
   --region value      AWS region (default: "eu-west-1") [$AWS_DEFAULT_REGION]
   --table_name value  DynamoDB table for locks (default: "dynolocker") [$DB_TABLE_NAME]
   --help, -h          show help
   --version, -v       print the version
```
