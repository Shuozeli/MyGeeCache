# MyGeeCache

MyGeeCache is a Go implementation inspired by groupcache, aiming to serve as a replacement for memcached in certain scenarios. Interestingly, the author of groupcache is also the author of memcached. Whether you're familiar with standalone caching or distributed caching, delving into the implementation of this library is highly meaningful.

## Features

- **Standalone Caching and HTTP-based Distributed Caching**: GeeCache supports both standalone caching for single instances as well as distributed caching over HTTP for larger deployments.
  
- **Least Recently Used (LRU) Cache Policy**: Implements the LRU cache eviction strategy to optimize cache performance and resource utilization.
  
- **Go Locking Mechanism to Prevent Cache Penetration**: Utilizes Go's locking mechanism to prevent cache penetration and maintain data consistency under high loads.
  
- **Consistent Hashing for Node Selection**: Implements consistent hashing for node selection, ensuring load balancing across distributed cache nodes.
  
- **Protocol Buffers (protobuf) for Optimized Binary Communication Between Nodes**: Enhances efficiency and speed of communication between cache nodes through the use of protobuf.

