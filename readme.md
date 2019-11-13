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

acme.sh --issue  -d laki_api.shitangdama.cn   --nginx

acme.sh --installcert -d laki_api.shitangdama.cn --key-file /etc/nginx/ssl/laki_api.shitangdama.cn.key --fullchain-file /etc/nginx/ssl/laki_api.cer --reloadcmd  "service nginx force-reload"

server {
        listen    80;
        listen    443 ssl;
        server_name drone.shitangdama.cn;

        ssl_certificate_key        /etc/nginx/ssl/drone.shitangdama.cn.key;
        ssl_certificate            /etc/nginx/ssl/drone.cer;

        location / {
                proxy_set_header X-Forwarded-For $remote_addr;
                proxy_set_header X-Forwarded-Proto $scheme;
                proxy_set_header Host $host;
                proxy_pass http://127.0.0.1:8080;
        }
}