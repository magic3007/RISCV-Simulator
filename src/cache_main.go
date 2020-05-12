package main

import (
	"cache"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"utils"
)

var (
	filepath = flag.String("f", "", "Filepath of the trace file")
)

func main(){
	utils.FlagInit([]string{"f"})
	log.SetFlags(log.LstdFlags | log.Llongfile)
	file, err := os.Open(*filepath)
	utils.PANIC_CHECK(err)
	defer file.Close()

	latency := cache.StorageLatency{
		HitLatency: 26,
	}
	mainMemory := cache.NewMemory(latency)

	latency = cache.StorageLatency{
		HitLatency: 4,
	}
	config := cache.CacheConfig{
		LineSize:        64,
		Associativity:   8,
		CacheSize:       8 * 1024 * 1024,
		WriteUpPolicy:   cache.WriteBack,
		WriteBackPolicy: cache.WriteAllocate,
	}
	L3Cache := cache.NewCache(latency, config, mainMemory)

	latency = cache.StorageLatency{
		HitLatency: 2,
	}
	config = cache.CacheConfig{
		LineSize:        64,
		Associativity:   8,
		CacheSize:       256 * 1024,
		WriteUpPolicy:   cache.WriteBack,
		WriteBackPolicy: cache.WriteAllocate,
	}
	L2Cache := cache.NewCache(latency, config, L3Cache)

	latency = cache.StorageLatency{
		HitLatency: 1,
	}
	config = cache.CacheConfig{
		LineSize:        64,
		Associativity:   8,
		CacheSize:       32 * 1024,
		WriteUpPolicy:   cache.WriteBack,
		WriteBackPolicy: cache.WriteAllocate,
	}
	L1Cache := cache.NewCache(latency, config, L2Cache)

	for {
		var op, addr_string string
		_, err := fmt.Fscanf(file,"%s %s\n", &op, &addr_string)
		if err == io.EOF { break }
		utils.PANIC_CHECK(err)
		addr, err := strconv.ParseInt(addr_string, 0, 64)
		utils.PANIC_CHECK(err)
		switch op {
		case "r": L1Cache.HandleRequest(uint64(addr), uint(1), cache.Read); break
		case "w": L1Cache.HandleRequest(uint64(addr), uint(1), cache.Write); break
		default:
			panic(fmt.Sprintf("Unknown operation: %s", op))
		}
	}
	fmt.Println("======================================================================================")
	fmt.Println(*filepath)
	fmt.Printf("L1 AccessCounter : %6d MissNum : %6d Miss Rate: %f \n", L1Cache.Stats.AccessCounter, L1Cache.Stats.MissNum, float32(L1Cache.Stats.MissNum) / float32(L1Cache.Stats.AccessCounter))
	fmt.Printf("L2 AccessCounter : %6d MissNum : %6d Miss Rate: %f \n", L2Cache.Stats.AccessCounter, L2Cache.Stats.MissNum, float32(L2Cache.Stats.MissNum) / float32(L2Cache.Stats.AccessCounter))
	fmt.Printf("L3 AccessCounter : %6d MissNum : %6d Miss Rate: %f \n", L3Cache.Stats.AccessCounter, L3Cache.Stats.MissNum, float32(L3Cache.Stats.MissNum) / float32(L3Cache.Stats.AccessCounter))
	fmt.Printf("Average Access Time: %f \n", float32(L1Cache.Stats.AccessTime) / float32(L1Cache.Stats.AccessCounter))
}