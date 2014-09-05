#Api说明

Created by: Yock.L

Created at: 2014.9.5

---

##基本规则

###协议

目前所有接口暂时使用http

`使用的框架不支持https，等我go技能更强大，会改为https的`

###Api URL与普通URL区分

1. 使用域名区分
2. 使用path前缀区分

示例：

```
	普通路径: http://example.com/path
	域名区分: http://api.example.com/path
	前缀区分: http://example.com/api/path
```

第一种方法应该比较好，不会用，所以现在用第二种方法

###Version

待定

###Auth

目前接口使用都需要Auth信息

在HEAD中增加以下内容，可参考`test/request_test.txt`

```
	Authorization: Basic YWRtaW46Z3Vlc3NtZQ==
```
###返回数据

#####状态码

```
- 200 OK - [GET]：服务器成功返回用户请求的数据，该操作是幂等的（Idempotent）。
- 201 CREATED - [POST/PUT/PATCH]：用户新建或修改数据成功。
- 204 NO CONTENT - [DELETE]：用户删除数据成功。
- 400 INVALID REQUEST - [POST/PUT/PATCH]：用户发出的请求有错误，服务器没有进行新建或修改数据的操作，该操作是幂等的。。
- 404 NOT FOUND - [*]：用户发出的请求针对的是不存在的记录，服务器没有进行操作，该操作是幂等的。
- 500 INTERNAL SERVER ERROR - [*]：服务器发生错误，用户将无法判断发出的请求是否成功。
```
####json

如果有内容返回，则使用json格式返回。

---

##App

####go struct
```
type App struct {
    Id        bson.ObjectId `json:"id"        bson:"_id,omitempty"`
    Name      string        `json:"name"      bson:"name"`
    Platforms []string      `json:"platforms" bson:"platforms"`
}
```

####json
```
{
	"id":			"$id",
	"name":			"$name",
	"platforms": 	["$p1", "$p2"],
}
```

####apis

* [Get App List](anchor-get_app_list)
* [Get App](anchor-get_app)
* [Post App (create)](anchor-post_app)
* [Put App (set)](anchor-put_app)
* [Delete App](anchor-delete_app)


###<a name="anchor-get_app_list" id="anchor-get_app_list">Get App List</a>

####功能
获得App列表

####request

#####Method
**_Get_**

#####Path
**_/api/app_**


####response
	
	status: 200
	body:
		[{app json 1},{app json 2},...]
		[] // 为空的情况

####示例

请求

```
curl -X GET \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	http://localhost:3000/api/app
```

返回

```
status: 200
body:
[{\"id\":\"540843152a936f110737225c\",\"name\":\"abc\",\"platforms\":[\"android\"]},{\"id\":\"540849ad2a936f110737225d\",\"name\":\"周末去哪儿\",\"platforms\":[\"android\"]},{\"id\":\"540924cd2a936f110737225e\",\"name\":\"周末\",\"platforms\":[\"android\"]},{\"id\":\"540925a82a936f110737225f\",\"name\":\"周末1\",\"platforms\":[\"android\"]},{\"id\":\"540925ec2a936f1107372260\",\"name\":\"周末2\",\"platforms\":[\"android\"]},{\"id\":\"540926312a936f1107372261\",\"name\":\"周末3\",\"platforms\":[\"android\"]}]
```

###<a name="anchor-get_app" id="anchor-get_app">Get App</a>

####功能
获得App详情

####request

#####Method
**_GET_**

#####Path
**_/api/app/:id_**


####response

	status: 200
	body:
		 {app json}

####示例

请求

```
curl -X GET \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	http://localhost:3000/api/app/:id
```

返回

```
status: 200
body:
{\"id\":\"540849ad2a936f110737225d\",\"name\":\"周末去哪儿\",\"platforms\":[\"android\"]}
```

###<a name="anchor-post_app" id="anchor-post_app">Post App</a>

####功能
创建App

####request

#####Method
**_POST_**

#####Path
**_/api/app_**

#####Body

**_{app json}_**

####response

	status: 201

####示例

请求

```
curl -X POST \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	-d '{
			"name": "周末",
			"platforms": ["android"]
		}' \
	http://localhost:3000/api/app
```

返回

```
status: 201
body:
	null
```

###<a name="anchor-put_app" id="anchor-put_app">Put App</a>

####功能
更新App

####request

#####Method
**_PUT_**

#####Path
**_/api/app/:id_**

#####Body

**_{app json}_**

####response

	status: 201

####示例

请求

```
curl -X PUT \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	-d '{
			"name": "周末",
			"platforms": ["android"]
		}' \
	http://localhost:3000/api/app/:id
```

返回

```
status: 201
body:
	null
```

###<a name="anchor-delete_app" id="anchor-delete_app">Delete App</a>

####功能
删除App

####request

#####Method
**_DELETE_**

#####Path
**_/api/app/:id_**

####response

	status: 204

####示例

请求

```
curl -X DELETE \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	http://localhost:3000/api/app/:id
```

返回

```
status: 204
body:
	null
```

##Version
待添加

##Channel
待添加