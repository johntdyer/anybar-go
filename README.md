# anybar-go
Go CLI for Anybar ( https://github.com/tonsky/AnyBar )


## Building

./make.sh

## Development
go get github.com/tools/godep
godep restore
godep run anybar-go


## Use

```shell
NAME:
   anybar-go - Anybar CLI

USAGE:
   anybar-go [global options] command [command options]

VERSION:
   0.0.1

AUTHOR:
  John Dyer - <johntdyer@gmail.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --port, -p '1738'    Port to connect to anybar [$ANYBAR_PORT]
   --address, -a 'localhost'  Address to send message.
   --msg, -m      Message to send to anybar
   --help, -h     show help
   --version, -v
```


### Sending message

#### Default

    ./anybar-go.osx --msg green
