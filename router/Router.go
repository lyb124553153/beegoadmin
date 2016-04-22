package router

import (
	"github.com/astaxie/beego"
	"github.com/beego/admin/src/rbac"
)
func router() {
	beego.Router("/", &rbac.MainController{}, "*:Index")
	beego.Router("/public/index", &rbac.MainController{}, "*:Index")
	beego.Router("/public/login", &rbac.MainController{}, "*:Login")
	beego.Router("/public/logout", &rbac.MainController{}, "*:Logout")
	beego.Router("/public/changepwd", &rbac.MainController{}, "*:Changepwd")

	beego.Router("/rbac/user/AddUser", &rbac.UserController{}, "*:AddUser")
	beego.Router("/rbac/user/UpdateUser", &rbac.UserController{}, "*:UpdateUser")
	beego.Router("/rbac/user/DelUser", &rbac.UserController{}, "*:DelUser")
	beego.Router("/rbac/user/index", &rbac.UserController{}, "*:Index")

	beego.Router("/rbac/node/AddAndEdit", &rbac.NodeController{}, "*:AddAndEdit")
	beego.Router("/rbac/node/DelNode", &rbac.NodeController{}, "*:DelNode")
	beego.Router("/rbac/node/index", &rbac.NodeController{}, "*:Index")

	beego.Router("/rbac/group/AddGroup", &rbac.GroupController{}, "*:AddGroup")
	beego.Router("/rbac/group/UpdateGroup", &rbac.GroupController{}, "*:UpdateGroup")
	beego.Router("/rbac/group/DelGroup", &rbac.GroupController{}, "*:DelGroup")
	beego.Router("/rbac/group/index", &rbac.GroupController{}, "*:Index")

	beego.Router("/rbac/role/AddAndEdit", &rbac.RoleController{}, "*:AddAndEdit")
	beego.Router("/rbac/role/DelRole", &rbac.RoleController{}, "*:DelRole")
	beego.Router("/rbac/role/AccessToNode", &rbac.RoleController{}, "*:AccessToNode")
	beego.Router("/rbac/role/AddAccess", &rbac.RoleController{}, "*:AddAccess")
	beego.Router("/rbac/role/RoleToUserList", &rbac.RoleController{}, "*:RoleToUserList")
	beego.Router("/rbac/role/AddRoleToUser", &rbac.RoleController{}, "*:AddRoleToUser")
	beego.Router("/rbac/role/Getlist", &rbac.RoleController{}, "*:Getlist")
	beego.Router("/rbac/role/index", &rbac.RoleController{}, "*:Index")

	beego.Router("/rbac/machine/AddMachine",&rbac.MachineController{},"*:AddMachine")
	beego.Router("/rbac/machine/index",&rbac.MachineController{},"*:Index")
	beego.Router("/rbac/machine/DelMachine",&rbac.MachineController{},"*:DelMachine")
	beego.Router("/rbac/machine/UpdateMachine",&rbac.MachineController{},"*:UpdateMachine")
	beego.Router("/rbac/machine/AddAndEdit", &rbac.MachineController{}, "*:AddAndEdit")
	beego.Router("/rbac/machine/MachineList", &rbac.MachineController{}, "*:MachineList")

	beego.Router("/rbac/strategy/AddAndEdit", &rbac.StrategyController{}, "*:AddAndEdit")
	beego.Router("/rbac/strategy/DelStrategy", &rbac.StrategyController{}, "*:DelStrategy")
	beego.Router("/rbac/strategy/StrategyToMachineList", &rbac.StrategyController{}, "*:StrategyToMachineList")
	beego.Router("/rbac/strategy/AddStrategyToMachine", &rbac.StrategyController{}, "*:AddStrategyToMachine")
	beego.Router("/rbac/strategy/Getlist", &rbac.StrategyController{}, "*:Getlist")
	beego.Router("/rbac/strategy/index", &rbac.StrategyController{}, "*:Index")



	beego.Router("/machine/AddMachine",&rbac.MachineController{},"*:AddMachine")
	beego.Router("/machine/DelMachine",&rbac.MachineController{},"*:DelMachine")
	beego.Router("/machine/MachineInfo",&rbac.MachineController{},"*:MachineInfo")

	beego.Router("/qrcode",&rbac.MachineController{},"*:Qrcodebuild")
	beego.Router("/permiss",&rbac.MainController{},"*:PermissRead")

}
