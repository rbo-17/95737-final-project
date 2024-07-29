ENV_FILE=$1
if [ -z "${ENV_FILE}" ]; then
  echo "Please provide a valid env file that contains Redis connection data. Exiting..."
  exit 1
fi

source ${ENV_FILE}

# Set redis as owner of redis dir
chown -R redis:redis /ephemeral/redis

# Stop service for maintenance
DEFAULT_REDIS_DIR='/var/lib/redis'
NEW_REDIS_DIR='/ephemeral/redis'

service redis stop

# Set password
sed -i -E "s|# requirepass foobared|requirepass ${REDIS_PASSWORD}|g" /etc/redis/redis.conf

# Change where data is store (use ephemeral storage)
sed -i -E "s|^dir ${DEFAULT_REDIS_DIR}|dir ${NEW_REDIS_DIR}|g" /etc/redis/redis.conf

# See: https://serverfault.com/questions/942328/redis-reports-read-only-filesystem-but-it-isnt
sed -i 's|/var/lib/redis|/ephemeral/redis|g' /etc/systemd/system/redis.service
sed -i "s|ReadWriteDirectories=-/etc/redis|ReadWriteDirectories=-${NEW_REDIS_DIR}|g" /etc/systemd/system/redis.service
systemctl daemon-reload

# Allow remote connections
sed -i -E "s|^bind 127.0.0.1|bind 0.0.0.0|g" /etc/redis/redis.conf

# Restart service
service redis start
