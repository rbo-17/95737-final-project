#!/bin/bash

apt update
apt upgrade

# Mount ephemeral storage
## See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/add-instance-store-volumes.html
mkfs -t xfs /dev/nvme1n1
mkdir /ephemeral
mount /dev/nvme1n1 /ephemeral

# Make folders for each database
mkdir -p /ephemeral/cassandra /ephemeral/mongodb /ephemeral/mysql /ephemeral/redis


# Setup MySQL
## See: https://stackoverflow.com/questions/1795176/how-to-change-mysql-data-directory
apt install -y mysql-server

## Stop service for maintenance
/etc/init.d/mysql stop

## Set generous folder permissions
DEFAULT_MYSQL_DIR='/var/lib/mysql'
NEW_MYSQL_DIR='/ephemeral/mysql'

chmod -R a+rwx ${NEW_MYSQL_DIR}

## Copy over db contents
cp -R -p "${DEFAULT_MYSQL_DIR}/." "${NEW_MYSQL_DIR}/"

## Change where data is store (use ephemeral storage)
sed -i "s|# datadir|datadir|g" /etc/mysql/mysql.conf.d/mysqld.cnf
sed -i "s|${DEFAULT_MYSQL_DIR}|${NEW_MYSQL_DIR}|g" /etc/mysql/mysql.conf.d/mysqld.cnf
sed -i "s|${DEFAULT_MYSQL_DIR}|${NEW_MYSQL_DIR}|g" /etc/apparmor.d/usr.sbin.mysqld

## Allow remote connections
sed -i -E 's|^bind-address\s+= 127.0.0.1|bind-address = 0.0.0.0|g' /etc/mysql/mysql.conf.d/mysqld.cnf

## Restart service
/etc/init.d/apparmor reload
/etc/init.d/mysql start

## TODO: IMPORTANT - The setup-mysql.sh file currently must be ran manually on the instance (run transfer-to-instance)


# Setup Redis
apt install -y redis

## Set generous folder permissions
NEW_REDIS_DIR='/ephemeral/redis'
chmod -R a+rwx ${NEW_REDIS_DIR}

# TODO: IMPORTANT - The setup-redis.sh file currently must be ran manually on the instance (run transfer-to-instance)


# Setup Cassandra
# See:
#   - https://cassandra.apache.org/doc/3.11/cassandra/getting_started/configuring.html
#   - https://cassandra.apache.org/doc/stable/cassandra/getting_started/installing.html

## Install prerequisites
apt install -y openjdk-8-jre-headless apt-transport-https

## Update repo list
echo "deb [signed-by=/etc/apt/keyrings/apache-cassandra.asc] https://debian.cassandra.apache.org 41x main" | sudo tee -a /etc/apt/sources.list.d/cassandra.sources.list
deb https://debian.cassandra.apache.org 41x main
curl -o /etc/apt/keyrings/apache-cassandra.asc https://downloads.apache.org/cassandra/KEYS
apt update

## Install cassandra
apt install -y cassandra net-tools

## Stop service for maintenance
service cassandra stop

DEFAULT_CASSANDRA_DIR='/var/lib/cassandra'
NEW_CASSANDRA_DIR='/ephemeral/cassandra'

## Set generous folder permissions
chmod -R a+rwx ${NEW_CASSANDRA_DIR}

## Replace default dir with ephemeral
sed -i -E "s|${DEFAULT_CASSANDRA_DIR}|${NEW_CASSANDRA_DIR}|g" /etc/cassandra/cassandra.yaml

## Allow connections from remote hosts
sed -i "s|rpc_address: localhost|rpc_address: 0.0.0.0|g" /etc/cassandra/cassandra.yaml
PRIVATE_IP=$(ifconfig | grep -Eo 'inet (addr:)?([0-9]*\.){3}[0-9]*' | grep -Eo '([0-9]*\.){3}[0-9]*' | grep -v '127.0.0.1')
sed -i "s|# broadcast_rpc_address: 1.2.3.4|broadcast_rpc_address: ${PRIVATE_IP}|g" /etc/cassandra/cassandra.yaml

## Restart service
service cassandra start

## TODO: IMPORTANT - The setup-cassandra.sh file currently must be ran manually on the instance (run transfer-to-instance)

# Setup MongoDB
# See: https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-ubuntu/

## Update repo list
curl -fsSL https://www.mongodb.org/static/pgp/server-7.0.asc | \
   gpg -o /usr/share/keyrings/mongodb-server-7.0.gpg \
   --dearmor
echo "deb [ arch=amd64,arm64 signed-by=/usr/share/keyrings/mongodb-server-7.0.gpg ] https://repo.mongodb.org/apt/ubuntu jammy/mongodb-org/7.0 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-7.0.list
apt update

## Install mongodb
apt install -y mongodb-org

## Set generous folder permissions
NEW_MONGODB_DIR='/ephemeral/mongodb'
DEFAULT_MONGODB_DIR='/var/lib/mongodb'

chmod -R a+rwx ${NEW_MONGODB_DIR}

## Replace default dir with ephemeral
sed -i "s|dbPath: ${DEFAULT_MONGODB_DIR}|dbPath: ${NEW_MONGODB_DIR}|g" /etc/mongod.conf

## Allow connections from remote hosts
sed -i "s|bindIp: 127.0.0.1|bindIp: 0.0.0.0|g" /etc/mongod.conf

## Restart service
sudo systemctl start mongod

# Write done message to home directory
touch /home/ubuntu/cloud-init-done.txt
