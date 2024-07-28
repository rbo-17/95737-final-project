ENV_FILE=$1
if ! test -f ${ENV_FILE}; then
  echo "Please provide a valid env file that contains MySQL connection data. Exiting..."
  exit 1
fi

source ${ENV_FILE}
echo 'db.createUser( { user: "'${MONGODB_USERNAME}'", pwd:"'${MONGODB_PASSWORD}'", roles: [ { role: "readWrite", db: "'${MONGODB_DB_NAME}'" } ] } );' > setup-mongodb.js
mongosh "${MONGODB_ADDRESS}:27017/CMU95737" setup-mongodb.js
