# Chapter 8: Complex Data Types
- key requirement for relational model is that the value of an attribute must be atomic (indivisible), but in real world, there are many data that are not atomic, for ex: phone number, address, name, etc.
- to represent those data, we can use complex data type, which allow us to store multiple value in a single attribute, or to store structured data in a single attribute

## 8.1 Semi-structured Data

#### 8.1.1.1 Flexible Schema
- Some database allow tuple to have different set of attribute (called wide coulumn)
- a more stricted way to archieve flexible is to have fixed but larger number of attribute, and allow null value for attribute that are not used in a particular tuple

#### 8.1.1.2 Multivalued Data Types
- for specialized support for arrays(spatial, raster data), postgres need extension like PostGIS

#### 8.1.1.3 Nested Data Types

#### 8.1.1.4 Knowledge Representation
- we use RDF (Resource Description Format) to represent knowledge in a graph form

### 8.1.2 JSON
- **BSON (Binary JSON)** is a binary representation of JSON, it is more efficient to store and query than JSON, but it is not human readable
- use `json_build_object` to create json object from key value pairs (key1, value1, key2, value2...)
- use `jso_agg` to aggregate, create a single json object from multiple json object
- `obj -> key`: get the value of the key in the json object
- `obj ->> key`: get the value of the key in the json object as text
- `obj -> index`: get the value of the index in the json array
- `obj #>> {key1, key2...}`: get the value of the nested key in the json object: key1.key2

### 8.1.4 RDF and Knowledge Graphs

#### 8.1.4.1 Triple Representation
- **RDF (Resource Description Framework)** represent knowledge in a graph form, it use triple to represent the relationship between entities, a triple form is `(subject, predicate, object)`, for ex: (Alice, knows, Bob)
- we use the keyword "resource" to indicate any thing in the real world that can be identified, it can be a person, a place, a thing, an idea...
- Uniform Resource Identifier only focus on identifying the resource, to connect a resource to another resource we also care about network system
- RDF only support binary relationship

#### 8.1.4.2 Graph Representation of RDF
- in graph, entities and attribute value are represented as node, and relationship and atribute names is represented as edge
- object show as oval, attribute value show as rectangle
- represent of RDF graph model is called knowledge graph
- while erd focus on storing and retrieving data efficiently, knowledge graph focus on connecting and semantic of data through semantic labels

#### 8.1.4.4 Representing N-ary Relationships
- in RDF predicate is not only a attribute name, it also a resource that can have its own attribute, semantic and relationship. By that its help reasoning

## 8.2 Object Orientation
- There are three main way to approach integrating object-oriented with the database:
    + **Object-Relational Mapping (ORM):** a technique to map object-oriented data model to relational data model
    + **Object-Oriented Database (OODB):** a database that support object-oriented data model natively (due to its complexity and lack of standardization, OODB is not widely used in practice)
    + **Object-Relational Database (ORDB):** a database that support both object-oriented and relational

### 8.2.1 Object-Relational Database Systems

#### 8.2.1.3 Type Inheritance
- **Type and Table Schema** is similar in concept, but different in level of abstraction
- Table Schema ~ Composite Type + Metadata + storage 
- In postgres dont support type inheritance, but we can use table inheritance to achieve similar effect
- about table inheritance in postgres, is is a "Structural Inheritance", which mean that the child table inherit the structure of the parent table, but not the data:
    + when query a parent table, it will also return the tuple in the child table (use "only" keyword to query only the parent table)
    + child table inherit all column from parent table, but it can also have its own column
    + check constraint, not null is inherited by child table
    + "default" is inherited by child table, but it can be override by child table
    + select, update, delete on parent table will also affect the child table
    + primary key, foreign key, unique constraint, index, trigger, rule, sequence, serial insert  is not inherited by child table
- if not effect global state (only on record, structural level) -> inherit, if effect global state -> not inherit
- in case of update, delete: careful with the cascade effect, because it can affect the child table, which can cause data inconsistency if not handle properly

#### 8.2.1.4 Reference Types in SQL

### 8.2.2 Object-Relational Mapping (ORM)

## 8.3 Textual Data
- **information retrieval:** the process of retrieving relevant information from a large collection of data, it is different from data retrieval because it focus on relevance rather than exact match
- **document:** a logical unit of information that can be retrieved, it can be a web page, a news article, a research paper, etc.

### 8.3.1 Keyword Queries
- in simplest form, an **IR (Information Retrieval)** use keyword to retrieve document which contain the keyword
- more advanced system estimate the relevance of a document to a query using occurences, hyperlink information, etc.
- keyword-based can also use with other type of data, such as image, video, audio, etc., that have descriptive keyword associated with them
- web search engine at core is an IR system that use keyword. To retrieve and store web page, it use web crawler to crawl the web and index the web page based on the keyword in the web page

### 8.3.2 Relevance Ranking

#### 8.3.2.1 Ranking Using TF-IDF
- **term:** is the result of tokenization, normalization
- term is a similar concept to keyword:
    + keyword focus on semantic of the word, use by human
    + term focus on transform keyword into a form that can be processed by computer, use by IR system by processof tokenization -> normalization -> stemming, lemmatization...
- **TF(d, t):** the relevance of a term t to a document d
- a way to measure TF(d, t): `log(1 + (n(d,t) / n(d)))`
    + `n(d)`: the total number of term in document d
    + `n(d,t)`: the number of term t in document d
    + `+1` is added for smoothing to ensure in (0,1]
    + `log` is used to dampen the effect of term frequency, because the relevance of a term does not increase linearly with its frequency in a document
- some term are more important than other term
- **IDF(t) = 1 / n(t):** `n(t)` is the number of document that contain term t (the higher, the less important the term is)

$$r(d, Q) = \sum_{t \in Q} TF(d,t) \times IDF(t)$$

- $r(d, Q)$: the relevance of a document d to a query Q
- $\sum_{t \in Q}$: the sum of the relevance of each term in the query Q to the document d
- if user can specify the importance of each term in the query, we can use $r(d, Q) = \sum_{t \in Q} w(t) \times TF(d,t) \times IDF(t)$, where w(t) is the weight of term t in the query Q
- combine TF and IDF we get **TF-IDF**
- **stop-word:** a term that is very common and does not carry much meaning, ex: "the", "is", "and", etc. We can remove them from the document and query to improve the relevance of the search result
- we also take account of proximity of the term in the document, because the closer the term is to each other, the more relevant they are to each other

#### 8.3.2.2 Ranking Using Hyperlinks
- the more hyperlink point to a document, the more relevant it is
- page that are pointed from a high rank page are more relevant than page that are pointed from a low rank page

$$\frac{\text{delta}}{N} + (1-\text{delta}) \times \sum_{i=1..N} T[i,j] \times P[i]$$

- `T[i,j]`: the probability of following the hyperlink from page i to page j.
- `T[i,j] = 1 / Ni`, where Ni is the number of hyperlink in page i
- $\sum_{i=1..N}$: for each page i that point to page j, we sum the contribution of page i to the relevance of page j
- `delta/N`: the probability of randomly jump to any page
- `(1-delta)`: the probability of following the hyperlink

### 8.3.3 Measuring Retrieval Effectiveness
- **precision:** the fraction of retrieved document that are relevant to the query
- **recall:** the fraction of relevant document that are retrieved by the query
- we use @k to denote the precision and recall at k, which is the precision and recall of the top k retrieved document

### 8.3.4 Keyword Querying on Structured Data and Knowledge Graphs

## 8.4 Spatial Data
- **geographic Information System (GIS):** a system that capture, store, manipulate, analyze, manage, and present spatial or geographic data
- **geometric data type:** a data type that represent geometric object, such as point, line, polygon, etc.

### 8.4.1 Representation of Geometric Information
- **line segment:** a line segment is defined by two point (x1, y1) and (x2, y2)
- **polyline:** a polyline is defined by a sequence of point (x1, y1), (x2, y2), ..., (xn, yn)
- **circular arc:** a circular arc is defined by three point factor: the center of the circle (cx, cy), the radius r, and the start and end angle (theta1, theta2) -> more optimal than using many polyline to approximate a circular arc
- **polygon:** a polygon is defined by a sequence of point (x1, y1),
    + polygon can devided into triangle (triangulation)

#### 8.4.3.2 Representation of Geographic Data
- **raster data:** a raster data is a grid of cell, each cell have a value that represent the attribute of the area covered by the cell (for ex: elevation, temperature, land use, etc.)
- **vector data:** a vector data is a collection of geometric object, such as point, line, polygon, etc. that represent the geographic feature (for ex: road, river, building, etc.)
- **topological data:** represent information in form of raster or vector data
    + it can also be triangulated to represent the geographic feature, called Triangulated Irregular Network (TIN)

### 8.4.4 Spatial Queries
- **spatial query:** a query that involve spatial data, it can be used to retrieve, manipulate, and analyze spatial data
- spatial query can be classified into: 
    + region queries
    + nearest queries
    + nearest neighbor queries
    + spatial graph queries
    + spatial join