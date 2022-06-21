<!-- markdownlint-disable -->

# Relational Database Management System (RDBMS)

Using the <a href="https://pkg.go.dev/database/sql">database/sql</a> package.

Requirement: The sql package **MUST** be used in conjuction with a database driver. see: <a href="https://github.com/golang/go/wiki/SQLDrivers">SQLDrivers</a>

```
MySQL: https://github.com/go-sql-driver/mysql/ [*]
MySQL: https://github.com/siddontang/go-mysql/ [**] (also handles replication)
MySQL: https://github.com/ziutek/mymysql [*]
```

We are using the first one for this course:

```bash
go get github.com/go-sql-driver/mysql/
```
