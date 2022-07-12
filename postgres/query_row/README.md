<!-- markdownlint-disable -->

# QueryRow

With `QueryRow` we get one record so we don't have to loop!

<pre><code>func (db *DB) QueryRow(query string, args ...interface{}) *Row</code></pre>
<br>

<h2>type Row</h2>

<pre><code>func (r *Row) Scan(dest ...interface{}) error</code></pre>
<br>

<h2>cURL request</h2>

<pre><code>curl -i localhost:8080/books/show?isbn=978-1505255607</code></pre>
