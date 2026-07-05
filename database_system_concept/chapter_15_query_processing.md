# Part 6: Query Processing and Optimization
# Chapter 15: Query Processing
- `Query processing` includes these activities:
    - parse and tranlation of queries from high level languages to expression (ussually an extention of relational-algebra) that can be used at physical level of file system
    - query optimizing transformations
    - query evaluation

## 15.1 Overview
![](../images/query_processing.png)
- for a query, there many ways to represent them as relational-algebra expression as well as execute and query them
- take example on the followwing query
```sql
select salary
from instructor
where salary < 75000;
```
- the query can be express as
  - $\Pi_{salary}(\sigma_{salary < 75000}(instructor))$
  - $\sigma_{salary<75000}(\Pi_{salary}(instructor))$
- after tranformation, query can be excuted in serveral ways. If there are $B^+$ index on the salary we'll use them for faster execution, otherwise we'll scan each tuple to find out if there any tuple math the predicate
- to specify how to evaluate query, not only we look at relational-algebra expression but we also look at `annotation`, an instruction to specify which algorithm is use, specific operation or which indexes is use
- `evaluation primitive` a evaluation primitive is a basic physical (or execution) unit that cannot be further divided during query planning; it specifies a relational-algebra operation annotated with instructions on how to evaluate it.
- a sequence of evaluation primitive called a `query-execution plan` or `query-evaluation plan`
- The `query-execution engine` takes a query-evaluation plan, executes
that plan, and returns the answers to the query.
![](../images/query_evaluation_plan.png)
> a query evaluation plan with "index 1" being used
- a different evaluation plan can have different `cost` and its is a job of system, not of the user to construct a query-evaluation plan that minimize the cost. This task called `query optimization`
- not all database follow exactly the same step, for example, some database use parse-tree representation instead of relational-algebra expression
- query optimization need to know the cost of each operation, although cost can depend on outside parameter such as memory available, a rough estimation can still get to calculated
- `pipeline` a evaluation where operations are grouped, instead of awaiting operation to fully completed before execute the next one, operation perform `overlapping execution` as consumer operation execute as the producer operation output tuples or block of tuples

## 15.2 Measure of Query Cost
- the (estimated) cost of query evaluation can be measures interm of these following resource:
  - number of disk I/O
  - CPU time to execute a query
  - cost of cimmunication (in distributed system)
- for database use magnetic disk, I/O cost from disk ussually dominate other costs
- for database use SSD, we must include CPU costs when computing cost of query evalation
- CPU costs includes: cpu_tuple_cost, cpu_index_tuple_cost, cpu_operator_cost
- database have default value for each of these cost and can be changed as configuration params
- the total time to I/O for in query evaluation can be calculated as follow
  $$ \text{Total Time} = b * t_T + S * t_S $$
  - $b$ number of blocks transfer
  - $t_T$ average time to transer a block of data
  - $S$ number of random I/O access
  - $t_S$ average block access-time (disk seek + rotational latency)
  
  |$(4\ kb/\text{block})$|HDD|SSD (SATA)|SSD (PCIe)|Main Memory|
  |-|-|-|-|-|
  |$t_T$|$0.1\ ms$|$10\ \mu s$|$2\ \mu s$|$<1\ \mu s$|
  |$t_S$|$4\ ms$|$90\ \mu s$|$20-60\ \mu s$|$<100\ ns$|
  |transfer rate|$40\ mb/s$|$400\ mb/s$|$2\ gb/s$|$>4\ gb/s$|
  |block reads per second|$250$|$10,000$|$50,000-15,000$|$>10,000,000$|
  > base on data of 2018 model
- database ideally perform a test to estimate $t_T$ and $t_S$ for specific system/storage device as part of software installation
- alternatively, user can input sprcific number in configration files
- to further refine cost estimation we can separate the cost of write and read operation
  - on HDD, write usually double the cost of read since disk read the sector back after write to verify write was succesfull
  - on SSD PCIe write through put is 50% less than read throughput due to physical limited
  - this different is masked with SATA interface because of the limited speed
  - the cost of wrting data back to disk when disk spill or materialized temp table,.. are not take into account when cost estimation and are calculated separately when required
- the cost of all algorithms depend on the size of the buffer in main memory, cost estimate use the ammount of memory available to an operator, $M$ as a parameter `work_mem`
- total memory available query called `effective_cache_size`, in postgres, this value is default to $4\ gb$
- the estimated cost can less than the actual one since block may already present in the memory buffer
- to calculate the cost of memory buffer, the cost of random page read is set to $1/10$ of the actual random page read to simulate the situation where 90% read are resident in cache
- most database also assume that all of the internal node of $B^+$ tree are presented in memory buffer
- `response time` (wall-clock time) is the actual time require to evaluate the query and could be use to measure the cost of the plan but its hard to estimate withour actually execute the query for reasons
  - the actual evaluation time repend on content of memory buffer at the time the query is actually evaluated
  - database are abstracted from actual hardware device and hard to estimate withour the detail knowledge of data layout
  - a plan may get better response time at the cost of extra resource compsumtion (parallel processing,...)
- instead of try to minimize response time, the optimizer generally try to minimize `resource consumption` of the query plan 
  
## 15.3 Selection Operation
- `file scan` is the lowest-level operation to access data
- file scans are ***search algorithm***, use to retrieve records which fulfill a selection condition
  
### 15.3.1 Selections Using File Scans and Indices
- **A1** (`linear search`): perform an initial seek and start scanning each block  in a file and test records in block if it match the selection condition
  - in case blocks of the file are not store sequentially, extra seek is required but we ignore it for simplicity
  - linear search algorithm can be applied to any file for implementing selection, regardless of file's ordering, index availability or selection operation's nature
  - other algorithms may not applicable in all cases
  - linear search generally slower than other algorithms
- cost estimate for linear scan and other algorithm are shown below

||Algorithm|Cost|Reason|
|-|-|-|-|
|A1|Linear Search|$t_S + b_r * t_T$|initial seek plus number of blocks in the file multiply by time to transfer a block|
|A1|Linear Search, Equality on Key|Average case $t_S + (\frac{b_r}{2}) * t_T$|scan can be terminated ass soon as required record is found, but in the worst case, $b_r$ block transfer is still required|
|A2|Clustering $B^+$ tree Index, Equality on Key|$(h_i + 1) * (t_T + t_S)$|$h_i$ denote the height of index, plus $1$ I/O to fetch the record, each of the I/O operation require a block seek + block transfer|
|A3|Clustering $B^+$ tree Index, Equality on Non-key|$h_i * (t_T + t_S) + t_S + b * t_T$|each levle of tree require a block seek and transfer, at leaf node level, we 1 more seek to find the first block, then sequentially transfer block since its a clustering index|
|A4|Secondary $B^+$ tree Index, Equality on Key|$(h_i + 1) * (t_T + t_S)$|similar to clustering index|
|A4|Secondary $B^+$ tree Index, Equality on Non-key|$(h_i + n) * (t_T + t_S)$|the part for tree height is similar to other algorithms, for each entry found, a seak and transfer is required since each entry is may located in different block|
|A5|Clustering $B^+$ tree Index, Comparison|$h_i * (t_T + t_S) + t_S + b * t_T$|similar to A3, equality on non-key|
|A6|Secondary $B^+$ tree Index, Comparison|$(h_i + n) * (t_T + t_S)$|similar to A4, equality on non-key|

- $h_i$ represent the height of the $B^+$ tree, most optimizer assume that all non-leaf node is already in the buffer since ussually, less than 1% of the nodes in the $B^+$ tree are non-leaf node 
- this assumption can be simplify by setting $h_i = 1$
- index structures are referred to as `access paths` beucause it provide a path (root $\rightarrow$ internal node $\rightarrow$ leaf) to locate and access the data
- search algorithms which use index are called `index scans`
- we use `selection predicate` for choosing which index to be used in query evaluation
- here aresome index scans:
  - A2 (clustering index, equality on key)
  - A3 (clustering index, equality on non-key)
  - A4 (secondary index, equality): 
    - in the worst case, each record needed to be fetched rely on different block and large number of record need to be fetched
    - if the in-memory buffer is large, we could expected estimation to also take account of a probability of the block containing record already in the buffer
    - the larger the buffer, the more cost are reduced
- in algorithm A2, using $B^+$ tree file organization can reduce the I/O access by 1 since records are store at leaf level
- this approach affect how secondary store and retrieving data
- secondary indices now store the search key values using in $B^+$ tree file organization and use it to find the records
- since it's more complex, we must modified the cost formulae for secondary indices approriatedly if such indices are used

### 15.3.2 Selections Involving Comparisons
- selection with the form $\sigma_{A \le v}(r)$ we can implement the selection either by by using linear search or using indices: 
  - A5 (clusteting index, comparison) for example, a clustering order index on attribute A:
    - when a selection inform of $A > v$ or $A \ge v$, this case is identical to A3, we start seaching for first tuple that match the condition, then linear scan consecutive block until value A dont match the above condition or end of file
    - when a selection inform of $A < v$ or $A \le v$, the index look up is not required, instead it scan file from the start to select out the tuple that met the predicate, it stop when met a tuple that dont match the condition or end of file
  - A6 (secondary index, comparison) 
    - we use the secondary index to find the lowest-level index block, either it contain smallest value ($<$, $\le$) or maximum value ($>$, $\ge$)
    - if there are meny record to be fetched, index scan might be more expensive than linear scan, therefore secondary index shoul be used if only a few record are selected
- to list out an *selectivity estimation* or *cardinality estimation* number of record that will be selected, database use *system catalog* and *statistics technique*
- when mathching tuples can not known accurately at compile time, this could lead query optimizer choosing an un-optimize access path $\rightarrow$ higher cost and bad performance
- `bitmap index scan` is a *hybrid algorithm* which postgres using to deal with above situation by doing the following
  - the bitmap index scan create a bitmap with number of bits as the number of blocks in the relation, all bit are set to 0
  - it then using the secondary index to indentify which tuple will be fetched
  - instead of immediately fetching it, we get the block number from the index entry and sets the corresponding bit in the bit map to 1
  - once all index entries are processed, the relation is then scanned linealy, if the correspond block is set to 1 in bit map $\rightarrow$ get fetched, if block is set to 0 in bitmap $\rightarrow$ skipped
- in the worst case (all block in the files need to be fetched + index bitmap overhead), bitmap index scan is slightly more expensive than linear scan and secondary index scan but in the best case, it is much cheaper
- a varient of algorithm is instead of using bitmap, we first:
  -  fetch all the entry of all record need to be fetched
  -  sort them by block id
  -  perform a linear scan and skip block which do not having matching entries
- using bitmap can be cheaper than above varient due to memory efficientcy in bitmap

### 15.3.3 Implementation of Complex Selections
- we continue consider more complex selection predicates:
  - `Conjunction`: a selection with the form:
  $$\sigma_{\theta_1 \land \theta_2 \land ... \land \theta_n}(r)$$
  
  - `Disjunction`: a selection with the form:
  $$\sigma_{\theta_1 \lor \theta_2 \lor ... \lor \theta_n}(r)$$
    - disjunction selection is simply a union of all $n$ individual set of records, which satisfy the predicate $\theta_1 ... \theta_n$ respectively
  
  - `Negation`: without $null$, the result of negation selection $\sigma_{\neg \theta}(r)$ is simply $r - \sigma_{\theta}(r)$

- A7 (conjunctive selection using one index)
  - to reduce cost, query optimizer choose algorithm from A1 to A6 and one of the simple condition $\theta_i$ which have the least cost for $\sigma_{\theta_i}(r)$
  - if the selected algorithm is using access path (A2...A6), the algorithm A7 start to evaluate by fetching the records and test where the retrieved record are satisfied the remaining simple conditions in the memory buffer
  - the cost of A7 is the cost of the choosen algorithm cost
- A8 (conjunctive selection using composite index): if selection specifies an equality on two or more attributes and a composite index exists on those attributes then it can classified as A8
  - the cost of A8 will be determined by which algorithm A2, A3, A4 will be used
- A9 (conjunctive selection by intersection of identifiers)
  - the algorithm require index for individual condition 
  - for each individual condition that have index, we scan it's index to retrieve set of pointers
  - we then intersect all set of pointers to create a new set of pointers
  - if indices are not available for all individual conditions, we use the result set to filter out the remaining conditions
  - the cost of A9 is sum of all indices scan plus the cost of retrieving record after inserting
  - the cost can reduce by sorting the pointer by block number 
    - all pointers in the same block will consecutively placed
    - reduce the cost of random I/O, block are now sequentially I/O
- A10 (disjunctive selection by union of identifiers): if access path are available on ***all*** individual condition of a disjuction selection we then
  - for each individual condition, we scan it's index and *yield* pointers back to parent operator (union), often use with bitmap (continue bitmap indices or)
  - union all retrieved pointers to create a new set of pointers
  - we then use the pointers to retrieve the actual record
- if 1 of the condition dont have access path, the have to perform linear scan of the relation to find tuples that match the condition $\rightarrow$ the index scan on other index will become meaningless

## 15.4 Sorting
Quick-sort (In-Memory) $\rightarrow$ External Merge Sort (Disk-Optimized) $\rightarrow$ Replacement Selection (I/O Pipelined) $\rightarrow$ Cache-Conscious/Radix Sort (CPU/Cache-Optimized)
- Sorting is important in database system for two reasons:
  - SQL queries can specify output to be sorted
  - several relational operation can be implemented efficiently if the input are sorted
- we can sort a relation logically by building an index on the sort key but reading the physical record need the random I/O which can be costly
- we desire to order records physically 
- in case of small relation that can fit in memory, we can use standard technique such as quick-sort
- for large relation that cannot fit in memory, we can use `external sort-merge` which is a disk-optimized sorting algorithm

### 15.4.1 External Sort-Merge Algorithm
- sorting a relation that is too large to fit in main memory is called `external sorting`
- most common external sorting algorithm is `external sort-merge` algorithm
- to describe the algorithm, we denote $M$ as the number of blocks that can fit in main memory buffer available for sorting
  1) in ther first stage `runs` are created
    - runs (chuỗi liên tục không đứt quãng) are sorted subset of the relation, each of size at most $M$ blocks
    ```
    i = 0;
    while (end of relation)
      read next M blocks of relation into memory buffer;
      sort the M blocks in memory buffer;
      write the sorted M blocks back to disk as run file Ri;
      i = i + 1;
    end
    ```
  2) in the second stage, with the precondition that total number of runs $N$ is less than $M$ (to hold one atleast oneblock for output) runs are *merged* as follows
  ```
  for each runs Ri of N run files
    read the first block of Ri into memory buffer;
  end
  while (not all runs are empty)
    take the first tuple (sorted order) among all block in memory buffer; 
    write the tuple to output buffer and remove it from the block;
    if any buffer block of Ri is empty and not end of file Ri
      read the next block of Ri into memory buffer;
    end
  end
   ``` 
- the output file is written in block as buffer block is full to reduce the number of I/O operation
- the standard in-memory sort-merge algorithm called `two-way merge sort`, meaning that it merge 2 runs at a time
- the generalization of two-way merge sort called `N-way merge sort` is the one the we described above
- to sort-merge algorithm to work, it is base on transitivity (bắc cầu) of the order relation, we sort a run, get the first block of each run, then compare the first tuple of each block to find the smallest one
- the choosen tuple is the smallest tuple among smallest block among all runs, which is the smallest tuple among all runs
- in case the relation is too large to fit in memory, we can use `multi-pass merge sort`. Merge operation is performed in multiple passes, each pass merge $M-1$ runs at a time to create new runs :
  1) in the first pass, we read $M$ blocks of the relation into memory buffer, sort them and write them back to disk as runs
  2) if the number of runs is greater than $M-1$, we perform a second pass to merge $M-1$ runs at a time to create new bigger runs, the number of new runs is reduced by the factor $M - 1$ 
  3) we repeat the process until the number of runs is less than $M-1$, then we perform a final pass to merge all runs into a single sorted file

![](../images/external_sort_merge.png)
> visualization of multi-passes external sort-merge algorithm
>> with one tuple fit in a block ($f_r = 1$) and memory hold up to 3 block ($M = 3$) for simplicity

### 15.4.2 Cost Analysis of External Sort-Merge