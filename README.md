# 基于ELK框架的静态/动态数据可视化系统

本项目所展示的可视化图表均已**额外**保存于`ELKDATA/docs`目录下。

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

- 可靠性：从https://archive.ics.uci.edu网站中下载的`csv`格式数据，名称为`Room Occupancy Estimation`。
- 复杂性：**1w+**行数据。![)RB(`4${SDZ((_B4RHU_ZS7.png](https://s2.loli.net/2023/09/08/3vKNBq7y5lPLcAo.png)
- 多样性：每行数据包含18个字段。![_6_H63_Q_N_FRW4D__0MBD3.png](https://s2.loli.net/2023/09/08/YKTJfIAeRH8zg92.png)

（数据源文件在项目`ELKDATA/data/statistic`目录下）

#### 数据介绍

在一个6m x 4.6m的房间内部署了占用估算的实验测试平台。该设置**由7个传感器节点和1个边缘节点组成**，采用星形配置，传感器节点每30秒使用无线收发器向边缘传输数据。在收集数据集时没有使用HVAC系统。

实验中使用了五种不同类型的非侵入式传感器:**温度、光、声、CO2和数字被动红外(PIR)**。二氧化碳、声音和PIR传感器需要人工校准。对于二氧化碳传感器，在首次使用之前，通过将其保持在清洁环境中超过20分钟，然后将校准引脚(HD引脚)拉低超过7秒，手动进行零点校准。声音传感器本质上是一个带有可变增益模拟放大器的麦克风。因此，该传感器的输出是由微控制器的ADC以伏特为单位读取的模拟量。调整与放大器增益相连的电位器以确保最高灵敏度。PIR传感器有两个微调器:一个用于调整灵敏度，另一个用于调整检测运动后输出保持高电平的时间。这两个都被调整到最高值。**传感器节点S1-S4由温度、光和声音传感器组成，S5有一个二氧化碳传感器，S6和S7有一个PIR传感器**，每个传感器以一定角度部署在天花板架上，以最大化传感器的运动检测视野。

**数据以一种受控的方式收集了4天，房间的占用率在0到3人之间变化。**房间里入住人数的真实情况是手工记录的。

#### 快速部署

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

![YJL_6XA8SYCR__QMI291B_0.png](https://s2.loli.net/2023/09/08/Og5x9BrvNRJLoI8.png)

`Discover`界面如图所示。这也是**基于时间戳的柱形图**。

![38X35I_WI_O9_57E9`K_T_M.png](https://s2.loli.net/2023/09/08/FPn7NyoK1p5zGgW.png)

#### 图表绘制与展示

当然，`Kibana`所提供的`Dashboard`才是数据可视化的精髓。接下来我将从不同维度绘制图表并做说明。

### 动态数据可视化

#### Web应用的后台数据可视化



#### 数据库可视化

目前企业开发Web应用所使用的数据库种类很多，如：`MySQL`、`Redis`、`MongoDB`、`PostgreSQL`等。基于`MySQL`开源免费、成熟稳定、易于使用、高性能和可扩展性强等特性，我们在此以`MySQL`为例进行数据可视化。