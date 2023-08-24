#!/bin/sh

# mac需要安装sshpass

LC_CTYPE="en_US.UTF-8"

# 本地项目路径
local_path="/Users/zouzhujia/Applications/github/bing-image/build"

# 打包的项目文件
service_file="bing"

# 使用的配置文件
config_file="config.yaml"

# 远程服务器信息
remote_host="49.232.222.252"
remote_user="root"
remote_path="/root/workspace/go/bing-image"

# 打包项目
echo "开始打包..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -ldflags "-w -s" -o "$local_path/$service_file"  ./cmd/

# 删除远程文件
echo "删除远程文件文件"
# shellcheck disable=SC2016
sshpass -p "$1" ssh $remote_user@$remote_host "cd $remote_path && rm -rf $service_file && rm -rf $config_file"

# 上传go执行文件
echo "上传后端文件..."
# shellcheck disable=SC2016
sshpass -p "$1" scp -r "$local_path/$service_file" $remote_user@$remote_host:$remote_path
# shellcheck disable=SC2016
sshpass -p "$1" scp -r "$local_path/$config_file" $remote_user@$remote_host:$remote_path

#删除打包文件（不删除前端，前端构建没脚本，手动执行的）
echo "删除后端打包文件"
# shellcheck disable=SC2115
rm -rf "$local_path/$service_file"