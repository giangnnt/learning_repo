# Concept Cheat Sheet: Routing & Dynamic Routing

## 1. What is Routing?
* **Concept:** The process of **finding and selecting the best path** for data packets to travel from a source device to a destination device across computer networks.
* **Analogy:** Think of it like a **road map or traffic intersection** that determines which direction a vehicle (data) needs to turn to reach the correct address (destination IP).

---

## 2. Routing Classification

| Criteria | Static Routing | Dynamic Routing |
| :--- | :--- | :--- |
| **How ​​it works** | Paths are **manually** defined by an administrator. | Routers **automatically exchange information** to learn and update paths. |
| **Flexibility** | Fixed; does not change automatically during network issues. | Automatically reroutes in the event of link failure or network congestion. |
| **Adding new networks** | Requires manual reconfiguration on each router. | Automatically updates new network ranges via protocols (BGP, OSPF). |
| **Use cases** | Small, simple networks with stable paths. | Enterprise networks, Cloud (e.g., Cloud VPN), large-scale networks. |

---

## 3. Dynamic Routing in Cloud VPN (BGP)
* **Role:** Facilitates automatic connection and synchronization of network ranges between **On-premises** (or AWS) and **Google Cloud VPC**.
* **Protocol used:** **BGP (Border Gateway Protocol)** running via **Cloud Router**.
* **Characteristics:** Mandatory for **HA VPN (99.99% SLA)**; uses Link-local IP ranges (`169.254.0.0/16`) to establish BGP sessions.

# Definition: Border Gateway Protocol (BGP)

## 1. What is BGP?
* **Definition:** Border Gateway Protocol (BGP) is the standard routing protocol of the Internet, designed to exchange routing and reachability information between different **Autonomous Systems (AS)**.
* **Analogy:** It functions as the **Global GPS for the Internet**, directing data traffic across independent network domains managed by ISPs, tech giants, and large enterprises.

---

## 2. Key Concepts

* **Autonomous System (AS):** A collection of connected IP routing prefixes under the control of a single administrative entity (e.g., Google, AWS, Viettel, AT&T) identified by a unique **AS Number (ASN)**.
* **eBGP (External BGP):** Used to exchange routing information **between different Autonomous Systems** (e.g., On-premises network to Google Cloud).
* **iBGP (Internal BGP):** Used to distribute BGP routing information **within the same Autonomous System**.

---

## 3. How BGP Works in Cloud Environments

* **Dynamic Network Discovery:** Routers establish a **BGP Session** (using link-local IP addresses) to automatically advertise and update available IP address prefixes.
* **Zero Downtime Updates:** Automatically detects and propagates network topology changes (e.g., adding a new subnet) without needing manual tunnel reconfigurations.
* **Failover & Redundancy:** Dynamically redirects traffic through active backup tunnels if a primary network path or VPN link fails.