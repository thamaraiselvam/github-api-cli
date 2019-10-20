# git-api-cli

gac is cli tool which fetches publicly available data from github.com

## Usage

### Fetch user information

```
gac info <username>
```

### Todo Commands:

```
- gac list pr <username>
- gac list followers <username>
- gac list following <username>
- gac list gist <username>
- gac list repo <username>
```

### Development

Run

```sh
$ git clone https://github.com/thamaraiselvam/git-api-cli.git
$ cd git-api-cli
$ go install
$ make compile # only build
$ make build # lint, test, compile
```
