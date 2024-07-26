#!/bin/bash

apt update
apt upgrade

# Install DBs
apt install -y redis mysql-server

# Mount ephemeral storage
# See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/add-instance-store-volumes.html
mkfs -t xfs /dev/nvme1n1
mkdir /ephemeral
mount /dev/nvme1n1 /ephemeral

#mkfs.ext4 -E nodiscard -m0 /dev/nvme1n1
#mount -o discard /dev/nvme1n1 /home/ubuntu/spda
#chown ubuntu:ubuntu /home/ubuntu/spda

## Install Redis
#
## Install MondoDB - see: https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-amazon/
#echo '[mongodb-org-7.0]
#name=MongoDB Repository
#baseurl=https://repo.mongodb.org/yum/amazon/2023/mongodb-org/7.0/x86_64/
#gpgcheck=1
#enabled=1
#gpgkey=https://pgp.mongodb.com/server-7.0.asc' > '/etc/yum.repos.d/mongodb-org-7.0.repo'
#
#sudo yum install -y mongodb-org
#
#
#
#
## Install Cassandra
#
## Install MySQL

