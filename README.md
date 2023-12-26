# go-mysql-client
pure Go mysql client

## Feature

- Single binary
- Complete table name
- Complete MySQL reserved words

## Demo

![](https://raw.githubusercontent.com/johejo/go-mysql-client/master/demo.gif)

## Install

```bash
# for someone who is newbe to go
git clone git@github.com:johejo/go-mysql-client.git
cd go-mysql-client
go get -u github.com/johejo/go-mysql-client/cmd/gomysql
cd /cmd/gomysql
go build
```

## Usage

execute SQL
```bash
gomysql -h [HOST] -P [PORT] -u [USER] -p [PASSWORD] [DATABASE] "[SQL]" 
```

prompt
```bash
gomysql -h [HOST] -P [PORT] -u [USER] -p [PASSWORD] [DATABASE] 
```

reading file
```bash
gomysql -h [HOST] -P [PORT] -u [USER] -p [PASSWORD] -f [SQL_FILE] [DATABASE] 
```
