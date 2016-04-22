{{template "../public/header.tpl"}}
<script type="text/javascript">
    var products = [
            {productid:'1',name:'写保护'},
        {productid:'2',name:'可读写'}
    ];
    var URL="/rbac/machine";
    $(function(){
        $("#treegrid").treegrid({
            url:URL+"/index",
            idField:"Id",
            treeField:"Paths",
            fitColumns:"true",
            columns:[[
                {field:'Paths',title:'目录',width:100,sortable:true,editor:'text'},
                {field:'Id',title:'ID',width:50},
                {field:'MachineNo',title:'设备编号',width:100,sortable:true,editor:'text'},
                {field:'ChargeName',title:'负责人',width:100,align:'center',editor:'text'},
                {field:'ContactWay',title:'联系方式',width:100,align:'center',editor:'text'},
                {field:'SiteName',title:'网站名',width:100,align:'center',editor:'text'},
                {field:'Registtime',title:'注册时间',width:100,align:'center',formatter:function(value,row,index){
                                                                                              if(value) return phpjs.date("Y-m-d H:i:s",phpjs.strtotime(value));
                                                                                              return value;
                                                                                          }},
                {field:'SiteUrl',title:'网站地址',width:100,align:'center',editor:'text'},
                {field:'Status',title:'状态',width:50,align:'center',
                    formatter:function(value){
                        for(var i=0; i<products.length; i++){
                            if (products[i].productid == value) return products[i].name;
                        }
                        return value;
                    },
                    editor:{
                        type:'combobox',
                        options:{
                            valueField:'productid',
                            textField:'name',
                            data:products,
                            required:true
                        }
                    }
                },
                {field:'Remark',title:'描述',width:150,editor:'text'}
            ]],
            onAfterEdit:function(c){
                if(vac.isEmpty(c)){
                    return;
                }
                vac.ajax(URL+'/AddAndEdit', c, 'POST', function(r){
                    if(!r.status){
                        vac.alert(r.info);
                    }else{
                                               vac.alert(r.info)
                                }
                        ;
                
                })
            },
            onDblClickRow:function(index,row){
                editrow();
            },
            onContextMenu:function(e, row){
                e.preventDefault();
                $(this).treegrid('select', row.Id);
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
  
});
    //新增行
    function addrow(){
        var Row = $("#treegrid").treegrid("getSelected");
        var data = [];
        data[0] = {Id:0,Title:'',Name:'',Remark:'',Status:'1',Pid:0};
        if(!vac.isEmpty(Row)){
            data[0].Pid =Row.Id;
            $("#treegrid").treegrid("expand",Row.Id);//展开节点
            if($("#treegrid").treegrid("getLevel",Row.Id) >2){
                vac.alert("不允许添加");
                return false;
            }
        }
        //如果没有数据，则从0行开始新增
        $("#treegrid").treegrid("append",{
            parent: (Row?Row.Id:null),
            data:data
        });
        $("#treegrid").treegrid("select",0);//选中
        $("#treegrid").treegrid("beginEdit",0);//编辑输入
    }
    //编辑
    function editrow(){
        var row = $("#treegrid").treegrid("getSelected");

        if(!row){
            vac.alert("请选择要编辑的行");
            return;
        }
        $("#treegrid").treegrid("beginEdit",row.Id);
    }
    //保存
    function saverow(){
        var row = $("#treegrid").treegrid("getSelected");
         console.log(row);
        if(!row){
            vac.alert("请选择要保存的行");
            return;
        }
        $("#treegrid").treegrid("endEdit",row.Id);
    }
    //取消
    function cancelrow(){
        var row = $("#treegrid").treegrid("getSelected");
        if(!row){
            vac.alert("请选择要取消的行");
            return;
        }
        $("#treegrid").treegrid("cancelEdit",row.Id);
    }
    //删除
    function delrow(){
        $.messager.confirm('Confirm','你确定要删除?',function(r){
            if (r){
                var row = $("#treegrid").treegrid("getSelected");
                if(!row){
                    vac.alert("请选择要删除的行");
                    return;
                }
                vac.ajax(URL+'/DelMachine', {Id:row.Id}, 'POST', function(r){
                    if(!r.status){
                        vac.alert(r.info);
                    }else{
                        $("#treegrid").treegrid("reload");
                    }
                })
            }
        });
    }
    //刷新
    function reloadrow(){
        $("#treegrid").treegrid("reload");
    }

</script>
<body>
<table id="treegrid" title="节点管理" class="easyui-treegrid" toolbar="#tb"></table>
<div id="tb" style="padding:5px;height:auto">
    <a href="#" icon='icon-add' plain="true" onclick="addrow()" class="easyui-linkbutton" >新增</a>
    <a href="#" icon='icon-edit' plain="true" onclick="editrow()" class="easyui-linkbutton" >编辑</a>
    <a href="#" icon='icon-save' plain="true" onclick="saverow()" class="easyui-linkbutton" >保存</a>
    <a href="#" icon='icon-cancel' plain="true" onclick="cancelrow()" class="easyui-linkbutton" >取消</a>
    <a href="#" icon='icon-cancel' plain="true" onclick="delrow()" class="easyui-linkbutton" >删除</a>
    <a href="#" icon='icon-reload' plain="true" onclick="reloadrow()" class="easyui-linkbutton" >刷新</a>
    <a href="../../qrcode" target="_blank" icon='icon-reload' plain="true" class="easyui-linkbutton" >生成二维码</a>
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
</body>
</html>
