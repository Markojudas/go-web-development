# Redirects

<!-- markdownlint-disable -->

On the web, we have a client/server architecture. Clients make requests, and servers write responses to thos clients. The request and respons are both just text that much conform to the rules of HTTP. Both the request & response have a start line, headers, and a body.

The request start line is called the "request line". it consists of:

`request-line = method SP request-target SP HTTP-version CRLF`

The response start line is called the "status line". it consists of:

`status-line = HTTP-version SP status-code SP reason-phrase CRLF`

You can see that the status line wants a status-code and a reason-phrase. For redirects we have several status-codes and reason-phrases from which we can choose, but primarily you should choose either

<ul>
    <li> 301 moved permantely</li>
    <li> 303 see other &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;//changes method to GET </li>
    <li> 307 temporary redirect &emsp;&emsp;//keeps same method</li>
</ul>
