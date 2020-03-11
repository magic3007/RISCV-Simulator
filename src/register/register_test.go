package register

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeap64(t *testing.T) {
	rh := Heap64{}
	for i := 0; i < RegNum; i++{
		rh.store(i, uint64(i+1))
		if i !=0 {
			assert.Equal(t, rh.load(i), uint64(i+1))
		}else{
			assert.Equal(t, rh.load(i), uint64(0))
		}
	}
	for i := 0; i < RegNum; i++{
		name := RegName[i][len(RegName[i])-1]
		rh.storeByName(RegName[i][len(RegName[i])-1], uint64(i+2))
		if i!=0 {
			assert.Equal(t,  rh.loadByName(name), uint64(i+2))
		}else{
			assert.Equal(t,  rh.loadByName(name), uint64(0))
		}
	}
}
