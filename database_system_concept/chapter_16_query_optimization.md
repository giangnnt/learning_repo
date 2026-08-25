# Chapter 16: Query Optimization
- `Query optimization`: a process of selecting the most efficient query-evaluation plan for a given query
- we dont expect user to write a query fot it to process efficiently, it is a job of the system to construct query-evaluation plan that minimize the cost of evalation
- the difference in cost between a good strategy and a bad strategy can be enormous so it is wothwhile for a system to spend an amount of time on selection for a good strategy

## 16.1 Overview 