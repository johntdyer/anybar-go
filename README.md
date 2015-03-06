# anybar-go

Small go app to send commands to [Anybar](https://github.com/tonsky/AnyBar).


## Building

* Build artifact `./make.sh`
* Copy to bin `cp builds/anybar-go.osx ~/bin/anybar-go`
* Profit


## Development
go get github.com/tools/godep
godep restore
godep run anybar-go


## Use

```shell
NAME:
   anybar-go - Anybar CLI

USAGE:
   anybar-go [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR:
  John Dyer - <johntdyer@gmail.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --port, -p '1738'    Port to connect to anybar [$ANYBAR_PORT]
   --address, -a 'localhost'  Address to send message.
   --help, -h     show help
   --version, -v    print the version

```


### Sending message

#### Defaults

    ./anybar-go.osx green

#### Setting port

    ./anybar-go.osx green -p 2000



