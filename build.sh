#!/usr/bin/env bash

set -x

# 获取程序版本
Version="1.0.1"

# 获取当前时间
BuildTime=$(date +'%Y-%m-%d %H:%M:%S')

# 将以上变量序列化至 LDFlags 变量中
LDFlags=" \
    -X 'gtp/version.version=${Version}' \
    -X 'gtp/version.buildTime=${BuildTime}' \
"

# 编译
go build -ldflags "${LDFlags}" -o ./release/gtp main.go
