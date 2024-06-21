这是个线虫品系存储数据库
需要存储线虫的品系信息
每个线虫有一个唯一不重复（但是可能不存在的）品系名（strain name），0或多个基因名（allele name），多条可能不同时间加上的注解，以及某些自定义的额外消息
基因名对应基因上的一个突变（用文本记录）或者其他的基因名或者两者皆有，初次之外还有注解以及额外信息

### 一：品系名

#### 1:新增品系名

##### Request

|             参数              | 类型          | 必填 | 说明                                  |
|:---------------------------:| ------------- | ---- |-------------------------------------|
|         strain_name         | string        | 否   | 品系名，可能为空                            |
|           number            | string        | 否   | 序列号                                 |
|         short_name          | array(string) | 否   | 简称，一个品系可能对应多个简称                     |
|       strain_annotate       | array(string) | 否   | 品系名注解， 一个品系可能对应多个注解                 |
|        strain_extra         | array(string) | 否   | 品系名额外信息，一个品系可能对应多个额外信息              |
|           allele            | array(object) | 否   | 基因                                  |
|     allele.allele_name      | string)       | 否   | 基因名，一个品系名可能对应多个基因名，品系名可能为空          |
| allele.allele_name_annotate | array(string) | 否   | 基因名注解，一个基因可能对应多个注解                  |
|  allele.allele_name_extra   | array(string) | 否   | 基因名额外信息，一个基因名可能对应多个额外信息             |
|     allele.genome_name      | string        | 否   | 基因修饰情况                              |
|        allele.serial        | array(string) | 否   | 对应第几条染色体，一般为罗马数字,一个基因修饰情况可能有多个染色体信息 |


##### Response

|   参数    | 类型     | 必填  | 说明           |
|:-------:|--------|-----|--------------|
|  code   | int    | 是   | 200为成功，其他为失败 |
| message | string | 是   | 消息内容         |

#### 2：展示品系名列表

##### Request

|    参数     | 类型     | 必填 | 说明                   |
|:---------:|--------| ---- |----------------------|
|  page_no  | int    | 否   | 查询第几页，默认1            |
| page_size | int    | 否   | 每页查询数量，默认10          |
|   field   | string | 否   | 查询条件，可为空             |
|   order   | string | 否   | 查询条件的降序/升序排列 desc/asc |
|    key    | string | 否   | 查询关键词                |

##### Response

|                   参数                    | 类型            | 必填  | 说明                                  |
|:---------------------------------------:|---------------|-----|-------------------------------------|
|                 page_no                 | int           | 是   | 第几页                                 |
|                page_size                | int           | 是   | 每页数量                                |
|                  total                  | int           | 是   | 总条数                                 |
|               strain_list               | array(object) | 是   | 品系列表                                |
|          strain_list.strain_id          | int           | 是   | 品系ID                                |
|         strain_list.strain_name         | string        | 否   | 品系名，可能为空                            |
|           strain_list.number            | string        | 是   | 序列号，一般#开头                           |
|         strain_list.short_name          | array(string) | 否   | 简称，一个品系可能对应                         |
|       strain_list.strain_annotate       | array(string) | 否   | 品系名注解， 一个品系可能对应多个注解                 |
|        strain_list.strain_extra         | array(string) | 否   | 品系名额外信息，一个品系可能对应多个额外信息              |
|           strain_list.allele            | array(object) | 否   | 基因                                  |
|    strain_list.allele.allele_name_id    | int           | 否   | 基因名ID                               |
|     strain_list.allele.allele_name      | string        | 否   | 基因名，一个品系名可能对应多个基因名，品系名可能为空          |
| strain_list.allele.allele_name_annotate | array(string) | 否   | 基因名注解，一个基因可能对应多个注解                  |
|  strain_list.allele.allele_name_extra   | array(string) | 否   | 基因名额外信息，一个基因名可能对应多个额外信息             |
|      strain_list.allele.genome_id       | string        | 否   | 基因修饰情况ID                            |
|     strain_list.allele.genome_name      | string        | 否   | 基因修饰情况                              |
|        strain_list.allele.serial        | array(string) | 否   | 对应第几条染色体，一般为罗马数字,一个基因修饰情况可能有多个染色体信息 |

#### 3：修改品系数据

##### Request

|             参数              | 类型          | 必填  | 说明                                  |
|:---------------------------:| ------------- |-----|-------------------------------------|
|          strain_id          | int        | 是   | 品系ID                           |
|         strain_name         | string        | 否   | 品系名，可能为空                            |
|         short_name          | array(string) | 否   | 简称，一个品系可能对应多个简称                     |
|       strain_annotate       | array(string) | 否   | 品系名注解， 一个品系可能对应多个注解                 |
|        strain_extra         | array(string) | 否   | 品系名额外信息，一个品系可能对应多个额外信息              |
|           allele            | array(object) | 否   | 基因                                  |
|     allele.allele_name      | string)       | 否   | 基因名，一个品系名可能对应多个基因名，品系名可能为空          |
| allele.allele_name_annotate | array(string) | 否   | 基因名注解，一个基因可能对应多个注解                  |
|  allele.allele_name_extra   | array(string) | 否   | 基因名额外信息，一个基因名可能对应多个额外信息             |
|     allele.genome_name      | string        | 否   | 基因修饰情况                              |
|        allele.serial        | array(string) | 否   | 对应第几条染色体，一般为罗马数字,一个基因修饰情况可能有多个染色体信息 |


##### Response

|   参数    | 类型     | 必填  | 说明           |
|:-------:|--------|-----|--------------|
|  code   | int    | 是   | 200为成功，其他为失败 |
| message | string | 是   | 消息内容         |

#### 4：删除品系数据

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

### 二：用户

#### 1：注册

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
|    data    | object | 是   | 数据         |
| data.token | string | 是   | token      |


#### 3：修改密码

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