# Session

<!-- markdownlint-disable -->

Store a unique ID in the cookie. On our server, we will associate each user with a unique ID. This wil allow us to identify every user visiting our website.

<hr>
<h2>Security</h2>

There are two factors which contrbute to the security of a session created using a cookie and a unique ID:

<ol>
    <li>Uniqueness of the ID</li>
    <li>Encryption in transit with HTTPS</li>
</ol>

You can use any unique ID that you would like: a Universall Unique Identifier (UUID) or database key. If you're using a database key, make sure it's not the key to your user but to a separate session table.

UUID cannot be intercepted in transit if we are using HTTPS
b

<hr>
<h2>UUID package</h2>

Using the <a href="github.com/satori/go.uuid">satori/go.uuid</a> package to generate UUID.
