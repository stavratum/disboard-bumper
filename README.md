# Disboard Bumper
A configurable Disboard client. All it requires is your browser cookies for `disboard.org` and internet connection to bump.

## Usage
Make sure you have `disboard.toml` or `bumper.toml` in the environment. Example:

```toml
[client-name]
# Disboard cookies
cookies = ""
 
# Servers to bump
servers = ["1280072122097602581"]
```

You can also have multiple clients configured:

```toml
[client1]
cookies = ""
servers = ["1280072122097602581"]

[client2]
cookies = ""
servers = ["1280703005703929896"]
```

## Installation
build urself

### Building yourself

Clone the repository

```sh
git clone https://github.com/stavratum/bumper.git
cd bumper
```

Build using Go

```sh
go build -o=bin
```

Build using [Gox](https://github.com/mitchellh/gox)

```sh
gox -output="bin/{{.Arch}}/{{.Dir}}-{{.OS}}" -osarch="linux/amd64 windows/amd64"
```
