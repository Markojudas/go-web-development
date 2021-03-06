<!-- markdownlint-disable -->

# Insert Records

Using:

<h2>type DB</h2>
<pre><code>func Open(driverName, dataSourceName string) (*DB, error)</code></pre>
<br>
<h2>Query & QueryRow</h2>
<pre><code>func (db *DB) Query(query string, args ...interface{}) (*Rows, error)</code></pre>
<pre><code>func (db *DB) QueryRow(query string, args ...interface{}) *Row</code></pre>
<br>

<h2>Exec</h2>
Using this to execute a SQL statement!
<br>
<pre><code>func (db \*DB) Exec(query string, args ...interface{}) (Result, error)</code></pre>

<h2>type Result</h2>
<pre>type Result interface {
    //LastInsertID returns the integer generated by the database
    //in response to a command. Typically this will be frm an
    //"auto increment" column when inserting a new row. Not all
    //databases support this feature, and the syntax of such
    //statements varies
    LastInsertId() (int64, error)</br>
    //RowsAffected returns the number of rows affected by an
    //update, insert, or delete. Not every database or database
    //drier may support this.
    RowsAffected() (int64, error)</br>
}</pre>
