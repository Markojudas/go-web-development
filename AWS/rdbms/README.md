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

<hr>
<h2>MYSQL</h2>

I am using the mysql CLI.

**INSTALL**

```bash
sudo dnf install mysql-server mysql
```

**START THE SERVICE**

```bash
sudo systemctl start mariadb
```

**Connect to AWS RDS**

```bash
mysql -h [endpoint] -P 3306 -u [masteruser] -p
```

When prompted for the password enter the master password for the RDS

**note:** pay mind at the security group of the database; in this example the SG is `webtier` and it will only accept connections from the instance with the same SG.
