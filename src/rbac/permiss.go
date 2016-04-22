package rbac

//获取操作设备权限
func (this *MainController) PermissRead() {
	uinfo := this.GetSession("accesslist")
	this.Data["json"] = &uinfo;
	this.ServeJSON();
}

func(this *MachineController) Qrcodebuild(){

	this.TplName =  "qrcode.html";

}