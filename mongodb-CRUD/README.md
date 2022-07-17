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

# Make a Database

<pre><code>use [db-name]</code></pre>

**note:** this is also use to "use" an existing database; be careful of typos!
<br>

# Show DBS

<pre><code>show dbs/databases</code></pre>
<br>

# Insert Document

<pre><code>db.[collection-name],insert({"name":"Jose"})</code></pre>

This method is deprecated!

use instead:

<pre><code>db.[collection-name].insertOne()
db.[collection-name].insertMany()
db.[collection-name].bulkWrite()</code></pre>

<h2>Example</h2>

<pre><code>db.customer.insertOne({"key":"value"})
db.customer.insertMany(
    [
        {"key":"value"},
        {"key":"value"},
        ...
        {"key":"value}
    ]
)</code></pre>

# view Collections

<pre><code>show collections</code></pre>

Pretty self explanatory :)
<br>

# See the Collection

<pre><code>db.[collection-name].find()
db.[collection-name].find().pretty()</code></pre>
<br>

# Drop DB

<pre><code>db.dropDatabase()</code></pre>

Also, pretty self explanatory! ;)
<br>
<br>

# Create Collections

## implicitly

<code>db.[collection-name].insertOne({"key":"value"})</code>
<br>
<br>

## explicitly

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

# Drop Collection

<pre><code>db.[collection-name].drop()</code></pre>
<br>

# Insert Many Documents

<pre><code>db.[collectionName].insertMany(
    [ < document 1 >, < document 2 >, ...],
    {
        writeConcern: < document >,
        ordered: < boolean >
    }
)</code></pre>

| Parameter      | Type     | Description                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| -------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `document`     | document | an array of documents to insert into the collection.                                                                                                                                                                                                                                                                                                                                                                                              |
| `writeConcern` | document | Optional. A document expressing the <a href="https://www.mongodb.com/docs/manual/reference/write-concern/">write concern</a>. Omit to use the default write concern. <br><br>Do not explicitly set the write concern for the operation if run in a transaction. To use write concern with transactions, see <a href="https://www.mongodb.com/docs/manual/core/transactions/#std-label-transactions-write-concern">Transactions and Write Concern. |
| `ordered`      | boolean  | Optional. A boolean specifying whether the `mongod` instance should perform an ordered or unordered insert. defaults to `true`.                                                                                                                                                                                                                                                                                                                   |

<br>

# Query

## find

<pre><code>db.[collection-name].find()</pre></code>
<br>

## find one

<pre><code>db.[collection-name].findOne()</pre></code>
<br>

## find specific

<pre><code>db.[collection-name],find({ "key", "value"})
db.customers.find({"name":"Bond"})
db.customers.find("name:"Bond"})</code></pre>

## AND

<pre><code>db.customers.find ( { $and: [ { role: "citizen" }, { age: {$gt: 40} } ] } )</code></pre>

## OR

<pre><code>db.customers.find(
    {
        $or: [            
            
            {age: { $lt: 40 }},
            {age: { $gt: 60 }}
            
        ]
    }
)</code></pre>
<pre><code>db.customers.find(
    {
        $or: [            
            
            {age: { $lt: 40 }},
            {age: { $gt: 60 }}
            
        ]
    }
)</code></pre>

## AND OR

<pre><code>db.customers.find(
    {
        $or: 
            [
                {
                    $and: 
                        [ { role: "citizen" }, { age: { $lt: 50 } } ]
                },
                {
                    $and: 
                        [ { role: "citizen" }, { age: { $gt: 60 } } ]
                }
            ]
    }
)</code></pre>

## Operators

| Operator            | Syntax              | Example                              |
| ------------------- | ------------------- | ------------------------------------ |
| equality            | {key:value}         | db.customers.find({"name": "bond"})  |
| less than           | {key:{$lt:value}}   | db.customers.find({age: {$lt: 40}})  |
| less than equals    | {key: {$lte:value}} | db.customers.find({age: {$lte: 40}}) |
| greater than        | {key: {$gt:value}   | db.customers.find({age: {$gt: 40}})  |
| greater than equals | {key: {$gte:value}} | db.customers.find({age: {$gte: 40}}) |
| not equals          | {key: {$ne: value}} | db.customers.find({age: {$ne: 40}})  |

<br>

## REGEX

<pre><code>db.customers.find (
    {
        name: { $regex: '^M' }
    }
)</code></pre>

<a href="./resources/regex.pdf">regex cheatsheet</a>
