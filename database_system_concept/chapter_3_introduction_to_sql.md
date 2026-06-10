# Chapter 3: Introduction to SQL
### 3.2.1 Basic Types
- `char(n)`: fixed-length character string with user specified length n
- `varchar(n)`: variable-length character string with maximum length n
- `int`: an integer (a finite subset of integers, machine dependent CPU architect, ram...) 
- `numeric(p,d)`: a fixed-point number with user-specified precision. the number consist of p digits (plus sign) and d digits are to the right of the decimal point
- `real`, `double precision`: floating point, double precision floating point with machine dependent precision
- `float(n)` floating-point number with precision of at least n digits (n is number of bit of mantissa)
- `Value = (-1)^S × 1.F × 2^(E - Bias)`
    + **S** represent sign (+,-)
    + **F** represent the precision of the number
    + **E** is biased exponent
    + **Bias** `= 2^(k-1) - 1`  (k is the number of bit of exponent, -1 inside is for balance between positive and negative, -1 outside is except 0)

- each ype include special value call "NULL", its indicate the absent of value
- when compare char and varchar, may the database add extra space to varchar type
- Unicode mapping between character -> code point (U+XXXX)
- Encoding (UTF-8, 16, 32) translate code point -> bytes
- ASCII is a subset of Unicode
- `nvarchar`: allow to store unicode value

### 3.2.2 Basic Schema Definition
- **primary-key:** non-null and unique
- **foreign-key:** must correspond to values of the primary key attributes of some tuple in referred relation
- **not null:** exclude the null value from the domain of that attribure
- alter table command is used to add attributes to an existing relation, all tuples in the relation are assigned null as the value for the new attribute
- adding non null attribute without "DEFAULT" will usually fail

## 3.3 Basic Structure of SQL Queries
- duplicate elimination is time-consuming, but we can eliminate them by using "DISTINCT" keyword

### 3.3.2 Queries  Multiple Relations

```sql
select name, instructor.dept name, building
from instructor, department
where instructor.dept name= department.dept name;
```
- above query represent joining two table, base on where clause, it can be cross join or inner join

## 3.4 Additional Basic Operations
**Note 3.1**
- SQL standard define how many copy of each tuple in the output base on duplicate of each tuple in the input
- To model this behavior, SQL use multi-set relational algebra
- Multi-set relational algebra is defined as a set that contain duplicates.
- If there are `c1` copies of `t1` in `r1` and `c2` copies of `t2` in `r2`, there `c1*c2` copies of `t1.t2` in `r1 x r2`
- Lower level representation of SQL queries perform query optimization and query evaluation base on relational-algebra

### 3.4.2 String 
- SQL standard specifies strings in a single quotes, for ex: `'Computer'`
- If in strings contain a quote, the quote can be represent by double quote, for ex: It’s right with `'It''s right'`.
- Strings evaluation is case sensitive, but some DBMS don't distinguish case sensitive (MySQL and SQL Server) but this behavior can be changed in the setting
- `%` charater matches any substring
- `_` character matches any character
- similar to quote, there also a way to express character `%`, using `\`
- for ex: `'ab\%cd%' -> 'ab%cd'`, `'ab\\cd' -> 'ab\cd'` 

### 3.4.4 Ordering the Display of Tuples
- `order by` clause lists item in asc order
- we can wrote `instructor.ID= teaches.ID` and `dept name = 'Biology'` -> `(instructor.ID, dept name) = (teaches.ID, 'Biology')`;

## 3.5 Set Operations
- use `union all` union operation automatically eliminates duplicates
- use `union all` if we want to retain all duplicates

### 3.5.2 The Intersect Operation
- intersect operation eliminated duplicate and similar to union, we can use `intersect all` to retain the duplicate
- intersect is similar to join if its meet these condition:
    + use distict keyword
    + both join table have the same schema
    + join condition is the whole tuple

### 3.5.3 The Except Operation
- similar to above operation, except choose tuple that appear in r1 but not in r2, eliminate duplicate, use `except all` to retain duplicate

## 3.6 Null Values
- result of arithmetic expression (+, -, *, /) is null if one of the input is null
- comparison to null treat as unknown
- it's create a third logic, other than true and false
- where clause only select true

- we can use `unknown` and `is not unknown` to test if the result known or unknown
- the null treatment in select distinct clause (null = null is true) is different from predicate (null = null is unknown)
- this approach also used in union, intersect and except

- **and:** The result of true and unknown is unknown, false and unknown is false, while unknown and unknown is unknown.
- **or:** The result of true or unknown is true, false or unknown is unknown, while unknown or unknown is unknown.
- **not:** The result of not unknown is unknown.

## 3.7 Aggregate Functions
- aggregate functions take a collection (a set or multiset) of values as input and return a single value. 
- there are five standard build-in aggregate functions: avg, min, max, sum, count
- sum and avg only operate on numbers
- we can use distinct keyword in aggregation for ex: `count(distinct ID)`
- only attributes that appears in  the select statement without being aggregated are presented in the group by clause

### 3.7.3 The Having Clause
- Having applies to each group constructed by the group by clause and after group have been formed
- we can use aggregate function in having clause cause group have been formed
- any attribute appear in having clause without present in aggregate function or group by clause in erroneous

- the predicate in where clause happen first, the place the result in to group by clause
- if the group by is absent, the whole relation is treated as one big group
- having clause applied to each group, if group not satisfy the having clause, it is removed

### 3.7.4 Aggression with Null and Boolean Values
- "Partial information is still information": dù thiếu 1 phần nhưng thông tin vẫn là thông tin
- all aggregate function except count(*) ignore null values
- count on empty set return 0
- other aggregation function return null 
- some: or(dis junction) on set of boolean value
- every: and(conjunction) on set of boolean value

## 3.8 Nested Sub queries 
- sub query is a select-from-where expression that is nested within another query

### 3.8.1 Set Membership
### 3.8.3 Test for Empty Relations
- correlated sub query is a query that depend on the outer query to use as parameter for where, select of inner query

### 3.8.5 Sub queries on the From 
- rename sub query and its attribute

```sql
select dept name, avg salary
from (select dept name, avg (salary)
from instructor
```

- sub query in the from clause is different from in the select/where clause (scalar)
- in the from clause, it separated from the process, happen first so it don't allow to see other value in the select clause (without "lateral" keyword)
- in the select/where clause, sub query happen each time its iterate though the defined relation ship and return scalar value so it can apply scoping rule

### 3.8.6 The With Clause
### 3.8.7 Scalar Sub queries
- **Scalar sub queries** is defined that return 1 tuple with 1 single attribute
- if a sub query can return more than one tuple in its result, run time error occurs
- SQL simply extract the relation of the scalar sub query and return the value

### 3.8.8 Scalar Without a From Clause
```sql
(select count (*) from teaches) / (select count (*) from instructor);
```

- the expression show a scalar sub query without the top level query
- this is legal in some dbms but not in the other

## 3.9 Modification of the Database
### 3.9.1 Deletion
- both select and delete clause act on a snapshot data from the

### 3.9.2 Insertion
- there are 2 ways to insert data into relation:
    + specify a tuple to be inserted 
    + write a query which result to be a set of tuple to be inserted
- attribute values for inserted tuple must be member of the corresponding attribute's domain
- inserted tuple must have the same number of attributes
- note that similar to delete statement, tuple or relation query result mus be evaluated first (to void looping condition)
- most dbms support bulk loader: a way to insert large set of tuple in to relation without insert statement

### 3.9.3 Update