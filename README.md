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
### postgres数据库运行
```bash
docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres  -e TZ=PRC -v pgdata:/var/lib/postgresql/data --restart=always -d postgres:12
```
# 我的卡组


### /api/getDesksById

#### GET
##### Summary:

查询卡组

##### Description:

查询卡组id，获取卡组信息

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| data | query | id | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | ok | [resp.SuccessResp](#resp.SuccessResp) & object |

### /api/getDesksList

#### GET
##### Summary:

查询卡组列表

##### Description:

查询用户的卡组列表

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| data | query | userId | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [resp.SuccessResp](#resp.SuccessResp) & object |

### /api/putDesks

#### POST
##### Summary:

保存卡组

##### Description:

保存卡组，如果卡组Id没有则创建。有id则覆盖数据库这条记录

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| data | body | 卡组信息 | Yes | [models.Decks](#models.Decks) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [resp.SuccessResp](#resp.SuccessResp) |

### Models


#### models.Cards

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| extra | [ integer ] |  | No |
| main | [ integer ] |  | No |
| side | [ integer ] |  | No |

#### models.Decks

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| cards | [models.Cards](#models.Cards) |  | No |
| id | string |  | No |
| name | string |  | No |
| user_id | integer |  | No |

#### resp.SuccessResp

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| data |  |  | No |
| status | integer |  | No |
