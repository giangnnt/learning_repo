# Chapter 13: Data Storage Structures
## 13.1 Database Storage Architecture
## 13.2 File Organization

- **file:** organized logically as a sequence of records, records are mapped onto disk blocks
- file also logically partitioned into "blocks", most dbms use fixed size block of 4KB or 8KB
- block may contain multiple record, in this case, we assume that is no bigger than a block, in reality, record can be larger if it is special type of record such as image

- one way to organize file is to store records that have the same length 
- another way to organize file is to store records that have different length

### 13.2.1 Fixed-Length Records
- to store fixed-length record, we first devide the block with record length, any remaining space are left unused
- it is undesirable to move records to occupy the unused space since it involve block I/O
    + to deal with problem of unused space, we create a link list of free space
    + file header contain a pointer to the first block of the free space, the first block of the free space contain a pointer to the next block of the free space, and so on
    + on insertion, we use the pointed block by the header to write the new record, and update the header to point to the next block of the free space
    + if head pointed to null, it means that there is no free space, we append record to the end of the file
    + this approach is simple is available for fix-length record, but it is not suitable for variable-length record since the size of the record can change

### 13.2.2 Variable-Length Records
- to store variable-length record, we must solve 2 problems:
    + how to extract individual record in block  
    + how to extract individual attribute in record

- the layout of a record can be divided into:
- initial part + variable-length attribute data
    + initial part: 
        * variable-length attributes are represented as "inital part" which contain the offset and length
    + offset is the distance from the beginning of the record to the beginning of the variable-length attribute
    + followed initial parts are the fixed-length attributes

- **null bitmap:** a part of the record that indicate whether an attribute is null or not: 
    + for ex: if there are 8 attribute, we can use 1 byte to indicate whether each attribute is null or not, if the first bit is 1, it means that the first attribute is null
    + in some representtion, null bitmap is set at the start of the record which can eliminate spare space 
    + the tradeoff is that we have to constantly check the null bitmap to determine whether an attribute is null or not, which can cause performance overhead 

- **slotted page structure:** a common way to organize variable-length record in a block, it consist of:
    + header: located at begining of the block, it contain the number of record entries, end of free space, array of location + size of each record 
    + records are allocated contiguously start from the end of the block
    + inserted record is allocated at the end of free space and the header add entries for the new record
    + when a record is deleted, its entry is set to deleted and the space is reclaimed by free space, other records are moved so that free space is between the header array and the first record again, end of free space is updated

### 13.2.3 Storing Large Objects
- $B^+$ tree fie organization allow us to efficiently store and access large objects 
- storing BLOB data internal coming with many problems such as efficiency or size for backup> 
    + one solution is to store BLOB data in file system and store the pointer to the BLOB data in the database, but it can cause consistency problem since the pointer in the database can point to a non-existing file in the file system
    + some database support file system integration such as constraint when delete a file 

## 13.3 Organization of Records in Files
- there are serveral ways to organize records in files:
    + heap file organization: records are stored anywhere in the file where there is space, no particular order
    + sequential file organization: store in sequential order based on "search key"
    + multitable file organization: records of differents relation are stored in the same file or even in the same block to reduce cost of certain join operations
    + $B^+$ tree file organization: records are stored in a $B^+$ tree structure, it allow efficient search, insertion, deletion of records based on the search key
    + hashing file organization: hash function is computed on some attribute of the record, then it will determine the block where the record is stored

### 13.3.1 Heap File Organization
- record is stored anywhere in the file, once placed, it is not usually moved
- one option is to insert record at the end of the file
- it important to use the for db to use the freed space from deleted record and to find block with enough space without sequential scan the whole file

- **free space map:** contains array of entries, each store 1 byte, each entry represent a fraction f if the block is free, to calculate the f 
    + for ex: if the block have 1000 byte space 
        * first we calculate the scale fmax = 1000 / 256 = 3.9 byte/unit
        * then we calculate the fraction f, for example if there is 400 byte free space, f = 400 / 3.9 = 102.5 unit
        * so the fraction f is 102.5 / 256 = 0.4
- free space map can still be slow for large file, there for we create a seconf layer of free space map for each 100 entries
- for evean larger file, we continue to create higher layer of free space map
- calculate free space on any updated would be expensive, so we can update the free space map periodically
    + such approach can cause outdated free space map, error can happen but further process is taken to deal with the error so it is acceptable

### 13.3.2 Sequential File Organization
- **sequential file:** processing of records in sorted order base on some "search key"
    + **search key:** an attribute or a set of attributes that determine the order of the records in the file, it is not necessarily be a primary key or super key
- we chain record together by pointer
- to minimize block access, we can physically store records that are close in order of search key

- to delete a record, we simply mark the record as free and update the pointer of the previous record to point to the next record
- to insert a record, we first find the position to insert
    + if there is enough free space in the block, we insert the record to **overflow block**
    + then we locate the record which came before and after the inserted record by using search key and update its pointer

- sequential file organization allow fast insertion but it force to read record in sequential order which does not match the physical order of the record on the disk
- after a while, the correspondence between the logical order of the record and the physical order of the record can become very bad, which can cause performance 
    + the file should be reorganized periodically 

### 13.3.3 Multitable Clustering File Organization
- **clustering** is a technique to store records of different relation in the same file or even in the same block by using "cluster key" to reduce cost of certain join operations
- it can also store cluste key as "header" to reduce storage overhead
- but it can slow other operations >

### 13.3.4 Partitioning
- **partitioning:** the process of dividing a large relation into smaller relations
    + provide a universal interface to access the data and translate the query into operations on the smaller relations
    + partition must not overlap
    + partition using lexicographical order so write a sub-partitioning if needed

## 13.4 Data Dictionary Storage
- dbms also store about info about relations, attributes, indexes, etc. called "metadata" in a special file called **data dictionary** or **system catalog**

![](../images/system_metadata.png)
> illustration of the data dictionary

- when a dbms need to retrieve records, it first access ***Relation_metadata*** to get location and storage organization of the relation
-  but ***Relation_metadata*** must be store elsewhere, might be in **database code** or **fixed location**
- system metadata is frequently accessed, so it is usually stored **in main memory** for fast access

## 13.5 Database Buffer
- a major goal of database is to **minimize** the number of disk I/O
- one way to achieve this goal is to keep as many block as possible in main memory 
- **buffer** is a part of main memory that is used to store block read from disk, we also need a **buffer manager** to manage if the buffer is outdated as the block is recently updated

### 13.5.1 Buffer Manager
- if `block is in buffer` -> **return** the block's pointer to requester
- if `block is not in buffer`:
    + **allocate** space for new block
    + **thown-out** other block if there is no free space in the buffer, **writen to disk** if the block is **modified**
    + **request** the needed block from disk and store it in the buffer
    + **return** the block's pointer to requester

#### 13.5.1.1 Buffer replacement strategy
- **Buffer manager** specifies which block is **least recently used (LRU)** to be **evicted** 

#### 13.5.1.2 Pinned blocks
- for a process to read data from a block it first must:
    + `pin` the block in the buffer to prevent it from being evicted 
    + `unpin` the block after it is done with the block
- multiple process can pin the same block for read, when a process pin a block, it also increase **pin count**
- block can only be evicted when its **pin count = 0**

#### 13.5.1.3 Shared and Exclusive Locks on Buffer
- database system provide **two types** of lock on buffer to prevent inconsistency when multiple process access the same block:
    + `shared lock` allow multiple process having shared lock to read on the same block
    + `exclusive lock` allow only one process to exclusively access to the block, and not allow any process to access to the block until the exclusive lock is released
- when a process want to **access** a block while another process have an exclusive lock, the request is **pending** until the exclusive lock is released
- process must **pin** block before acquiring **lock** and release it before **unpin**

#### 13.5.1.4 Output of blocks
- we should not wait until the buffer space is needed to to **output** block to disk
- instead, we can **output** block to disk periodically or when the block is modified, the process to **output** block must acquire **shared lock**

#### 13.5.1.5 Forced output of blocks
- **forced output** is the process of writing modified block to the disk to esure consistency of the data on disk
- before write actual block to disk, we write **log** to disk first to ensure transaction have enough information to recover the data in case of failure

### 13.5.2 Buffer-Replacement Strategy
- database is able to predict partern of block access accurate than general purpose OS, base one information about the **data** and **query** being processed
- an example is **toss-immediately** strategy: given the following SQL
```sql
select *
from instructor natural join department;
```
- the block of **instructor** relation is read once and not needed anymore, so we can **toss** the block immediately after reading it, which can free up buffer space for other block
- another example is **most recently used (MRU)** strategy: for each **instructor**, we need to read the corresponding **department** record as follow
$$ R1 -> R2 -> R3 -> R4 -> R5 $$
- we about to read **R6**, but buffer is full, we should not evict **R1 (LRU)**, instead, we should evict **R5 (MRU)** since it is most likely the last one to be accessed in the next loop
- buffer manager can use **statistical information** to know which block is likely to be accessed in the future
- **statistical information** can be collected by:
    + `ANALYZE` command: it collect information about record count, block count, selectivity
    + **monitoring & query history**
    + **system catalogs**
- ***index*** for a file may be accessed more frequently than the file itself, so we can give higher priority to index block than data block
- use `EXPLAIN (ANALYZE, BUFFERS)` to see the buffer usage of a query
    + `Shared Hit`: number of blocks is already in the buffer
    + `Shared Read`: number of blocks is read from disk
    + `Shared Dirtied`: number of blocks is modified in the buffer
    + `Shared Written`: number of blocks is written to disk for to free up buffer space
$$CacheRatio = \frac{\text{ShareHit}}{\text{ShareHit} + \text{SharedRead}} \times 100\$$
- in **concurrency environment**: if a **request is delayed** cause other block is **locked** by another transaction, **buffer manager** can consider **evicting the other non-locked** block of the needed by the **delayed request**

### 13.5.3 Reordering of Writes and Recovery
- ***log disk***: a separate disk to store log, log is written before the actual block is written to disk in sequential order, not re-ordered (source of truth for recovery)
- ***journaling file systems***:  a file system that keep a log of changes to the file system, can be implemented without separate log disk
- ***systemd journaling***: a service that provide a logging for kernal, service and applications
- file write by application are not ussually written to log disk, DBMS instead implement its **own logging mechanism**

## 13.6 Column-Oriented Storage
- **column-oriented** is suitable for analytical query which process **many rows** but only access **few attribute**:
    + **reduced I/O**: only read the column that are needed for the query, not the whole row
    + **improved CPU cache performance**
        * `cache line`: a unit of data that is transferred between main memory and CPU cache, it typically contain 64 bytes
        * when a **CPU** need data, it will read form L1 if `Cache Miss`, it continue to read from L2, L3, until `Cache Hit`
        * if even L3 miss, it will read from main memory, which is much slower than reading from cache
        * when using `row-oriented`, adjacent byte contain attribute that are not needed for the query
        * in contrast, when using `column-oriented`, CPU cache line contain byte of the same attribute, which can improve cache performance
    + **better compression** `compression` can significantly reduce the storage space and time taken to to retrieve data. Storing same type of data together in column-oriented storage can improve compression ratio
    + **vector processing**: a technique that allow CPU to process multiple data element in a single instruction using larger register size 125-512 bit (to be continue Vector Processing)
- **indexing** and **query processing** in column-oriented storage can be more complex than row-oriented storage and must be designed carefully to achieve good performance
- **column-oriented storage** also have some drawbacks:
    + **Cost of tuple reconstruction**: reconstructing a tuple from column-oriented storage can be expensive since it multiple I/O
    + **cost of tuple deletion and update**: 
        * deleting or updating a tuple in column-oriented storage can result in **recalculating** the compressed data for the column, which can be expensive
        * in data warehouse, old tuple is **bulk deleted** and new tuple **appended** to the end of relation
    + **cost of decompression**: in ***simplest compression representations*** (without indexing,...) fetching a record need to **read** all the data **from the beginning** of the file. Which un-suitable for transaction processing query since it only need to **fetch a few record**, but it can be suitable for analytical query since it need to **fetch many record**
- even analytical query dont need to **access all records** which are not selected by the query
    + the solution is to use **skipping technique**, where we keep track group of records for example 1000 record
    + for example, we need to find record at 2400, we can skip the first 2 group of record and read the third group of record
- ***ORC***: a technique of organizing records in file into **stripes**
    + a file contain serveral stripes, each stripe is up to about 250mb
    + **row data** contain compressions for each attribute of records and place it consecutively space 
    + **index data** **ORC*** offen store recors into **row group**, for example 10000 records
        * **index data** contain the **minimum** and **maximum** value of each attribute for each row group, it allow us to skip stripe that do not contain the record we need
        * it also contain **pointer** to the row group in the row data
- some columnar storage allow to store **group of columns** that often accessed together to be stored together
- there also some **hybrid storage** that allow collumn-oriented storage without using compression within a block and block contains set of tuples        
    + it allow retrieving a tuple **without multiple disk accesses**, while provide **effient** memory access, cache and vector processing
    + it's drawback is that, when we need to access only a few attribute of a tuple, we still need to read the **whole block** which 
    + another drawback is that it can't archive benefits of **compression**

## 13.7 Storage Organization in Main-Memory Databases
 - ***main-memory database*** is that store all data in main memory, remove entirely the **buffer manager**
 - it can significantly improve performance since it remove the cpu cycle needed of **check** if the record in the block, **finding** where is it located in the block and if not **read** the block from disk
- **slotted memory** structure is not suitable for main-memory database since it require **2 layer of indirection** to access a record and **locking** the block to ensure **not reading** records when its **moving arround** due to **update** and **delete**
- **direct access** to record is **not allow moving record** around since doing it, we must update all pointer to the record, which can be expensive
- the db must ensore that main-memory is not **fragmented over time** either by using **suitably** designed memory or performing **compaction** periodically 
- column-oriented can store all value of an attribute in consecutive memory location. But if we need to append new record, existing record may need to be moved which can be expensive, so we can use **indirection table** to store the pointer to the physical array location of the attribute value