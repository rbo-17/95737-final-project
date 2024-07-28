#!/bin/bash

apt update
apt upgrade

# Install DBs
#apt install -y redis mysql-server

# Mount ephemeral storage
# See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/add-instance-store-volumes.html
mkfs -t xfs /dev/nvme1n1
mkdir /ephemeral
mount /dev/nvme1n1 /ephemeral

#mkfs.ext4 -E nodiscard -m0 /dev/nvme1n1
#mount -o discard /dev/nvme1n1 /home/ubuntu/spda
#chown ubuntu:ubuntu /home/ubuntu/spda

mkdir -p /ephemeral/cassandra /ephemeral/mongodb /ephemeral/mysql /ephemeral/redis

# Setup MySQL
# See: https://stackoverflow.com/questions/1795176/how-to-change-mysql-data-directory
apt install -y mysql-server

/etc/init.d/mysql stop

sed -i -E 's/\# datadir/datadir/g' /etc/mysql/mysql.conf.d/mysqld.cnf
sed -i -E 's/\/var\/lib\/mysql/\/ephemeral\/mysql/g' /etc/mysql/mysql.conf.d/mysqld.cnf
sed -i -E 's/\/var\/lib\/mysql/\/ephemeral\/mysql/g' /etc/apparmor.d/usr.sbin.mysqld

/etc/init.d/apparmor reload
/etc/init.d/mysql restart

# Setup Redis
apt install -y redis


# Setup Cassandra
# See:
#   - https://cassandra.apache.org/doc/3.11/cassandra/getting_started/configuring.html
#   - https://cassandra.apache.org/doc/stable/cassandra/getting_started/installing.html

apt install -y openjdk-8-jre-headless
apt install -y apt-transport-https

echo "deb [signed-by=/etc/apt/keyrings/apache-cassandra.asc] https://debian.cassandra.apache.org 41x main" | sudo tee -a /etc/apt/sources.list.d/cassandra.sources.list
deb https://debian.cassandra.apache.org 41x main
curl -o /etc/apt/keyrings/apache-cassandra.asc https://downloads.apache.org/cassandra/KEYS

apt update
apt install -y cassandra

service cassandra stop
DEFAULT_CASSANDRA_DIR='/var/lib/cassandra'
DEFAULT_CASSANDRA_DIR_ESC='\/var\/lib\/cassandra'
NEW_CASSANDRA_DIR='/ephemeral/cassandra'
NEW_CASSANDRA_DIR_ESC='\/ephemeral\/cassandra'
sed -i -E "s/${DEFAULT_CASSANDRA_DIR_ESC}/${NEW_CASSANDRA_DIR_ESC}/g" /etc/cassandra/cassandra.yaml
service cassandra start

#wget -q -O cassandra-keys https://www.apache.org/dist/cassandra/KEYS | sudo apt-key add -
#echo "deb http://www.apache.org/dist/cassandra/debian 40x main" | sudo tee -a /etc/apt/sources.list.d/cassandra.sources.list deb http://www.apache.org/dist/cassandra/debian 40x main

#echo "CONFIG SET protected-mode no" | redis-cli
#
#echo "CONFIG SET dir /ephemeral/redis" | redis-cli
#
#echo "CONFIG REWRITE" | redis-cli
#echo "CONFIG REWRITE" | redis-cli


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

