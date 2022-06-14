# URL values

<!-- markdownlint-disable -->

You can always append values to a URL

Anything after the "?" is the query string - the area where values are stored

<h2>Breakdown of a URL</h2>
https://video.google.co.uk:80/videoplay?docid=-7246927612831078230&hl=en#00h02m30s
<br>
<br>
<strong>protocol:</strong>      https://<br>
<strong>subdomain:</strong>     video.</br>
<strong>domain name:</strong>   google.co.uk<br>
<strong>port:</strong>          :80<br>
<strong>path:</strong>          /videoplay<br>
<strong>query:</strong>         ?<br>
<strong>parameters:</strong>    docid=-7246927612831078230&hl=en<br>
<strong>fragment:</strong>      #00h02m30s<br>
<br>
The vales are stored in a <em>identifier=value</em> fashion.<br>
You can have multiple <em>idenfifier=value</em> by separating them with the "&" ampersand.
<br>
<hr>
<h2>Retriving values</h2>
While there are multiple ways to retrieve values, we will stick with:

func (\*Request) FormValue

```go
func (r *Request) FormValue(key string) string
```
