<!-- markdownlint-disable -->

# QUERY: SELECT

Using the following functions:

From database/sql package:

<h2>Connect to the database</h2>

<code>func Open(driverName, dataSourceName string) (\*DB, error)</code><br>

This will return a pointer to a database needed to perform queries:

<h2>Query -> pointer to rows</h2>

<code>func (db \*DB) Query(query string, args ...interface{}) ( \*Rows, error)</code><br>

query == the SQL statement

<h5><strong>not passing any arguments</strong></h5>

<code>func (rs \*Rows) Close() error</code><br>

<h2>Next() & Scan()</h2>

A loop iterating through the ROW returned by the QUERY and SCAN into a variable the desired piece of data.<br>
<code>func (rs \*Rows) Next() bool</code><br>
<code>func (rs \*Rows) Scan(dest ...interface{}) error</code>
