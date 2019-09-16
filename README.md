# go-mysql-client
pure Go mysql client

## Feature

- Single binary
- Complete table name
- Complete MySQL reserved words

## Install

```bash
go get -u github.com/johejo/go-mysql-client/cmd/gomysql
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
