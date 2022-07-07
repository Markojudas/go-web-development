<!-- markdownlint-disable -->

# MongoDB

<h2>Install MongoDB</h2>
<a href="https://markojudas.github.io/chronicles/blog/MongoDB-Fedora.html">Read this little write-up on how to install MongoDB on Fedora</a>

<br>
<hr>
<br>

<h2>Driver for Mongo</h2>

```bash
go install gopkg.in/mgo.v2
go install gopkg.in/mgo.v2/bson
```

<br>
<hr>
<br>

<h2>CLI commands once server is running</h2>

**Create user and add to database/ POST:**<br>

<code>curl -X POST -H "Content-Type: application/json" -d '{"name":"Miss Moneypenny","gender":"female","age":27}' http://localhost:8080/user</code>
<br>
<br>
**See user / GET:**<br>

<code>curl http://localhost:8080/user/[enter-user-id-here]</code>
<br>
<br>
**Delete User:**<br>
<code>curl -X DELETE http://localhost:8080/user/[enter-user-id-here]</code>
