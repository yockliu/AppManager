#Api说明

>Created by: Yock.L
>
>Created at: 2014.9.5
>
>Updated at: 2014.9.9

---

##目录

* [基本规则](#anchor-基本规则)
* [App](#anchor-App)
* [Version](#anchor-Version)
* [Channel](#anchor-Channel)
* [AppBuild](#anchor-AppBuild)

---

##<a name="anchor-基本规则" id="anchor-基本规则">基本规则</a>

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
- 400 INVALID REQUEST - [POST/PUT/PATCH]：用户发出的请求有错误，服务器没有进行新建或修改数据的操作，该操作是幂等的。
- 404 NOT FOUND - [*]：用户发出的请求针对的是不存在的记录，服务器没有进行操作，该操作是幂等的。
- 500 INTERNAL SERVER ERROR - [*]：服务器发生错误，用户将无法判断发出的请求是否成功。
```
####json

- GET 返回GET的数据（如果是列表且列表没有数据，则返回"[]"）
- POST, PUT 返回创建或修改后的数据
- DELETE 不返回内容

---

##<a name="anchor-App" id="anchor-App">App</a>

####go struct
```
type App struct {
	Id        bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Name      string        `json:"name"      bson:"name"`
	Platforms []string      `json:"platforms" bson:"platforms"`
	Created   time.Time     `json:"created"   bson:"created"`
	Updated   time.Time     `json:"updated"   bson:"updated,omitempty"`
	Forbidden bool          `json:"forbidden" bson:"forbidden"`
	Validate  bool          `json:"validate"  bson:"validate"`
}
```

####json
```
{
	"id":			"$id",
	"name":			"$name",
	"platforms": 	["android", "ios"],
	"Validate":		true
}
```

####apis

* [Get App List](#anchor-get_app_list)
* [Get App](#anchor-get_app)
* [Post App (create)](#anchor-post_app)
* [Put App (set)](#anchor-put_app)
* [Delete App](#anchor-delete_app)


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
[{"id":"540843152a936f110737225c","name":"周末去哪儿","platforms":["android","ios"]}]
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
	body: {app json}

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
{"id":"540849ad2a936f110737225d","name":"周末去哪儿","platforms":["android"]}
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
	body: {app json}

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
body: {"name": "周末","platforms":["android"]}
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
	body: {app json}

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
body: {"name": "周末","platforms":["android"]}
```

###<a name="anchor-delete_app" id="anchor-delete_app">Delete App</a>

####功能
删除App

`App的删除实际没有真正删除记录，只是把validate置成了false`

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
```

##<a name="anchor-Version" id="anchor-Version">Version</a>

####go struct
```
type Version struct {
	Id       bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Code     string        `json:"code"      bson:"code"`
	Name     string        `json:"name"      bson:"name"`
	Platform string        `json:"platform"  bson:"platform"`	GitTag   string        `json:"git_tag"   bson:"git_tag,omitempty"`
	GitIndex string        `json:"git_index" bson:"git_index,omitempty"`
 	Created  time.Time     `json:"created"   bson:"created"`
	Updated  time.Time     `json:"updated"   bson:"updated,omitempty"`
}
```

####json
```
{
	"id":			"$id",
	"code":			"$code"
	"name":			"$name",
	"platforms": 	"android",
}
```

####apis

* [Get Version List](#anchor-get_version_list)
* [Get Version](#anchor-get_version)
* [Post Version (create)](#anchor-post_version)
* [Put Version (set)](#anchor-put_version)
* [Delete Version](#anchor-delete_version)


###<a name="anchor-get_version_list" id="anchor-get_version_list">Get Version List</a>

####功能
获得Version列表

####request

#####Method
**_Get_**

#####Path
**_/api/app/:appid/version_**


####response
	
	status: 200
	body:
		[{version json 1},{version json 2},...]
		[] // 为空的情况

####示例

请求

```
curl -X GET \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	http://localhost:3000/api/app/:appid/version
```

返回

```
status: 200
body:
[{"id":"540843152a936f110737225c","code":"1","name":"0.0.1","platforms":"android"}]
```

###<a name="anchor-get_version" id="anchor-get_version">Get Version</a>

####功能
获得Version详情

####request

#####Method
**_GET_**

#####Path
**_/api/app/:appid/version/:id_**


####response

	status: 200
	body: {version json}

####示例

请求

```
curl -X GET \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	http://localhost:3000/api/app/:appid/version/:id
```

返回


###<a name="anchor-post_version" id="anchor-post_version">Post Version</a>

####功能
创建Version

####request

#####Method
**_POST_**

#####Path
**_/api/app/:appid/version_**

#####Body

**_{version json}_**

####response

	status: 201
	body: {version json}

####示例

请求

```
curl -X POST \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	-d '{
			"code": "1"
			"name": "0.0.1",
			"platform": "android"
		}' \
	http://localhost:3000/api/app/:appid/version
```

返回

```
status: 201
body: {"id": "...","code": "1","name": "0.0.1","platforms":"android"}
```

###<a name="anchor-put_version" id="anchor-put_version">Put Version</a>

####功能
更新Version

####request

#####Method
**_PUT_**

#####Path
**_/api/app/:appid/version/:id_**

#####Body

**_{version json}_**

####response

	status: 201
	body: {version json}

####示例

请求

```
curl -X PUT \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	-d '{
			"code": "1"
			"name": "0.0.1",
			"platforms": "android"
		}' \
	http://localhost:3000/api/app/:appid/version/:id
```

返回

```
status: 201
body: {"code": "1","name": "0.0.1","platforms":"android"}
```

###<a name="anchor-delete_version" id="anchor-delete_version">Delete Version</a>

####功能
删除Version

####request

#####Method
**_DELETE_**

#####Path
**_/api/app/:appid/version/:id_**

####response

	status: 204

####示例

请求

```
curl -X DELETE \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	http://localhost:3000/api/app/:appid/version/:id
```

返回

```
status: 204
```

##<a name="anchor-Channel" id="anchor-Channel">Channel</a>

####go struct
```
type Channel struct {
	Id        bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Code      string        `json:"code"      bson:"code"`
	Name      string        `json:"name"      bson:"name"`
	Platform  string        `json:"platform"  bson:"platform"`
	Created   time.Time     `json:"created"   bson:"created"`
	Updated   time.Time     `json:"updated"   bson:"updated,omitempty"`
}
```

####json
```
{
	"id":			"$id",
	"code":			"$code",
	"name":			"$name",
	"platforms": 	"android"
}
```

####apis

* [Get Channel List](#anchor-get_channel_list)
* [Get Channel](#anchor-get_channel)
* [Post Channel (create)](#anchor-post_channel)
* [Put Channel (set)](#anchor-put_channel)
* [Delete Channel](#anchor-delete_channel)


###<a name="anchor-get_channel_list" id="anchor-get_channel_list">Get Channel List</a>

####功能
获得Channel列表

####request

#####Method
**_Get_**

#####Path
**_/api/app/:appid/channel_**


####response
	
	status: 200
	body:
		[{channel json 1},{channel json 2},...]
		[] // 为空的情况

####示例

请求

```
curl -X GET \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	http://localhost:3000/api/app/:appid/channel
```

返回

```
status: 200
body:
[{"id":"540843152a936f110737225c","code":"and-a0","name":"测试","platforms":"android"}]
```

###<a name="anchor-get_channel" id="anchor-get_channel">Get Channel</a>

####功能
获得Channel详情

####request

#####Method
**_GET_**

#####Path
**_/api/app/:appid/channel/:id_**


####response

	status: 200
	body: {channel json}

####示例

请求

```
curl -X GET \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	http://localhost:3000/api/app/:appid/channel/:id
```

返回

```
status: 200
body:
{"id":"540849ad2a936f110737225d","code":"and-a0","name":"测试","platforms":"android"}
```

###<a name="anchor-post_channel" id="anchor-post_channel">Post Channel</a>

####功能
创建Channel

####request

#####Method
**_POST_**

#####Path
**_/api/app/:appid/channel_**

#####Body

**_{channel json}_**

####response

	status: 201
	body: {channel json}

####示例

请求

```
curl -X POST \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	-d '{
			"code": "and-a0"
			"name": "测试",
			"platforms": "android"
		}' \
	http://localhost:3000/api/app/:appid/channel
```

返回

```
status: 201
body: {"id":"540849ad2a936f110737225d","code":"and-a0","name":"测试","platforms":"android"}
```

###<a name="anchor-put_channel" id="anchor-put_channel">Put Channel</a>

####功能
更新Channel

####request

#####Method
**_PUT_**

#####Path
**_/api/app/:appid/channel/:id_**

#####Body

**_{channel json}_**

####response

	status: 201
	body: {channel json}

####示例

请求

```
curl -X PUT \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	-d '{
			"code": "and-a0"
			"name": "测试",
			"platform": "android"
		}' \
	http://localhost:3000/api/app/:appid/channel/:id
```

返回

```
status: 201
body: {"id":"540849ad2a936f110737225d","code":"and-a0","name":"测试","platforms":"android"}
```

###<a name="anchor-delete_channel" id="anchor-delete_channel">Delete Channel</a>

####功能
删除Channel

####request

#####Method
**_DELETE_**

#####Path
**_/api/app/:appid/channel/:id_**

####response

	status: 204

####示例

请求

```
curl -X DELETE \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	http://localhost:3000/api/app/:appid/channel/:id
```

返回

```
status: 204
```

##<a name="anchor-AppBuild" id="anchor-AppBuild">AppBuild</a>

####apis

* [build](#anchor-build)
* [build status](#anchor-build_status)
* [package download urls](#anchor-package_download_url)

###<a name="anchor-build" id="anchor-build">build</a>

####Request

Method

```
POST
```

Path

```
/api/build
```

Body

```
{
	"appid": "$appid",
	"versionid": "$versionid",
	"channels": [$channel_code_arrays]
}
```

####Response

```
status: 200
```

###<a name="anchor-build_status" id="anchor-build_status">build status</a>

####Request

Method

```
GET
```

Path

```
/api/build/status/:appid
```

####Response

```
status: 200
body: {"running":$isrunning} // isrunning = [true | false]
```

###<a name="anchor-package_download_url" id="anchor-package_download_url">package download urls</a>

待定
