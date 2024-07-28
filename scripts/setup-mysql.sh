# Check if MySQL is installed
if ! mysql --version &> /dev/null
then
    echo "MySQL is not installed. Please install before running this command. Exiting..."
    exit 1
fi

ENV_FILE=$1
if ! test -f ${ENV_FILE}; then
  echo "Please provide a valid env file that contains MySQL connection data. Exiting..."
  exit 1
fi

source ${ENV_FILE}
mysql -u root -p -e 'CREATE DATABASE '${MYSQL_DB_NAME}';'
mysql -u root -p -e "CREATE USER '"${MYSQL_USERNAME}"'@'localhost' IDENTIFIED BY '"${MYSQL_PASSWORD}"';"
mysql -u root -p -e "GRANT ALL PRIVILEGES ON ${MYSQL_DB_NAME}.* TO '"${MYSQL_USERNAME}"'@'localhost';"
mysql -u ${MYSQL_USERNAME} -p${MYSQL_PASSWORD} < ./setup-mysql.sql