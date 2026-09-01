# Chapter 16: Query Optimization
- `Query optimization`: a process of selecting the most efficient query-evaluation plan for a given query
- we dont expect user to write a query fot it to process efficiently, it is a job of the system to construct query-evaluation plan that minimize the cost of evalation
- the difference in cost between a good strategy and a bad strategy can be enormous so it is wothwhile for a system to spend an amount of time on selection for a good strategy

## 16.1 Overview 
- given a `relational-algebra expression`, it is the job of the system to construct a `query-evaluation plan` that minimize the cost of evaluation, this require the following step:
  - generate alternative expressions that logically equivalent to the original expression
  - annotating expressions to generate query-evaluation plans
  - estimating the cost of each evaluation plan and choosing the plan with the lowest cost
- generation of query evaluation plan and annotating are **interleaved** (back & bound algorithm)
- some expressions are generated and annotated then further generation and annotation of sub-expressions are done and so on 
- as evalation plan are generated, the system will make expression using statical information of the system 
- to generate a equivalent expression, the system use *equivalence rules*
- we then use the statical information to estimate the cost of each individual operation and then combine them to determine the cost of the whole expression
- the cost is estimated so the evalation plan might not be the least costly, but if the estimate is good, it is likely to be the least costly one
- materialized views help to speed up processing of certain queries, we then learn how to maintain and use materialized views

## 16.2 Transformation of Relational Expressions
- the two expressions are said to be **equivalent** if they generate the same set of tuples (orders of tupels are not important)
- the input and output when evaluate the SQL is **multiset**
- so two expression are equivalent in the multiset version of relational algebra if they generate the same **multiset** of tuples (equal on the values and the multiplicity of tuples)

### 16.2.1 Equivalence rules
- 