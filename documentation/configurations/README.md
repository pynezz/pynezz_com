# Configurations

Here's the configurations of the different 3rd party services included in the project.

- [Configurations](#configurations)
  - [OpenResty](#openresty)
    - [nginx.conf](#nginxconf)
    - [pynezz.com.conf](#pynezzcomconf)
    - [pynezz.dev.conf](#pynezzdevconf)
  - [Certificates](#certificates)


## OpenResty

### nginx.conf

```zsh
$ cat /usr/local/openresty/nginx/conf/nginx.conf
```

```nginx
user www-data;
worker_processes 1;

error_log  logs/error.log;
pid        logs/nginx.pid;

events {
    worker_connections  1024;
}


http {
    server_tokens off;
    include       mime.types;
    default_type  application/octet-stream;
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

    access_log  logs/access.log  main;
    sendfile        on;

    #keepalive_timeout  0;
    keepalive_timeout  65;

    gzip  on;

    server {
        listen       80;
        server_name  localhost;

        #charset koi8-r;

        #access_log  logs/host.access.log  main;

        location / {
            root   html;
            index  index.html index.htm;
        }

        #error_page  404              /404.html;

        # redirect server error pages to the static page /50x.html
        #
        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }

        # deny access to .htaccess files, if Apache's document root
        # concurs with nginx's one
        #
        location ~ /\.ht {
            deny  all;
        }
    }

    include ../sites/*;
}

```

### pynezz.com.conf

Here, we'll need to forward/proxy pass the requests from the OpenResty server to the Go server when it's up and running.

```zsh
$ cat /usr/local/openresty/nginx/sites/pynezz.com.conf
```

```nginx

server {
    listen 80;
    listen [::]:80;

    root /usr/local/openresty/nginx/html/pynezz;

    index index.html index.htm;

    server_name www.pynezz.com pynezz.com;

    # just redirect to https
    return 301 https://$server_name/$request_uri;
}

server {
    listen 443 ssl;
    listen [::]:443 ssl;

    root /usr/local/openresty/nginx/html/pynezz;

    ssl_certificate /etc/letsencrypt/live/pynezz.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/pynezz.com/privkey.pem;

    index index.html;

    server_name www.pynezz.com pynezz.com;

    location / {
        try_files $uri $uri/ =404;
    }

    error_page 500 502 503 504  /50x.html;
    location = /50x.html {
        root /usr/local/openresty/nginx/html;
    }

    # allow for security.txt to be served
    location ^~ /.well-known/ {
                allow all;
                default_type text/plain;
    }

    location = /robots.txt {
        access_log off;
        log_not_found off;
    }

    # deny access to paths and files starting with a dot (.env/.git/.htaccess etc.)
    # they shouldn't be accessible anyways, but it's defense in depth
    location ~ /\. {
        deny  all;
        access_log off;
        log_not_found off;
    }
}
```


### pynezz.dev.conf

*identical to `pynezz.com.conf`*

## Certificates

Set up with [Certbot](https://certbot.eff.org/) with `--webroot` parameter like so:

```zsh
$ certbot certonly --webroot -d pynezz.com -d www.pynezz.com
```

```zsh
$ certbot certonly --webroot -d pynezz.dev -d www.pynezz.dev
```

and added the root of the site when requested, which is `/usr/local/openresty/nginx/html/pynezz` in this case.

**Output from adding a certificate:**

```zsh
$ sudo certbot certonly --webroot -w "/usr/local/openresty/nginx/html/pynezz" -d www.pynezz.com                                                      13:06:38
Saving debug log to /var/log/letsencrypt/letsencrypt.log
Plugins selected: Authenticator webroot, Installer None
Requesting a certificate for www.pynezz.com

IMPORTANT NOTES:
 - Congratulations! Your certificate and chain have been saved at:
   /etc/letsencrypt/live/www.pynezz.com/fullchain.pem
   Your key file has been saved at:
   /etc/letsencrypt/live/www.pynezz.com/privkey.pem
   Your certificate will expire on 2024-09-18. To obtain a new or
   tweaked version of this certificate in the future, simply run
   certbot again. To non-interactively renew *all* of your
   certificates, run "certbot renew"
 - If you like Certbot, please consider supporting our work by:

   Donating to ISRG / Let's Encrypt:   https://letsencrypt.org/donate
   Donating to EFF:                    https://eff.org/donate-le
```

After the certificates are added, we need to add the paths to the certificates in the `pynezz.com.conf` and `pynezz.dev.conf` files.

```nginx
ssl_certificate /etc/letsencrypt/live/www.pynezz.com/fullchain.pem;

ssl_certificate_key /etc/letsencrypt/live/www.pynezz.com/privkey.pem;
```

and then we reload the OpenResty server:

```zsh
sudo openresty -s reload
```

Result:
![result](../assets/img/image.png)

> :heavy_check_mark: **Success**<br/>
> The OpenResty server is up and running with the certificates added.
