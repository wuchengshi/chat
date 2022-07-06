// url 接口地址 type请求的类型 data数据(不写的话 返回的请求数据)
function get(url,type='GET',data=""){
    // resolve 成功的参数 reject 失败的参数
    return new Promise((resolve,reject)=>{

        //创建XMLHttpRequest对象
        // XMLHttpRequest对象向服务器发异步请求，从服务器获得数据,
        let xhr = new XMLHttpRequest();
        // open()方法有3个参数：1、method 请求方式；2、url 接口地址；3、boolean true/false；
        xhr.open(type,url)
        // 发送请求
        xhr.send()

        // onreadyStateChange事件是在readyState属性发生改变时触发的
        xhr.onreadystatechange=function(){
            // readyState的值表示了当前请求的状态
            // readyState==4说明请求已经完成 status==200 状态码为200表示成功
            // readyState有五种可取值0：尚未初始化，1：正在加载，2：加载完毕，3：正在处理；4：处理完毕
            if (xhr.readyState == 4 ){
                if(xhr.status >= 200 && xhr.status < 300 ){
                    // 在成功的方法中 将请求的数据转换成数组
                    resolve(JSON.parse(xhr.response))
                }
            }
        }
    })
}

function post(uri,data,fn){
    // 创建一个Promise对象
    return new Promise((resolve, reject) =>{
        var xhr = new XMLHttpRequest();
        xhr.open("POST","//"+location.host+"/"+uri, true);
        // 添加http头，发送信息至服务器时内容编码类型
        xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
        xhr.onreadystatechange = function() {
            if (xhr.readyState == 4 ){
                if(xhr.status >= 200 && xhr.status < 300 ){
                    // 在成功的方法中 将请求的数据转换成数组
                    resolve(JSON.parse(xhr.response))
                }else{
                    console.error(xhr.status)
                    reject(xhr.status)
                }

            }
        };
        var _data=[];
        if(!!userId() && !data.userid){
            data["userid"] = userId();
        }
        for(var i in data){
            _data.push( i +"=" + encodeURI(data[i]));
        }
        xhr.send(_data.join("&"));
    })
}

function upload(dom){
    uploadfile("attach/upload",dom,function(res){
        if(res.code==0){
            app.imgval='';
            app.sendpicmsg(res.data,5)
        }

    })
}

function uploadfile(uri,dom,fn){
    var xhr = new XMLHttpRequest();
    xhr.open("POST","//"+location.host+"/"+uri, true);
    // 添加http头，发送信息至服务器时内容编码类型
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4 && (xhr.status == 200 || xhr.status == 304)) {
            fn.call(this, JSON.parse(xhr.responseText));
        }
    };
    var _data=[];
    var formdata = new FormData();
    if(!! userId()){
        formdata.append("userid",userId());
    }
    formdata.append("file",dom.files[0])
    xhr.send(formdata);
}
function uploadblob(uri,blob,filetype,fn){
    var xhr = new XMLHttpRequest();
    xhr.open("POST","//"+location.host+"/"+uri, true);
    // 添加http头，发送信息至服务器时内容编码类型
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4 && (xhr.status == 200 || xhr.status == 304)) {
            fn.call(this, JSON.parse(xhr.responseText));
        }
    };
    var _data=[];
    var formdata = new FormData();
    formdata.append("filetype",filetype);
    if(!! userId()){
        formdata.append("userid",userId());
    }
    formdata.append("file",blob)
    xhr.send(formdata);
}
function uploadaudio(uri,blob,fn){
    uploadblob(uri,blob,".mp3",fn)
}
function uploadvideo(uri,blob,fn){
    uploadblob(uri,blob,".mp4",fn)
}

function loadcommunitys(){
    var that = this;
    post("contact/loadcommunity",{userid:userId()},function(res){
        that.communitys = res.rows ||[];
    })
}