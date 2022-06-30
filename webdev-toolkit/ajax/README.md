<!-- markdownlint-disable -->

# Asynchronous JavaScript and XML (AJAX)

<a href="https://developer.mozilla.org/en-US/docs/Web/Guide/AJAX">Documentation</a><br>

Ajax is not a techonology but uses a number of exisiting technologies together, including `HTML`, `XHTML`, `CSS`, `JavaScript`, `DOM`, `XML`, `XSLT`, and most importantly the `XMLHttpRequest` object. When these technologies are combined in the Ajax model, web applicatons are able to make quick, incremental updates to the user interface without reloading the entire browser page. This makes the application faster and more responsive to user actions.

JSON and XML are used for packaging information in the Ajax model.</br>
<br>

<hr>
<h1>XMLHttpREquest</h1>

<a href="https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest">Documentation</a><br>

`XMLHttpRequest` (XHR) objects are used to interact with servers. You can retrieve data from a URL without having to do a full page refresh. This enables a Web page to update just part of a page without disrupting what the user is doing.

`XMLHttpRequest` is used heavily in AJAX programming. Despite its name, it can be used to retrieve any type of data, just XML.

<h2>Constructor</h2>

`XMLHttpRequest()`

The constructor initializes an XMLHttpRequest. It must be called before any other method calls.

**Syntax**

```JavaScript
new XMLHttpReqest()
```

**Return value**

A new `XMLHttpRequst` object. the object must be prepared by at least calling `open()` to initialize it before calling `send()` to send the request to the server.

<h2>Open()</h2>

The `XMLHttpRequest` method `open()` initializes a newly-created request, or re-initializes an existing one (equivalent of calling `abort()`)

**Syntax**

```JavaScript
open(method, url)
open(method, url, async)
open(method, url, async, user)
open(method, url, async, user, password)
```

**Params**

`method`<br>
The HTTP request method to use, such as `"GET"`, `"POST"`, `"PUT"`, `"DELETE"`, etc... Ignored for non-HTTP(S) URLs.<br>
<br>
`url`<br>
A string representing the URL to send the request to<br>
<br>
`async` (optional)<br>
An optional Boolean parameter, defaulting to `true`, indicating whether or not to perform the operation asynchronously. If this value is `false`, the `send()` method does not return until the response is received. If `true`, notification of a completed transaction is provided using event listeners. This must be true if the `multipart` attribute is `true`, or an exception will be thrown.<br>
<br>
`user` (optional)<br>
The optional user name to use for authentication purposes; by default, this is the `null` value.<br>
<br>
`password` (optional)<br>
The optional user name to use for authentication purposes; by default, this is the `null` value.<br>
<br>
