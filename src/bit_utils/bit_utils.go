package bit_utils

import "math/bits"

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

func signI64(x int64) int64{
	if x>0 {return 1}
	if x<0 {return -1}
	return 0
}

func absI64(x int64) uint64{
	if signI64(x)<0{
		return uint64(-x)
	}else{
		return uint64(x)
	}
}

func MulI64I64(x int64, y int64) (int64, int64){
	f:= signI64(x) * signI64(y)
	ax, ay := absI64(x), absI64(y)
	if f>=0 {
		hi, lo := bits.Mul64(ax, ay)
		return int64(hi), int64(lo)
	}else{
		hi, lo := bits.Mul64(ax, ay)
		hi = ^hi
		if lo == 0{hi += 1}
		lo = -lo
		return int64(hi), int64(lo)
	}
}

func MulI64U64(x int64, y uint64) (int64, int64){
	ax := absI64(x)
	if signI64(x)>=0 {
		hi, lo := bits.Mul64(ax, y)
		return int64(hi), int64(lo)
	}else{
		hi, lo := bits.Mul64(ax, y)
		hi = ^hi
		if lo == 0{hi += 1}
		lo = -lo
		return int64(hi), int64(lo)
	}
}
