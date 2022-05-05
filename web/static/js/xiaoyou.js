// 自己封装的一些工具类
// 自己封装的函数
const xy = {
    net: {
        server: '',
        request(url, method, data,head = []) {
            return new Promise((resolve, reject) => {
                $.ajax(this.server + url, {
                    type: method,
                    headers: head,
                    contentType: "application/json; charset=utf-8",
                    data: method === 'GET' ? data : JSON.stringify(data),
                    success(data) {
                        resolve(data)
                    },
                    error(err) {
                        reject(err.responseJSON.message)
                    }
                })
            })
        }
    },
    window: {
        getScrollTop(){let scrollTop=0;if(document.documentElement&&document.documentElement.scrollTop){scrollTop=document.documentElement.scrollTop}else if(document.body){scrollTop=document.body.scrollTop}return scrollTop},
        getScrollHeight() {return Math.max(document.body.scrollHeight, document.documentElement.scrollHeight)},
        getClientHeight(){let clientHeight=0;if(document.body.clientHeight&&document.documentElement.clientHeight){clientHeight=Math.min(document.body.clientHeight,document.documentElement.clientHeight)}else{clientHeight=Math.max(document.body.clientHeight,document.documentElement.clientHeight)}return clientHeight},
        onScroll(event){window.onscroll=()=>{if(this.getScrollTop()+this.getClientHeight()===this.getScrollHeight()){event()}}},
        destory(){window.onscroll=null}
    },
    tools: {
        // cookie域名
        domain: '*',
        // 更新参数
        updateQueryStringParameter(uri,key,value){if(!value){return uri}const re=new RegExp("([?&])"+key+"=.*?(&|$)","i");const separator=uri.indexOf('?')!==-1?"&":"?";if(uri.match(re)){return uri.replace(re,'$1'+key+"="+value+'$2')}else{return uri+separator+key+"="+value}},
        // 获取cookie数据
        // 参考 https://developer.mozilla.org/zh-CN/docs/Web/API/Document/cookie
        getCookie:function(sKey){return decodeURIComponent(document.cookie.replace(new RegExp("(?:(?:^|.*;)\\s*"+encodeURIComponent(sKey).replace(/[-.+*]/g,"\\$&")+"\\s*\\=\\s*([^;]*).*$)|^.*$"),"$1"))||null},
        // 设置cookie
        setCookie:function(sKey,sValue,exDays,sPath,sDomain,bSecure){if(!sKey||/^(?:expires|max\-age|path|domain|secure)$/i.test(sKey)){return false}let d=new Date();d.setTime(d.getTime()+(exDays*24*60*60*1000));let sExpires='';if(exDays!==0){sExpires="; expires="+d.toUTCString()}document.cookie=encodeURIComponent(sKey)+"="+encodeURIComponent(sValue)+sExpires+(sDomain?"; domain="+sDomain:"")+(sPath?"; path="+sPath:"")+(bSecure?"; secure":"");return true},
        // 删除cookie
        removeCookie:function(sKey,sPath,sDomain){if(!sKey||!this.hasCookie(sKey)){return false}document.cookie=encodeURIComponent(sKey)+"=; expires=Thu, 01 Jan 1970 00:00:00 GMT"+(sDomain?"; domain="+sDomain:"")+(sPath?"; path="+sPath:"");return true},
        // cookie是否存在
        hasCookie:function(sKey){return(new RegExp("(?:^|;\\s*)"+encodeURIComponent(sKey).replace(/[-.+*]/g,"\\$&")+"\\s*\\=")).test(document.cookie)},
        // 获取所有的cookie
        getAllCookie:function(){const aKeys=document.cookie.replace(/((?:^|\s*;)[^\=]+)(?=;|$)|^\s*|\s*(?:\=[^;]*)?(?:\1|$)/g,"").split(/\s*(?:\=[^;]*)?;\s*/);for(let nIdx=0;nIdx<aKeys.length;nIdx++){aKeys[nIdx]=decodeURIComponent(aKeys[nIdx])}return aKeys},
        // 获取URL参数
        getQueryVariable:function(variable){const query=window.location.search.substring(1);const vars=query.split("&");for(let i=0;i<vars.length;i++){const pair=vars[i].split("=");if(pair[0]===variable){return pair[1]}}return false}
    },
    validate:{
        // 判断邮箱格式是否正确
        checkEmail(email){const re=/^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/;return re.test(email)}
    }
}