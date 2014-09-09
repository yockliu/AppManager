#Api Request Example

---

##App api

###GET app list
```
curl -X GET \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	http://localhost:3000/api/app
```

###POST app detail
```
curl -X POST \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	-d '{
			"name": "周末去哪儿",
			"platforms": ["android", "ios"]
		}' \
	http://localhost:3000/api/app
```

###GET app detail
```
curl -X GET \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	http://localhost:3000/api/app/540ea6b0421e44d184000001
```

###PUT app detail
```
curl -X PUT \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	-d '{
			"platforms": ["ios"]
		}' \
	http://localhost:3000/api/app/540ea6b0421e44d184000001
```

###DELETE app detail
```
curl -X DELETE \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	http://localhost:3000/api/app/540ea6b0421e44d184000001
```

---

##Version api

###GET version list
```
curl -X GET \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	http://localhost:3000/api/app/540ea6b0421e44d184000001/version
```

###POST version detail
```
curl -X POST \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	-d '{
		"code": "2",
		"name": "0.0.1",
		"platform": "android"
	}' \
	http://localhost:3000/api/app/540ea615421e44d11e000001/version
```

###GET version detail
```
curl -X GET \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	http://localhost:3000/api/app/540ea615421e44d11e000001/version/540ebd3e421e44d696000002
```

###PUT version detail
```
curl -X PUT \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	-d '{
			"code": "2",
			"name": "0.1.1",
			"platform": "android"
		}' \
	http://localhost:3000/api/app/540ea615421e44d11e000001/version/540ebd3e421e44d696000002
```

###DELETE version detail
```
curl -X DELETE \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	http://localhost:3000/api/app/540ea615421e44d11e000001/version/540ebcbb2a936f1107372268
```

---

##Channel api

###GET channel list
```
curl -X GET \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	http://localhost:3000/api/app/540ea615421e44d11e000001/channel
```

###POST channel detail
```
curl -X POST \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	-d '{
		"code": "and-a0",
		"name": "测试"
	}' \
	http://localhost:3000/api/app/540ea615421e44d11e000001/channel
```

###GET channel detail
```
curl -X GET \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	http://localhost:3000/api/app/540ea615421e44d11e000001/channel/540ec18f421e44d794000002
```

###PUT channel detail
```
curl -X PUT \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	-d '{
			"code": "ios-a0",
			"name": "测试",
			"platform": "ios"
		}' \
	http://localhost:3000/api/app/540ea615421e44d11e000001/channel/540ec0c12a936f1107372269
```

###DELETE channel detail
```
curl -X DELETE \
	-H "Authorization: Basic YWRtaW46Z3Vlc3NtZQ==" \
	http://localhost:3000/api/app/540ea615421e44d11e000001/channel/540ec0f1421e44d794000001
```



















