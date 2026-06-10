# Chapter 7: Relational Database Design
## 7.1 Features of a Good Relational Design
### 7.1.1 Decomposition
- **decomposition:** the process of breaking down a relation into smaller relation to eliminate redundancy and update anomaly
- a good decomposition should be lossless and dependency preserving
- **lessless decompossion** ensure that we can reconstruct the original relation from the decomposed relation by using natural join
- to obtain lossless decomposition we must ensure that the common attribute between the decomposed relation is a super key for at least one of the decomposed relation

### 7.1.3 Normalization Theory
- **normalization:** the process of organizing the attributes and relation of a database to minimize redundancy and dependency
- **normal form:** a property of a relation that indicate the level of normalization
- there are several normal form: 1NF, 2NF, 3NF

## 7.2 Decomposition Using Functional Dependencies
### 7.2.1 Notational Conventions
### 7.2.2 Keys and Functional Dependencies
- it is called "**Functional Dependency a -> b**" (a uniquely determine b) if for any two tuple in any legal instance of r(R), if they have the same value on attribute a (t1[a] = t2[a]), then they must have the same value on attribute b (t1[b] = t2[b])

- value on attribute Y (f(x) = y, the same x must have the same y)
- instance of r(R) equal to r (R is the schema, r is the relation (actual data))
- functional dependency is holds (được giữ quy tắc) when every "legal" instance of r(R) satisfy the functional dependency

- **K is superkey** of r(R) if K -> R holds on r(R): t1[K] = t2[K] => t1[R] = t2[R] (t1 = t2(the same tuple))

- we use functional dependency to:
    + to test whether a relation is "satisfy" a given set F of functional dependency (static checking)
    + to specify whether a relation is "holds" a given set F of functional dependency (not only satisfy but also that, from this point on, it will always satisfy the functional dependency) (dynamic checking)
- functional dependency is said to be "trivial" (tầm thường, hiển nhiên) if a->b and b is a subset of a (for ex: a->a, a b -> a)

- we can also inference new functional dependency from a given set of functional dependency (for ex: if a->b and b->c, then we can inference a->c)
- we use notation F+ to denote the closure of a set F of functional dependency(the set of all functional dependency that can be inference from F)

### 7.2.3 Lossless Decomposition and Functional Dependencies
- we obtain lossless decomposition if R1 ∩ R2 → R1 or R1 ∩ R2 → R2
- in other word, R1 ∩ R2 must be a superkey for at least one of the decomposed relation

- but to ensure that the decomposition is actually lossless we need to impose that:
    + R1 ∩ R2 is primary key for R1 (superkey only ensure theoretically, but primary key is the acctual implementation that ensure the lossless decomposition)
    + R1 ∩ R2 is foreign key for R2 referencing R1 (foreign key also ensure the lossless decomposition because it ensure that the value of R1 ∩ R2 in R2 must exist in R1)

## 7.3 Normal Forms
#### 7.3.3.1 Definition
- **BCNF (Boyce-Codd Normal Form):** a relation is in BCNF mean it is eliminate all redundancy and anomoly
- relation can archieve BCNF if it satisfy the following condition, for all functional dependencies in F+ at least one must holds:
    + a -> b is trivial: 
    + a is superkey for R: at least one the FD of form a->b must have a as superkey or there will be redundancy and anomaly