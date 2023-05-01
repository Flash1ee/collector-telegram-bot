FROM migrate/migrate

COPY ./db /db

ENTRYPOINT []
CMD migrate -path '/db' -database postgres://$USER:$PSWD@$HOST/$DB?sslmode=disable up \
    && echo "y" | migrate -path '/db' -database postgres://$USER:$PSWD@$HOST/$DB?sslmode=disable down
