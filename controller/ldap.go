package controller

import (
    "github.com/gin-gonic/gin"
    "ucore/model"
    "ucore/lib"
    "ucore/logs"
)

/*登录*/
func Auth( c *gin.Context )  {
    name := c.Param("cn")
    passwd := c.PostForm("pwd")
    if name == "" || passwd == "" {
        c.JSON( lib.StatusUnauthorized , ulog.ULogs.GinJson( lib.StatusUnauthorized , "用户名/密码错误" , nil ) )
        return
    }

    searchresult, err := model.SearchPlatFormUser( name  )

    if err != nil || len( searchresult ) != 1  {
        c.JSON( lib.StatusBadRequest , ulog.ULogs.GinJson( lib.StatusBadRequest , "无权登录" , nil ) )
    } else {
        err , conn := (*lib.LibConfigParms.ConfLdap()).Open()
        defer (*lib.LibConfigParms.ConfLdap()).Close()
        
        if err != nil {
            c.JSON( lib.StatusInternalServerError , ulog.ULogs.GinJson( lib.StatusInternalServerError , err.Error() , nil ) )
        } else {
            for _,_user := range searchresult {
                err := conn.Bind( _user.DN , passwd )
                if err != nil {
                    c.JSON( lib.StatusUnauthorized , ulog.ULogs.GinJson( lib.StatusUnauthorized , "用户名/密码错误" , nil ) )
                } else {
                    _user.DN = ""
                    c.JSON( lib.StatusOK , ulog.ULogs.GinJson( lib.StatusOK , "登录成功" , _user ) )
                }
                return
            }
        }
    }
}

/*获取所有部门*/
func AllGroups( c *gin.Context )  {
    searchresult, err := model.AllGroups()
    if err != nil {
        c.JSON( lib.StatusBadRequest , ulog.ULogs.GinJson( lib.StatusBadRequest , err.Error() , nil ) )
    } else {
        c.JSON( lib.StatusOK , ulog.ULogs.GinJson( lib.StatusOK , "" , searchresult ) )
    }
}

/*根据部门id获取部门信息*/
func GetGroupById( c *gin.Context )  {
    gid := c.Param("gid")
    if gid == "" {
        c.JSON( lib.StatusBadRequest , ulog.ULogs.GinJson( lib.StatusBadRequest , "请求错误" , nil ) )
    } else {
        searchresult, err := model.GetGroupById( gid )
        if err != nil {
            c.JSON( lib.StatusBadRequest , ulog.ULogs.GinJson( lib.StatusBadRequest , err.Error() , nil ) )
        } else {
            c.JSON( lib.StatusOK , ulog.ULogs.GinJson( lib.StatusOK , "" , searchresult ) )
        }
    }
}

/*获取所有用户*/
func AllUsers( c *gin.Context )  {
    searchresult, err := model.AllUsers()
    if err != nil {
        c.JSON( lib.StatusBadRequest , ulog.ULogs.GinJson( lib.StatusBadRequest , err.Error() , nil ) )
    } else {
        c.JSON( lib.StatusOK , ulog.ULogs.GinJson( lib.StatusOK , "" , searchresult ) )
    }
}

/*根据用户id获取用户信息*/
func GetUserById( c *gin.Context )  {
    uid := c.Param("uid")
    if uid == "" {
        c.JSON( lib.StatusBadRequest , ulog.ULogs.GinJson( lib.StatusBadRequest , "请求错误" , nil ) )
    } else {
        searchresult, err := model.GetUserById( uid )
        if err != nil {
            c.JSON( lib.StatusBadRequest , ulog.ULogs.GinJson( lib.StatusBadRequest , err.Error() , nil ) )
        } else {
            c.JSON( lib.StatusOK , ulog.ULogs.GinJson( lib.StatusOK , "" , searchresult ) )
        }
    }
}

/*不受平台限制，根据用户登录名获取用户信息*/
func GetUserByCn( c *gin.Context )  {
    cn := c.Param("cn")
    if cn == "" {
        c.JSON( lib.StatusBadRequest , ulog.ULogs.GinJson( lib.StatusBadRequest , "请求错误" , nil ) )
    } else {
        searchresult, err := model.GetUserByCn( cn )
        if err != nil {
            c.JSON( lib.StatusBadRequest , ulog.ULogs.GinJson( lib.StatusBadRequest , err.Error() , nil ) )
        } else {
            c.JSON( lib.StatusOK , ulog.ULogs.GinJson( lib.StatusOK , "" , searchresult ) )
        }
    }
}


