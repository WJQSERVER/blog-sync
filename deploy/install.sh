# /bin/bash

# 检查是否安装前置hugo
if ! command -v hugo &> /dev/null; then
  echo "Hugo 未安装，请先安装hugo"
  exit 1
fi

# 创建目录
mkdir -p /root/data/blog-sync/config
mkdir -p /root/data/blog-sync/log
mkdir -p /root/data/blog-sync/web

# 拉取主程序
VERSION=$(curl -s https://raw.githubusercontent.com/WJQSERVER/blog-sync/main/VERSION)
wget -O /root/data/blog-sync/blog-sync https://github.com/WJQSERVER/blog-sync/releases/download/$VERSION/blog-sync
chmod +x /root/data/blog-sync/blog-sync

# 配置文件
if [ ! -f /root/data/blog-sync/config/config.toml ]; then
    wget -O /root/data/blog-sync/config/config.toml https://raw.githubusercontent.com/WJQSERVER/blog-sync/main/config/config.toml
fi

# 拉取 systemd unit 文件
wget -O /etc/systemd/system/blog-sync.service https://raw.githubusercontent.com/WJQSERVER/blog-sync/main/deploy/blog-sync.service

# 启动服务
systemctl enable blog-sync
systemctl start blog-sync
