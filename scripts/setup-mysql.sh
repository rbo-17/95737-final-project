# Check if MySQL is installed
if ! mysql --version &> /dev/null
then
    echo "MySQL is not installed. Please install before running this command. Exiting..."
    exit 1
fi

ENV_FILE=$1
if [ -z "${ENV_FILE}" ]; then
  echo "Please provide a valid env file that contains MySQL connection data. Exiting..."
  exit 1
fi

source ${ENV_FILE}

# Set root password
# See: https://stackoverflow.com/questions/7534056/mysql-root-password-change
service mysql stop
mysqld_safe --skip-grant-tables
service mysql start
mysql -u root -e 'USE mysql;';
mysql -u root -e "ALTER USER 'root'@'localhost' IDENTIFIED WITH caching_sha2_password BY '"${MYSQL_PASSWORD}"';";
mysql -u root -p${MYSQL_PASSWORD} -e "FLUSH PRIVILEGES";

# Create database
mysql -u root -p${MYSQL_PASSWORD} -h ${MYSQL_ADDRESS} -e 'CREATE DATABASE '${MYSQL_DB_NAME}';'

# Create user
mysql -u root -p${MYSQL_PASSWORD} -h ${MYSQL_ADDRESS} -e "CREATE USER '"${MYSQL_USERNAME}"'@'%' IDENTIFIED BY '"${MYSQL_PASSWORD}"';"
mysql -u root -p${MYSQL_PASSWORD} -h ${MYSQL_ADDRESS} -e "GRANT ALL PRIVILEGES ON ${MYSQL_DB_NAME}.* TO '"${MYSQL_USERNAME}"'@'%';"

# Create table
mysql -u ${MYSQL_USERNAME} -p${MYSQL_PASSWORD} -h ${MYSQL_ADDRESS} < ./setup-mysql.sql
