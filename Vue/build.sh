#!/bin/bash

set -e

dockerDomain="registry.cn-shenzhen.aliyuncs.com"

dockerNamespace="unliar"

projectName="app-deploy"

read -t 30 -p "输入dockerhub的用户名:" dockerhubUser

read -t 30 -p "输入对应用户密码:" dockerhubPass

docker login -u $dockerhubUser -p $dockerhubPass $dockerDomain

read -t 30 -p "输入项目的tag:" projectTag

FullRemoteImageName="$dockerDomain/$dockerNamespace/$projectName:$projectTag"

FullLocalImageName="$dockerNamespace/$projectName:$projectTag"

echo "完整本地IMAGE名称:$FullLocalImageName"
echo "完整远程IMAGE名称:$FullRemoteImageName"

echo "构建本地"
docker build -t $FullLocalImageName .

echo "tag标记"
docker tag $FullLocalImageName $FullRemoteImageName

echo "远程推送"
docker push $FullRemoteImageName

echo "清理本地镜像"

docker rmi $FullRemoteImageName

docker rmi $FullLocalImageName
