# sqlc-buf-pg-template
SQL query driven backend template with sqlc, gRPC and postgres and also visualized by pgweb and grafana.

## HOW TO USE
First, write proto files and edit buf yamls (`buf.work.yaml`, `buf.gen.yaml` and `buf.yaml`)

Second, write the configuration for sqlc in `sqlc.yaml`, SQL Table Schema in `schema.sql` and SQL query for the services in `query.sql`.
See [sqlc メモ](https://zenn.dev/voluntas/scraps/f7f3f7e419af31) [query annotations](https://docs.sqlc.dev/en/stable/reference/query-annotations.html).

Finally generate pb files in `/gen` and an autogenerated package for handling SQL by
```
docker-compose -f docker-compose.autogen.yml up --force-recreate
```




## TODO
 - `buf lint`
 - `buf breaking`
cf. https://zenn.dev/y16ra/articles/23e0cd68251554
