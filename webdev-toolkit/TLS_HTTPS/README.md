<!-- markdownlint-disable -->

# TLS & HTTPS

`func ListenAndServeTLS`

```go
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
```

ListenAndServeTLS acts identically to ListenAndServe, except that it expects HTTPS connections.
Additionally, files containing a certificate and matching private key for the server must be provided.
If the certificate is signed by a certificate authority, the certFile should be the concatenation of
the server's certificate, any intermediates, and the CA's certificate.

<hr>
<h2>To Create a Cert & Key file</h2>

```bash
go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=somedomainname.com
```

<hr>
<h2>autocert</h2>

<a href="https://pkg.go.dev/golang.org/x/crypto/acme/autocert">golang.org/x/crypto/acme/autocert</a>

A work in progress package that provides automatic access to certificates from Let's Encrypt and any other ACME-based CA.
No examples provided. :(
