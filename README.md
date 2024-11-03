# Disboard Bumper
A configurable Discord selfbot.

## Usage
Make sure you have `disboard-bumper.toml` in the environment. Example:

```toml
[client-name]
# Discord auth token
token = ""
 
# Channels to bump
channels = ["1302406154223227115"]
```

You can also have multiple selfbots configured:

```toml
[user1]
token = ""
channels = ["1302406154223227115"]

[user2]
token = ""
channels = ["1302578838894673920"]
```

## Installation
build urself

### Building yourself

Clone the repository

```sh
git clone https://github.com/stavratum/disboard-bumper.git
cd disboard-bumper
```

Build using Go

```sh
go build
```

Build using [Gox](https://github.com/mitchellh/gox)

```sh
gox -output="bin/{{.Arch}}/{{.Dir}}-{{.OS}}" -osarch="linux/amd64 windows/amd64"
```
