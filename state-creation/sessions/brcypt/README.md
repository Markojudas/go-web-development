# Passowrd Encryption

<!-- markdownlint-disable -->

Using the package: <a href="golang.org/x/crypto/bcrypt">x/crypto/bcrypt</a>. We are changing:

<ol>
    <li>Changing the user struct's password field to that data type []byte</li>
    <li>Use bcrypt to encrypt the password before storing it</li> 
</ol>
