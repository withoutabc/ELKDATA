# 基于ELK框架的静态/动态数据可视化系统

本项目所展示的可视化图表均已**额外**保存于`ELKDATA/docs/dashboard`目录下。

## ELK介绍

### 组成

ELK是`Elasticsearch`、`Logstash`、`Kibana`三大开源框架首字母大写简称。

#### Elasticsearch——搜索引擎

Elasticsearch是Elastic Stack核心的分布式搜索和分析引擎,是一个基于Lucene、分布式、通过Restful方式进行交互的近实时搜索平台框架。Elasticsearch为所有类型的数据提供近乎实时的搜索和分析。无论您是结构化文本还是非结构化文本，数字数据或地理空间数据，Elasticsearch都能以支持快速搜索的方式有效地对其进行存储和索引。

#### Logstash——数据输送

Logstash是免费且开放的服务器端数据处理管道，能够从多个来源采集数据，转换数据，然后将数据发送到指定的“存储库”中。Logstash能够动态地采集、转换和传输数据，不受格式或复杂度的影响。利用Grok从非结构化数据中派生出结构，从IP地址解码出地理坐标，匿名化或排除敏感字段，并简化整体处理过程。

#### Kibana——可视化

Kibana是一个针对Elasticsearch的开源分析及可视化平台，用来搜索、查看交互存储在Elasticsearch索引中的数据。使用Kibana，可以通过各种图表进行高级数据分析及展示。并且可以为 Logstash 和 ElasticSearch 提供的日志分析友好的 Web 界面，可以汇总、分析和搜索重要数据日志。还可以让海量数据更容易理解。它操作简单，基于浏览器的用户界面可以快速创建仪表板（dashboard）实时显示Elasticsearch查询动态。

## ELK用于数据可视化的优势

1. **统一的数据源**：ELK整合了Elasticsearch、Logstash和Kibana，可以通过Logstash将各种数据源的数据进行收集和处理，然后通过Elasticsearch进行索引和存储。这意味着可以从不同的数据源中获取数据，并将其统一存储在一个集中的位置，方便进行统一的数据可视化。
2. **多样化的可视化选项**：Kibana作为ELK的可视化工具，提供了丰富多样的图表和仪表盘选项。你可以通过Kibana创建各种类型的图表（如柱状图、折线图、饼图等）、地图、数据表格等，以及自定义仪表盘，来展示和分析数据。
3. **实时数据展示**：ELK具备实时处理和分析日志数据的能力。你可以实时查看最新的数据，并将其实时展示在可视化界面上。
4. **大数据支持**：ELK基于分布式架构，可以处理大规模的数据集。Elasticsearch的分片和副本机制保证了系统的高性能和高可用性，在处理大量数据时仍能保持较高的响应速度。

## 数据可视化

### 静态数据可视化

#### 数据源

- 可靠性：从 https://archive.ics.uci.edu 网站中下载的`csv`格式数据，名称为`Room Occupancy Estimation`。
- 复杂性：共有**1w**行数据。![](./docs/images/datasource1.png)
- 多样性：每行数据包含18个字段。![](./docs/images/data2.png)

（数据源文件在项目`ELKDATA/data/statistic`目录下）

#### 数据介绍

在一个6m x 4.6m的房间内部署了占用估算的实验测试平台。该设置**由7个传感器节点和1个边缘节点组成**，采用星形配置，传感器节点每30秒使用无线收发器向边缘传输数据。在收集数据集时没有使用HVAC系统。

实验中使用了五种不同类型的非侵入式传感器:**温度、光、声、CO2和数字被动红外(PIR)**。二氧化碳、声音和PIR传感器需要人工校准。对于二氧化碳传感器，在首次使用之前，通过将其保持在清洁环境中超过20分钟，然后将校准引脚(HD引脚)拉低超过7秒，手动进行零点校准。声音传感器本质上是一个带有可变增益模拟放大器的麦克风。因此，该传感器的输出是由微控制器的ADC以伏特为单位读取的模拟量。调整与放大器增益相连的电位器以确保最高灵敏度。PIR传感器有两个微调器:一个用于调整灵敏度，另一个用于调整检测运动后输出保持高电平的时间。这两个都被调整到最高值。**传感器节点S1-S4由温度、光和声音传感器组成，S5有一个二氧化碳传感器，S6和S7有一个PIR传感器**，每个传感器以一定角度部署在天花板架上，以最大化传感器的运动检测视野。

数据**以一种受控的方式收集了4天，房间的占用率在0到3人之间变化**。房间里入住人数的真实情况是手工记录的。

#### 部署

这里用`docker compose`实现`Elasticsearch`和`Kibana`的快速部署。

```yml
# 提前创建容器网络
# docker network create elk

version: "2"
services:
  elasticsearch:
    image: elasticsearch:7.13.1
    container_name: elasticsearch
    volumes:
      - "./elasticsearch_data:/bitnami/elasticsearch"
    ports:
      - "9200:9200"
      - "9300:9300"
    environment:
      - TZ=Asia/Shanghai
      - discovery.type=single-node
      - "ES_JAVA_OPTS=-Xms512m -Xmx512m"
    privileged: true
    ulimits:
      nofile:
        soft: 65536
        hard: 65536
    networks:
      - elk
  kibana:
    image: kibana:7.13.1
    container_name: kibana
    environment:
      - elasticsearch.hosts=http://elasticsearch:9200
      - TZ=Asia/Shanghai
    restart: always
    ports:
      - "5601:5601"
    volumes:
      - "./kibana_data:/bitnami/kibana"
    depends_on:
      - elasticsearch
    networks:
      - elk
networks:
  elk:
    external: true
```

#### 数据导入

访问5601端口，进入web界面，导入数据

![](./docs/images/dis1.png)

`Discover`界面如图所示。这也是**基于时间戳的柱形图**。

![](./docs/images/dis2.png)

#### 数据可视化

当然，`Kibana`所提供的`Dashboard`才是数据可视化的精髓。接下来我将从不同维度绘制图表并做说明。

### 动态数据可视化

#### Web应用的后台数据可视化

##### 数据源

由于本赛道主题与Web应用关系不大，我只用`golang`模拟了几个接口，并简单地编写了前端界面。这里只做简单介绍：

|   路径   |             作用              |               备注               |
| :------: | :---------------------------: | :------------------------------: |
| `/visit` |          进入web界面          |    请求此接口后自动请求`/ip`     |
|  `/ip`   |      获取登录主机的地址       | 有时请求信息速度较慢，请耐心等待 |
| `/slow`  | 模拟耗时较长的业务，睡500毫秒 |  点击`/visit`界面的按钮即可访问  |

我已用Docker将代码打包成镜像部署在服务器上，评委若感兴趣可以访问：http://49.7.114.49:5888/visit 

您的访问**将会以日志形式记录，并被ELK框架读取分析**。

换句话说，**这里的数据来自对接口的访问**。

##### 数据介绍

一共有三类日志格式：

1. `level`+`msg`+`time`：记录请求过程中产生的错误信息（msg格式不固定，可能是任何信息）。

   ```log
   {"level":"error","msg":"请求失败:Get \"http://ip-api.com/json/125.86.165.54?fields=61439\u0026lang=zh-CN\": read tcp 192.168.128.2:40052-\u003e208.95.112.1:80: read: connection timed out","time":"2023-09-12T18:41:40+08:00"}
   ```

2. `level`+`msg`+`time`，`msg`含有`country`+`region`+`city`+`latitude`+`longitude`：记录客户端地址。

   ```log
   {"level":"info","msg":"country:中国,region:河南,city:郑州市,latitude:34.747200,longitude:113.625000","time":"2023-09-15T22:13:27+08:00"}
   ```

3. `timestamp`+`status_code`+`client_ip`+`latency`+`method`+`path`：记录请求信息和响应结果。

   ```log
   timestamp:2023-09-12 18:24:06,status_code:200,client_ip:125.86.165.54,latency:5.672637623s,method:GET,path:/ip
   timestamp:2023-09-12 18:24:14,status_code:200,client_ip:125.86.165.54,latency:351.117µs,method:GET,path:/visit
   ```

- 以上字段信息的获取都借助了一些框架或API接口，这里不赘述。

##### 数据收集

基于之前的静态数据可视化（`elasticsearch`+`kibana`），现在额外部署`logstash`，用于不断将日志信息收集处理并发送至`elasticsearch`中。

```yml
# docker-compose.yml中添加  
  logstash:
    image: docker.elastic.co/logstash/logstash:7.13.1
    container_name: logstash
    volumes:
      - "./logstash.conf:/usr/share/logstash/pipeline/logstash.conf"
      - "./elk_data/log/:/home/withoutabc/elk/elk_data/log/"
      - "./jdbc_driver/mysql-connector-java-8.0.26.jar:/jdbc_driver/mysql-connector-java-8.0.26.jar"
    environment:
      - "XPACK_MONITORING_ENABLED=false"
      - TZ=Asia/Shanghai
    depends_on:
      - elasticsearch
    networks:
      - elk
```

其中，`"./logstash.conf:/usr/share/logstash/pipeline/logstash.conf"`用于将宿主机上的`logstash`配置文件挂载到容器中。

```ruby
# logstash.conf
input {
  file {
    path => "/home/withoutabc/elk/elk_data/log/*.log"
    start_position => "beginning"
    sincedb_path => "/dev/null"
    tags => ["file"]
  }
}

filter {
  if "client_ip" in [message] {
    grok {
      match => { "message" => "timestamp:%{TIMESTAMP_ISO8601:timestamp},status_code:%{NUMBER:status_code},client_ip:%{IP:client_ip},latency:%{NUMBER:latency_value}%{DATA:latency_unit},method:%{WORD:method},path:%{URIPATH:url_path}" }
    }
    date {
        match => [ "timestamp", "yyyy-MM-dd HH:mm:ss" ]
        target => "@timestamp"
    }
    mutate {
      add_field => {
        "index_name" => "visit"
      }
    }
  } else if "country" in [message] {
    grok {
    match => { "message" => "{\"level\":\"%{WORD:level}\",\"msg\":\"country:%{DATA:country},region:%{DATA:region},city:%{DATA:city},latitude:%{NUMBER:latitude},longitude:%{NUMBER:longitude}\",\"time\":\"%{TIMESTAMP_ISO8601:timestamp}\"}" }

    }
    date {
        match => [ "timestamp", "yyyy-MM-dd'T'HH:mm:ssZ" ]
        target => "@timestamp"
    }
    mutate {
      add_field => {
        "index_name" => "ip"
      }
    }
    mutate {
    convert => {
      "latitude" => "float"
      "longitude" => "float"
    }
  }

    mutate {
    add_field => {
      "location" => "%{[latitude]},%{[longitude]}"
    }
  }
  } else if "level" in [message] and "country" not in [message] {
    grok {
        match => { "message" => "{\"level\":\"%{WORD:level}\",\"msg\":\"%{GREEDYDATA:msg}\",\"time\":\"%{TIMESTAMP_ISO8601:timestamp}\"}" }

        }
        date {
            match => [ "timestamp", "yyyy-MM-dd'T'HH:mm:ssZ" ]
            target => "@timestamp"
        }
        mutate {
          add_field => {
            "index_name" => "log"
          }
        }
  } else {
    drop {}
  }
}

output {
    if "file" in [tags] {
      elasticsearch {
          hosts => ["elasticsearch:9200"]
          index => "%{index_name}"
        }
    }
    stdout {
            codec => rubydebug
    }
}
```

- 配置文件使用`ruby`语言进行逻辑处理。
- 通过对不同格式的日志信息解析，将数据导入`elasticsearch`。
- 用3个索引分别存储了以上三种格式的日志：`log`，`ip`，`visit`。

现在访问`Kibana`界面，并`Create index pattern`，可以在`Discover`中看到数据关于日期的分布。

1. `log`![](./docs/images/dis-log.png)
2. `ip`![](./docs/images/dis-ip.png)
3. `visit`![](./docs/images/dis-visit.png)

##### 数据可视化

用`Kibana`的`Dashboard`进行数据可视化，虽然无法短时间获取大量数据，但我会尽量描述清楚它们的作用。

1. **监测错误日志**
   - 横轴：时间戳，纵轴：日志记录（过滤掉**不是错误**的日志记录）
   - 作用：方便开发人员观测错误日志数的走向和趋势、及时排查问题。

![](./docs/images/log-error.png)

2. **错误日志数**

   - 和上一项搭配，统计过去24小时产生的错误日志。
   - 作用：直观清晰地反映是否有错误待排查。

   ![](./docs/images/log-metric.png)

3. **访问来源前10名**

   - 作用：了解用户受众、根据用户偏好改进内容、区域的市场扩展

   ![](./docs/images/ip-top10.png)

#### 数据库可视化

目前企业开发Web应用所使用的数据库种类很多，如：`MySQL`、`Redis`、`MongoDB`、`PostgreSQL`等。基于`MySQL`开源免费、成熟稳定、易于使用、高性能和可扩展性强等特性，我们在此以`MySQL`为例进行数据可视化。