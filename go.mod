module doit

go 1.17

require (
	github.com/Jeffail/tunny v0.1.4
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-redis/redis/v8 v8.11.5
	github.com/segmentio/kafka-go v0.4.31
	github.com/willf/bloom v2.0.3+incompatible
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/klauspost/compress v1.14.2 // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/willf/bitset v1.2.2 // indirect
)

replace github.com/willf/bitset v1.2.2 => github.com/bits-and-blooms/bitset v1.2.2
