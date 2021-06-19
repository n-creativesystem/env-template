# variable-template

テンプレートエンジンを使って環境変数もしくは変数をセットします。

## Sample

Environment Variables

```sh
docker run \
    --rm \
    --workdir /github/workspace \
    -v $PWD/:/github/workspace/ \
    -e INPUT_INPUT=Sample.tpl.yml \
    -e INPUT_OUTPUT=Sample.yml \
    -e version=alpine \
    -e DEBUG=true variable-template
```

| Key     | Value          |
| :------ | :------------- |
| INPUT   | Sample.tpl.yml |
| OUTPUT  | Sample.yml     |
| version | alpine         |

Input Template file

```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: nginx
spec:
  replicas: 1
  template:
    metadata:
      labels:
        editor: vscode
    spec:
      containers:
        - name: nginx
          image: nginx:${{version}}
          resources: {}
          volumeMounts:
            - name: config
              mountPath: /etc/nginx/conf.d/
              readOnly: true
      volumes:
        - name: config
          configMap:
            name: nginx-config
            items:
              - key: default.conf
                path: default.conf
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: nginx-config
data:
  default.conf: |
    server {
        server_tokens off;
        listen 443 ssl http2;
        server_name localhost;
        rewrite_log on;
        charset utf-8;
        client_max_body_size    2G;
        ssl_certificate /etc/nginx/ssl/cert.csr;
        ssl_certificate_key /etc/nginx/ssl/key.pem;
        proxy_set_header    Host    $host;
        proxy_set_header    X-Real-IP    $remote_addr;
        proxy_set_header    X-Forwarded-Host       $host;
        proxy_set_header    X-Forwarded-Server    $host;
        proxy_set_header    X-Forwarded-For    $proxy_add_x_forwarded_for;
        real_ip_header 	X-Forwarded-For;
        gzip on;
        gzip_types text/css text/javascript application/json application/font-woff application/font-tff image/gif image/png image/jpeg application/octet-stream text/plain font/woff2;
        location / {
            root /usr/share/nginx/html;
            try_files $uri $uri/ /index.html;
        }
    }
```

Output file

```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: nginx
spec:
  replicas: 1
  template:
    metadata:
      labels:
        editor: vscode
    spec:
      containers:
        - name: nginx
          image: nginx:alpine
          resources: {}
          volumeMounts:
            - name: config
              mountPath: /etc/nginx/conf.d/
              readOnly: true
      volumes:
        - name: config
          configMap:
            name: nginx-config
            items:
              - key: default.conf
                path: default.conf
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: nginx-config
data:
  default.conf: |
    server {
        server_tokens off;
        listen 443 ssl http2;
        server_name localhost;
        rewrite_log on;
        charset utf-8;
        client_max_body_size    2G;
        ssl_certificate /etc/nginx/ssl/cert.csr;
        ssl_certificate_key /etc/nginx/ssl/key.pem;
        proxy_set_header    Host    $host;
        proxy_set_header    X-Real-IP    $remote_addr;
        proxy_set_header    X-Forwarded-Host       $host;
        proxy_set_header    X-Forwarded-Server    $host;
        proxy_set_header    X-Forwarded-For    $proxy_add_x_forwarded_for;
        real_ip_header  X-Forwarded-For;
        gzip on;
        gzip_types text/css text/javascript application/json application/font-woff application/font-tff image/gif image/png image/jpeg application/octet-stream text/plain font/woff2;
        location / {
            root /usr/share/nginx/html;
            try_files $uri $uri/ /index.html;
        }
    }
```
