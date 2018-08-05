# Cron expression parser
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

<p align="center">
  <img src="assets/583e8cae50ccd12d000083e1.png">
</p>

## Presumptions

* Cron expression consists of only numbers. `mon` | `january` | `@monthly` are not supported
* `*` is passed as an argument backslashed `\*`
* The next argument after the expression is taken as the `command`
* Rules are tested according this source https://crontab.guru

## Usage

### Compile

Have [go](https://golang.org/doc/install) workspace [setup](https://www.ardanlabs.com/blog/2016/05/installing-go-and-your-workspace.html) on your machine

```bash
cd $GOPATH/src
git clone git@github.com:apzuk/cron.git cron
cd cron
go build -o cron cmd/main.go
./cron args...
```

### Docker

Alternatively, if you have docker on your machine, run it on Docker

```bash
docker run apzuk/cron args...
```

## A few examples

```
docker run apzuk/cron -f --debug \* \* \* \* \*  

docker run apzuk/cron -f --debug */4 \* 8-14 \* \*

docker run apzuk/cron -f --debug */4 \* 8-14 3-7/2 \* /usr/bin/find
```

For more examples see cron_test.go