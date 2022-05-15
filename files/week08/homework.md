10 20 50 100 200 1k 5k
// 10 bytes"
redis-benchmark -h 192.168.1.201 -p 6379 -q script load "redis.call('set','test_10','1')" -d 10 
redis-benchmark -h 192.168.1.201 -p 6379 -q script load "redis.call('get','test_10','1')" -d 10 

// 20 bytes
redis-benchmark -h 192.168.1.201 -p 6379 -q script load "redis.call('set','test_20','1')" -d 20 
redis-benchmark -h 192.168.1.201 -p 6379 -q script load "redis.call('get','test_20','1')" -d 20 

// 50 bytes
redis-benchmark -h 192.168.1.201 -p 6379 -q script load "redis.call('set','test_50','1')" -d 50 
redis-benchmark -h 192.168.1.201 -p 6379 -q script load "redis.call('get','test_50','1')" -d 50 

// 100 bytes
redis-benchmark -h 192.168.1.201 -p 6379 -q script load "redis.call('set','test_100','1')" -d 100
redis-benchmark -h 192.168.1.201 -p 6379 -q script load "redis.call('get','test_100','1')" -d 100 

// 200 bytes
redis-benchmark -h 192.168.1.201 -p 6379 -q script load "redis.call('set','test_200','1')" -d 200
redis-benchmark -h 192.168.1.201 -p 6379 -q script load "redis.call('get','test_200','1')" -d 200

// 1024 bytes
redis-benchmark -h 192.168.1.201 -p 6379 -q script load "redis.call('set','test_1024','1')" -d 1024
redis-benchmark -h 192.168.1.201 -p 6379 -q script load "redis.call('get','test_1024','1')" -d 1024

// 32 * 1024 bytes
redis-benchmark -h 192.168.1.201 -p 6379 -q script load "redis.call('set','test_32768','1')" -d 32768
redis-benchmark -h 192.168.1.201 -p 6379 -q script load "redis.call('get','test_32768','1')" -d 32768

# 第二问
info memory 

used_memory:13490096 //数据占用了多少内存（字节）

used_memory_human:12.87M //数据占用了多少内存（带单位的，可读性好）


假如缓存数据小于4GB，就使用32位的Redis实例。因为32位实例上的指针大小只有64位的一半，它的内存空间占用空间会更少些。 这有一个坏处就是，假设物理内存超过4GB，那么32位实例能使用的内存仍然会被限制在4GB以下。 要是实例同时也共享给其他一些应用使用的话，那可能需要更高效的64位Redis实例，这种情况下切换到32位是不可取的。 不管使用哪种方式，Redis的dump文件在32位和64位之间是互相兼容的， 因此倘若有减少占用内存空间的需求，可以尝试先使用32位，后面再切换到64位上。

尽可能的使用Hash数据结构。因为Redis在储存小于100个字段的Hash结构上，其存储效率是非常高的。所以在不需要集合(set)操作或list的push/pop操作的时候，尽可能的使用Hash结构。比如，在一个web应用程序中，需要存储一个对象表示用户信息，使用单个key表示一个用户，其每个属性存储在Hash的字段里，这样要比给每个属性单独设置一个key-value要高效的多。 通常情况下倘若有数据使用string结构，用多个key存储时，那么应该转换成单key多字段的Hash结构。 如上述例子中介绍的Hash结构应包含，单个对象的属性或者单个用户各种各样的资料。Hash结构的操作命令是HSET(key, fields, value)和HGET(key, field)，使用它可以存储或从Hash中取出指定的字段。

设置key的过期时间。一个减少内存使用率的简单方法就是，每当存储对象时确保设置key的过期时间。倘若key在明确的时间周期内使用或者旧key不大可能被使用时，就可以用Redis过期时间命令(expire,expireat, pexpire, pexpireat)去设置过期时间，这样Redis会在key过期时自动删除key。 假如你知道每秒钟有多少个新key-value被创建，那可以调整key的存活时间，并指定阀值去限制Redis使用的最大内存。

回收key。在Redis配置文件中(一般叫Redis.conf)，通过设置“maxmemory”属性的值可以限制Redis最大使用的内存，修改后重启实例生效。 也可以使用客户端命令config set maxmemory 去修改值，这个命令是立即生效的，但会在重启后会失效，需要使用config rewrite命令去刷新配置文件。 若是启用了Redis快照功能，应该设置“maxmemory”值为系统可使用内存的45%，因为快照时需要一倍的内存来复制整个数据集，也就是说如果当前已使用45%，在快照期间会变成95%(45%+45%+5%)，其中5%是预留给其他的开销。 如果没开启快照功能，maxmemory最高能设置为系统可用内存的95%。

当内存使用达到设置的最大阀值时，需要选择一种key的回收策略，可在Redis.conf配置文件中修改“maxmemory-policy”属性值。 若是Redis数据集中的key都设置了过期时间，那么“volatile-ttl”策略是比较好的选择。但如果key在达到最大内存限制时没能够迅速过期，或者根本没有设置过期时间。那么设置为“allkeys-lru”值比较合适，它允许Redis从整个数据集中挑选最近最少使用的key进行删除(LRU淘汰算法)。Redis还提供了一些其他淘汰策略，如下：

volatile-lru：使用LRU算法从已设置过期时间的数据集合中淘汰数据。
volatile-ttl：从已设置过期时间的数据集合中挑选即将过期的数据淘汰。
volatile-random：从已设置过期时间的数据集合中随机挑选数据淘汰。
allkeys-lru：使用LRU算法从所有数据集合中淘汰数据。
allkeys-random：从数据集合中任意选择数据淘汰
no-enviction：禁止淘汰数据。
通过设置maxmemory为系统可用内存的45%或95%(取决于持久化策略)和设置“maxmemory-policy”为“volatile-ttl”或“allkeys-lru”(取决于过期设置)，可以比较准确的限制Redis最大内存使用率，在绝大多数场景下使用这2种方式可确保Redis不会进行内存交换。倘若你担心由于限制了内存使用率导致丢失数据的话，可以设置noneviction值禁止淘汰数据。
