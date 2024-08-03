ENV_FILE=$1
if [ -z "${ENV_FILE}" ]; then
  echo "Please provide a valid env file that contains Cassandra connection data. Exiting..."
  exit 1
fi

source ${ENV_FILE}

# Fix Python version. See:
#  - https://stackoverflow.com/a/78422557
#  - https://askubuntu.com/questions/682869/how-do-i-install-a-different-python-version-using-apt-get
add-apt-repository -y ppa:deadsnakes/ppa
apt update
apt install -y python3.9

# See: https://docs.datastax.com/en/cassandra-oss/3.x/cassandra/architecture/archDataDistributeReplication.html
python3.9 /usr/bin/cqlsh.py -u ${CASSANDRA_USERNAME} -p ${CASSANDRA_PASSWORD} -e "CREATE KEYSPACE IF NOT EXISTS \"${CASSANDRA_DB_NAME}\" WITH REPLICATION = { 'class': 'SimpleStrategy', 'replication_factor': 1};"
#python3.9 /usr/bin/cqlsh.py -u ${CASSANDRA_USERNAME} -p ${CASSANDRA_PASSWORD} -e "CREATE ROLE IF NOT EXISTS ${CASSANDRA_USERNAME} WITH PASSWORD = '${CASSANDRA_PASSWORD}' AND LOGIN = true;"
#python3.9 /usr/bin/cqlsh.py -u ${CASSANDRA_USERNAME} -p ${CASSANDRA_PASSWORD} -e "GRANT ALL PERMISSIONS on KEYSPACE \"${CASSANDRA_DB_NAME}\" to ${CASSANDRA_USERNAME};"

# Note: Path pf setup-cassandra.cql may change depending on where this script is run
python3.9 /usr/bin/cqlsh.py -u ${CASSANDRA_USERNAME} -p ${CASSANDRA_PASSWORD} < ./setup-cassandra.cql
