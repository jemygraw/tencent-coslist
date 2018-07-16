# coslist

### 简介

coslist 是利用腾讯云存储API开发的空间文件列表获取工具，以便于根据文件的列表备份或者迁移数据。

### 下载

可以直接下载预编译好的二进制可执行文件使用。

|版本|支持平台|链接|
|------|-----|------|
|coslist v1.0|Linux, Mac, Windows|[http://devtools.qiniu.com/coslist-v1.0.zip](http://devtools.qiniu.com/coslist-v1.0.zip)|


### 使用
该工具是一个命令行工具，需要指定相关的参数来运行。

```
Usage of coslist:
  -appId string
    	app id
  -secretId string
    	secret id
  -secretKey string
    	secret key
  -bucket string
    	bucket name
  -result string
    	result file
  -dir string
    	dir (default "/")
  -prefix string
    	prefix
```

|参数|描述|是否可选|
|--------|--------|-------|
|appId|腾讯云存储的appId，从腾讯云存储后台获取|必需|
|secretId|腾讯云存储的secretId，从腾讯云存储后台获取|必需|
|secretKey|腾讯云存储的secretKey，从腾讯云存储后台获取|必需|
|bucket|腾讯云存储的空间，从腾讯云存储后台获取|必需|
|result|获取的列表保存在本地的文件名|必需|
|dir|需要获取文件列表的目录，可以不填，默认为`/`，表示获取全部目录文件列表|可选|
|prefix|需要获取的文件列表的名称前缀，可以使用该参数过滤指定前缀文件名称列表|可选|

### 示例

```
1. 获取空间中的所有文件列表

$ coslist -appId '10032293' -secretId 'AKIDNgGxB0OLDBVN8ZuqlyOPndWTsa9vIvTU' -secretKey 'QO60fjSYYKCF4RVNQ7gq48EDwpFx6c3h' -bucket 'smile' -result 'list.txt'
 
2. 获取空间中的指定目录下的文件列表

$ coslist -appId '10032293' -secretId 'AKIDNgGxB0OLDBVN8ZuqlyOPndWTsa9vIvTU' -secretKey 'QO60fjSYYKCF4RVNQ7gq48EDwpFx6c3h' -bucket 'smile' -dir 'hello/world' -result 'list.txt'

3. 获取空间中指定目录下，特定文件前缀的文件列表
$ coslist -appId '10032293' -secretId 'AKIDNgGxB0OLDBVN8ZuqlyOPndWTsa9vIvTU' -secretKey 'QO60fjSYYKCF4RVNQ7gq48EDwpFx6c3h' -bucket 'smile' -dir 'hello/world' -prefix 'jemy' -result 'list.txt' 

```

备注：如果是在Windows下面使用，请从终端运行工具，另外命令行参数两边都不用加单引号。

获取的文件列表结果如下格式：

```
http://smile-10032292.cos.myqcloud.com/hello/pig.jpg  http://smile-10032292.file.myqcloud.com/hello/pig.jpg 34230
http://smile-10032292.cos.myqcloud.com/hello/pili.png http://smile-10032292.file.myqcloud.com/hello/pili.png  14437
http://smile-10032292.cos.myqcloud.com/hello/service-req.jpg  http://smile-10032292.file.myqcloud.com/hello/service-req.jpg 126846
http://smile-10032292.cos.myqcloud.com/pool_bear.png  http://smile-10032292.file.myqcloud.com/pool_bear.png 174524
```

第一列为源站访问地址，第二列为CDN访问地址，第三列为文件大小的字节数。


