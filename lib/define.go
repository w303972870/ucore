package lib

import(
    "github.com/gin-gonic/gin"
    "gopkg.in/ini.v1"
    "github.com/go-ldap/ldap/v3"
)

/*整体系统配置变量信息*/
type ConfigParms struct {
     /*ini对象*/    
    IniPkg *ini.File

     /*Gin对象*/
    SysGin *gin.Engine
    
    /*控制并行数*/
    Cores int

    /*配置文件*/
    ConfFile string

    /*系统平台类型*/
    SysType string

    /*ini配置文件内容*/
    Config *uIni

    /*当前请求的平台标识*/
    RequestPlatform string

    /*支持验证的平台标识*/
    Platform []string

    /*支持验证的平台标识key*/
    PlatformKey string

    /*ip白名单*/
    WhiteIps []string
}

/*配置文件*/
type uIni struct {

    /*是否后台运行*/
    Daemon bool

    /*监听ip*/
    Host string

    /*监听端口*/
    Port int

    /*日志目录*/
    LogsDir string

    /*日志名称*/
    LogsName string

    /*ldap配置部分*/
    uLdap *LdapConf
}

/*配置文件:ldap配置*/
type LdapConf struct {
    /*ldap服务器*/
    Host string

    /*ldap端口*/
    Port int

    /*绑定管理员*/
    BindDn string

    /*管理员密码*/
    DnPassword string

    /*Ldap对象*/
    Connection *ldap.Conn

    /*每一个ldap服务的具体配置*/
    EachLdap map[string]*LdapEach
}

/*配置文件:ldap具体配置*/
type LdapEach struct {

    /*部门Base*/
    GroupBase string

    /*部门过滤*/
    GroupFilter string

    /*部门名称key*/
    GroupNameKey string

    /*部门id key*/
    GroupIdKey string

    /*用户base*/
    UserBase string

    /*用户过滤*/
    UserFilter string

    /*代表名称的字段*/
    UserNameKey string

    /*代表用户id的字段*/
    UserIdKey string

    /*代表用户邮箱的字段*/
    UserMailKey string

    /*代表用户手机的字段*/
    UserMobileKey string

    /*代表用户所属部门id的字段*/
    UserGroupKey string

    /*代表用户登录用户名的字段*/
    UserLoginKey string

    /*用户信息附加字段*/
    UserAppendKey []string
}

/*Ldap部门*/
type LdapDepart struct {
    DN string
    Name string
    Id string
}

/*Ldap用户*/
type LdapUser struct {
    DN string
    CN string
    DisplayName string
    Id string
    Mail string
    Mobile string
    GidNumber string
    Append map[string]interface{}
    //List * ldap.Entry
}

/*支持验证的平台标识*/
func ( c * ConfigParms ) SetPlatform( l []string ) {
    c.Platform = make( []string , len(l) )
    c.Platform = l
    c.Config.uLdap.EachLdap = make( map[string]*LdapEach , len(l) )
    for _ , name := range l {
        c.Config.uLdap.EachLdap[string(name)] = new(LdapEach)
    }
}

/*ip白名单*/
func ( c * ConfigParms ) SetWhiteIps( l []string ) {
    c.WhiteIps = make( []string , len(l) )
    c.WhiteIps = l
}

/*配置详情*/
func ( c * ConfigParms ) ConfLdap() **LdapConf {
    return &c.Config.uLdap
}

var LibConfigParms ConfigParms

/*初始化*/
func init() {
    LibConfigParms.Config = &uIni{}
    LibConfigParms.Config.uLdap = &LdapConf{}
    gin.SetMode( gin.ReleaseMode )
    LibConfigParms.SysGin = gin.New()
}
