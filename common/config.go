package ucore
import (
    "gopkg.in/ini.v1"
    "os"
    "strings"
    "ucore/lib"
    "ucore/logs"
)

func init() {
    if len(os.Args) < 2 {
        ulog.ULogs.Error( "未传入配置文件" )
    }
    lib.LibConfigParms.ConfFile = os.Args[1]
    
    in , err := os.Open( lib.LibConfigParms.ConfFile )
    if  err != nil {
       ulog.ULogs.Error( "配置文件打开错误" )
    }
    in.Close()

    cfg , err := ini.Load( lib.LibConfigParms.ConfFile )
    if err != nil {
        ulog.ULogs.Error( "配置文件读取错误" )
    }
    lib.LibConfigParms.IniPkg = cfg
}

func InitConf() {
    
    /* Daemon */
    lib.LibConfigParms.Config.Daemon = lib.LibConfigParms.IniPkg.Section("ucore").Key("daemon").MustBool(true)

    /* LogDir */
    lib.LibConfigParms.Config.LogsDir = lib.LibConfigParms.IniPkg.Section("log").Key("dir").String()
    if is_path , _ := UTools.Exist( lib.LibConfigParms.Config.LogsDir ) ; is_path != true {
        ulog.ULogs.Info( UTools.Str( lib.LibConfigParms.ConfFile , "\t" , lib.LibConfigParms.Config.LogsDir ,  "\t日志目录错误，系统将自动创建" ) )
        UTools.MkDir( lib.LibConfigParms.Config.LogsDir )
    }
    /*logName*/
    lib.LibConfigParms.Config.LogsName = lib.LibConfigParms.IniPkg.Section("log").Key("logname").MustString("access-ucore.log")

    /* Port */
    lib.LibConfigParms.Config.Port = lib.LibConfigParms.IniPkg.Section("ucore").Key("port").MustInt(80)
    
    /* Host */
    ip := lib.LibConfigParms.IniPkg.Section("ucore").Key("host").MustString("")
    if ip == "" || UTools.IsIp( ip ) == true {
        lib.LibConfigParms.Config.Host = ip
    } else {
        ulog.ULogs.Error( UTools.Str( lib.LibConfigParms.ConfFile , "\t" , ip ,  "\tIp错误" ) )
    }

    /*Ldap*/
    (*lib.LibConfigParms.ConfLdap()).Host = lib.LibConfigParms.IniPkg.Section("ldap").Key("host").MustString("")
    if (*lib.LibConfigParms.ConfLdap()).Host == "" {
        ulog.ULogs.Error( UTools.Str( lib.LibConfigParms.ConfFile , "\tLdap Host错误" ) )
    }

    (*lib.LibConfigParms.ConfLdap()).BindDn = lib.LibConfigParms.IniPkg.Section("ldap").Key("bind_dn").MustString("")
    if (*lib.LibConfigParms.ConfLdap()).BindDn == "" {
        ulog.ULogs.Error( UTools.Str( lib.LibConfigParms.ConfFile , "\tLdap Dn错误" ) )
    }

    (*lib.LibConfigParms.ConfLdap()).DnPassword = lib.LibConfigParms.IniPkg.Section("ldap").Key("password").MustString("")
    if (*lib.LibConfigParms.ConfLdap()).DnPassword == "" {
        ulog.ULogs.Error( UTools.Str( lib.LibConfigParms.ConfFile , "\tLdap DnPassword错误" ) )
    }

    /*ldap:port*/
    (*lib.LibConfigParms.ConfLdap()).Port = lib.LibConfigParms.IniPkg.Section("ldap").Key("port").MustInt(389)

    /*platform*/
    lib.LibConfigParms.SetPlatform( strings.Split( lib.LibConfigParms.IniPkg.Section("platform").Key("platform").MustString("") , ",") )
    if len( lib.LibConfigParms.Platform ) <= 1 && lib.LibConfigParms.Platform[0] == "" {
        ulog.ULogs.Error( UTools.Str( lib.LibConfigParms.ConfFile , "\tplatform配置错误或未配置，将无法验证" ) )
    } else {
        for _ , v := range lib.LibConfigParms.Platform {
            _ , err := lib.LibConfigParms.IniPkg.SectionsByName(v)
            if err != nil {
                ulog.ULogs.Error( UTools.Str( lib.LibConfigParms.ConfFile , "\t" , v ,"缺失" ) )
            }
            (*lib.LibConfigParms.ConfLdap()).EachLdap[v].UserAppendKey = strings.Split( lib.LibConfigParms.IniPkg.Section( v ).Key("append_user_keys").MustString("") , "," )
            (*lib.LibConfigParms.ConfLdap()).EachLdap[v].GroupBase = lib.LibConfigParms.IniPkg.Section( v ).Key("group_base").MustString("")
            (*lib.LibConfigParms.ConfLdap()).EachLdap[v].GroupFilter = lib.LibConfigParms.IniPkg.Section( v ).Key("group_filter").MustString("")
            (*lib.LibConfigParms.ConfLdap()).EachLdap[v].GroupNameKey = lib.LibConfigParms.IniPkg.Section( v ).Key("group_name_key").MustString("")
            (*lib.LibConfigParms.ConfLdap()).EachLdap[v].GroupIdKey = lib.LibConfigParms.IniPkg.Section( v ).Key("group_id_key").MustString("")
            (*lib.LibConfigParms.ConfLdap()).EachLdap[v].UserBase = lib.LibConfigParms.IniPkg.Section( v ).Key("user_base").MustString("")
            (*lib.LibConfigParms.ConfLdap()).EachLdap[v].UserFilter = lib.LibConfigParms.IniPkg.Section( v ).Key("user_filter").MustString("")
            (*lib.LibConfigParms.ConfLdap()).EachLdap[v].UserNameKey = lib.LibConfigParms.IniPkg.Section( v ).Key("user_name_key").MustString("")
            (*lib.LibConfigParms.ConfLdap()).EachLdap[v].UserIdKey = lib.LibConfigParms.IniPkg.Section( v ).Key("user_id_key").MustString("")
            (*lib.LibConfigParms.ConfLdap()).EachLdap[v].UserMailKey = lib.LibConfigParms.IniPkg.Section( v ).Key("user_mail_key").MustString("")
            (*lib.LibConfigParms.ConfLdap()).EachLdap[v].UserMobileKey = lib.LibConfigParms.IniPkg.Section( v ).Key("user_mobile_key").MustString("")
            (*lib.LibConfigParms.ConfLdap()).EachLdap[v].UserGroupKey = lib.LibConfigParms.IniPkg.Section( v ).Key("user_group_key").MustString("")
            (*lib.LibConfigParms.ConfLdap()).EachLdap[v].UserLoginKey = lib.LibConfigParms.IniPkg.Section( v ).Key("user_login_key").MustString("")
            if (*lib.LibConfigParms.ConfLdap()).EachLdap[v].GroupBase == "" || (*lib.LibConfigParms.ConfLdap()).EachLdap[v].UserBase == "" {
                ulog.ULogs.Waring( UTools.Str( lib.LibConfigParms.ConfFile , "\t" , v ,"\tBase配置缺失" ) )
            }
        }
    }
    lib.LibConfigParms.PlatformKey = lib.LibConfigParms.IniPkg.Section("platform").Key("platform_key").MustString("")
    if lib.LibConfigParms.PlatformKey == "" {
        ulog.ULogs.Error( UTools.Str( lib.LibConfigParms.ConfFile , "\tLdap platform_key错误" ) )
    }

    /*IP白名单*/
    lib.LibConfigParms.SetWhiteIps( strings.Split( lib.LibConfigParms.IniPkg.Section("platform").Key("ips").MustString("") , ",") )

    if len( lib.LibConfigParms.WhiteIps ) <= 1 && lib.LibConfigParms.WhiteIps[0] == "" {
        ulog.ULogs.Error( UTools.Str( lib.LibConfigParms.ConfFile , "\tplatform配置错误或未配置，将无法验证" ) )
    }
    
    lib.LibConfigParms.IniPkg = nil
}