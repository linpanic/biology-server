### 简介

每个实验室的管理都难以避免的一个问题是实验动物信息的记录。尤其是当实验室保存的品系过多、实验室人员操作多，单纯的纸质或者excel表格难以维护，更无法解决同步和远程的问题。因此解决实验室品系记录是一个迫在眉睫的问题。

这是一个本意对应于动物品系存储数据库，但是应该也可以适用于其他的场景。您可以使用本系统来集中化的存储您需要的动物的品系信息。



### 品系的新增

> 考虑到实验室内部分品系创建过程中可能不会命名，因此我们余留了很多可以空置的字段供使用者后期再添加。唯一要求添加的就是实验室内编号。

序列号是需要从后端获取的，会自动为您寻找到一个没有利用过的从小到大的顺延号码，如果您有需求可以手动指定，序列号不可重复。

品系名是一个品系的最基础属性，不应重复，但是允许空置。

您可以添加多条简称，注释与简称类似，您也可以用同样的方法添加多条。

额外信息是为了能够方便您使用的一些键值对储存区。

基因信息是该品系中包含的基因信息的合集，其中的每一条都代表着一个基因修饰情况。



### 搜索与排序

支持全字段搜索（包括品系信息和基因信息全部游泳字段），且可对所选字段进行正向或逆向的排序。



##### 使用的技术

数据库: *sqlite* (如果数据量大建议换成pgsql/mysql) 使用的驱动为gorm.io/driver/sqlite,故而无法交叉编译

数据库框架: *gorm*

web框架: *gin*

鉴权库: *jwt-go*

日志框架: *logrus*



##### 编译方式
```shell
./build.bat
```



#### 如果需要自己搭建前端，可以根据下方的接口文档去搭建，如果想直接用现成的，可以参考

https://github.com/sugarbecomer/biology-vue



### 接口文档

### 一：用户

#### 1：注册

##### Post http://127.0.0.1:10080/register

##### Request

|    参数    | 类型     | 必填  | 说明                             |
|:--------:|--------|-----|--------------------------------|
| username | string    | 是   | 用户名                            |
| password | string | 是   | 密码                             |
|   time   | int    | 是   | 10位数字时间戳                       |
|   sign   | string    | 是   | MD5(username+time+password)转大写 |


##### Response

|     参数     | 类型     | 必填  | 说明         |
|:----------:|--------|-----|------------|
|    code    | int    | 是   | 200为成功，其他为失败 |
|  message   | string | 是   | 消息内容       |



#### 2：登陆

##### Post http://127.0.0.1:10080/login

##### Request

|    参数    | 类型     | 必填  | 说明                                                 |
|:--------:|--------|-----|----------------------------------------------------|
| username | string    | 是   | 用户名                                                |
| password | string | 是   | MD5(MD5(密码))转大写                                    |
|   time   | int    | 是   | 10位数字时间戳                                           |
|   sign   | string    | 是   | MD5(username+time+password)转大写,这里的password为2次MD5后的 |

##### Response

|     参数     | 类型     | 必填  | 说明         |
|:----------:|--------|-----|------------|
|    code    | int    | 是   | 200为成功，其他为失败 |
|  message   | string | 是   | 消息内容       |
|    data    | object | 是   | 数据         |
| data.token | string | 是   | token      |

### 以下所有请求都需要在http header 带上x-token字段,值为返回的data.token



#### 3：修改密码(todo)

##### Post http://127.0.0.1:10080/change_password

##### Request

|      参数      | 类型     | 必填  | 说明                                 |
|:------------:|--------|-----|------------------------------------|
| old_password | string | 是   | 旧密码                                |
| new_password | string | 是   | 新密码                                |
|     time     | int    | 是   | 10位数字时间戳                           |
|     sign     | string    | 是   | MD5(old_password+time+new_password)转大写 |

##### Response

|     参数     | 类型     | 必填  | 说明         |
|:----------:|--------|-----|------------|
|    code    | int    | 是   | 200为成功，其他为失败 |
|  message   | string | 是   | 消息内容       |



### 二：品系

#### 1:新增品系名

##### Post http://127.0.0.1:10080/biology/strain_add

##### Request

|            参数            | 类型            | 必填  | 说明                        |
|:------------------------:|---------------|-----|---------------------------|
|       strain_name        | string        | 否   | 品系名，可能为空                  |
|          number          | string        | 是   | 品系序列号，序列号需要从接口获取          |
|        short_name        | array(string) | 否   | 简称，一个品系可能对应多个简称           |
|     annotate      | array(string) | 否   | 品系名注解， 一个品系可能对应多个注解       |
|       extra       | array(object) | 否   | 品系额外信息，一个品系可能对应多个额外信息     |
|     extra.extra_key      | string        | 否   | 品系额外信息key                 |
|    extra.extra_value     | string| 否   | 品系额外信息value               |
|          allele          | array(object) | 否   | 基因,一个品系可能对应多个基因   |
|    allele.name    | string       | 否   | 基因名                       |
|  allele.annotate  | array(string) | 否   | 基因注解，一个基因可能对应多个注解         |
|      allele.genome       | string        | 否   | 基因修饰情况                    |
|      allele.serial       | string        | 否   | 对应第几条染色体，一般为罗马数字          |
|       allele.extra       | array(object) | 否   | 基因的额外信息，一个基因可能对应多个额外信息 |
|  allele.extra.extra_key  | string        | 否   | 基因的额外信息key             |
| allele.extra.extra_value | string| 否   | 基因的额外信息value          |


##### Response

|   参数    | 类型     | 必填  | 说明           |
|:-------:|--------|-----|--------------|
|  code   | int    | 是   | 200为成功，其他为失败 |
| message | string | 是   | 消息内容         |



#### 2：展示品系名列表

##### Post http://127.0.0.1:10080/biology/strain_list


##### Request

|    参数     | 类型     | 必填 | 说明                   |
|:---------:|--------| ---- |----------------------|
|  page_no  | int    | 否   | 查询第几页，默认1            |
| page_size | int    | 否   | 每页查询数量，默认10          |
|   field   | string | 否   | 查询条件，可为空             |
|   order   | string | 否   | 查询条件的降序/升序排列 desc/asc |
|    key    | string | 否   | 查询关键词                |

##### Response

|              参数               | 类型            | 必填  | 说明                      |
|:-----------------------------:|---------------|-----|-------------------------|
|            page_no            | int           | 是   | 第几页                     |
|           page_size           | int           | 是   | 每页数量                    |
|             total             | int           | 是   | 总条数                     |
|             data              | object        | 是   | 数据                      |
|       data.strain_list        | array(object) | 是   | 品系列表                    |
|     strain_list.strain_id     | int           | 是   | 品系ID                    |
|       strain_list.name        | string        | 否   | 品系名，可能为空                |
|      strain_list.number       | string        | 是   | 序列号，一般#开头               |
|    strain_list.short_name     | array(string) | 否   | 简称，一个品系可能对应             |
|     strain_list.annotate      | array(string) | 否   | 品系注解， 一个品系可能对应多个注解      |
|       strain_list.extra       | array(object) | 否   | 品系的额外信息，一个品系可能对应多个额外信息  |
|  strain_list.extra.extra_key  | string        | 是   | 品系额外信息key               |
| strain_list.extra.extra_value | string        | 是   | 品系额外信息value             |
|      strain_list.allele       | array(object) | 否   | 基因，一个品系可能对应多个基因，品系名可能为空 |
|     allele.id     | int           | 否   | 基因ID                    |
|    allele.name    | string        | 否   | 基因名                     |
|  allele.annotate  | array(string) | 否   | 基因注解，一个基因可能对应多个注解       |
|   allele.genome   | string        | 否   | 基因修饰情况                  |
|   allele.serial   | string        | 否   | 对应第几条染色体，一般为罗马数字        |
|         allele.extra          | array(object) | 否   | 基因的额外信息，一个基因可能对应多个额外信息  |
|    allele.extra.extra_key     | string        | 是   | 基因的额外信息key              |
|   allele.extra.extra_value    | string        | 是   | 基因的额外信息value            |



#### 3：修改品系数据

##### Post http://127.0.0.1:10080/biology/strain_update

##### Request

|            参数            | 类型            | 必填  | 说明                     |
|:------------------------:|---------------|-----|------------------------|
|            id            | int           | 是   | 品系ID                   |
|       strain_name        | stri![img.png](img.png)ng        | 否   | 品系名，可能为空               |
|          number          | string        | 否   | 品系序列号                  |
|        short_name        | array(string) | 否   | 简称，一个品系可能对应多个简称        |
|     strain_annotate      | array(string) | 否   | 品系名注解， 一个品系可能对应多个注解    |
|       strain_extra       | array(object) | 否   | 品系额外信息，一个品系可能对应多个额外信息 |
|  strain_extra.extra_key  | string        | 是   | 品系额外信息key              |
|  strain_extra.extra_value  | string| 是   | 品系额外信息value            ||
|          allele          | array(object) | 否   | 基因,一个品系可能对应多个基因                             |
|        allele.id         | int           | 否   | 基因ID                              |
|       allele.name        | string        | 否   | 基因名                               |
|     allele.annotate      | array(string) | 否   | 基因注解，一个基因可能对应多个注解                 |
|      allele.genome       | string        | 否   | 基因修饰情况                            |
|      allele.serial       | string        | 否   | 对应第几条染色体，一般为罗马数字 |
|       allele.extra       | array(object) | 否   | 基因额外信息，一个基因可能对应多个额外信息             |
|  allele.extra.extra_key  | string        | 是   | 基因额外信息key                         |
| allele.extra.extra_value | string| 是   | 基因额外信息value                       |

##### Response

|   参数    | 类型     | 必填  | 说明           |
|:-------:|--------|-----|--------------|
|  code   | int    | 是   | 200为成功，其他为失败 |
| message | string | 是   | 消息内容         |



#### 4：删除品系数据

##### Post http://127.0.0.1:10080/biology/strain_delete

##### Request

|             参数              | 类型          | 必填  | 说明                                  |
|:---------------------------:| ------------- |-----|-------------------------------------|
|          strain_id          | int        | 是   | 品系ID                           |


##### Response

|   参数    | 类型     | 必填  | 说明           |
|:-------:|--------|-----|--------------|
|  code   | int    | 是   | 200为成功，其他为失败 |
| message | string | 是   | 消息内容         |



#### 5：获取序列号

##### Post http://127.0.0.1:10080/biology/get_number

##### Request

|  参数  | 类型     | 必填  | 说明            |
|:----:|--------|-----|---------------|
| time | int    | 是   | 10位时间戳        |
| sign | string | 是   | 10位时间戳的MD5转大写 |

##### Response

|   参数    | 类型     | 必填  | 说明           |
|:-------:|--------|-----|--------------|
|  code   | int    | 是   | 200为成功，其他为失败 |
| message | string | 是   | 消息内容         |
|  data   | object | 是   | 数据 |
|  data.number  | string | 是   | 序列号          |



#### 6：修改基因数据

##### Post http://127.0.0.1:10080/biology/allele_update

##### Request

|   参数   | 类型            | 必填  | 说明                                |
|:------:|---------------|-----|-----------------------------------|
|   Id   | int           | 是   | 品系ID                              |
| allele | array(object) | 否   | 基因,一个品系可能对应多个基因                             |
|    allele.id    | int           | 否   | 基因ID                              |
|   allele.name   | string        | 否   | 基因名                               |
| allele.annotate | array(string) | 否   | 基因注解，一个基因可能对应多个注解                 |
|   allele.genome   | string        | 否   | 基因修饰情况                            |
|     allele.serial      | string        | 否   | 对应第几条染色体，一般为罗马数字 |
|            allele.extra            | array(object) | 否   | 基因额外信息，一个基因可能对应多个额外信息             |
|       allele.extra.extra_key       | string        | 是   | 基因额外信息key                         |
|      allele.extra.extra_val      | string| 是   | 基因额外信息value                       |

##### Response

|   参数    | 类型     | 必填  | 说明           |
|:-------:|--------|-----|--------------|
|  code   | int    | 是   | 200为成功，其他为失败 |
| message | string | 是   | 消息内容         |



#### 7：搜索基因列表

##### Post http://127.0.0.1:10080/biology/allele_search

##### Request

|    参数    | 类型            | 必填  | 说明                                 |
|:--------:|---------------|-----|------------------------------------|
|   name   | string        | 否   | 基因名                               |

##### Response

|      参数       | 类型            | 必填  | 说明                    |
|:-------------:|---------------|-----|-----------------------|
|     code      | int           | 是   | 200为成功，其他为失败          |
|    message    | string        | 是   | 消息内容                  |
|     data      | object        | 是   | 基因数据                  |
|  data.allele  | array(object) | 否   | 基因,一个品系可能对应多个基因       |
|   allele.id   | int           | 否   | 基因ID                  |
|  allele.name  | string        | 否   | 基因名                  |
| allele.genome | string        | 否   | 基因修饰情况                |
| allele.serial | string        | 否   | 对应第几条染色体，一般为罗马数字      |



### 三：基因

#### 1：基因查询

##### Request

##### Post http://127.0.0.1:10080/biology/allele_list


|    参数     | 类型     | 必填 | 说明                   |
|:---------:|--------| ---- |----------------------|
|  page_no  | int    | 否   | 查询第几页，默认1            |
| page_size | int    | 否   | 每页查询数量，默认10          |
|   field   | string | 否   | 查询条件，可为空             |
|   order   | string | 否   | 查询条件的降序/升序排列 desc/asc |
|    key    | string | 否   | 查询关键词                |

##### Response

|        参数         | 类型            | 必填  | 说明                      |
|:-----------------:|---------------|-----|-------------------------|
|      page_no      | int           | 是   | 第几页                     |
|     page_size     | int           | 是   | 每页数量                    |
|       total       | int           | 是   | 总条数                     |
|       data        | object        | 是   | 基因数据                    |
|    data.allele    | array(object) | 否   | 基因，一个品系可能对应多个基因，品系名可能为空 |
|     allele.id     | int           | 否   | 基因ID                    |
|    allele.name    | string        | 否   | 基因名                     |
|  allele.annotate  | array(string) | 否   | 基因注解，一个基因可能对应多个注解       |
|   allele.genome   | string        | 否   | 基因修饰情况                  |
|   allele.serial   | string        | 否   | 对应第几条染色体，一般为罗马数字        |
|       extra       | array(object) | 否   | 基因的额外信息，一个基因可能对应多个额外信息  |
|  extra.extra_key  | string        | 是   | 基因的额外信息key              |
| extra.extra_value | string        | 是   | 基因的额外信息value            |



#### 2：新增基因

##### Request

##### Post http://127.0.0.1:10080/biology/allele_add


|    参数     | 类型     | 必填 | 说明                   |
|:---------:|--------| ---- |----------------------|
|    name    | string        | 否   | 基因名                     |
|  annotate  | array(string) | 否   | 基因注解，一个基因可能对应多个注解       |
|   genome   | string        | 否   | 基因修饰情况                  |
|   serial   | string        | 否   | 对应第几条染色体，一般为罗马数字        |
|       extra       | array(object) | 否   | 基因的额外信息，一个基因可能对应多个额外信息  |
|  extra.extra_key  | string        | 是   | 基因的额外信息key              |
| extra.extra_value | string        | 是   | 基因的额外信息value            |

##### Response

|   参数    | 类型     | 必填  | 说明           |
|:-------:|--------|-----|--------------|
|  code   | int    | 是   | 200为成功，其他为失败 |
| message | string | 是   | 消息内容         |

#### 3：修改基因

##### Request

##### Post http://127.0.0.1:10080/biology/allele_update

|        参数         | 类型         | 必填  | 说明     |
|:-----------------:|------------|-----|--------|
|        id         | int | 是   | 基因ID |
|       name        | string        | 否   | 基因名                     |
|     annotate      | array(string) | 否   | 基因注解，一个基因可能对应多个注解       |
|      genome       | string        | 否   | 基因修饰情况                  |
|      serial       | string        | 否   | 对应第几条染色体，一般为罗马数字        |
|       extra       | array(object) | 否   | 基因的额外信息，一个基因可能对应多个额外信息  |
|  extra.extra_key  | string        | 是   | 基因的额外信息key              |
| extra.extra_value | string        | 是   | 基因的额外信息value            |

##### Response

|   参数    | 类型     | 必填  | 说明           |
|:-------:|--------|-----|--------------|
|  code   | int    | 是   | 200为成功，其他为失败 |
| message | string | 是   | 消息内容         |

#### 4：删除基因

##### Request

##### Post http://127.0.0.1:10080/biology/allele_delete

| 参数  | 类型         | 必填  | 说明     |
|:---:|------------|-----|--------|
| id  | int | 是   | 基因ID |

##### Response

|   参数    | 类型     | 必填  | 说明           |
|:-------:|--------|-----|--------------|
|  code   | int    | 是   | 200为成功，其他为失败 |
| message | string | 是   | 消息内容         |

## 鸣谢

感谢 [JetBrains](https://www.jetbrains.com/?from=biology-server) 提供的 free JetBrains Open Source license

![JetBrains-logo](https://account.jetbrains.com/static/images/jetbrains-logo-minimal.svg)