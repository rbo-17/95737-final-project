# Comparing Database Performance For Small Datasets

## Test Setup

### Data
Benchmarks will consider three different data sizes:
1. Small text (100-200B)
2. Large text (10-20KB)
3. Image (1-2MB)
This will enable considerations between lookup times and data transfer speeds

Additionally, benchmarks will consider three different read/write scenarios:
1. Write-heavy scenarios (10/90 reads/writes)
2. Balanced (50/50 reads/writes)
3. Read-heavy scenarios (90/10 reads/writes)


This will cover a variety of different use cases including:
* Caching 
* Mostly static content

### Benchmark

## Test Design




## Database Setup

### Redis
Amazon Machine Image (AMI): Amazon Linux 2023 AMI 
Arch: 64-bit Arm
Instance Type: 
Storage: 

### MongoDB

### Cassandra

### MySQL

### Results


## Conclusion

