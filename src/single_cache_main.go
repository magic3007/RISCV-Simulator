package main

import (
	"cache"
	"flag"
	"fmt"
	"github.com/inhies/go-bytesize"
	"io"
	"log"
	"os"
	"strconv"
	"utils"
)

var (
	filepath = flag.String("f", "", "Filepath of the trace file")
	outfile = flag.String("o", "", "output file")
)

func foo(LineSize uint, Associativity uint, CacheSize uint, WriteUpPolicy cache.WriteUpPolicyType,
	WriteBackPolicy cache.WriteBackPolicyType, writer io.Writer){
	file, err := os.Open(*filepath)
	utils.PANIC_CHECK(err)
	defer file.Close()

	latency := cache.StorageLatency{
		HitLatency: 10,
	}
	mainMemory := cache.NewMemory(latency)

	latency = cache.StorageLatency{
		HitLatency: 1,
	}
	config := cache.CacheConfig{
		LineSize:        LineSize,
		Associativity:   Associativity,
		CacheSize:       CacheSize,
		WriteUpPolicy:   WriteUpPolicy,
		WriteBackPolicy: WriteBackPolicy,
	}
	Cache := cache.NewCache(latency, config, mainMemory)

	for {
		var op, addr_string string
		_, err := fmt.Fscanf(file,"%s %s\n", &op, &addr_string)
		if err == io.EOF { break }
		utils.PANIC_CHECK(err)
		addr, err := strconv.ParseInt(addr_string, 0, 64)
		utils.PANIC_CHECK(err)
		switch op {
		case "r": Cache.HandleRequest(uint64(addr), uint(1), cache.Read); break
		case "w": Cache.HandleRequest(uint64(addr), uint(1), cache.Write); break
		default:
			panic(fmt.Sprintf("Unknown operation: %s", op))
		}
	}
	WriteUpPolicyToString := func(x cache.WriteUpPolicyType) string{
		switch x {
		case cache.WriteThrough: return "WriteThrough"
		case cache.WriteBack: return "WriteBack"
		}
		return "Unknown Policy"
	}
	WriteBackPolicyToString := func(x cache.WriteBackPolicyType) string{
		switch x {
		case cache.WriteAllocate: return "WriteAllocate"
		case cache.NoWriteAllocate: return "NoWriteAllocate"
		}
		return "Unknown Policy"
	}

	fmt.Println("=====================================================")
	fmt.Println(config)
	fmt.Printf("Miss Rate: %f \n", float32(Cache.Stats.MissNum) / float32(Cache.Stats.AccessCounter))
	fmt.Printf("Average Access Time: %f \n", float32(Cache.Stats.AccessTime) / float32(Cache.Stats.AccessCounter))
	fmt.Fprintf(writer, "%d,%d,%s,%s,%s,%f,%f\n",
		LineSize,
		Associativity,
		bytesize.New(float64(CacheSize)),
		WriteUpPolicyToString(WriteUpPolicy),
		WriteBackPolicyToString(WriteBackPolicy),
		float32(Cache.Stats.MissNum)/float32(Cache.Stats.AccessCounter),
		float32(Cache.Stats.AccessTime)/float32(Cache.Stats.AccessCounter))
}

func main(){
	utils.FlagInit([]string{"f", "o"})
	log.SetFlags(log.LstdFlags | log.Llongfile)
	outfile, err := os.Create(*outfile)
	utils.PANIC_CHECK(err)
	defer outfile.Close()

	fmt.Fprint(outfile, "LineSize,Associativity,CacheSize,WriteUpPolicy,WriteBackPolicy,Miss Rate,AMAT\n")
	//LineSizes := []uint{32, 64, 128, 256,512,1024,2048,4096}
	LineSizes := []uint{256}
	//Associativities := []uint{1,2,4,8,16,32}
	Associativities := []uint{8}
	//Associativities := []uint{16}
	//CacheSizes := []uint{32 * 1024, 64 * 1024, 128 * 1024, 256 * 1024, 512 * 1024,  1024 * 1024,
	//	2 * 1024 * 1024, 4 * 1024 * 1024, 8 * 1024 * 1024, 16 * 1024 * 1024, 32 * 1024 * 1024,
	//	}
	//CacheSizes := []uint{32 * 1024, 64 * 1024, 128 * 1024, 256 * 1024, 384 * 1024, 512 * 1024,  640 * 1024, 768 * 1024,
	//	896 * 1024, 1024 * 1024,
	//	2 * 1024 * 1024, 4 * 1024 * 1024, 8 * 1024 * 1024, 16 * 1024 * 1024, 32 * 1024 * 1024,
	//	}
	CacheSizes := []uint{512 * 1024}
	//WriteUpPolicies := []cache.WriteUpPolicyType{cache.WriteBack}
	//WriteBackPolicies := []cache.WriteBackPolicyType{cache.WriteAllocate}
	WriteUpPolicies := []cache.WriteUpPolicyType{cache.WriteBack, cache.WriteThrough}
	WriteBackPolicies := []cache.WriteBackPolicyType{cache.WriteAllocate, cache.NoWriteAllocate}

	for _, LineSize := range LineSizes{
		for _, Associativity := range Associativities{
			for _, CacheSize := range CacheSizes{
				for _, writeUpPolicy := range WriteUpPolicies{
					for _, WriteBackPolicy := range WriteBackPolicies{
						foo(LineSize, Associativity, CacheSize, writeUpPolicy, WriteBackPolicy, outfile)
					}
				}
			}
		}
	}
}