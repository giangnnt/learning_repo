# Chapter 2
## 2.1 Structure of Relational Databases
- A relation (table) is a set of tuples; each tuple is an ordered list of values and corresponds to a row in the table.
- The domain is set of permitted value for that attribute
- A domain is considered atomic when its elements not divisible, for ex: address can divide to house number, road, city...
- But more important is how we use the domain, don't divide domain too small, for ex: phone number divide to single unit number...
- Null value signifies the missing, absent of value
- We need to consider between strict structure and efficiencies (pros ans cons)

## 2.2 Database Schema
- Relation correspond to instance of an object while Relation schema correspond to the type of the object its self

## 2.3 Keys
- Keys identify tuple uniquely (allow tuple to have same value for all attributes except key)
- Super key is a set of attribute, allow identify tuple uniquely
- Candidate key is super key that don't exist subset also a super key (minimal super key)
- Primary key is a chosen candidate key
- Two individual tuple are prohibited from having the same key attribute at the same time (primary key constraints)

- Foreign key A for each tuple in Relation r1 must also be the value of B (primary key) for some tuple in r2 (foreign-key constraint)

- Foreign Key Constraints is a special case of Referential Integrity Constraint, where referred attribute is a primary key
- Also note that only FK Constraint is support by dbms

## 2.4 Schema Diagrams
## 2.5 Relational Query Languages
- Query language is language user use to request data from the database
- Query language usually on a higher level than standard programming language
- Can be categorized as: Imperative, Functional, Declarative

- **Imperative query language:** user instruct the system to perform specific sequence of operation, in order to achieve the result
- **Functional query language:** can be call a subset of imperative (procedure based) language, include function call which is side effect free to achieve result
- **Declarative query language:** in the other hand, achieve result without the specific instruction, user focus on "WHAT" will be the result and the system DBMS will handle "HOW" result will be obtained

- **Relational Algebra:** will classified as functional language
- **Tuple relational calculus and domain relational calculus:** are classified as declarative

## 2.6 Relational Algebra
- Relational Algebra is a collection of operation, which take 1 or many Relation to produce a new Relation (for simple, math form of query language)
- Its operation can be Unary, Binary
- Database don't allow user to use relational algebra as query language

### 2.6.1 The Select Operation (Sigma)
- The "SELECT" select tuples that satisfy given 
```math
σdept name = “Physics” ∧ salary>90000 (instructor)
```

### 2.6.2 The Project Operation (Pi)
- mapping value input -> output like a function
- The "PROJECT" produce new relation with certain attributes left out
- Any duplicate row is eliminated (use "DISTINCT" in SQL)

### 2.6.3 Composition of Relational Operations
- result of Relational Operation is an another Relation
```math
Πname (σdept name = “Physics” (instructor))
```
- We can, instead given a name as argument of uprojection, we use a relation that evaluate to a relation
- Relational Algebra operations can be composed into a Relational Algebra expression by arithmetic operations (+,-,*,/)

### 2.6.4 The Cartesian Product Operation (Cross)
- `(a1,b1), (a1,b2), (a2,b1), (a2,b2)` from Cartesian product in math evaluate to row (tuple) in database
- `r = r1 x r2`, r in constructed by each possible pair of relation 1 to relation 2
- if r1 have n rows and r2 have m rows, r will have m * n rows

### 2.6.5 The Join Operation
```math
σinstructor.ID = teaches.ID (instructor × teaches)
```
- represent a subset of Cartesian product, satisfies some "Predicate" (a boolean condition) 
```math
r ⋈θ s = σθ (r × s) 
```
- r and s are relation, θ is predicate, ⋈ is join operation

### 2.6.6 Set Operations
- We say that relation is equivalent to function in math, because both is a collection of tuple
```math
Πcourse id (σsemester = “Fall” ∧ year=2017 (section)) ∪ Πcourse id (σsemester = “Spring” ∧ year=2018 (section))
```
- The Union (∪) merge two relation into 1 relation vertically
- To Union relations, realtion must be "compatible" which conditions are:
    + relatations must have the same "arity" (in math, arity refer to number of argument of a function)
    + the type of ith attribute of both relation must be the same for each i

- the "Intersect" (∩) produce a relation containing tuple that appear in r1 as well in r2
- the "Set-difference" (−) produce a relation containing tuple that appear in 1 relation but not the other

### 2.6.7 The Assignment Operation
- Assign result relation into temp relation variable. denoted by (←), work like assignment in programming language
- Allow query can be written in sequential, consisting of a series assignment (CTE, subquery)

### 2.6.8 The Rename Operation (Rho)
```math
ρx (E)
```
- return the result of expression E in name x
- A relation is considered a relational-algebra expression
```math
ρx(A1 ,A2 ,…,An ) (E)
```
- this form of rename operation can be used to give names to attributes in the results of relational algebra

### 2.6.9 Equivalent Queries
- There is often more than one way to write queries, those queries still output the same result. Those are called "Equivalent Queries"
- Query optimizer look what result an expression computes and find efficient way of computing result
- The "Algebraic structure" of relational algebra makes it easy to find efficient but equivalent alternative expressions