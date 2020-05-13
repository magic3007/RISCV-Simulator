package main

import (
	"cache"
	"debug/elf"
	"debugger"
	"flag"
	"fmt"
	"log"
	"memory"
	"pipedebug"
	"pipeline"
	"simulator"
	"utils"
)

var (
	verbose 	= flag.Bool("v", true, "Display verbose info of ELF file. To disable displaying these info, type \"-v=0\"")
	CacheFlag	= flag.Bool("cache", false, "Use Cache")
	filepath 	= flag.String("f", "", "Filepath of the ELF file")
	mode 		= flag.String("m", "debug", "Simulation mode. Valid modes are: debug, pipeline")
)

func statElf(file *elf.File){
	utils.RecursiveStatField(*file, 0)
	for index, prog := range file.Progs{
		fmt.Println("===========================================")
		fmt.Printf("Porgram %d:\n", index)
		utils.RecursiveStatField(*prog, 2)
	}
}

func NewNoCachedStorage() (storage cache.Storage, StatFunc func()){
	latency := cache.StorageLatency{
		HitLatency: 40,
	}
	mainMemory := cache.NewMemory(latency)
	return mainMemory, nil
}

func NewCachedStorage() (storage cache.Storage, StatFunc func()){
	latency := cache.StorageLatency{
		HitLatency: 40,
	}
	mainMemory := cache.NewMemory(latency)

	latency = cache.StorageLatency{
		HitLatency: 20,
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
		HitLatency: 8,
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
	storage = L1Cache
	StatFunc = func(){
		fmt.Printf("L1 AccessCounter : %6d MissNum : %6d Miss Rate: %f \n", L1Cache.Stats.AccessCounter, L1Cache.Stats.MissNum, float32(L1Cache.Stats.MissNum) / float32(L1Cache.Stats.AccessCounter))
		fmt.Printf("L2 AccessCounter : %6d MissNum : %6d Miss Rate: %f \n", L2Cache.Stats.AccessCounter, L2Cache.Stats.MissNum, float32(L2Cache.Stats.MissNum) / float32(L2Cache.Stats.AccessCounter))
		fmt.Printf("L3 AccessCounter : %6d MissNum : %6d Miss Rate: %f \n", L3Cache.Stats.AccessCounter, L3Cache.Stats.MissNum, float32(L3Cache.Stats.MissNum) / float32(L3Cache.Stats.AccessCounter))
		fmt.Printf("Average Access Time: %f \n", float32(L1Cache.Stats.AccessTime) / float32(L1Cache.Stats.AccessCounter))
	}
	return
}

func main() {
	utils.FlagInit([]string{"f"})
	log.SetFlags(log.LstdFlags | log.Llongfile)
	file, err := elf.Open(*filepath)
	utils.PANIC_CHECK(err)
	defer file.Close()
	if *verbose == true {
		statElf(file)
	}
	var storage cache.Storage
	var StatFunc func()
	if *CacheFlag {
		storage, StatFunc = NewCachedStorage()
	}else{
		storage, StatFunc = NewNoCachedStorage()
	}
	sim := simulator.NewSimulator(*filepath, memory.NewMemory(storage))
	sim.LoadMemory()
	switch *mode{
	case "debug":
		var d debugger.Debugger
		d.RunPrompt(sim)
	case "pipeline":
		var d pipedebug.PipeDebugger
		pipe := pipeline.Pipe{Simulator: sim}
		d.RunPrompt(&pipe)
	default:
		err := fmt.Errorf("unvalid mode")
		fmt.Print(err.Error())
	}
	if StatFunc != nil {
		StatFunc()
	}
}
