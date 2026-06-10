# Chapter 12: Physical Storage Systems

## 12.1 Overview of Physical Storage Media
- **medium:** a physical material, environment which store, transfer data
- **drive:** the part that control, read and write data to the medium
- **cache:** the fastest and most expensive storage
- **main memory:** the second fastest and second most expensive storage, it is volatile, which means that the data will be lost when the power is off
- **flash memory:** a non-volatile storage, widely used for storage in devices such as cellphones, cameras. SSD use flash internally.
    + flash provide a interface similar to magnetic disk called block oriented interface
- **magnetic-disk storage:** also a non-volatile storage, it is used for long-term storage of data, it is slower than flash but cheaper. Also refer to as "hard disk drive" (HDD)
    + to read or write data to a magnetic disk, the system must first move data to the main memory, after performing operations, data is written back to the disk
- **optical storage:** a DVD is an example of optical storage, it is a non-volatile storage, using laser light to read and write data
    + some DVD can be read only, some can be write once, WORM (write once read many)  
    + suitable for large amount of data that need to be stored for a long time but not need to be updated, such as backup
    + **jukebox:** a system contain multiple DVD and a robotic arm to load DVD to the drive
- **tape storage:** refer to sequential storage which data must be read at the beginning of the tape <> direct access storage which data can be read at any point (SSD, HDD)

- **magnetic tape and optical disk jukeboxes** are referred to as tertiary storage, or offline storage.

## 12.2 Storage Interfaces
- **Serial ATA (SATA):** a common interface for connecting storage devices such as SSD and HDD to a computer, it use AHCI protocal, usually used for personal devices and half-duplex (only write or read at a time)
- **AHCI protocal** run on SATA interface
- **Serial Attached SCSI (SAS):** a high-speed interface for connecting storage devices, it faster and usually used for enterprise storage systems, it support full-duplex (write and read at the same time)

- **Non-Volatile Memory Express (NVMe):** a high-performance  protocal for accessing non-volatile storage media, it run on PCIe bus, it provide low latency and high throughput

- disk can also be connected to the computer through a high-speed network, such as Storage Area Network (SAN)
- **RAID (Redundant Array of Independent Disks):** a storage organization technique to give server a logical view of a large, reliable disk

- **SCSI:** a protocal for connecting and transferring data between computers and peripheral devices
- **iSCSI:** a protocal that allow to use SCSI over a network, it is commonly used for SAN

- **Network attached storage (NAS):** its like SAN but instead of network storage appear to be a large disk, it provide file system interface instead of access through block of data
    + it use network file system (NFS) or Common Internet File System (CIFS)
    + **DFS** is an abstract layer of NFS and CIFS protocols, it allow us to store and access files across multiple machine as if they were stored on a single machine

## 12.3 Magnetic Disks

### 12.3.1 Physical Characteristics of Disks
- **platter:** made from metal or glass, a flat, circular shape, 2 surfaces are coated with magnetic device
- drive motor can rotate the platter at a constant high speed
- read/write head place just above the surface of the platter
- platter is logically divided into "tracks" and subdivided into "sectors"
- **sector:** the smallest unit of data that can be read or written to a disk, typically 512 bytes
- disk can contain 2-24 bilion sectors
- inner track has smaller circumference than outer track, so it can store less data than outer track (less sector per track)
- disk contain multiple platters
- read/write head is attached to an "disk arm" and move together
- **cylinder:** a set of tracks that are vertically aligned across multiple platters, it form a cylinder 
- read write a kept as close as possible to the surface of the platter to maximize the "storage density"
- spinning platter create "air cushion" that keep the read/write head from touching the surface of the platter
- **disk controller:** the interface to interact disk, implemented within the disk drive unit
- when write data to each sector, the disk controller also write "checksum" (computed value from data written) to the sector, when read data from the sector, the disk controller will compute the checksum and compare it with the checksum stored in the sector to detect if there is an error in the data
- if there is error, the disk controller will try to read the sector serveral times, if it still fail, it will signal as read failure
- **remapping of bad sector:** when the disk controller detect a bad sector while formatting the disk, write to a sector fail:
    + mark the sector as bad
    + remap the sector to a spare sector that is reserved for this purpose
    + remapping is transparent to the user -> redirect the read/write request to the remapped sector
    + save the mapping information in a non-volatile memory within the disk drive unit

### 12.3.2 Performance Measures of Disks
- **access time:** the time from read/write request is issued -> data transfer begin

- **seek time:** the time to move the read/write head to the track where the data is located 
- **average seek time:** the average time to move the read/write head to a random track, it is typically one-half of worst-case seek time (around 4-10 ms) depending on the disk model

- **rotational latency:** the time to rotate the disk to the position the desired sector is under the read/write head.
- on average its is one-half of the time for a full rotation

- `access time = seek time + rotational latency`

- **data-transer rate:** is the rate at which data can be retrieved from or written to the disk in an amount of time
- transfer rate in the inner track are lower than the outer track since theu have less sector

- request for disk I/O mainly generated by the file system, some generated by the database system
- request species address called "disk number"
- **disk block:** a logical unit of data that is read/written to the disk, size range from 4 -> 16 kilo byte
- data transfer to main memory is done in unit of disk block (often called "page" in database system)

- **sequential access:** successive (liên tục) request for sucessive block number (which are likely to be stored in the same track or adjacent (liền kề) track) 
    + not require block seek except for the first block and adjacent block (faster seek time)

- **I/O operations per second (IOPS):** the number of random block access that can be performed in one second
- data transfer rate in random access is much lower than sequential access since it require seek time and rotational latency for each block access

- **mean time to failure (MTTF):** the average time until a disk fail

## 12.4 Flash Memory
- There are two types of flash memory:
- **NOR flash:** it is faster for read and random access, but it is slower for write and erase, it is more expensive than NAND flash, it is commonly used for code storage and execution in embedded systems

- **NAND flash:** it is faster for write and erase, but it is slower for read and random access, it is cheaper than NOR flash, it is commonly used for data storage in devices such as SSD, USB flash drive, etc.
    + reading and writing data to NAND flash is done in unit of "page" (typically 4096 bytes)
    + to erase data from NAND flash, it must be done in unit of "erase block" (typically 128-256 pages)
    + there also a limit on the number of write/erase cycle for each block, after that error on storing bit are likely to happen

- to minimize the impact of both slow erase and update, flash memory use mapping called "translation table to" map logical page number to physical page number
    + when a logical page is updated, the new data is written to already erased physical page and the original physical page is mark as deleted and erased later
    + each physical page have a small section of storage to store the logical page number

- if block having multiple deleted pagesm it will be set up for a garbage collection process
    + first, the valid page in the block will be copied to another block
    + then the block will be erased and all page in the block will be marked as free
    + translation table will be updated to reflect the new mapping of logical page number to physical page number

- **cold data:** data that is rarely updated
- **hot data:** data that is frequently updated
- **wear leveling:** a technique to distribute write and erase cycles evenly across the flash memory by: assign cold data to block have been updated many times, and assign hot data to block have been updated few times

- all above action are carried out in **flash translation layer (FTL)**
- above FTL, flash storage provide identical interface to magnetic disk

- **SSD performance:**
    + support 10,000 IOPS with 4 kb blocks
    + 32 parallel channel to read/write data (100,000 random with SATA, 350,000 random with NVMe PCIe)
    - data transfer rate 400-500 mb/s for SATA, 2000-3000 mb/s for NVMe PCIe

## 12.5 RAID

### 12.5.1 Improvement of Reliability via Redundancy
- **redurdency:** the use of additional disk to store redundant data
    + **mirroring:** a logical disk consist of two physical disk, write data to both disk, if one disk fail, the other disk can still provide the data

- **mean time to failure (MTTF):** the average time until a disk fail
- **mean time to repair (MTTR):** the average time to repair a failed disk
- **mean time to data loss (MTTDL):** `(2 * delta) "probility of 2 disk failure" * (delta * MTTR) "the other disk fail during the repair time of the first disk"`
    + `(delta = 1/MTTF)`

- if power fail befor both disk finish writing block, it can cause data inconsistency
    + to deal with this problem, we can write one copy first, then write the other copy

### 12.5.2 Improvement of Performance via Parallelism
- **block level stripping:** this technique is similar to partitioning data, it divide data into block and distribute the block across multiple disk
    + improve performance by allowing multiple disk to read/write data in parallel

### 12.5.3 RAID Levels
- **parity block:** a block that store the parity information, computed by applying XOR operation to the data block, it can be used to recover data in case of disk failure
- the best way to update parity block is: `P(new) = P(old) XOR A(old) XOR A(new)`

- **RAID levels:**
    + **RAID 0:** block level stripping without redundancy (mirroring or parity)
    + **RAID 1:** mirroring with block level stripping
    + some vendors use the term "RAID 10" to refer to what RAID 1 actually is (they use RAID 1 to refer to mirroring without block level stripping, write data to disk until it is full, then write to the other disk)
    + **RAID 5:** block level stripping with distributed parity, it use N + 1 disk to store data and parity for each set of N data block
    + parity is at disk k mod 5 
    + where k is a stripe number
    + a stripe is a set of block that are stored across multiple disk, it consist of N data block and 1 parity block
    + dont store parity block and its data on the same disk, it can cause data loss if the disk fail
    + **RAID 6:** block level stripping with double distributed parity, it can tolerate two disk failure using Reed Solomon containing two parity block (P, Q) for each set of N data block

### 12.5.4 Hardware Issues
- **software RAID:** implemented with no change in hardware level, using only software modification
- **hardware RAID:** specialized hardware to support RAID
- **scrubbing:** a process of periodically reading all data and parity block to detect and correct any error in the data
- **hot swapping:** the ability to replace a failed disk without shutting down the system, it can significantly reduce the MTTR

- RAID system might also have a hot spare disk, which is a disk that is not used for storing data but can be automatically used to replace a failed disk
- we also should implement power supplies in case of power failure

### 12.5.5 Choice of RAID Level
- there are factors to consider when choosing a RAID level:
    + monetary cost
    + performance in I/O operation
    + performance when a disk fail
    + performance during the repair of a failed disk

- RAID 0 is suitable for applications that require high performance but can tolerate data lossS
- RAID 1 provide best write performance 
- RAID 5 is suitable for applications where high read and written rarely (except for sequential write)
- RAID 6 is similar to RAID 5 but with better fault tolerance, it is important to notice that latent sector error can happen during the repair of a failed disk, which can cause data loss in RAID 5 but not in RAID 6

- mirroring 3 copies of data is commonly used in large scale distributed storage system but not in single machine storage system since chance of a machine failure is much higher than chance of a disk failure

## 12.6 Disk-Block Access
- request for disk I/O are generated by the database system
    + case direct access to disk: request specify disk id and logical block number
    + case access through file system: specify file identifier and block number

- there are some techniques to improve the random access performance of disk:
    + **buffering:** the use of main memory to store recently accessed data

    + **read-ahead:** block on the same track are likely to be accessed together in case of sequential access, so we pre-read the consecutive block

    + **scheduling:** attempt to reorder the disk I/O request to minimize the seek time and rotational latency, for ex: elevator algorithm

    + **file organization:** we expect block in a way correspond to the way data is accessed, for ex: store a file data in the same track or adjacent track
        * but it is not always possible to store data in the same track or adjacent track, for ex: when the file is larger, its hard to find a large enough contiguous space on the disk, or when the file is updated and grow in size, it may need to be moved to another location on the disk, which can cause fragmentation
        * because the unknown nature of the file size and update, we can use "extent", a consecutive block of disk
        * over time, a sequential file has multiple small append can become fragmented, to reduce fragmentation, we can use "backup & restore" or utilities, basically it re-organize the file on the disk

    + **non-volatile write buffers:** for heavily write application, we can use non-volatile write buffer to store the data and immediatly notifies the OS before writing to the disk and 

- reordering disk I/O request can become inconsistent