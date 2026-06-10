# Chapter 10: Big Data
- **big data:** a term used to describe the large volume of data that is generated and collected by organizations, it is characterized by the 3Vs: volume, velocity, and variety

- company use big data for multiple purpose such as:
    + decision making: big data can provide insights and support decision making by analyzing large volume of data (post making, news, advertisement, etc.)
    + optimize user experience
    + measuring and improving the impact of decision

- there are some source of big data:
    + user interaction data: data generated from user interaction with website, mobile app, social media, etc.
    + transaction data: data generated from transaction such as purchase, payment, etc.
    + sensor data: data generated from sensor of the devices connected to the internet (IoT)
    + metadata: data that describe other data

### 10.1.2 Querying Big Data
- building a big data system is hard because of the scale and complexity of big data, it require a different approach than traditional database system
- there are two categories of big data system:

- **transaction-processing system (high scalability):** focus on handling high volume of transaction
    + for it to do that, the system need to be fast and scalable (trade off with feature that normal database system have)
    + one approach is to use a key value store
    + join problem can be solved by using view materialization (precompute the join result) or by using applicaltion level join (join in the application code instead of in the database)

- **query-processing system (high scalability, need to support non relational data)** 
    + 1st: it need to support store large volume of data
    + 2nd: it need to support query on non relational data

- Big Data application often require processing data that is non-relational, such as text, image, video, etc. 
- with the rapid growth of big data, there is a need for processing being done in parallel across multiple machine, which require a different approach and is not easy to archieve

## 10.2 Big Data Storage Systems
- to handle such large volume of data, big data storage system need to be distributed across multiple computing and storage node
- such a system includes:
    + distributed file system: a layer of abstraction that allow us to store and access files across multiple machine as if they were stored on a single machine, run above the local file system (for ex: HDFS)
    + sharding accros multiple databases: sharding is a technique to partition a large database into smaller
    + key-value storage systems: sometimes call NoSQL because they dont support sql and not a full feature database system
    + parallel and distributed databases: provide a database interface but store data accross multiple machine, query processing in parallel across multiple machine

### 10.2.1 Distributed File Systems
- **distributed file system** design to store and access files, whose size can be from megabyte to hundred gigabytes, across multiple machine
- file are broken into block and stored across multiple machine, each block is replicated (typically 3) to ensure fault tolerance and availability
- DFS contain:
    + directory system: allow us to organize files into dir ans subdir
    + a mapping system: map file name to sequence of identifier of the block that store the file
    + ability to store and retrieve with specified identifier, its provide block identifier to store and retrieve file, in case of distributed and replicated file system, each block identifier will also provide a set of machine identifier
- **NameNode:** a master node that manage the metadata of the file system contain directory system, for each file, it store a list block id and the machine id that store the block and its copies
- **DataNode:** a worker node that store the actual data blocks
- to read a file, client call NameNode to get the block id and the machine id that store the block, then call assigned DataNode to retrieve the block
- to write a file, the call NameNode to get the block id and the machine id that will store the block, then client send the block id and block to the assign machine to store the block 
- hdfs can also be connected as local file system

### 10.2.2 Sharding
- **sharding** is a technique to partition a large database into smaller, more manageable pieces called shard
- patition can be done by range of value, by hash of the value
- sharding can be done at the application
    + this approach is simple way to scale application, but it require the application to handle the complexity of sharding, such as routing query to the correct shard, handling data consistency across shard, etc.

### 10.2.3 Key-Value Storage Systems
- As application grow, the amount of files that need to be stored also grow. Its pose problem such as metadata overhead and optimization for small file
- a relational database is ideal for those above problem, but designing a relational database that can handle dbms features in a distributed environment is hard
- serveral key-value storage system require to stored data in a specific format (called document store)
- **key-value storage system** is not a full feature database system, it only provide basic operations such as put, get, delete, etc.
- some document store additionally support limited forms of query on data values
- key-value system also not support transaction, which can lead to data inconsistency if not handle properly
- partitioning can be done base on the value of specified attribute

- **Bigtable** is a key value storage system:
    + it require key to be follow a specific format: `(RowKey, ColumnFamily:ColumnQualifier, Timestamp) → Value`
    + the reason for this format is to: Sparse, distributed, persistent multi-dimensional sorted map

### 10.2.4 Parallel and Distributed Databases
- **parallel database system** can be classified into two categories:
    + using for transaction processing: suported only a few machine in a cluster
    + using for analytical processing: support a large number of machine in a cluster from 10 to 100 machine
- in such system, if a query failed, we can retry the query on the same node machine with data from the replica since the probability of failure is low
- in case of extreme system with large number of machine, retrying the query on the same machine may not be effective. In such case, we must "partial recovery" which will be discussed later
- however doing so cause significant overhead

### 10.2.5 Replication and Consistency
- there a serveral problems with replication:
    + how to justify the point of transaction commit: transaction is said to be commit all when all replica have commit the transaction
    + how to handle update: if the machine contain the replica failed, how can we update the replica

## 10.3 The MapReduce Paradigm
- **MapReduce** is a programming model and an associated implementation for processing and generating large data sets with a parallel, distributed algorithm on a cluster
- The MapReduce model consists of two main functions: 
    + **Map:** takes a set of input key-value pairs and produces a set of intermediate key-value pairs. 
    + **Reduce:** takes the intermediate key-value pairs produced by the Map function and merges them together to produce a smaller set of output key-value pairs. 

- for pair of `(mapKey, mapValue)` in the input data, map function will produce a list of `-> (reduceKey, reduceValue)` pairs
- then the system will group, shuffle and sort the intermediate key-value pairs by reduceKey `-> (reduceKey, [reduceValue1, reduceValue2, ...])`
- finally, for each reduceKey, the reduce function will take the list of reduceValue and produce a list of output key-value pairs `-> (outputKey, outputValue)`

### 10.3.4 Parralel Processing of MapReduce Tasks
- Part i: input partion i 
- Map i: the node that process the map function for Part i
- Reduce i: the node that process the reduce function for output of Map i

- Step 1: Master Node send copy of map() and reduce() function to each worker node
- Step 2: Map i start processing Part i `-> (reduceKey, reduceValue)` pairs
- Step 3: `(reduceKey, reduceValue)` going through a filter called "Partitione Function" to determine which Reduce i will process the reduceKey
    + common partition function: `Reducer_ID = hash(reduceKey) mod R`, where R is the number of reduce node
    + `mod R` ensure that R_ID is in the range of 0 to R-1, which is the index of the reduce node
- Reduce i will actively fetch the output of Map i
- Step 4: Reduce i start apply merge and sort on the intermediate key-value pairs
- Step 5: Reduce i apply reduce() function on the intermediate key-value pairs to produce the final output

### 10.3.5 Mapreduce in Hadoop
- **combine():** also called "combiner" or "semi-reducer", it is an optimization technique that allow us to perform a local reduce on the output of the map function before sending through network it to the reduce function 
    + it can significantly reduce the amount of data that need to be sent through network
    + but it is not always possible to use combine function, it only work for certain type of reduce function that is commutative and associative (for ex: sum, count, max, min, etc.)

### 10.3.6 SQL on MapReduce
- **select:** can be implemented as a map function that filter the input data based on the condition in the select statement

- **aggregate function:** can be implemented as follows:
    + map function: for each input record, emit a key-value pair where the key is the group by attribute and the value is the value of the aggregate function for that record
    + reduce function: for each key, apply the aggregate function to the list of values and emit a key-value pair

- **join:** can be implemented as follows:
    + map function: for each input record, emit a key-value pair where the key is the join attribute and the value is the record itself
    + reduce function: for each key, apply the join condition to the list of records and emit the joined record if the condition is satisfied

## 10.4 Beyond MapReduce: Algebraic Operations
- modern data system processing data reuse the same idea of take data input, apply a series of algebraic operation on the data, and produce output data
- but the algebraic operation is not limited columns with atomic data types as in relational algebra, it can also operate on complex data and more complex expression

### 10.4.2 Algebraic Operations in Spark
- **Resilient Distributed Dataset (RDD):** is a equivalent of a relation in relational algebra

## 10.5 Streaming Data

### 10.5.1 Applications of Streaming Data

### 10.5.2 Querying Streaming Data
- **window** is a way to deal with unbound nature of streaming data, it allow us to compute with a finite amount of data at a time
- another is instead of re quering data from the beginning, we can add the arriving data to the previous query result, this is called incremental processing

- base on the two above approach, we develop to serveral approaches:
    + **continuous query:** a query that is continuously running and produce output as new data arrive
    + **streaming query languages:** extesion of sql that support querying streaming data using window 
    + **algegraic operators on the streams:** define "user functions" to apply to a stream of data, great on maintain internal state of the stream
    + **pattern matching:** define a pattern of "events" that we want to detect in the stream, and the system will emit an "action" when the pattern is detected

- stream-processing system may not guarantee for persistence data storage
    + we approach this problem by using machanism such as copy data, separate storage and query processing
    + this approach consequently lead to 2 problems: we have to write a data twice, streaming query may not have access to all data

#### 10.5.2.1 Stream Extensions to SQL
- instead of a "work-around" to query streaming data by adding a timestamp to the data and use window function, we can approach this problem by streaming query language:
    + **tumbling window:** window that is non-overlapping and contiguous, it is defined by a fixed "size"
    + **hopping window:** after a "hop", create a new window with "size", for smoothing data 
    + **sliding window:** define a new window for each new data arrive
    + **session window:** define using first time user is active and a "gap" of inactivity. When user is active, the window keep expanding, when user is inactive for a "gap" time, the window is closed and a new window will be created when user is active again

- we must defferentiate between stream and relation
    + **stream:** unbounded sequence of data, implicit timestamp
    + **relation:** data is fixed at any point

- join a relatin with a stream
- join a stream with a stream create a new stream, (but we have to specify the window for the join, avoiding the unbound nature of the stream)

### 10.5.3 Algebraic Operations on Streams
- we approach streaming process with a Algebraic operations model
- the model provide: 
    + user defined function
    + predefined algebraic operations

- the key task of operating a model like this is to create a fault-tolerance system

- the logical routing can be represented by a directed acyclic graph (DAG), where each node represent an operation and each edge represent the data flow between the operation, the exist point is sink node where data can be store elsewhere
- Apache Storm use topology config file to define the DAG
    + data source node is called "spout"
    + data processing node is called "bolt"

- in algebra operations system, we use publish-subscribe mechanism to handle the data flow between the operation
    + **pub-sub system** can decouple operators
    + **fan-out:** the output of an operator can be sent to multiple downstream operator, which can be use for other purpose such as backup, monitoring, etc.
    + create a "buffer" between operators which allow fault-tolerance
- data entry considered as "document" and tagged with "topic", operators can subscribe to topics to receive the data, and publish to topic to send data to other operator
- published document are retained for retention period

- there is issue when dealing with stream, some operaters only process until the stream is consume which lead to unbounded time.
    + spark approach this problem by using discretized stream which treat the stream asfile or relation
    + out put of these operations in treated as a stream again

## 10.6 Graph Databases
- there are two approaches when come to parallel processing of graph data:
    1. Map-reduce and algebraic framework: basically represent the graph data as a relation and use map-reduce or algebraic operation to process the graph data, but this approach is not efficient
    2. Bulk synchronous processing frameworks
- we focus on three things:
    + **mind set:** focus only on a single vertex (node) and its neighbor, interact through message passing, and do not have global view of the graph
    + **superstep:** follow these step rules:
        * receive message from previous superstep
        * compute new value for the vertex and send message to neighbor
        * synchronize with other vertex (wait until all vertex have finish the superstep)
    + **halting condition:** the computation will halt when all vertex have halt, and there is no message in transition