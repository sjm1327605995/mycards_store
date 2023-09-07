# mycards_store
## docker 运行
postgres启动
```bash
docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres  -e TZ=PRC -v pgdata:/var/lib/postgresql/data --restart=always -d postgres:12
```
mycards_store启动
```bash
 docker run -it --rm -e postgre="host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai" 1327605995/mycard-store:1.0
```
