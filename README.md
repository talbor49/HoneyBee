# HoneyBee
The Honey Bee DB behaves like a beehive. A key-value database built in GOLANG



DB concepts / things to notice while building:

1. Compress data - save pointers to data, etc.
2. RAM is the cache, everything is saved to memory eventually
3. Distributing the DB into multiple machines. Split the data, split the tasks, synchronize.
4. Make it stable & durable - have replica, backup data, keep logs, avoid single point of failures. 
