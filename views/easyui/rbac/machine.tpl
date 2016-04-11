{{template "../public/header.tpl"}}
<script type="text/javascript">
var statuslist = [
    {statusid:'1',name:'禁用'},
    {statusid:'2',name:'启用'}
];
var URL="/rbac/machine";
$(function(){
    //用户列表
    $("#datagrid").datagrid({
        title:'machine',
        url:URL+'/index',
        method:'POST',
        pagination:true,
        fitColumns:true,
        striped:true,
        rownumbers:true,
        singleSelect:true,
        idField:'Id',
        pagination:true,
        pageSize:20,
        pageList:[10,20,30,50,100],
        columns:[[
            {field:'Id',title:'ID',width:50,sortable:true},
            {field:'MachineNo',title:'设备编号',width:100,sortable:true},
            {field:'ChargeName',title:'负责人',width:100,align:'center',editor:'text'},
            {field:'ContactWay',title:'联系方式',width:100,align:'center',editor:'text'},
            {field:'SiteName',title:'网站名',width:100,align:'center',editor:'text'},
            {field:'SiteUrl',title:'网站地址',width:100,align:'center',editor:'text'},
            {field:'Status',title:'状态',width:50,align:'center',
                            formatter:function(value){
                                for(var i=0; i<statuslist.length; i++){
                                    if (statuslist[i].statusid == value) return statuslist[i].name;
                                }
                                return value;
                            },
                            editor:{
                                type:'combobox',
                                options:{
                                    valueField:'statusid',
                                    textField:'name',
                                    data:statuslist,
                                    required:true
                                }
                            }
                        },
               {field:'Registtime',title:'注册时间',width:100,align:'center',
                                        formatter:function(value,row,index){
                                            if(value) return phpjs.date("Y-m-d H:i:s",phpjs.strtotime(value));
                                            return value;
                                        }
                                    }

        ]],
                  onAfterEdit:function(index, data, changes){
                      if(vac.isEmpty(changes)){
                          return;
                      }
                      changes.Id = data.Id;
                      vac.ajax(URL+'/UpdateMachine', changes, 'POST', function(r){
                          if(!r.status){
                              vac.alert(r.info);
                          }else{
                              $("#datagrid").datagrid("reload");
                          }
                      })
                  },
                  onDblClickRow:function(index,row){
                      editrow();
                  },
                  onRowContextMenu:function(e, index, row){
                      e.preventDefault();
                      $(this).datagrid("selectRow",index);
                      $('#mm').menu('show',{
                          left: e.clientX,
                          top: e.clientY
                      });
                  },
                  onHeaderContextMenu:function(e, field){
                      e.preventDefault();
                      $('#mm1').menu('show',{
                          left: e.clientX,
                          top: e.clientY
                      });
                  }



    });
    //创建设备添加窗口
    $("#dialog").dialog({
        modal:true,
        resizable:true,
        top:150,
        closed:true,
        buttons:[{
            text:'保存',
            iconCls:'icon-save',
            handler:function(){
                $("#form1").form('submit',{
                    url:URL+'/AddMachine',
                    onSubmit:function(){
                        return $("#form1").form('validate');
                    },
                    success:function(r){
                        var r = $.parseJSON( r );
                        if(r.status){
                            $("#dialog").dialog("close");
                            $("#datagrid").datagrid('reload');
                        }else{
                            vac.alert(r.info);
                        }
                    }
                });
            }
        },{
            text:'取消',
            iconCls:'icon-cancel',
            handler:function(){
                $("#dialog").dialog("close");
            }
        }]
    });


})

function editrow(){
    if(!$("#datagrid").datagrid("getSelected")){
        vac.alert("请选择要编辑的行");
        return;
    }
    $('#datagrid').datagrid('beginEdit', vac.getindex("datagrid"));
}
function saverow(index){
    if(!$("#datagrid").datagrid("getSelected")){
        vac.alert("请选择要保存的行");
        return;
    }
    $('#datagrid').datagrid('endEdit', vac.getindex("datagrid"));
}
//取消
function cancelrow(){
    if(! $("#datagrid").datagrid("getSelected")){
        vac.alert("请选择要取消的行");
        return;
    }
    $("#datagrid").datagrid("cancelEdit",vac.getindex("datagrid"));
}
//刷新
function reloadrow(){
    $("#datagrid").datagrid("reload");
}

//添加用户弹窗
function addrow(){
    $("#dialog").dialog('open');
    $("#form1").form('clear');
}



//删除设备
function delrow(){
    $.messager.confirm('Confirm','你确定要删除?',function(r){
        if (r){
            var row = $("#datagrid").datagrid("getSelected");
            if(!row){
                vac.alert("请选择要删除的行");
                return;
            }
            vac.ajax(URL+'/DelMachine', {Id:row.Id}, 'POST', function(r){
                if(r.status){
                    $("#datagrid").datagrid('reload');
                }else{
                    vac.alert(r.info);
                }
            })
        }
    });
}
</script>
<body>
<table id="datagrid" toolbar="#tb"></table>
<div id="tb" style="padding:5px;height:auto">
    <a href="#" icon='icon-add' plain="true" onclick="addrow()" class="easyui-linkbutton" >新增</a>
    <a href="#" icon='icon-edit' plain="true" onclick="editrow()" class="easyui-linkbutton" >编辑</a>
    <a href="#" icon='icon-save' plain="true" onclick="saverow()" class="easyui-linkbutton" >保存</a>
    <a href="#" icon='icon-cancel' plain="true" onclick="delrow()" class="easyui-linkbutton" >删除</a>
    <a href="#" icon='icon-reload' plain="true" onclick="reloadrow()" class="easyui-linkbutton" >刷新</a>
    <a href="#" icon='icon-reload' plain="true" onclick="reloadrow()" class="easyui-linkbutton" >生成二维码</a>
    <a href="#" icon='icon-edit' plain="true" onclick="updateuserpassword()" class="easyui-linkbutton" >修改用户密码</a>
</div>
<!--表格内的右键菜单-->
<div id="mm" class="easyui-menu" style="width:120px;display: none" >
    <div iconCls='icon-add' onclick="addrow()">新增</div>
    <div iconCls="icon-edit" onclick="editrow()">编辑</div>
    <div iconCls='icon-save' onclick="saverow()">保存</div>
    <div iconCls='icon-cancel' onclick="cancelrow()">取消</div>
    <div class="menu-sep"></div>
    <div iconCls='icon-cancel' onclick="delrow()">删除</div>
    <div iconCls='icon-reload' onclick="reloadrow()">刷新</div>
    <div class="menu-sep"></div>
    <div>Exit</div>
</div>
<!--表头的右键菜单-->
<div id="mm1" class="easyui-menu" style="width:120px;display: none"  >
    <div icon='icon-add' onclick="addrow()">新增</div>
</div>
<div id="dialog" title="添加用户" style="width:400px;height:400px;">
    <div style="padding:20px 20px 40px 80px;" >
        <form id="form1" method="post">
            <table>
                <tr>
                    <td>设备编号：</td>
                    <td><input name="MachineNo" class="easyui-validatebox" required="true"/></td>
                </tr>
                <tr>
                    <td>网站名称：</td>
                    <td><input name="SiteName" class="easyui-validatebox" required="true"  /></td>
                </tr>
                <tr>
                <td>网站域名：</td>
                   <td><input name="SiteUrl" class="easyui-validatebox" required="true"  /></td>
                </tr>
                <tr>
                     <td>负责人：</td>
                    <td><input name="ChargeName" class="easyui-validatebox" required="true"  /></td>
                 </tr>
                  <tr>
                       <td>联系方式：</td>
                        <td><input name="ContactWay" class="easyui-validatebox" required="true"  /></td>
                      </tr>
                    <tr>
                                       <td>状态：</td>
                                       <td>
                                           <select name="Status"  style="width:153px;" class="easyui-combobox " editable="false" required="true"  >
                                               <option value="2" >启用</option>
                                               <option value="1">禁用</option>
                                           </select>
                                       </td>
                                   </tr>

            </table>
        </form>
    </div>
</div>
</body>
</html>