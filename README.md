# mycards_store

## docker 运行

```bash
docker run -it --rm -e DB_INIT="true" -p 8080:8080 1327605995/mycard-store:1.0
```

| 环境变量    | 默认值    | 描述                           |
|---------|--------|------------------------------|
| DB_INIT | false  | 数据库表初始化true开启                |
| DB_TYPE | sqlite | 数据库类型(postgres,mysql,sqlite) |
| DB_NAME |        | 数据库表名                        |
| DB_HOST |        | 数据库主机ip                      |
| DB_USER |        | 数据库用户                        |
| DB_PWD  |        | 数据库密码                        |
| DB_PORT |        | 数据库端口                        |
| API     | false  | swagger请求文档                  |
### postgres数据库运行
```bash
docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres  -e TZ=PRC -v pgdata:/var/lib/postgresql/data --restart=always -d postgres:12
```
# swagger文档
添加API参数启动即可访问
访问地址 http://localhost:8080/swagger/index.html
```bash
docker run -it --rm -e DB_INIT="true" -e API="true" -p 8080:8080 1327605995/mycard-store:1.0
```
![img.png](./docs%2Fimg.png)