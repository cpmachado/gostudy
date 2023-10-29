# database-access

Tutorial: Accessing a relational database

URL: <https://go.dev/doc/tutorial/database-access>

## DB Migration

The script is in the scripts folder.

Command to run:
```shell
docker exec -i database-access_db_1 sh -c 'exec mysql -uroot -p"$MYSQL_ROOT_PASSWORD"' < scripts/migration.sql
```
