FROM postgres:14.5

# Copy the init.sql file into the Docker container
COPY database/init.sql /docker-entrypoint-initdb.d/

# Set the database user and password
ENV POSTGRES_USER=${POSTGRES_USER}
ENV POSTGRES_PASSWORD=${POSTGRES_PASSWORD}
ENV POSTGRES_DB=${POSTGRES_DB}

# Start the database using the environment variables
ENTRYPOINT ["docker-entrypoint.sh"]

CMD ["postgres"]
