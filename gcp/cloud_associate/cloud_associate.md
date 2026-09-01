# VCP
- `Region` A geographical area that contains at least 3 zones, where Google hosts its infrastructure.
- `Zone` A deployment area within a region, representing one or a cluster of data centers with independent power, cooling, and networking
- `Point of Presence (PoP)` is a physical location/edge facility where a cloud provider's global network connects to external networks, such as the public internet or private enterprise networks
- `Project` is the fundamental organizational entity in Google Cloud. It holds all your cloud resources, isolates them logically, and acts as the boundary for IAM access management and billing.
- `VPC Network` is a global, logically isolated virtual network construct on Google Cloud that holds subnets across multiple regions, defining the overall networking environment, routes, and firewall rules for your cloud resources, VPC itself dont have an IP address
- `Subnet` is a logical network segment within a VPC network, belonging to a region.
  - subnet can not have overlapping IP ranges
- `Internal IP Address`
    - Mandatory: Every VM instance created on GCP is required to have an internal IP address.
    - Used for internal communication: Used for VM instances to communicate with each other within the same VPC network (even across different subnets or regions) with high speed and security.
    - Scope: Operates strictly within the VPC network and cannot be directly accessed from the public internet.
    - There is a DNS server running on each VM that can be used to resolve internal hostnames within the VPC.
    - We can assign a range of internal IP addresses to a single VM instance, this is useful when we run multiple services on the same VM
- `External IP Address`
    - Optional: A VM instance may or may not have an external IP address.
    - Used for internet access: Required if the VM needs to receive direct inbound traffic from the internet (e.g., serving a public web server) or initiate outbound connections without Cloud NAT.
    - 1:1 NAT Mapping: On GCP, the external IP is managed by a 1:1 NAT mechanism at Google's infrastructure layer and mapped to the VM's internal IP, rather than being assigned directly to the VM's OS network interface (vNIC).
- for a packet to reach a VM instance, it must go through `Route` and must have a matching `Firewall rule`
- `Route` can be thought as traffic sign which indicate the `destination` and the `next hop`
- only VM that have the same `tag` with the route can see and go through it
- `Firewall` is a collection of rules that control the traffic at the instance level
- Network will scale at 2Gbps for each CPU core except instance with 2-4 CPU

# VM
## VM Lifecycle
![](../../images/vm_lifecycle.png)
> `VM Lifecycle`

| State | Detailed Description | Billing (CPU/RAM)? |
| :--- | :--- | :--- |
| **PROVISIONING** | GCP is allocating hardware resources (CPU, RAM, Network) for the VM. The VM is not yet running. | **No** |
| **STAGING** | Resources are allocated. GCP is preparing basic configurations and boot setup for the VM. | **No** |
| **RUNNING** | The VM is booting the Operating System (OS) or running normally. SSH/RDP connections are available. | **Yes** |
| **SUSPENDING** | The system is saving the full memory state (RAM) of the VM to disk to prepare for suspension. | **Yes** (until process completes) |
| **SUSPENDED** | The VM is fully suspended (similar to Hibernate). Memory state (RAM) is preserved on disk. | **No** (only charged for disk storage & RAM image) |
| **PENDING_STOP** / **STOPPING** | The VM is undergoing a graceful OS shutdown process. | **Yes** (while in PENDING_STOP) |
| **TERMINATED** | The VM has completely stopped. CPU/RAM resources are released, but VM configurations and Persistent Disks are retained. | **No** (only charged for attached storage disks) |
| **REPAIRING** | The VM encountered a hardware fault or is being automatically repaired/recovered by GCP. | **No** |

## Compute Engine machine families 
### 1.General-purpose machine family

| Series | Workload | Applications |
| :--- | :--- | :--- |
| **E2** | (Cost-optimized) Day-to-day computing at a lower cost | • Web serving<br>• App serving<br>• Back office apps<br>• Small-medium databases<br>• Microservices<br>• Virtual desktops<br>• Development environments |
| **N2, N2D, N1** | (Balanced) Balanced price/performance across a wide range of VM shapes | • Web serving<br>• App serving<br>• Back office apps<br>• Medium-large databases<br>• Cache<br>• Media/streaming |
| **Tau T2D, Tau T2A** | (Scale-out optimized) Best performance/ cost for scale-out workloads | • Scale-out workloads<br>• Web serving<br>• Containerized microservices<br>• Media transcoding<br>• Large-scale Java applications |

### 2.Compute-optimized machine family

| Series | Workload | Applications |
| :--- | :--- | :--- |
| **C2** | Ultra high performance for compute-intensive workloads | • Compute-bound workloads<br>• High-performance web serving<br>• Gaming (AAA game servers)<br>• Ad serving<br>• High-performance computing (HPC)<br>• Media transcoding<br>• AI/ML |
| **C2D** | Ultra high performance for compute-intensive workloads | • Memory-bound workloads<br>• Gaming (AAA game servers)<br>• High-performance computing (HPC)<br>• High-performance databases<br>• Electronic Design Automation (EDA)<br>• Media transcoding |
| **H3** | Ultra high performance for compute-intensive workloads | • High-performance computing (HPC)<br>• Electronic Design Automation (EDA) |

### 3.Memory-optimized machine family

| Series | Workload | Applications |
| :--- | :--- | :--- |
| **M1** | Ultra high-memory workloads | • Medium in-memory databases such as SAP HANA<br>• Tasks that require intensive use of memory with higher memory-to-vCPU ratios than the general-purpose high-memory machine types.<br>• In-memory databases and in-memory analytics, business warehousing (BW) workloads, genomics analysis, SQL analysis services.<br>• Microsoft SQL Server and similar databases. |
| **M2** | Ultra high-memory workloads | • Large in-memory databases such as SAP HANA<br>• In-memory databases and in-memory analytics, business warehousing (BW) workloads, genomics analysis, SQL analysis services, etc. |
| **M3** | Ultra high-memory workloads | • OLAP and OLTP SAP workloads<br>• Memory Intense Electronic Design Automation |

### 4.Accelerator-optimized machine family

| Series | Workload | Applications |
| :--- | :--- | :--- |
| **A2** | Optimized for high-performance computing workloads | • CUDA-enabled ML training and inference<br>• HPC<br>• Massive parallelized computation |
| **G2** | Optimized for high-performance computing workloads | • Video transcoding<br>• Remote visualization workstation |

## Pricing
![alt text](image.png)
![alt text](image-1.png)

- A **Preemptible VM** (or **Spot VM** in modern GCP terminology) is a short-lived, cost-effective Compute Engine virtual machine instance. GCP offers these instances at a **60% to 90% discount** compared to standard on-demand pricing by utilizing excess capacity in their data centers.

- A **Sole-Tenant Node** in GCP Compute Engine is a dedicated physical Compute Engine server allocated exclusively for your use. It allows you to run VM instances on hardware isolated from other Google Cloud customers.

- A **Shielded VM** is a Compute Engine instance hardened by cryptographic security features. It protects workloads from enterprise-level threats such as rootkits, bootkits, and unauthorized firmware modification by verifying the boot integrity of the Virtual Machine.

- A **Confidential VM** is a GCP Compute Engine instance built on hardware-based Confidential Computing technology (such as AMD SEV) that automatically encrypts all memory (RAM) in real time while data is actively being processed (Data-in-Use), preventing unauthorized access from hypervisors, cloud admins, or co-located workloads without requiring any code changes.

# Disk
- A **Persistent Disk (PD)** is a durable, high-performance network-attached block storage service provided by Google Cloud Platform for Virtual Machines (Compute Engine and GKE). It exists independently of the VM lifecycle, ensuring data remains intact even if the VM is stopped, restarted, or deleted.

- A **Local SSD** is a high-performance, temporary block storage drive physically attached to the host server hosting a GCP Compute Engine instance. It provides ultra-high IOPS and extremely low latency for demanding workloads, but its data is ephemeral and does not persist when the VM is stopped or deleted.

- RAM disk use **tmpfs** is a temporary virtual file system in Linux that stores files directly in volatile memory (RAM and dynamic swap space). It provides ultra-fast read/write operations with near-zero latency, making it ideal for temporary or ephemeral data, though all contents are lost when the system powers off or reboots.

# Storage service
![alt text](image-2.png)
> overview of storage service

![alt text](image-3.png)
> how to choose service

## Cloud Storage
- `Cloud Storage`: An Object Storage service that stores immutable data as objects via API/HTTP, offering virtually unlimited scalability at extremely low cost on a global scale.

![alt text](image-4.png)
- Object uploaded to the bucket have a storage class of the bucket unless specify it
- Can't change location type of the bucket

### Syntax

| Operation | Command Syntax |
| --- | --- |
| Create Bucket | `gcloud storage buckets create gs://BUCKET_NAME --location=asia-southeast1` |
| List Buckets | `gcloud storage buckets list` |
| Get Bucket Info | `gcloud storage buckets describe gs://BUCKET_NAME` |
| Delete Bucket | `gcloud storage buckets delete gs://BUCKET_NAME` |
| Upload Single File | `gcloud storage cp file.txt gs://BUCKET_NAME/` |
| Upload Directory | `gcloud storage cp -r /path/to/folder gs://BUCKET_NAME/` |
| Download Single File | `gcloud storage cp gs://BUCKET_NAME/file.txt ./` |
| Download Directory | `gcloud storage cp -r gs://BUCKET_NAME/folder/ ./` |
| List Files | `gcloud storage ls gs://BUCKET_NAME/` |
| Rename / Move File | `gcloud storage mv gs://BUCKET_NAME/old.txt gs://BUCKET_NAME/new.txt` |
| Delete Single File | `gcloud storage rm gs://BUCKET_NAME/file.txt` |
| Delete Directory on Bucket | `gcloud storage rm -r gs://BUCKET_NAME/folder/` |
| Synchronize Data (Sync) | `gcloud storage rsync -r /path/local gs://BUCKET_NAME/` |
| View IAM Policy | `gcloud storage buckets get-iam-policy gs://BUCKET_NAME` |
| Grant Read Permission to User | `gcloud storage buckets add-iam-policy-binding gs://BUCKET_NAME --member="user:email@gmail.com" --role="roles/storage.objectViewer"` |
| Grant Write Permission to User | `gcloud storage buckets add-iam-policy-binding gs://BUCKET_NAME --member="user:email@gmail.com" --role="roles/storage.objectAdmin"` |
| Make Bucket Public | `gcloud storage buckets add-iam-policy-binding gs://BUCKET_NAME --member="allUsers" --role="roles/storage.objectViewer"` |
| Check Total Size | `gcloud storage du -s gs://BUCKET_NAME/` |
| Generate Temporary Link (1h) | `gcloud storage sign-url gs://BUCKET_NAME/file.txt --duration=1h` |
| Enable Object Versioning | `gcloud storage buckets update gs://BUCKET_NAME --versioning` |

## Filestore
- `Filestore` A File Storage (NFS) service organized according to a POSIX-compliant directory tree structure, allowing multiple servers or containers to mount the storage for shared read/write access with ultra-low latency.

## Cloud SQL
![alt text](image-5.png)
> cloud sql HA configuration

- the config is made up by `primary instance` and `standby instance`
- for each operation made to the primary instance, it will be replicated to the standby instance through zone persistent disk
- the transaction for the record is `commited` only when the record is replicated to the standby instance
- in event of zone failure, the persist disk is attatched to the standby instance and that instance become the primary ones (`failover`)
- support `Point in Time` recovery, database import export use file dump, import export CSV

## Spanner 
![alt text](image-6.png)