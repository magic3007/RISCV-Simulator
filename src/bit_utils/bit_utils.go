package bit_utils

/**
* @return value[to:from] (interval closed both at the left and the right)
 */
func U32Bits(value uint32, from uint32, to uint32) uint32{
	length := to - from + 1
	mask := (uint32(1)<< length) - 1
	return (value >> from) & mask
}

func ConcatenateSignExtBits(value uint32, intervals ...[2]uint32) int32{
	rv, l := int32(0), uint32(0)
	for _ , interval := range intervals {
		from, to := interval[0], interval[1]
		rv |= int32(U32Bits(value, from, to))<< l
		l += to - from + 1
	}
	rv = (rv<< (32 - l)) >> (32 - l)
	return rv
}

func BitMask(length uint8) uint64{
	return uint64(1)<<length - 1
}

func SetBit(value uint64, index uint32, targetBit uint64) uint64{
	target := uint64(1) << index
	mask := ^ target
	rv := value & mask
	if targetBit == 1 {
		rv |= target
	}
	return rv
}

func SignExtU64(x int32) uint64{
	return uint64(int64(x))
}

func UnSignExtU64(x int32) uint64{
	return uint64(uint32(x))
}