$env:GOPROXY = "https://goproxy.io"

先从helm里面拉出数据
然后用client部署

使用gin来写数据中心

https://blog.csdn.net/weixin_42506905/article/details/93135684

go get k8s.io/client-go@master
测试
lint
打包
部署

acme.sh --issue -d laki-api.shitangdama.cn --nginx

acme.sh --installcert -d laki-api.shitangdama.cn --key-file /etc/nginx/ssl/laki-api.shitangdama.cn.key --fullchain-file /etc/nginx/ssl/laki-api.cer --reloadcmd  "service nginx force-reload"


chan底层实现 协程调度算法