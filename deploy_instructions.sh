### Create /etc/systemd/system/chatroom.service:
### copy from service/chatroom.service

### build
```shell
go build -o server cmd/server/main.go
```

### Then copy it to the deployment location:

```shell
sudo mkdir -p /opt/chatroom
sudo cp server /opt/chatroom/
sudo mkdir -p /opt/chatroom/chatdata
```

### Create a dedicated user for running the service:
```shell
sudo useradd -r -s /bin/false chatroom
sudo chown -R chatroom:chatroom /opt/chatroom
```

### Enable and start the service:
```shell
sudo systemctl enable chatroom
sudo systemctl start chatroom
```

### Check that it's running:
```shell
sudo systemctl status chatroom
```
### You can view logs with:
```shell
sudo journalctl -u chatroom -f
```
### The -f flag follows the logs in real-time, similar to tail -f.


## With docker
### Build the docker image:
```shell
docker build -t chatroom .
docker run -p 9000:9000 -v $(pwd)/chatdata:/root/chatdata chatroom
```

### if you have docker compose
```shell
docker-compose up -d
```