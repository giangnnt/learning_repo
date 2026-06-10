# Chapter 5: Advanced SQL
## 5.1 Accessing SQL from a Programming Language
- There are 2 approaches to access database from application program:
    + **Dynamic SQL:** sql statement is constructed as string at runtime and sent to the dbms for execution
    + **Embedded SQL:** sql statement is embedded in the host programming language and precompiled before compilation of the host program
- Embedded SQL is more efficient than dynamic sql because sql statement is precompiled
- Dynamic SQL is more flexible than embedded sql because sql statement can be constructed at runtime base on
- sql return relation, so when we want to use the result in application program, we need to convert it to data structure in the programming language (for ex: array, list, map...)

### 5.1.1 JDBC
- **JDBC** stand for Java Database Connectivity
- JDBC is an api that allow java program to access database
- orm -> api -> driver -> database
- JDBC dynamically load the driver class at runtime base on the database being used
- Driver support protocol that translate JDBC api call to database specific protocol

#### 5.1.1.2 Shipping SQL Statements to the Database system
- **statement object** is used to invoke method that send sql statement to the database

#### 5.1.1.3 Exception and Resource Management
#### 5.1.1.4 Retrieving the Result of a Query
- the **Next() method** of the ResultSet object is used to iterate through the result of a query
- its test whether there is another row in the result

#### 5.1.1.5 Prepared Statements
- We can compile statments, replace real argument with placeholder ($)
- When execute the prepared statement, we provide the real argument 

#### 5.1.1.6 Callable Statements
- **Callable statement** is used to call procedure or function in the database

## 5.2 Functions and Procedures
- Functions are useful with specialized data types such as images and geometric objects
- The pros of using functions and procedures: 
    + allow change in case business without chanigng other parts of application program

### 5.2.1 Declaring and Invoking SQL Functions and Procedures
- despite remain difference, table function can also be thought as parameterized view, because it return a relation base on the input parameter
- sql permits procedure to have the same name as long as their parameter list is different (overloading)
- functions can also be overloaded, as long as they have different parameter list

### 5.2.2 Language Constructs for Procedures and Functions
- **SQL/PSM (Persistent Stored Modules)** is a standard language for writing stored procedure and function in SQL
- Its provides: DECLARE, SET, IF, THEN, ELSE, CASE, LOOP, WHILE, REPEAT, LEAVE, ITERATE, RETURN
- in postgres, we can also use PL/pgSQL, which is similar to SQL/PSM but with some additional feature and different syntax
- its also support signal and exit handler
    + **signal:** raise an exception with a specific sql state and message
    + **exit handler:** specify a block of code to be executed when a specific exception is raised
- SQL allow us define function in programming language
- function defined in programming language can be called from sql statement, but it come with overhead and security concern
- **sandbox** is a security mechanism to restrict the access of function defined in programming language to certain resource (file system, network, environment variable...) 

## 5.3 Triggers
- **trigger** is a statement that automatically execute in response to certain event on a particular table or view
- to define a trigger, we need to specify:
    + the event that will activate the trigger (insert, update, delete)
    + the action that will be executed when the trigger is activated

### 5.3.1 Need for Triggers
- **for each row clause:** trigger will execute for each row that is affected by the event
- **referencing new row as nrow:** create a variable nrow that contain the new value of the row being inserted or updated (iterate through each row being affected)
- a trigger can be disabled or enabled using alter table command
- **referencing table** usually used in statement level trigger, it create a variable that contain the relation of the rows being affected by the event (for ex: all rows being inserted)
- we cannot use referencing table with before trigger because the relation of the rows being affected is not yet available before the event happen
- some problem can also be found with above approach, for ex it will cause infinite loop or memory and performance problem

### 5.3.3 When Not to Use Triggers
- using `on delete cascade` instead of trigger to maintain referential integrity
- using materialized view instead of trigger to maintain summary data
- using replication feature of dbms instead of trigger to maintain data in sub table (delta table)
    + to sync data between two table, when table a is updated, we trigger to write a log to delta table, then we have a process that read the log, delete log table and update table b, this approach is call "**Change Data Capture (CDC)**"
- when there is a used of backup and restore, trigger will be activated and cause data inconsistency, so we need to disable trigger before backup and restore, then enable it again after that
    + or we we can specify not for replication in the trigger definition, so it will not be activated during backup and restore
- many trigger can be replaced by using procedure

## 5.4 Recursive Queries
- **Transitive closure:** the set of all nodes that can be reached from a given node in a graph
- **Recursive query:** a query that refer to itself in its definition
- Recursive query is useful for querying hierarchical data, such as organizational chart, bill of materials,

### 5.4.1 Trasitive Closure Using Iteration
- in postgres, recursive with CTE is a bit different than in programming language
- the result of first part of the query is reutrn to the cte then the second part use cte as input. Then the result is used as input for the second part of the query until there is no new tuple being added to the result

- closure make sure that we can reach all node that can be reached from the given node, nothing more, nothing less
- `create temporary table`: create temp table that only available during the session, it will be automatically dropped at the end of the session
- if there is two "create temporary table" running cocurrently, each get its own copy

### 5.4.2 Recursion in SQL
- **recursive view** must define a relation that is "union" of 2 relation:
    + base query
    + recursive query
- the flow is understood as follow:
    + first, the base query is evaluated and its result is stored in the view relation
    + then, the recursive query is evaluated using the view relation as input, and its result is added to the view relation
    + this process is repeated until no new tuple is added to the view 
    + **fix point:** the point at which the view relation does not change anymore, so the recursion can stop

- the recursive query must be monotonic:
    + result of recursive query must be a superset of the input relation: if r1 is superset of r2 => recursive query result on r1 is superset of recursive query result on r2

- because recursive query depend on the previous result, it can not use non-monotonic operation, for ex:
    + aggregate function: because aggregate function is not monotonic, it can produce result that shrink when the input relation grow
    + not exist
    + except 

## 5.5 Advanced Aggregation Features
### 5.5.1 Ranking
- in postgres, null value is treated as the highest value
- using `nulls first` or `nulls last` to change the default behavior of null value in order to treat it as the lowest value
- `rank()` reset to 1 for each partition
- ranking apply after group by

- `percent_rank()` is defined as 
$$\frac{\text{rank}() - 1}{\text{number of tuples in partition} - 1}$$
- it return a value between 0 and 1, inclusive
    + it represent the relative distance of position of the tuple in the partition, compare to the start and the end -> thats why its domain is [0,1]

- `cume_dist()` is defined as
$$\frac{\text{rank}()}{\text{number of tuple in the partition}}$$
- it represent the relative position of current tuple compare to how many tuple its already pass in the partition, compare to the total number of tuple in the partition -> thats why its domain is (0,1] (include the first tuple)

- `ntile(n)` devide the partition into n group, and assign a number to each tuple to indicate which group it belong to if number of tuple in the partition is not divisible by n, its priority to assign more tuple to the higher group

### 5.5.2 Windowing
- **windowing** is different from partitioning, because it does not divide the relation into group, but it define a "window" of tuples, the window is slide through the relation, and the aggregate function is applied to the tuples in the window
- original windowing compute aggregate base on physical tuple, the window can enxpand and shrink as it slide through the relation
- use **"unbound" keyword** to define window that start from the first tuple or end at the last tuple
- use **"preceding" and "following" keyword** to define window that start from the current tuple and expand to the preceding or following tuple
- use **"range" keyword** to define window take value of the current tuple (not physical tuple) as reference 

### 5.5.3 Pivoting
- **pivoting** is a technique to transform rows value into columns, it is useful for reporting and data visualization

### 5.5.4 Rollup and Cube
- **rollup** is a extension of group by that produce subtotals and grand total in the result
- its create many group by using the same group by attribute, but with different level of aggregation
- **cube** is a extension of rollup that produce subtotals for all combination of group by attribute
- use **"grouping set"** to specify which combination of group by attribute to be included in the result
- **grouping() function** is used to determine whether a tuple in the result of grouping (1) or part of the original relation (0)