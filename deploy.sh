#!bin/bash
#Tao file bash de tao bien moi truong
export GOROOT=/usr/local/go
echo $GOROOT
echo "Da cap nhat GOROOT"

export GOPATH=/home/go
echo $GOPATH
echo "Da cap nhat GOPATH"

export PATH="$GOROOT/bin:$PATH"
echo $PATH
echo "Da cap nhat lai PATH"

echo "Chuan bi deploy he thong"

go build service-deploy.go

echo "Da build xong project!"

systemctl stop resident

cp -f service-deploy /usr/sbin/resident-service

echo "Da copy file he thong vao sbin"

cp -f configs/app.yaml /usr/sbin/configs/app.yaml

echo "Da copy config vao sbin"

systemctl start resident

systemctl status resident
