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
<br>

# Update

<a href="https://www.mongodb.com/docs/v4.4/reference/method/db.collection.updateOne/#mongodb-method-db.collection.updateOne">db.collection.updateOne()</a><br>
<a href="https://www.mongodb.com/docs/v4.4/reference/method/db.collection.updateMany/#mongodb-method-db.collection.updateMany">db.collection.updateMany()</a><br>
<a href="https://www.mongodb.com/docs/v4.4/reference/method/db.collection.replaceOne/#mongodb-method-db.collection.replaceOne">db.collection.replaceOne()</a><br>
<a href="https://www.mongodb.com/docs/v4.4/reference/update-methods/#std-label-additional-updates">additional methods</a><br>

## Usage:

<pre><code>db.customers.updateOne (
    { _id:ObjectId("62d445d84d723c4925c40f28") },
    { $set:
        { role: "double-zero" }
    }
)
db.customers.updateOne (
    {name:"Moneypenny"},
    {$set:
        { 
            role:"citizen",
            name: "Miss Moneypenny"}
    }
)
//update multiple ones (optional option)
db.customers.updateMany(
    { age: { $gt: 35 } },
    { $set:
        { role: "double-zero"}
    }
)
//update ALL
db.customers.updateMany (
    {},
    { $set:
        { role: "citizen" }
    }
)</code></pre>
<br>

# Delete

<a href="https://www.mongodb.com/docs/manual/reference/method/db.collection.deleteOne/#mongodb-method-db.collection.deleteOne">db.collection.deleteOne()</a><br>
<a href="https://www.mongodb.com/docs/manual/reference/method/db.collection.deleteMany/#mongodb-method-db.collection.deleteMany">db.collection.deleteMany()</a><br>

## Usage:

<pre><code>db.customers.deleteOne (
    { name: "M"}
)
//Delete all
db.customers.deleteMany (
    {}
)</code></pre>
<br>

# Limit

<br>

<pre><code>//prints only 2 (first) records
db.customers.find().limit(2)</code></pre>
<br>

# Sort

<pre><code>//limit & sort

//limit to 10 records
db.oscars.find().limit(10)

//using projection, limit to 10, and sort in ascending order
//projection only prints the selected fields (year, title)
//_id:0 excludes the _id field from printing.
db.oscars.find({}, {_id:0, year:1, title:1}).limit(10).sort({year:1})

//the same but descending order
db.oscars.find({}, {_id:0, year:1, title:1}).limit(10).sort({year:-1})

//the same but with a boolean field
//see that it prints the last 10 records in descending order
//if we sort {year:1} it'll print the first 10 records in ascending order
db.oscars.find({releaseYear: {$lte:1960}},{_id:0,title:1,year:1}).limit(10).sort({year:-1})</code></pre>
<br>

# Create Index

<code>db.[collection-name].createIndex( { [field to index]: 1/-1 })</code><br>

<pre><code>db.oscars.createIndex({title:1})
db.oscars.getIndexes()</code></pre>
<br>

<a href="https://www.mongodb.com/docs/manual/reference/method/db.collection.createIndex/#db.collection.createIndex">db.collection.createIndex()</a><br>
<a href="https://www.mongodb.com/docs/manual/reference/method/db.collection.createIndexes/">db.collection.createIndexes()</a><br>
<a href="https://www.mongodb.com/docs/manual/reference/method/db.collection.getIndexes/#mongodb-method-db.collection.getIndexes">db.collection.getIndexes()</a><br>
<br>

# Aggregation

## example:

Creating a new collection and adding some records:<br>

<pre><code>db.orders.insertMany (
    [
        {"cust_id":"A123","amount":500,"status":"A"},
        {"cust_id":"A123","amount":250,"status":"A"},
        {"cust_id":"B212","amount":200,"status":"A"},
        {"cust_id":"A123","amount":300,"status":"D"}
    ]
)</code></pre>
<br>

Aggregation and grouping: getting all that `match` with `status:"A"`, grouped by `cust_id`, and adding the `amount` together.<br>

<pre><code>db.orders.aggregate (
    [
        { $match: { status: "A" } },
        { $group: { 
                    _id: "$cust_id,
                    total: { $sum: "$amount } } 
                }
    ]
)</code></pre>
<br>

Printing out the distinct `cust_id`'s:

<pre><code>db.orders.distinct (
    "cust_id"
)</code></pre>
<br>

<a href="https://www.mongodb.com/docs/manual/aggregation/">Aggregation Operations Documentation</a>
