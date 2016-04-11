package rbac
import (


)
import "encoding/json"
//获取操作设备权限
func (this *MainController) PermissRead() {
	uinfo := this.GetSession("userinfo")
	/*user_auth_type, _ := strconv.Atoi(beego.AppConfig.String("user_auth_type"))
	rbac_auth_gateway := beego.AppConfig.String("rbac_auth_gateway")
	var accesslist map[string]bool
	if user_auth_type > 0 {
		params := strings.Split(strings.ToLower(UrlFor(this)), "/")

		if uinfo == nil {
			this.Redirect(302, rbac_auth_gateway)
		}
		//admin用户不用认证权限
		adminuser := beego.AppConfig.String("rbac_admin_user")
		if uinfo.(m.User).Username == adminuser {
			this.Rsp(true, "验证")
		}

		if user_auth_type == 1 {
			listbysession := this.GetSession("accesslist")
			if listbysession != nil {
				accesslist = listbysession.(map[string]bool)
			}
		} else if user_auth_type == 2 {

			accesslist, _ = permiss.GetAccessList(uinfo.(m.User).Id)
		}

		ret := permiss.AccessDecision(params, accesslist)
		if !ret {
			this.Rsp( false, "权限不足")
		}
	}	else{
        this.Rsp(true,"跳过验证")
}*/
	userindo,_ := json.Marshal(uinfo)
	this.Rsp(true,string(userindo))
}
