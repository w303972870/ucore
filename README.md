# 本说明使用Editor.md编辑

![](https://pandao.github.io/editor.md/images/logos/editormd-logo-180x180.png)

![](https://img.shields.io/github/stars/pandao/editor.md.svg) ![](https://img.shields.io/github/forks/pandao/editor.md.svg) ![](https://img.shields.io/github/tag/pandao/editor.md.svg) ![](https://img.shields.io/github/release/pandao/editor.md.svg) ![](https://img.shields.io/github/issues/pandao/editor.md.svg) ![](https://img.shields.io/bower/v/editor.md.svg)

# 项目背景

- 需要接入ldap的项目较多，但是每一个系统不需要都写一套对接ldap的代码，甚至有的时候不同的开发语言实现ldap对接会比较复杂，所以才写了这套ldap统一用户对接中心

# 配置文件
> 源码的config目录中有一个示例ini

```
[ucore]
daemon = False #daemon方式后台运行，默认true
port = 80      #监听端口，默认80
host = 192.168.0.211 #监听ip，默认为空(监听的时候“:port”)

[log]
dir = /data/logs/ #日志目录
logname = access-ucore.log #日志名称，默认access-ucore.log

[ldap]
host = LDAP的地址
port = 389 # 默认389
uid = cn # 默认cn
bind_dn = 绑定账号
password = 绑定密码

[platform]
platform = cmdp,workcenter # 支持验证的平台标识，所有的验证都包含这个
platform_key = staffAccess # 包含验证的平台标识的ldap中的key，即Ldap中的过滤规则：($platform_key=$platform)，请求url会包含$platform，登录验证时用到
ips = 192.168.0.211,192.168.0.74 # IP请求白名单,不在白名单的禁止请求

[cmdp]
group_base =  # 部门的BaseDn
group_filter = (objectClass=posixGroup) # 部门的过滤规则
group_name_key = cn         # 部门的名称字段
group_id_key = gidNumber         # 部门的id字段

user_base =  # 用户的BaseDn
user_filter = (staffActive=1) # 用户的过滤规则
user_name_key = displayName     # 代表名称的字段
user_id_key = uidNumber         # 代表用户id的字段
user_mail_key = mail            # 代表用户邮箱的字段
user_mobile_key = mobile        # 代表用户手机的字段
user_group_key = gidNumber      # 代表用户所属部门id的字段
user_login_key = cn             # 代表用户登录用户名的字段
append_user_keys = staffAccess,staffActive,sn,givenName

[workcenter]
group_base =  # 部门的BaseDn
group_filter = (objectClass=posixGroup) # 部门的过滤规则
group_name_key = cn
group_id_key = gidNumber

user_base =  # 用户的BaseDn
user_filter = (staffActive=1) # 用户的过滤规则
user_name_key = displayName     # 代表名称的字段
user_id_key = uidNumber         # 代表用户id的字段
user_mail_key = mail            # 代表用户邮箱的字段
user_mobile_key = mobile        # 代表用户手机的字段
user_group_key = gidNumber      # 代表用户所属部门id的字段
user_login_key = cn             # 代表用户登录用户名的字段
append_user_keys = staffAccess,staffActive,sn,givenName   # 用户查询接口返回用户信息会附加这些字段
```
# 编译
```
env GOOS=linux GOARCH=amd64 go build ucore.go
```

# 启动
```
ucore /etc/ucore.ini 
```

# 接口
- 文档用例使用的ip是192.168.0.211，文档说明也使用该ip进行举例
- 文档中**“:platform”**参数，不同平台定义不同的标识，例如：oa系统我们随便给起了一个标识叫oa，ldap中的用户可以访问oa的都会赋予这个属性，所以接口都会使用这个参数

## 1. 部门
### 1.1 获取所有部门
`http://192.168.0.211/v1/:platform/ldap/group/all`
##### 请求方式
- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|:platform |是  |string |平台代码   |


##### 返回示例 

``` 
{
	"code": 200,
	"data": {
		"23766": {
			"DN": "",
			"Name": "经营策略中心",
			"Id": "23766"
		},
		"501": {
			"DN": "",
			"Name": "技术部",
			"Id": "501"
		},
		"502": {
			"DN": "",
			"Name": "策划部",
			"Id": "502"
		},
		"503": {
			"DN": "",
			"Name": "客户一部",
			"Id": "503"
		},
		"504": {
			"DN": "",
			"Name": "客户二部",
			"Id": "504"
		},
		"505": {
			"DN": "",
			"Name": "客户三部",
			"Id": "505"
		},
		"506": {
			"DN": "",
			"Name": "客户四部",
			"Id": "506"
		},
		"507": {
			"DN": "",
			"Name": "创意设计部",
			"Id": "507"
		},
		"508": {
			"DN": "",
			"Name": "行政部",
			"Id": "508"
		},
		"509": {
			"DN": "",
			"Name": "人力资源部",
			"Id": "509"
		},
		"510": {
			"DN": "",
			"Name": "财务部",
			"Id": "510"
		},
		"511": {
			"DN": "",
			"Name": "创新事业部",
			"Id": "511"
		},
		"512": {
			"DN": "",
			"Name": "媒介部",
			"Id": "512"
		}
	},
	"msg": ""
}
```

##### 返回参数说明 

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|        |string   |键值：是部门id  |
|DN      |string   |暂时没有用处，以后若涉及更深层次操作时可能会用到，但可能性不大，在这里做冗余处理  |
|Name    |string   |部门名称  |
|Id    |string   |部门id  |
### 1.2 根据ID获取部门
`http://192.168.0.211/v1/:platform/ldap/group/:gid/`
##### 请求方式
- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|:platform |是  |string |平台代码   |
|:gid |是  |string |部门id   |


##### 返回示例 

``` 
{
	"code": 200,
	"data": {
		"501": {
			"DN": "",
			"Name": "技术部",
			"Id": "501"
		}
	},
	"msg": ""
}
```

##### 返回参数说明 

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|        |string   |键值：是部门id  |
|DN      |string   |暂时没有用处，以后若涉及更深层次操作时可能会用到，但可能性不大，在这里做冗余处理  |
|Name    |string   |部门名称  |
|Id    |string   |部门id  |
## 2. 用户
### 2.1 获取所有用户
`http://192.168.0.211/v1/:platform/user/all`
##### 请求方式
- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|:platform |是  |string |平台代码   |


##### 返回示例 

``` 
{
	"code": 200,
	"data": {
		"1000": {
			"DN": "",
			"CN": "ceshi",
			"DisplayName": "测试",
			"Id": "1000",
			"Mail": "test@mtad.cn",
			"Mobile": "",
			"GidNumber": "501",
			"Append": {
				"givenName": "试",
				"sn": "测",
				"staffAccess": "gitlab",
				"staffActive": "1"
			}
		}
	},
	"msg": ""
}
```

##### 返回参数说明 

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|        |string   |键值：是用户id  |
|DN      |string   |暂时没有用处，以后若涉及更深层次操作时可能会用到，但可能性不大，在这里做冗余处理  |
|CN    |string   |登录用户名  |
|DisplayName    |string   |显示名称  |
|Id    |string   |用户id  |
|Mail    |string   |邮箱  |
|Mobile    |string   |手机号  |
|GidNumber    |string   |用户所属用户组id  |
|Append    |string   |一些附加属性  |
|- - - -  givenName    |string   |名  |
|- - - -  sn    |string   |姓  |
|- - - -  staffAccess    |string   |具有哪些平台标识的访问权限  |
|- - - -  staffActive    |string   |账户是否激活  |
### 2.2 根据用户id获取用户
`http://192.168.0.211/v1/:platform/user/:uid/`
##### 请求方式
- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|:platform |是  |string |平台代码   |
|:uid |是  |string |用户id   |

##### 返回示例 

``` 
{
	"code": 200,
	"data": {
		"1000": {
			"DN": "",
			"CN": "ceshi",
			"DisplayName": "测试",
			"Id": "1000",
			"Mail": "test@mtad.cn",
			"Mobile": "",
			"GidNumber": "501",
			"Append": {
				"givenName": "试",
				"sn": "测",
				"staffAccess": "gitlab",
				"staffActive": "1"
			}
		}
	},
	"msg": ""
}
```

##### 返回参数说明 

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|        |string   |键值：是用户id  |
|DN      |string   |暂时没有用处，以后若涉及更深层次操作时可能会用到，但可能性不大，在这里做冗余处理  |
|CN    |string   |登录用户名  |
|DisplayName    |string   |显示名称  |
|Id    |string   |用户id  |
|Mail    |string   |邮箱  |
|Mobile    |string   |手机号  |
|GidNumber    |string   |用户所属用户组id  |
|Append    |string   |一些附加属性  |
|- - - -  givenName    |string   |名  |
|- - - -  sn    |string   |姓  |
|- - - -  staffAccess    |string   |具有哪些平台标识的访问权限  |
|- - - -  staffActive    |string   |账户是否激活  |
### 2.3 根据登录账号获取用户
`http://192.168.0.211/v1/:platform/:cn/user`
##### 请求方式
- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|:platform |是  |string |平台代码   |
|:cn |是  |string |用户登录名   |

##### 返回示例 

``` 
{
	"code": 200,
	"data": {
		"1000": {
			"DN": "",
			"CN": "ceshi",
			"DisplayName": "测试",
			"Id": "1000",
			"Mail": "test@mtad.cn",
			"Mobile": "",
			"GidNumber": "501",
			"Append": {
				"givenName": "试",
				"sn": "测",
				"staffAccess": "gitlab",
				"staffActive": "1"
			}
		}
	},
	"msg": ""
}
```

##### 返回参数说明 

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|        |string   |键值：是用户id  |
|DN      |string   |暂时没有用处，以后若涉及更深层次操作时可能会用到，但可能性不大，在这里做冗余处理  |
|CN    |string   |登录用户名  |
|DisplayName    |string   |显示名称  |
|Id    |string   |用户id  |
|Mail    |string   |邮箱  |
|Mobile    |string   |手机号  |
|GidNumber    |string   |用户所属用户组id  |
|Append    |string   |一些附加属性  |
|- - - -  givenName    |string   |名  |
|- - - -  sn    |string   |姓  |
|- - - -  staffAccess    |string   |具有哪些平台标识的访问权限  |
|- - - -  staffActive    |string   |账户是否激活  |
### 2.4 【登录】验证账号密码
`http://192.168.0.211/v1/:platform/:cn/auth`
##### 请求方式
- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|:platform |是  |string |平台代码   |
|:cn |是  |string |用户登录名   |

|POST参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|pwd |是  |string |登录密码   |

##### 返回示例 

``` 
{
	"code": 200,
	"data": {
		"DN": "",
		"CN": "kaifaji",
		"DisplayName": "开发机",
		"Id": "61851",
		"Mail": "",
		"Mobile": "",
		"GidNumber": "0",
		"Append": {
			"givenName": "机",
			"sn": "开发",
			"staffAccess": "workcenter",
			"staffActive": "1"
		}
	},
	"msg": "登录成功"
}
```

##### 返回参数说明 

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|DN      |string   |暂时没有用处，以后若涉及更深层次操作时可能会用到，但可能性不大，在这里做冗余处理  |
|CN    |string   |登录用户名  |
|DisplayName    |string   |显示名称  |
|Id    |string   |用户id  |
|Mail    |string   |邮箱  |
|Mobile    |string   |手机号  |
|GidNumber    |string   |用户所属用户组id  |
|Append    |string   |一些附加属性  |
|- - - -  givenName    |string   |名  |
|- - - -  sn    |string   |姓  |
|- - - -  staffAccess    |string   |具有哪些平台标识的访问权限  |
|- - - -  staffActive    |string   |账户是否激活  |

# 附一个PHP请求接口的测试例子
```
<?php

function _post( $url , $parm ) {
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL,$url);
    curl_setopt($ch, CURLOPT_POST, 1);
    curl_setopt($ch, CURLOPT_POSTFIELDS , $parm ) ;
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);

    $server_output = curl_exec($ch);

    curl_close ($ch);
    return $server_output;
}

echo "\n=======================获取所有部门===============================\n";
$url = "http://192.168.0.211/v1/oa/group/all";
$auth_post = "";
print_r( _post($url,$auth_post ));


echo "\n=======================根据id获取部门===============================\n";
$url = "http://192.168.0.211/v1/oa/group/501/";
$auth_post = "";
print_r( _post($url,$auth_post ));


echo "\n=======================获取所有用户===============================\n";
$url = "http://192.168.0.211/v1/oa/user/all";
$auth_post = "";
print_r( _post($url,$auth_post ));


echo "\n=======================根据id获取用户===============================\n";
$url = "http://192.168.0.211/v1/oa/user/65178/";
$auth_post = "";
print_r( _post($url,$auth_post ));



echo "\n=======================根据登录名获取用户===============================\n";
$url = "http://192.168.0.211/v1/oa/pangbo/user";
$auth_post = "";
print_r( _post($url,$auth_post ));


echo "\n=======================登录===============================\n";
$url = "http://192.168.0.211/v1/oa/kaifaji/auth";
$auth_post = "pwd=123456";
print_r( _post($url,$auth_post ));


echo "\n";
```
