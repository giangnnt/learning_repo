# Part 2: Database Design
## 6.1 Overview of the Design Process
### 6.1.1 Design Phases
- **initial phase:** interact wih domain expert to understand the requirement and the domain. 
    + the outcome is specification of user requirement
- **conceptual design phase:** translate user requirement into conceptual schema. It should provide a detailed overview of the enterprise (detail of bussiness, overview of the technical aspect)
    + entity relationship model should contain the following: entites, their relationship and constraints 
    + the outcome is ERD (Entity Relationship Diagram): a graphical representation of the conceptual schema
    + focus on describing the data and its relationship, not on how to implement it in database
- **specification of functional requirement phase:** describe kinds of operation (transaction) that will be performed. Ensure the schema can support the required operation on data
- **logical design phase:** translate the conceptual schema into implementation data model (a genetal logical structure, use to talk to computer, to manipulate data by sql statement)
    + relational model is the most common logical design: it use relation (table) to represent entity and relationship, and use attribute (column) to represent property of entity and relationship

### 6.1.2 Design Alternative
- in design database schema, we must avoid two pitfalls:
    + **redundancy:** the same information is stored in more than one place, which can lead to inconsistency and update anomaly
    + **incompleteness:** the schema does not capture all the necessary information, which can lead to loss of information and inability to support required operation

### 6.2.1 Entity Sets
- **entity:** a thing or object in the real world that is distinguishable from other objects (for ex: student a, student b, course a, course b...)
- **entity set:** a collection of entity (abstract, not refer to paticular entities) that share the same attributes (for ex: student, course...)
- **extension of an entity set:** actual collection of entites belonging to the entity set
- **enity set:** represent by a rectangle, with the name of the entity set inside

### 6.2.2 Relationship Sets
- **relationship set:** represent by a diamond, with the name of the relationship set inside
- **relationship:** an association among two or more entities (for ex: student a take course a, student b take course b...)
- **relationship set:** a collection of relationship (abstract, not refer to particular relationship) that share
- **relationship instance:** actual collection of relationship belonging to the relationship set
- theoretically, relationship set is a subset of the cartesian product of the entity set that participate in the relationship
- **recursive relationship set:** a entity paticipate in the relationship more than once (for ex: employee a is the manager of employee b, employee b is the manager of employee c...). In this case we need to use role name to distinguish the different participation of the same entity set in the relationship
- **descriptive attribute:** an attribute that describe the relationship itself, not the participating entity (for ex: grade in the relationship "take" between student and course)
- attribute of relationship set is represented by an undevided rectangle, link with a dashed line to the relationship set
- attrubutes of entity should only appear on the first time the entity appear in the ERD
- it is possible for the same entities sets to have more than one relationship set
- problems: if we wish students can have multiple grade for the same course, we can consider to have grade as multivalued attribute 
- **binary relationship set:** relationship set that involve only two entity set
- **ternary relationship set:** relationship set that involve three entity set
- **value set or domain:** is the set of possible value for an attribute, it can be specified by a set of value or by a data type (for ex: integer, string...)

- attributes can be categorized into:
    + **simple attribute:** an attribute that cannot be divided into smaller subpart (for ex: name, age...)
    + **composite attribute:** an attribute that can be divided into smaller subpart (for ex: name can be divided into first name and last name)
    + **single-valued attribute:** an attribute that have a single value for a particular entity
    + **multi-valued attribute:** entity can have multiple value for an attribures(for ex: having multiple phone numbers) by using data type like array
    + **derived attributes:** the value of this attribute type can be calculated from other attribute, (for ex: age can be calculated from birth year and current year)

## 6.4 Mapping Cardinalities
- **mapping cadinalities:** express the number of entities to which another entity can be associated in a relationship set
- there are: one to one, one to many, many to one, many to many
    + undirected line between two entity set indicate this relation is many to many
    + directed line from a to b (a --> b) indicate b can have many a but a can only have one b

- **participation of E in relationship set R** is said to be total if every entity in E in at least one relationship R (for ex: every employee must be under one dept)
- it is said to be partial if it is not necessary for every entity in E to paticipate in R 
- we use double line to indicate total, and single line for partial
- for further more complex constraint about participation's in the relationship set we use form l..h. where l is the minimum and h is the maximum carnality

## 6.5 Primary Key
### 6.5.1 Entity Sets
### 6.5.2 Relationship Sets
- **Many to many relationship set:** primary key is the combination of the primary key of the participating entity set
- **One to many relationship set:** primary key of the many side is the minimal primary key 
- **weak entity set:** an entity set that its primary key is partially or totally derived from the primary key of another entity set (for ex: section entity set, its primary key is derived from the course entity set and the section number)
- **One to one relationship:** primary key can be either the primary key of one of the participating entity set
- **Non-binary relationship set:**
    + no cardinality constraint: primary key is the combination of the primary key of all participating entity set
    + with cardinality constraint: (depenend on the constraint and semantic of the relationship set)

- **ERD for weak entity set** follow these rules:
    + weak entity set is represented by a double rectangle
    + discriminator is underlined dashed
    + identifying relationship set is represented by a double diamond
    + double line from weak entity set to identifying relationship set
    + arrow from identifying relationship set to the strong entity set

## 6.6 Removing Redudant Attributes in Entity Sets
- we should not have add attribute that represent relatioship while doing ERD, we should focus on describing the data and its attribute to avoid assumption and redundancy

## 6.7 Reducing E-R Diagrams to Relational Schemas
### 6.7.1 Representation of Strong Entity Sets
### 6.7.2 Representation of Strong Entity Sets with Complex Attributes
- **flattening composite attribute** by separate it into simple attribute
- **flattening multi-valued attribute** by creating a new entity set and a relationship set between the new entity set and the original entity set

### 6.7.3 Representation of Weak Entity Sets
- combine the primary key of the strong entity set and the discriminator to form the primary key of the weak entity set

### 6.7.4 Representation of Relationship Sets
- for many to many relationship set, we create a new relation with the primary key is the combination of the primary key of the participating entity set
- for one to many relationship set, the primary key is the many side, and we add a foreign key to each entity participate in the relationship

### 6.7.6 Combination of Schemas
- representation of relationship set between strong and weak entity is not necessary