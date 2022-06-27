<!-- markdownlint-disable -->

# Keyed-Hash Message Authentication Code(HMAC)

<a href="https://en.wikipedia.org/wiki/HMAC" target="_blank" rel="noopener noreferrer">Wiki: HMAC</a>

In cryptography, an HMAC (sometimes expanded as either keyed-hash message authentication code or hash-based message authentication code) is a specific type of message authentication code (MAC) involving a cryptographic hash function and a secret cryptographic key. As with any MAC, it may be used to simultaneously verify both the data integrity and authenticity of a message.

HMAC can provide authentication using a shared secret instead of using digital signatures with asymmetric cryptography. It trades off the need for a complex public key infrastructure by delegating the key exchange to the communicating parties, who are responsible for establishing and using a trusted channel to agree on the key prior to communication.

<h2>Packages:</h2>

To implement HMAC on Go, use the package(s):<br>

<code>golang.org/x/cyrpto</code><br>
<code>golang.org/x/cyrpto/hmac</code><br>
<code>golang.org/x/cyrpto/sha256</code><br>
