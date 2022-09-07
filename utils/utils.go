package utils

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
)

func IsPathExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func CreatDir(dirPath string) error {
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.Chmod(dirPath, 0777)
	if err != nil {
		return err
	}
	return nil
}

func RemoveFile(filePath string) error {
	return os.Remove(filePath)
}

func RemoveAll(dirPath string) error {
	return os.RemoveAll(dirPath)
}

func GetPrevDir(dirPath string) string {
	substr := func(s string, pos, length int) string {
		runes := []rune(s)
		l := pos + length
		if l > len(runes) {
			l = len(runes)
		}
		return string(runes[pos:l])
	}
	cnt := len(dirPath) - 1
	if dirPath[cnt:] == "/" {
		dirPath = dirPath[:cnt]
	}
	return substr(dirPath, 0, strings.LastIndex(dirPath, "/"))
}

func Contain(a []interface{}, e interface{}) bool {
	cnt := len(a)
	for i := 0; i < cnt; i++ {
		if a[i] == e {
			return true
		}
	}
	return false
}

func ContainEx(a interface{}, f func(predicate interface{}) bool) bool {
	src := reflect.ValueOf(a)
	switch src.Kind() {
	case reflect.Slice, reflect.Array:
		count := src.Len()
		for i := 0; i < count; i++ {
			e := src.Index(i).Interface()
			if f(e) {
				return true
			}
		}
	default:
	}
	return false
}

func ReadAll(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func WriteAll(path string, data []byte) error {
	return os.WriteFile(path, data, 0777)
}

func Cmd(cmd string) (string, error) {
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(result)), nil
}

func CheckNetworkCableConn(deviceName string) bool {
	if res, err := Cmd(fmt.Sprintf("cat /sys/class/net/%s/carrier", deviceName)); err == nil && res == "1" {
		return true
	}
	return false
}

func InSlice(ss []string, s string) bool {
	for _, v := range ss {
		if s == v {
			return true
		}
	}
	return false
}

// 获取磁盘总大小(KB)
func GetDiskTotalSize() (uint64, error) {
	s, _ := Cmd("df | awk '{print $2}' | awk '{sum+=$1}END{print sum}'")
	size, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return size, nil
}

// 获取磁盘已用大小(KB)
func GetDiskUsedSize() (uint64, error) {
	s, _ := Cmd("df | awk '{print $3}' | awk '{sum+=$1}END{print sum}'")
	size, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return size, nil
}

// 获取磁盘可用大小(KB)
func GetDiskFreeSize() (uint64, error) {
	s, _ := Cmd("df | awk '{print $4}' | awk '{sum+=$1}END{print sum}'")
	size, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return size, nil
}

func GetDiskFreePercent() (string, error) {
	freeSize, err := GetDiskFreeSize()
	if err != nil {
		return "", err
	}
	usedSize, err := GetDiskUsedSize()
	if err != nil {
		return "", err
	}
	fp := (freeSize * 100) / (freeSize + usedSize)
	return fmt.Sprintf("%d%%", fp), nil
}

func SuitableDisplaySize(size int64) string {
	if size > (1 << 30) {
		return strconv.FormatInt(size>>30, 10) + "GB"
	} else if size > (1 << 20) {
		return strconv.FormatInt(size>>20, 10) + "MB"
	} else if size > (1 << 10) {
		return strconv.FormatInt(size>>10, 10) + "KB"
	} else {
		return strconv.FormatInt(size, 10) + "B"
	}
}

// 获取指定目录大小(KB)
func GetSizeOfDir(dirPath string) (int64, error) {
	s, err := Cmd(fmt.Sprintf("du --max-depth=0 %s | awk '{print $1}'", dirPath))
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	i <<= 10
	return i, nil
}
