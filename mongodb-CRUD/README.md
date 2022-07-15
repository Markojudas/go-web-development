<!-- markdownlint-disable -->

# MongoDB

In this module we will go deeper (to an extend) with MongoDB to **Create, Read, Update, Delete (CRUD)**. MongoDB is a **schema-less** database (NoSQL, or nonrelational database).

**this is just a very basic introduction to NoSQL**

<h2>Organization</h2>

Since there aren't any schemas to follow, we are 100% responsible for the organization of the data. A bunch of non-related data can be stored together. The developer needs to impose the schema.

<h2>CAP THEOREM</h2>

**Consistency, Availability, Partitionability**<br>
The theorem says you can have any tow of these three.

**SQL:** Consistent & Available
**NoSQL:** Availability & Partitionability

<h2>Different types of NoSQL</h2>

<ul>
    <li>Key-Value</li>
    <li>Document (MongoDB)</li>
    <li>Column</li>
    <li>Graph</li>
</ul>

<h2>NoSQL Hierarchy</h2>

A NoSQL database has COLLECTIONS which have Documents.

**COLLECTION**<br>
Equivalent to a SQL able

**DOCUMENT**<br>
A collection of key-value pairs

<h2>Advantages over RDBMS</h2>

<ol>
    <li>Documents can vary</li>
    <li>Store objects as they are in your program</li>
    <li>You don't have to do join queries</li>
    <li>Scalability performance</li>
</ol>
<br>

## Make a Database

<pre><code>use [db-name]</code></pre>

**note:** this is also use to "use" an existing database; be careful of typos!
<br>

## Show DBS

<pre><code>show dbs/databases</code></pre>
<br>

## Insert Document

<pre><code>db.[collection-name],insert({"name":"Jose"})</code></pre>

This method is deprecated!

use instead:

<pre><code>db.[collection-name].insertOne()
db.[collection-name].insertMany()
db.[collection-name].bulkWrite()</code></pre>

## view Collections

<pre><code>show collections</code></pre>

Pretty self explanatory :)
<br>

## See the Collection

<pre><code>db.[collection-name].find()
db.[collection-name].find().pretty()</code></pre>
<br>

## Drop DB

<pre><code>db.dropDatabase()</code></pre>

Also, pretty self explanatory! ;)
<br>
<br>

## Create Collections

**implicitly**<br>
<code>db.[collection-name].insertOne("key":"value")</coe>
<br>
<br>

**explicitly**<br>
<code>db.createCollection(\<name\>, {\<optional options\>})</code>
<br>
<br>
**optional options**:
| option | type | description |
| --- | --- | --- |
| capped | bool | caps the size |
| size | number | sets size of cap in bytes |
| max | bool | maximum number of documents allowed in capped collection |

<a href="https://www.mongodb.com/docs/manual/reference/method/db.createCollection/">more details</a>
<br>
<br>

**Examples:**

<pre><code>db.createCollection("customers")</code></pre>

<pre><code>db.createCollection("crs", {capped:true, size:65536, max:1000000})</code></pre>
<br>

## Drop Collection

<pre><code>db.[collection-name].drop()</code></pre>
