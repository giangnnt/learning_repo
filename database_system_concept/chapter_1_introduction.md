# Chapter 1
### 1.3.1 Data Models
- **Overview**
    + **What:** ways to describe data relationships, data semantics, consistency and constrains, data properties; blueprint when design database
    + **When:** start to design database system

- **4 Categories:**
    + **Relational Model:** represent data and relationship among data. table = relation, row = tuple, column = attribute. relational model is record base model. record base model: data organize in to 'records', record have a fix structure, database contains many record types
    + **Entity Relationship Model:** collection of entities (things or objects) in real world. entities are distinguishable from others
    + **Semi-structured Data Model:** permit that individual data items of same type have different set of attributes
    + **Object-Based Data Model:** integrated in relational db

- **Procedure:**
    + collection of instructions
    + focus on process, not output
- **Function:**
    + logic component that: transform, mapping Input -> Output
    + no side effect
- **Reason database separate function and procedure:**
    + pure function (no side effect) suitable for: optimize query, execution plan and cache data. Cant optimize if function can cause side effect (for example to reorganize query, insert, update, delete, will cause faulty data if not organize in correct order)
    + pure function dont have transaction (reorganize transaction might break ACID)
    + procedure permit to have side effect and transaction control, because it will not be reorganize (cannot be optimize by optimizer as expression)

### 1.3.3 Data Abstraction
There are 3 levels of abstraction:
- **Physical level**
    + the lowest level of abstraction, its describe 'HOW' data stored on physical devices
- **Logical level**
    + next-higher level of abstraction describe 'WHAT' is stored in database, relationship between these data
- **View level**
    + highest level of abstraction show only necessary data for user
    + system may provide many views for the same database

### 1.3.4 Instances and Schemas
- **Instance:** is a collection of information stored in database at particular moment
- **Schema:** overall design of the database
    + there physical schema, logical schema and view level schema (subschemas)

## 1.4 Database Languages
- **Data-definition language (DDL)** -> specify database schema
- **Data-manipulation language (DML)** -> express database queries and updates

### 1.4.1 Data-definition language
DDL provides facility to determine constraints of data. But constrain might costly to test so its only provide constrain can be easy to test (minimal overhead)
- **Domain constraints**
    + include: Data types, Check, NULL, default value
    + basic kind of constraint, easy to determine
- **Referential integrity**
    + ensure value in one relation appear in another relation
    + prevent modification cause violation in referential integrity
- **Authorization**
    + Can assign the user to all, none or a combination of CRUD operation
- Data dictionary is a special table, contain metadata of other data table, could only access via system (not regular user)

### 1.4.3 Data-manipulation language
- **DML** is a programming language that help user to access, manipulate data stored in the database: Retrieval, Insertion, Deletion and Modification
- There are 2 types of DML:
    + **Procedural DMLs:**
        * **What:** user point out "HOW" to get the expected output
        * **When:** logic depend on instruction order, complicated condition, express loop and branching
        * **Why:** when raw SQL cant handle the complicated logic, process data row by row
    + **Declarative DMLs:**
        * **What:** user focus on "WHAT", not "HOW" to get the expected output(data)
        * **When:** when logic can be written in declarative style, highly optimize performance
        * **Why:** remain abstraction, distinct from how data is actually stored in hardware, automatically optimize by DBMS, reduce error

### 1.4.4 The SQL Data-Manipulation Language
- SQL query language is non-procedural
- Takes several tables and output 1

### 1.4.5 Database access from application programs
- SQL is not so powerful so Application programs help database with communication over network, action from user (user input, user output)
- To access database, DML need to be sent from host to the database, through application-program interface
- The reason for calling api as set of procedures is that api contains function (procedures) to do something: 
```text
connect() -> execute() -> close()
```

## 1.5 Database Design 
## 1.6 Database Engine
- Database system can be partitioned into modules, including: storage manager, query processor and transaction management component
- **The storage manager**
    + Data move from disk is slow compare to CPU, its important that the database system organize, structure the data so as to minimize move data between disk and memory
- **The query processor**
    + Allow users to obtain good performance while working at view level
    + Translate queries written in non-procedural at logical level into efficient sequence of operation at physical level
- **The transaction manager**
    + Allow user to treat sequence of database access as if they were a single unit, happen entirely or not at all (atomic)

### 1.6.1 Storage Manager
Responsible with interaction between application program and file manager (low-level data store) 
Storage manager components include: 
- Authorization and integrity manager
- Transaction manager
- File manager
- Buffer manager: responsible for fetching data from the disk into main memory, cache data in main memory...

Several data structures implemented in the storage manager:
- Data files: store the database
- Data dictionary: store metadata about the structure of the database (schema)
- Indices: provide fast access to data item (indexing)

### 1.6.3 Transaction Management
- **Atomicity:** transaction happen all or none
- **Consistency:** correctness of the transaction
- **Durability:** value must persist despite the possibility of system failure

- During execution of transaction, it may necessary to allow inconsistency because nothing is really atomic (1 must happen before the other)

Recovery manager ensure database stored the state before transaction executing, if transaction fail, the database roll back to the previous state (failure recovery)

Transaction manager include concurrency-control manager and recovery manager

## 1.8 Database Users and Administrators
### 1.8.2 Database Administrator
- A database administrator (DBA) can:
    + Schema definition
    + Storage structure and access-method definition: creating indexing
    + Schema and physical-organization modiﬁcation: update indexing, change schema
    + Granting of authorization for data access: grating authorization for user to access parts of database
    + Routine maintenance: backing up database, checking disk space, monitoring jobs