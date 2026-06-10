# Chapter 11: Data Analytics
- **dataware house** might collect data from multiple source, and store it in a central repository, under unified schema which efficient for analysis and reporting
- datawarehouse facing with
    + non-relational data: today warehouse also collect and store data from non-relational source 
    + data source error: data source have error data can be detected
    + data duplication: data source may have duplicate data since they are collected from multiple source

- **online analytical processing (OLAP):** a category of software tools that provide analysis of data stored in a data warehouse, it allow users to perform complex queries and analysis on the data, such as aggregation, slicing and dicing in small amount of time
- **decision support:** use SQL to query large amounts of data to support decision making, it is different from transaction processing because it focus on analysis rather than data manipulation

## 11.2 Data Warehousing

### 11.2.1 Component of Data Warehouse
- there are issue when building a warehouse
- **when and how to gather data:** 
    + **source-driven architecture:** data is send to the warehouse continuously or periodically
    + **destination-driven architecture:** warehouse will send request to the source to gather data
    + usually "yesterday's data" is good enough for analysis, but for some case we also need for real-time data, which use stream processing to gather data
- **what schema to use:** data schema in the warehouse is different from the schema in the source, can be thought as a "view" of the source schema
- **data transformation and cleaning:** 
    + **data cleansing, fuzzy lookup:** the process of detecting and correcting, removing inconsistent, inaccurate, incomplete, or irrelevant data from the data source
    + **deduplication, merge-purge operation:** the process of identifying and removing duplicate data from the data source
    + **householding:** a technique to group together records that belong to the same entity (for ex: grouping individuals have the same address into a household)
    + **transformed:** a process of converting data from one format or structure to another, it can be used to standardize data from different source
- **how to propagate updates:** update in the source can be propagated to the warehouse in different way, we focus on "view-maintenance" problem
- **what data to summarize:** instead of storing all data in the warehouse, we can also store aggregation of the data 

- modern generation of data warehouse support user defined function, mapreduce so the step are **ELT (Extract, Load, Transform)** instead of **ETL (Extract, Transform, Load)**

### 11.2.2 Multidimensional Data and Warehouse Schemas
- **fact table:** a table record information about a individual event
    + **measure attribute:** a attribute that contain the value of the event (ex: quantity, price, etc.)
    + **dimension attribute:** a attribute that contain the context of the event (ex: time, location, product, etc.)
- **dimension table:** a table that contain the values of the dimension attributes

- schema with fact table reference to dimension table is called **star schema**
- if dimension table reference to another dimension table (giving layer detail), it is called **snowflake schema**

### 11.2.3 Database Support for Data Warehouses
- key difference between data warehouse (OLAP) and traditional database system (OLTP) is that OLTP focus on transaction processing, involve updates in addition to reads, OLAP need to process fewer queries but larger amount of data
- reasons for data warehouse can perform better at read and store large amount of data:
    + never update data: only insert and delete, data never updated once in the warehouse 
    + no concurrency control: since there is no update, there is no need for concurrency control, which can significantly improve the performance of the warehouse

- **row-oriented storage:** store all attribute of a tuple together and tuples are stored sequentially in a file
    + its provide fast insert and update since all attribute of a tuple are stored together, but it is not efficient for read because irrelevant attribute of a tuple are also fetch, cause waiting cache space and memory
- **column-oriented storage:** each attribute of a relation is stored in a separate file, attribute of a tuple i can be read by (i-1) * size of attribute i (m), we read m space for each attribute
    + it provide fast read since require attribute all store in the same place
    + storing value of the same type together increase the compression -> reduce the storage space and improve the performance of read
    + suitable for aggregate query

### 11.2.4 Data Lakes
- **data lake:** a repository that allow to store unstructured data, it do not require up-front effort but it require more effort to creating queries

## 11.3 Online Analytical Processing 

### 11.3.1 Aggregation on Multidimensional Data
- **cross-tabulation(cross-tab):** a table that show the aggregation of data based on two or more dimension attribute, it is also called pivot table
- **data cube:** is a generalization of cross-tab which can present n-dimension 
    + each cell in data cube is identified by values for three dimension
- with a set of n dimensions, we can have 2^n subset
- **pivoting:** the process of transforming a cross-tab to another cross-tab by swapping the row and column dimension
- **slicing, dicing:** the process of selecting a subset of the data cube by fixing the value of one or more dimension
- when cross tab is use for viewing multidimensional cube, the value that are not show in the cross-tab is show above the cross-tab. Dicing and Slicing simply is just adjusting those value

- **rollup** aggregate data by level of aggregation instead of each subset of dimension
- **drill down:** the process of going from a higher level of aggregation to a lower level of aggregation

### 11.3.2 Relational Representation of Cross-Tab
- **cube:** generate all possible combination of dimension
- **rollup:** generate combination base on the level of aggregation -> hierarchy of dimension
- **grouping sets:** generate specific combination of dimension

### 11.3.4 Reporting and Visualization Tools
- **report generators** support the creation of reports (formatted text, graphical), query database
- **data visualization** go beyond basic chart, it can help examine large data and detect pattern by combining:
    + compactly visual display 
    + encoded data with visual attributes such as position, size, color, density, etc.
    -> help create hypothesis, link data together, overview of data

- there also development of web based visualization tool, which allow to display data in multiple device and platform
- interaction is also important for data visualization

## 11.4 Data Mining
- **data mining:** the process of discovering patterns and data relationship
- **statistical analysis:** analyze data, find out if the pattern, relationship is "correct" or just a coincidence, and make inference about the data
- **machine learning:** learn from data, make prediction about the future data
- **rules:** a rule is a statement that describe a relationship between two or more attribute
- **model:** knowledge discover from machine learning, it can be used to make prediction about future data

### 11.4.1 Types of Data Mining Tasks
- **prediction:** the task is to carry out a prediction about future data base on attributes, for ex: predict the wehther a person will repay a loan...
- **descriptive pattern:** the task is to find out a pattern in the data, it sub categories are:
    + **clustering:** group of data share the same attribute value together -> help find out pattern in the data, for ex: group of civilization share the same disease
    + **association:** find out relationship between attributes, for ex: find out that people who buy bread also buy butter, new realease medicine raise a strange pattern of disease, etc.

### 11.4.2 Classification
- **classification** is a sub set of prediction, where the task is to predict a categorical value (class) base on the value of other attribute

#### 11.4.2.1 Decision-Tree Classifiers
- we create a tree, each leaf node associated with a class and internal node has a predicate (function) that test the value of an attribute

#### 11.4.2.2 Bayesian Classifiers
$$P(C|X) = \frac{P(X|C) \times P(C)}{P(X)}$$

- $P(C|X)$: the probability of instance with attribute value X belong to class C
- $P(X|C)$: the probability of instance with class C have attribute value X
- $P(C)$: the past probability of appearance of class C
- $P(X)$: the past probability of appearance of instance that have attribute value X

- in general, a instance can have multiple attribute, there for calculating all posibility of attribute value combination p(d|cj) can be very expensive
    + there for we make an assumption that all attribute are independent, which is called "**naive Bayesian classifier**"
    + $p(d|cj) = p(d1|cj) \times p(d2|cj) \times ... \times p(dn|cj)$
    + to deal with range value attribute, we devide the range into interval and calulate the probability base on the interval

### 11.4.3 Support Vector Machine Classifiers
- to classify a instance, we can find a "maximum margin line" that separate the instance from the other class
- more generally, we can find a "dividing plane" that separate the instance from the other class in higher dimension space
- using kernel function, we can also find a "non-linear curve" that separate the instance from the other class in higher dimension space

#### 11.4.2.4 Neural Network Classifiers
- **neural network** contain multiple layer of "neuron", each neuron is a function that take a weighted sum of its input and apply an activation function to produce an output
- initially, the weight of each neuron is set randomly, if the prediction of the neural network is not correct, we can adjust the weight of each neuron using a technique called "backpropagation"
- **deep learning:** a machine learning technique that use multiple layer of neural network and large amount of instance to learn

### 11.4.3 Regression
- **regression** is a sub set of prediction, where the task is to predict a continuous value base on the value of other attribute
- **linear regression:** is a regression technique that assume a linear relationship between the independent variable and the dependent variable

$$Y = a_0 + a_1 \times X_1 + a_2 \times X_2 + ... + a_n \times X_n$$

- $Y$: the dependent variable (the value we want to predict)
- $X_1, X_2, ..., X_n$: the independent variable (the attribute that we use to predict Y)
- $a_0, a_1, a_2, ..., a_n$: the coefficient that we need to learn from the data
- **curve fitting:** is a task of finding a curve that best fit the data

### 11.4.4 Association Rules
- **population:** the set of all instance in the data
- **instance:** a individual data record in the population
    * choosing population and instance is important, because it can affect the result of the analysis
    * each analysis target a specific population and instance

- **support:** the chance of having the antecedent and the consequent happen together in the population
- **confidence:** the chance of having the consequent if the antecedent happened

### 11.4.5 Clustering
- **clustering:** the task is finding clusters of points in the given data
    + clustering gather group of instance that share the same attribute value together, it help find out pattern in the data, predict the value of an instance

### 11.4.5 Text Mining
- **sentiment analysis:** the task is to determine the sentiment of a text, such as positive, negative, or neutral base on the words and phrases in the text
- **information extraction:** the task is to extract structured information from unstructured text, such as extracting entities, relationships, and events from a text
- **entity recognition:** the task is to identify and classify named entities in a text
- **disambiguation:** the task is to determine the correct meaning of a word or phrase in a text using the context of the text
- information extraction can be use for analyze customer review, extract information about a specific aspect of a product, etc.
- extracted information can be represented in a graph, called "**knowledge graph**"

# Part 5: Storage Management and Indexing