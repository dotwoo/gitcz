# gitcz

Golang 版本 Git Commitizen，commit 规范工具

该项目 clone 自 [xiaoqidun](https://github.com/xiaoqidun/gitcz),

1. 添加了提交正文和引用部分
2. 去掉 git 包含的 release 部分

# 快速安装

go get -u github.com/dotwoo/gitcz

# 编译安装

```
git clone https://github.com/dotwoo/gitcz.git
cd gitcz
go build gitcz.go
```

# 手动安装

1. 根据系统架构下载为你编译好的[二进制文件](https://github.com/dotwoo/gitcz/releases)
2. 将下载好的二进制文件重命名为 gitcz 并保留后缀
3. 把 gitcz 文件移动到系统 PATH 环境变量中的目录下
4. windows 外的系统需使用 chmod 命令赋予可执行权限

# 使用说明

```shell
git add .
# 简单提交
gitcz
# 包含详细信息
# gitcz -f
# 包含引用的提交
# gitcz -r
git push
```

# 规范文档

gitcz 使用：[Git 提交信息样式指南](https://github.com/udacity/frontend-nanodegree-styleguide-zh/blob/master/%E5%89%8D%E7%AB%AF%E5%B7%A5%E7%A8%8B%E5%B8%88%E7%BA%B3%E7%B1%B3%E5%AD%A6%E4%BD%8D%E6%A0%B7%E5%BC%8F%E6%8C%87%E5%8D%97%20-%20Git.md)
