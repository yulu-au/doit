# 模板
```
es没有规定一定要在插入文档的时候声明类型,很多时候都是插入文档,es自动推断类型.
但有些情况是不可以使用es自动推断,比如插入"ip"类型的数据,es会想当然的认为是text,于是我们就要纠正这种情况.
如何纠正,使用模板

PUT _index_template/template1
{
  "index_patterns": [
    "my*"//这个模板生效的索引是哪些
  ],
  "template": {
    "settings": {
      "number_of_shards": 1
    },
    "mappings": {
      "_source": {
        "enabled": true
      },
      "properties": {
        "name": {
          "type": "keyword"//设置这个字符串只要keyword类型
        },
        "sip": {
          "type": "ip"//这里就是我们规定的类型
        }
      }
    }
  }
}
```
# 查看索引的元信息
```
get /{索引名}/_mapping
get /{索引名}/_setting
```
# 查询
```
post /{索引名}/_search
{request_body}

query->term/terms/wildcard

```

# 插入文档
```
post /{索引名}/_doc
{request body}
```

# keyword vs text
```
查看索引的mapping就知道,每一个字符串都声明了两种类型 keyword+text

keyword 不会分词 支持聚合
text 分词 不支持聚合
```

# 特殊类型的查询
```
ip类型:term "ip/掩码位数"
```

# 聚合
```
terms聚合: 去重
基数聚合: cardinality 得到聚合数据的总条数
聚合多个并列,也可以嵌套
```
# 聚合后的分页及排序
```
terms聚合分页使用嵌套聚合bucket_sort
  "aggs": {
    "cc": {
      "cardinality": {
        "field": "sip"//得到sip这个字段去重之后的数量
      }
    },
    "aa":{
      "terms": {
        "field": "sip",
        "size": 10,//这个是terms聚合出的文档数量
        "order": {
          "_count": "asc"//聚合后的字段根据doc_count正序,还可以根据_key,或者其他字段排序
        }
        //exclude字段可以排除一些我们不感兴趣的数据,比如空的字符串,或者也可以写正则表达式
      }
      , "aggs": {
        "bb": {
          "bucket_sort": {
            "from": 7,//从第8个文档处,下标从零开始的
            "size": 10
          }
        }
      }
      }
  }

```

# 调优
```
同一类型的索引按照日期做切分,这样搜索的时候先过滤一遍索引名,找到指定的索引
使用索引别名,它是一个快捷方式,相当于做了一层抽象,这样要更改索引的字段属性就比较方便
自定义路由+冷热节点
倒排索引机制，能 keyword 类型尽量 keyword
forcemerge以减少集群的segment个数和清理已删除或更新的⽂档
禁止使用 swap 空间

堆内存设置为：Min（节点内存/2, 32GB)
线程池+队列大小根据业务需要做调整
```
# 路由机制
```
当插入某个文档的时候,要决定放到哪个分片里(ES有很多分片),这个算法是放到哪个分片=hash(文档ID)%主分片数量,还可以自定义某个文档放到哪个分片里

POST routetest/_doc/c?routing=key1//routing就是一个参数,表明使用这个参数计算分到哪个分片
{
  "name":"name key1",
  "age":18
}
自定义某些文档分到哪些分片有什么好处,不自定义的话文档会被均匀的分到所有分片里,查询的时候所有分片都要查询.自定义的话,查询的时候可以只查一个分片

对于超大数据量的搜索，routing再配合hot&warm的架构，是非常有用的一种解决方案
```

# 文档写入ES的过程
```
写入请求到某个节点,根据文档ID计算主分片节点(路由算法),发到对应节点写入主分片,然后写入副本分片,都写入成功之后给协调节点响应,随后协调节点响应这个写入请求
```

# 搜索文档的过程
```
query阶段:搜索请求发给了协调节点,协调节点广播到主分片或者副本分片,这些分片所在的节点计算命中的局部有序文档集合(不是数据的全集,只是必要的ID和搜索条件),发给协调节点
fetch阶段:协调节点综合全部的数据搞出一个全局有序的集合,根据这个集合从不同的分片拉去数据,响应搜索请求
```

# ES删除/更新文档
```
es中文档不可变,所以更新就是标记删除,并且写了一份新的文档代替.删除肯定也是标记删除了.
搜索的时候会检测到删除标记,不记录这些删除的文档
什么时候会真正删除,段合并的时候
```

# elasticsearch 是如何实现 master 选举
```
候选主节点才能成为master(master:true),最小节点数(min_master_nodes)是为了防止脑裂
选举流程是 
1候选主节点的数量达标 
2每个节点对自己已知的节点ID排序,选择小的那个投票
3得票数最多的并且它自己也不要脸的自荐就成为master 

ps:es节点分类 master节点 data节点 client节点(接受请求)
```