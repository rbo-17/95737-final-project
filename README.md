# Comparing Database Performance For Small Datasets

# Description

# Relevant Common Use Cases
- Caching
- Comment data store (e.g. social media post, comment section post)
- Static text (e.g. frequent referenced books snippets, blog and forum posts)
- Image data store 

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

### Note Regarding Optimizations
* Fo simplicities' sake, an EC2 instance type of <insert-instance-type> was used for all tests. While it is widely 
recommended for database workloads, it may not be the ideal choice in all cases. For example, M series instances with the proper EBS configurations may be a better choice for Redis

## Results & Analysis
### Query Limitations
Redis: 
* All data must fit in memory
* For good performance, queries must be get operations
MongoDB:
Cassandra:
MySQL:

### Interesting Findings

## Conclusion

####
- Latency ms (avg)
- Latency boxplot
- Bytes/sec
- Req/sec

```html
<h2>Example of code</h2>

<pre>
    <div class="container">
        <div class="block two first">
            <h2>Your title</h2>
            <div class="wrap">
            //Your content
            </div>
        </div>
    </div>
</pre>
```