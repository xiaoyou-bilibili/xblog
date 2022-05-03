# 前端版本


> 如果想前后端使用同一个域名需要设置伪静态，内容如下（前提是后端端口是2333且和前段部署在同一个服务器上）

```
underscores_in_headers on;
location /{
 try_files $uri $uri/ /index.php?$args;
}
location /admin{
 try_files $uri $uri/ /admin/index.html?$args;
}
location ~^/(api|assets|upload|plugins){
    proxy_pass http://127.0.0.1:2333;
    proxy_set_header X-Real-IP $remote_addr;
    add_header Cache-Control no-cache;
}
location /ws{
   proxy_pass http://127.0.0.1:2333;
   proxy_http_version 1.1;
   proxy_set_header Upgrade $http_upgrade;
   proxy_set_header Connection "Upgrade";
   proxy_set_header X-Real-IP $remote_addr;
}
```

域名啥的都在`index.php`里面，自己按照要求改吧