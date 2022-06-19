<!-- markdownlint-disable -->

# Expiring Sessions

<h2><strong>Step 1:</strong></h2>
Change <code>dbSessions</code>. Now the value is TYPE session. Creating a new type session:

```go
type session struct{
    username string
    lastActivity time.Time
}
```

```go
var dbSessions = map[string]session{}
```

<h2><strong>Step 2:</strong></h2>

Creating a new variable `dbSessionsCleaned`

Clean up the dbSessions on logout

```go
if time.Since(dbSessionsCleaned) > (time.Second * 30){
    go cleanSessions()
}
```

**note:**

```go
time.Since()
```

is a shorthand for:

```go
time.Now().Sub(t)
```
