#!/bin/bash

dnf -y update
dnf -y upgrade

# Install Redis

# Install MondoDB - see: https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-amazon/
echo '[mongodb-org-7.0]
name=MongoDB Repository
baseurl=https://repo.mongodb.org/yum/amazon/2023/mongodb-org/7.0/x86_64/
gpgcheck=1
enabled=1
gpgkey=https://pgp.mongodb.com/server-7.0.asc' > '/etc/yum.repos.d/mongodb-org-7.0.repo'

sudo yum install -y mongodb-org




# Install Cassandra

# Install MySQL

