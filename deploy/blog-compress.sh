#!/bin/bash

# 克隆 Git 仓库
cd /root/data/hugo_compress
git clone git@github.com:WJQSERVER/blog.git

#构建
cd /root/data/hugo_compress/blog
hugo

# 打包文件
tar -czvf /root/data/hugo_download/hugo.tar.gz -C /root/data/hugo_compress/blog/public .