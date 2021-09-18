package model
import(
    "fmt"
    "ucore/lib"   
)

/*根据用户名找出对应平台权限中的用户*/
func SearchPlatFormUser( name string )(map [string]*lib.LdapUser , error) {
    filter := fmt.Sprintf( 
        "(&%s(%s=%s)(%s=%s))" , 
        (*lib.LibConfigParms.ConfLdap()).EachLdap[lib.LibConfigParms.RequestPlatform].UserFilter , 
        (*lib.LibConfigParms.ConfLdap()).EachLdap[lib.LibConfigParms.RequestPlatform].UserLoginKey , name ,
        lib.LibConfigParms.PlatformKey , lib.LibConfigParms.RequestPlatform )
    
    return (*lib.LibConfigParms.ConfLdap()).SearchUser( lib.LibConfigParms.RequestPlatform,filter , true  )
}

/*获取所有部门*/
func AllGroups() (  map [string]*lib.LdapDepart , error ) {
    return (*lib.LibConfigParms.ConfLdap()).SearchGroup( 
        lib.LibConfigParms.RequestPlatform,(*lib.LibConfigParms.ConfLdap()).EachLdap[lib.LibConfigParms.RequestPlatform].GroupFilter , false )
}

/*根据部门id获取部门信息*/
func GetGroupById( gid string )(  map [string]*lib.LdapDepart , error )  {
    filter := fmt.Sprintf( 
        "(&%s(%s=%s))" , 
        (*lib.LibConfigParms.ConfLdap()).EachLdap[lib.LibConfigParms.RequestPlatform].GroupFilter , 
        (*lib.LibConfigParms.ConfLdap()).EachLdap[lib.LibConfigParms.RequestPlatform].GroupIdKey , 
        gid )
    
    return (*lib.LibConfigParms.ConfLdap()).SearchGroup( lib.LibConfigParms.RequestPlatform,filter , false  )
}

/*获取所有用户*/
func AllUsers()(map [string]*lib.LdapUser , error)  {
    return (*lib.LibConfigParms.ConfLdap()).SearchUser( 
        lib.LibConfigParms.RequestPlatform,(*lib.LibConfigParms.ConfLdap()).EachLdap[lib.LibConfigParms.RequestPlatform].UserFilter , false )
}

/*根据用户名找出对应平台权限中的用户*/
func GetUserById( uid string )(map [string]*lib.LdapUser , error) {
        filter := fmt.Sprintf( 
            "(&%s(%s=%s))" , 
            (*lib.LibConfigParms.ConfLdap()).EachLdap[lib.LibConfigParms.RequestPlatform].UserFilter , 
            (*lib.LibConfigParms.ConfLdap()).EachLdap[lib.LibConfigParms.RequestPlatform].UserIdKey , 
            uid )
        
    return  (*lib.LibConfigParms.ConfLdap()).SearchUser( lib.LibConfigParms.RequestPlatform,filter , false  )
}

/*不受平台限制，根据用户登录名获取用户信息*/
func GetUserByCn( cn string )(map [string]*lib.LdapUser , error)  {
    filter := fmt.Sprintf( 
        "(&%s(%s=%s))" , 
        (*lib.LibConfigParms.ConfLdap()).EachLdap[lib.LibConfigParms.RequestPlatform].UserFilter , 
        (*lib.LibConfigParms.ConfLdap()).EachLdap[lib.LibConfigParms.RequestPlatform].UserLoginKey , 
        cn )
    
    return (*lib.LibConfigParms.ConfLdap()).SearchUser( lib.LibConfigParms.RequestPlatform,filter , false  )
}
