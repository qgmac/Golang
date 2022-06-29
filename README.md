# Golang demo

##### 2022-06-29
```shell
# 更新go package
# 更新所有包 
go get -u
# 更新指定包名
go get 包名

# git 打标签
git tag -a "v0.0.5" -m ""

# 推送标签到远程服务器 
git push --tags

# git 修改源的三种方法
  # 1. 通过命令直接修改远程仓库地址
    git remote 查看所有远程仓库
    git remote xxx 查看指定远程仓库地址
    git remote set-url origin 你新的远程仓库地址
  # 2. 先删除在添加你的远程仓库
      git remote rm origin
      git remote add origin 你的新远程仓库地址
  # 3. 修改本地 .git 配置文件
      vi config 
    
```
