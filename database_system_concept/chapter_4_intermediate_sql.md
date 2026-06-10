# Chapter 4: Intermediate SQL
## 4.1 Join Expressions
- **natural join** calculate relation from tables base on attributes that have the same name
- **join using** is a short for join on: instead of defining table alias, both attribute, join using only specify which table to join on. 
- the condition for using keyword is that both table must have the same attribute

### 4.1.2 Join Conditions
### 4.1.3 Outer Joins
- **outer-join** preserve tuples which would be lost in a join by create a tuple in the result containing null values
- **full outer join** preserves tuple in both relation
- step to produce a left outer join:
    + compute inner join
    + for every tuple t in left table that does not match any tuple in the right hand side relation in the inner join, add a tuple r to the result
    + tuple r constructed from: all attribute from the left hand table + null value for attributes from right hand table
- **full outer join** is a union from left outer join and right outer join (note that inner join duplicate will be eliminated)

- on and where behave differently on outer join 
    + **on** added a null padded for tuples that don't match the predicate
    + **where clause** only choose tuple which meet the predicate

### 4.1.4 Join Types and 
## 4.2 Views
- **Views** is used when there is a security concern, or using to visualize higher layer of business instead of table
- the view relation is recomputed whenever we use it

### 4.2.3 Materialized Views
- **materialized views:** certain database systems allow view relation to be stored instead of recomputed each time being use
- the action of keeping the materialized view update to date is call view maintenance
- view maintenance behave differently on different dbms, some will actively recompute the view if relation being change, some lazily recompute, other depend on how db admins configuring
- db admin must consider between the benefit of pre-compute view and its storage cost, overhead for updates

### 4.2.4 Update of a View
- modification on view are not permitted on view relations, except limited case due to data inconsistency problems.
- view can be update-able (but the still problems) when its meet these constraint
    + from clause use 1 relation
    + select clause don't contain expression, aggregates, distinct
    + attribute listed in select are null-able, and not part of primary key
    + query don't contain group by or having clause
- there is an option that view update is reject if tuple inserted not satisfy the where clause conditions
- more preferable way to modify a view is to use "trigger"

## 4.3 Transactions
- **transaction:** contain instruction and begin implicitly when sql statement execute
- transaction must end with: commit work or rollback work
- in case of power outage or system crash, rollback occurs when the system restarts
- if programs terminates without execute any command, its either committed or rollback depend on the implementation
- in many dbms implementation, each sql statement is taken to be a transaction on its own, auto commit of individual sql statement must be turn off if there are many sql statement will executed at given time

## 4.4 Integrity Constraints
- integrity constraints are usually identified as part of db schema when its created as create table command being used
- it can also be defied in alter table add constraint command

### 4.4.1 Constraints on Single Relation
### 4.4.3 Unique Constraint
- **unique** says that no two tuple have the same value on unique attribute
- null = some value is not true

### 4.4.4 The Check Clause
- a check clause is satisfied if its not false, so unknown is evaluate as valid in check
- when a foreign key is violated, instead of rejecting the action, system decide what to do base on associated clause:
    + `on delete cascade`: also delete tuples that referred to the deleted one
    + `on update cascade`: also update tuples that referred to the updated one
    + we can replace cascade with "set null" or "set default" for different behavior

### 4.4.5 Referential Integrity 
- by default, foreign key referenced to primary key
- we can also specified attribure which is refferred to, as long as the the attribute is unique and not null
- we eventually can make foreign key reference a non candidate key(not unique, null allowed), but its not recommended and not supported by many dbms
- forreign key and referenced key must compatible (same domain, same number of attribute)
- if cascading can not be handled(self referenced), the action will be rejected
- attributes of foreign key can be null, if any attribute is null, the foreign key constraint is not enforced (null is not equal to any value, so it can't violate the constraint)

### 4.4.6 Assigning Names to Constraints
### 4.4.7 Integrity Constraints Violation During Transaction
- integrity constraint is checked at the end of the transaction, not during the transaction
- so its possible that during the transaction, the constraint is violated temporarily, but at the end of the transaction, the constraint is satisfied again
- **initial deferred:** constraint is checked at the end of transaction
- **deferable:** constraint can be set to immediate or deferred mode at the start of transaction
- we can walk around the constraint check by using null value if the attribute permit null

### 4.4.8 Complex Check Conditions and Assetions
- not many dbms support subquery in check condition
- if there is subquery in check condition, its evaluated each time the target relation and refferred relation is modified
- **assertion** is a general condition that involve more than one relation
- not many dbms support assertion due to its complexity and cost to evaluate

## 4.5 SQL Data Types and Schemas
### 4.5.1 Date and Time Types
- `date`: store date value (year, month, day)
- `time`: store time value (hour, minute, second, fraction of second, time zone)
- `timestamp`: store both date and time value
- sql define functions to get current date and time: `current_date`, `current_time`, `current_timestamp`
- `interval`: represent a duration of time (difference between 2 date/time/timestamp)

### 4.5.2 Type Conversion and Formatting Functions
- **implicit conversion:** automatic conversion between compatible data type
- **explicit conversion:** user use cast operator to convert data type
- **formatting function:** convert data to string with certain format
- **coalease** indicate that all the value in the argument must be the same type

### 4.5.3 Default Values
### 4.5.4 Large Object Types
- DBMS support `clob`, `blob` for large object storage (photo, video, document...)
- **lob:** large object
- it's costly to process large object, so dbms usually store lob in separate location and only store pointer in the relation

- **temporal validity**
    + there is a need to keep track of time period that a fact is true in the real world
    + the solution is to add 2 attribute: valid time start and valid time end to the relation
    + but this approach also came with problems: the same tuple (same primary key) can appear multiple time with overlapping time period
    + the solution for postgres is to use EXCLUDE constraint to prevent overlapping time period for the same primary key

### 4.5.5 User-Defined Types
- At the system level, some field are just the same (for ex: phone number, id number...), but at the logical level, they have different meaning
- user-defined type help to distinguish those field at the logical level 
- user-defined type also help to prevent accidental mix-up between those field at the system level
- we can create domain by "Create Domain" command
- **type:** define how data is stored
- **domain:** define constrain (range of valid values) on top of type

### 4.5.6 Generating Unique Key Values
- when the always option is used, user cannot specify value for the attribute during insertion, the system will automatically generate value

### 4.5.7 Create Table Expressions
- `create table like`: create a new table with the same structure as an existing table
- `create table as`: create a new table with the structure and data from the result of a select query

### 4.5.8 Schemas, Catalogs and Environments
## 4.6 Index Definition in SQL
- **index** is a data structure that improve the speed of data retrieval operation on a database table at the cost of additional writes and storage space to maintain the index data structure
- `create index` tradeoff between read and write performance, space cost
- when submit a sql query that can benefit from index, the query optimizer will choose to use the index to speed up data retrieval
- the "benefit" is refer to the cost of using index is lower than cost of scanning the whole table

## 4.7 Authorization
- Authorization on data include: select, insert, update, delete. These are called privileges

### 4.7.1 Granting and Revoking Privileges
- A user who creates a relation automatically have all privileges on that relation

### 4.7.2 Roles
- **role** is a named group of related privileges
- role help to manage authorization when there are many users

### 4.7.3 Authorization on Views
- authorization on view is similar to relation
- it is useful when we want to restrict user access to certain attribute or tuple in the relation
- user create view does not receive all privileges on that view
- the excecute privilege is granted on function or procedure
- function and procedure have all the privilege that the creator have
- if function definition contain sql security invoker clause, the function will have the privilege of the user who call the function instead of the creator 
- there are two type of sql security: invoker and definer
    + **invoker:** function run with the privilege of the user who call the function
    + **definer:** function run with the privilege of the user who create the function

### 4.7.4 Authorization on Schemas
### 4.7.5 Transfer of Privileges
- The creattor of a relation/view/role holds all privileges on that object, including the privilege to grant those privileges to other users
- User have privilege if and only if there is a line between root (DBA) to that user in the privilege graph
- the restrict keyword is specified in order to prevent cascading revoke of privilege
- we can use granted by role to indicate that the privilege is granted by a role instead of user

### 4.7.7 Row-Level Authorization
- DBA create a function that generate predicate base on user
- when user query the relation, it will automatically add the predicate to the where clause
- this approach is call **Row-Level Security (RLS)**, in oracle call **Virtual Private Database (VPD)**

