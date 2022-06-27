<!-- markdownlint-disable -->

# BASE64 encoding

They are only certain values that we can store in a cookie, per RFC 2109, 2965, 6265 2616. You can pick from alphanumeric plus `` !#$%&'*+-.^_`|~ ``

<h2>Package</h2>

`encoding/base64` from the standard library

<h2>Standard Encode</h2>

<code>s64 := base64.StdEncoding.EncodetoString([]byte(s))</code><br>

You can either use the standard encode or define your own:

**example:**

<code>encodeStd := "ABCDEFGHIJKLMNOPQRStUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"</code><br>

<code>s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))</code><br>
where `s` is the string you wish to encode.
