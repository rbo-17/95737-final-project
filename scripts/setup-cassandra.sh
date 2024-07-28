ENV_FILE=$1
if ! test -f ${ENV_FILE}; then
  echo "Please provide a valid env file that contains Cassandra connection data. Exiting..."
  exit 1
fi

source ${ENV_FILE}

# See: https://docs.datastax.com/en/cassandra-oss/3.x/cassandra/architecture/archDataDistributeReplication.html
cqlsh -u cassandra -p cassandra -e "CREATE KEYSPACE IF NOT EXISTS \"${CASSANDRA_DB_NAME}\" WITH REPLICATION = { 'class': 'SimpleStrategy', 'replication_factor': 1};" ${CASSANDRA_ADDRESS}
cqlsh -u cassandra -p cassandra -e "CREATE ROLE IF NOT EXISTS ${CASSANDRA_USERNAME} WITH PASSWORD = '${CASSANDRA_PASSWORD}' AND LOGIN = true;" ${CASSANDRA_ADDRESS}
cqlsh -u cassandra -p cassandra -e "GRANT ALL PERMISSIONS on KEYSPACE \"${CASSANDRA_DB_NAME}\" to ${CASSANDRA_USERNAME};" ${CASSANDRA_ADDRESS}

cqlsh -u ${CASSANDRA_USERNAME} -p ${CASSANDRA_PASSWORD} ${CASSANDRA_ADDRESS} < ./scripts/setup-cassandra.cql
