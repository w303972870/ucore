package lib

import(
    "github.com/go-ldap/ldap/v3"
    "fmt"
)
const ObjectCategory_OU string = "organizationalUnit"
const ObjectCategory_Group string = "group"
const ObjectCategory_Person string = "user"

/*创建链接*/
func ( c *LdapConf ) Open() ( error , *ldap.Conn ) {
    l, err := ldap.Dial("tcp", fmt.Sprintf( "%s:%d" ,  c.Host ,  c.Port ) )
    if err != nil {
        return err , nil
    }
    err = l.Bind(  c.BindDn ,  c.DnPassword )
    if err != nil {
        return err , nil
    }
    c.Connection = l
    return nil , l    
}

/*关闭链接*/
func ( c *LdapConf ) Close() {
    if c.Connection != nil {
        c.Connection.Close()
        c.Connection = nil
    }
}

/*查询用户*/
func ( c *LdapConf ) LoginAuth( which string , search string , needDn bool ) ( map [string]*LdapUser , error) {
    err , conn := c.Open()
    if err != nil {
        return nil, err
    }
    defer c.Close()

    result , err := conn.Search( ldap.NewSearchRequest( c.EachLdap[ which ].UserBase , ldap.ScopeWholeSubtree , ldap.NeverDerefAliases , 0 , 0 , false , search , []string{} , nil ) )
    if err != nil {
        return nil , err
    }
    users := make( map [string]*LdapUser  )
    if len( result.Entries ) > 0 {
        for _, item := range result.Entries {
            uid := item.GetAttributeValue( c.EachLdap[ which ].UserIdKey )
            users[ uid ] = new( LdapUser )
            
            if needDn == true {
                users[ uid ].DN = item.DN
            }
            if c.EachLdap[ which ].UserLoginKey != "" {
                users[ uid ].CN = item.GetAttributeValue( c.EachLdap[ which ].UserLoginKey )
            }
            if c.EachLdap[ which ].UserNameKey != "" {
                users[ uid ].DisplayName = item.GetAttributeValue( c.EachLdap[ which ].UserNameKey )
            }
            if c.EachLdap[ which ].UserIdKey != "" {
                users[ uid ].Id = uid
            }
            if c.EachLdap[ which ].UserMailKey != "" {
                users[ uid ].Mail = item.GetAttributeValue( c.EachLdap[ which ].UserMailKey )
            }
            if c.EachLdap[ which ].UserMobileKey != "" {
                users[ uid ].Mobile = item.GetAttributeValue( c.EachLdap[ which ].UserMobileKey )
            }
            if c.EachLdap[ which ].UserGroupKey != "" {
                users[ uid ].GidNumber = item.GetAttributeValue( c.EachLdap[ which ].UserGroupKey )
            }
            users[ uid ].Append = make( map[string]interface{} )
            for _ , k := range c.EachLdap[ which ].UserAppendKey {
                if k != "" {
                    vs := item.GetAttributeValues( k )
                    if len( vs ) > 1 {
                        users[ uid ].Append[ k ] = vs
                    } else if len( vs ) == 1  {
                        users[ uid ].Append[ k ] = vs[0]
                    }
                }
            }
            //users[ uid ].List = item
        }
    }
    return users , err
}

/*查询部门*/
func ( c *LdapConf ) SearchGroup( which string , search string , needDn bool ) (  map [string]*LdapDepart , error ) {
    err , conn := c.Open()
    if err != nil {
        return nil, err
    }
    defer c.Close()
    result , err := conn.Search( ldap.NewSearchRequest( c.EachLdap[ which ].GroupBase , ldap.ScopeWholeSubtree , ldap.NeverDerefAliases , 0 , 0 , false , search , []string{} , nil ) )
    if err != nil {
        return nil, err
    }
    groups := make( map [string]*LdapDepart  )
    if len( result.Entries ) > 0 {
        for _, item := range result.Entries {
            groups[ item.GetAttributeValue( c.EachLdap[ which ].GroupIdKey ) ] = new( LdapDepart )
            if needDn == true {
                groups[ item.GetAttributeValue( c.EachLdap[ which ].GroupIdKey ) ].DN = item.DN
            }
            groups[ item.GetAttributeValue( c.EachLdap[ which ].GroupIdKey ) ].Name = item.GetAttributeValue( c.EachLdap[ which ].GroupNameKey )
            groups[ item.GetAttributeValue( c.EachLdap[ which ].GroupIdKey ) ].Id = item.GetAttributeValue( c.EachLdap[ which ].GroupIdKey )
        }
    }
    return groups,err
}

/*查询用户*/
func ( c *LdapConf ) SearchUser( which string , search string , needDn bool ) ( map [string]*LdapUser , error) {
    err , conn := c.Open()
    if err != nil {
        return nil, err
    }
    defer c.Close()

    result , err := conn.Search( ldap.NewSearchRequest( c.EachLdap[ which ].UserBase , ldap.ScopeWholeSubtree , ldap.NeverDerefAliases , 0 , 0 , false , search , []string{} , nil ) )
    if err != nil {
        return nil , err
    }
    users := make( map [string]*LdapUser  )
    if len( result.Entries ) > 0 {
        for _, item := range result.Entries {
            uid := item.GetAttributeValue( c.EachLdap[ which ].UserIdKey )
            users[ uid ] = new( LdapUser )
            
            if needDn == true {
                users[ uid ].DN = item.DN
            }
            if c.EachLdap[ which ].UserLoginKey != "" {
                users[ uid ].CN = item.GetAttributeValue( c.EachLdap[ which ].UserLoginKey )
            }
            if c.EachLdap[ which ].UserNameKey != "" {
                users[ uid ].DisplayName = item.GetAttributeValue( c.EachLdap[ which ].UserNameKey )
            }
            if c.EachLdap[ which ].UserIdKey != "" {
                users[ uid ].Id = uid
            }
            if c.EachLdap[ which ].UserMailKey != "" {
                users[ uid ].Mail = item.GetAttributeValue( c.EachLdap[ which ].UserMailKey )
            }
            if c.EachLdap[ which ].UserMobileKey != "" {
                users[ uid ].Mobile = item.GetAttributeValue( c.EachLdap[ which ].UserMobileKey )
            }
            if c.EachLdap[ which ].UserGroupKey != "" {
                users[ uid ].GidNumber = item.GetAttributeValue( c.EachLdap[ which ].UserGroupKey )
            }
            users[ uid ].Append = make( map[string]interface{} )
            for _ , k := range c.EachLdap[ which ].UserAppendKey {
                if k != "" {
                    vs := item.GetAttributeValues( k )
                    if len( vs ) > 1 {
                        users[ uid ].Append[ k ] = vs
                    } else if len( vs ) == 1  {
                        users[ uid ].Append[ k ] = vs[0]
                    }
                }
            }
            //users[ uid ].List = item
        }
    }
    return users , err
}