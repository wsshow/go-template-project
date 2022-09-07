package utils

import (
	"log"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestContain(t *testing.T) {
	ss := []interface{}{"123", "456", "7", "8", "9"}
	assert.Equal(t, Contain(ss, "7"), true)
	assert.Equal(t, Contain(ss, "0"), false)
}

func TestContainEx(t *testing.T) {
	type st struct {
		Name string
	}
	sts := []*st{
		{Name: "123"}, {Name: "456"},
	}
	assert.Equal(t, ContainEx(sts, func(u interface{}) bool { return u.(*st).Name == "123" }), true)
	assert.Equal(t, ContainEx(sts, func(u interface{}) bool { return u.(*st).Name == "789" }), false)
}

func TestGetPrevDir(t *testing.T) {
	dirPath := "./lpcap/"
	assert.Equal(t, GetPrevDir(dirPath), "./lpcap")
}

func TestCheckNetworkCableConn(t *testing.T) {
	assert.Equal(t, CheckNetworkCableConn("ens33"), true)
	assert.Equal(t, CheckNetworkCableConn("ens36"), true)
}

func TestGetDiskSize(t *testing.T) {
	totalSize, err := GetDiskTotalSize()
	assert.Equal(t, err, nil)
	usedSize, err := GetDiskUsedSize()
	assert.Equal(t, err, nil)
	freeSize, err := GetDiskFreeSize()
	assert.Equal(t, err, nil)
	log.Printf("TotalSize:%dGB UsedSize:%dGB FreeSize:%dGB", totalSize>>20, usedSize>>20, freeSize>>20)
}
