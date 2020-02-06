function replace_sym(re_str) {
    return re_str.replace(/\\"/g, '"').replace(/\\'/g, "'").replace(/\\\\/g, '\\');
}
var $ = function(id){return document.getElementById(id);};
var up_but = $('up');
var up_index = 50;
var main = $('content');
up_but.style.display = 'none';
function up(){
    var scrollTop = document.documentElement.scrollTop||document.body.scrollTop;
    var sup = setInterval(function(){
        up_index += 3;
        scrollTop -= up_index;
        if (scrollTop>0) {
            window.scrollTo(0, scrollTop);
        } else {
            window.scrollTo(0, 0);
            up_index = 50;
            clearInterval(sup);
        }   
    },20);
}
var timer = null;
window.onscroll = function(){
    if (timer) {
        clearTimeout(timer)
    }
    timer = setTimeout(function(){
        if (document.documentElement.scrollTop||document.body.scrollTop >500 && up_but.style.display=='none') {
            up_but.style.display='block';
        } else if(document.documentElement.scrollTop||document.body.scrollTop <500 && up_but.style.display=='block'){
            up_but.style.display='none';
        }
    },80);
}
var f = {};
// 绑定事件
f.addEve = function(node, even, fun){
    if(window.addEventListener){
        node.addEventListener(even,function(){
            fun();
        }, false);
    } else { // IE
        node.attachEvent('on'+even, function(){
            fun();
        });
    }
}
// ajax获取
f.ajax = function(options) {
    options = options || {};
    options.type = (options.type || "GET").toUpperCase();
    options.dataType = options.dataType || "json";
    var arr = [];
    for (var name in options.data) {
        arr.push(encodeURIComponent(name) + "=" + encodeURIComponent(options.data[name]));
    }
    arr.push(("v=" + Math.random()).replace(".",""));
    var params = arr.join("&");
    if (window.XMLHttpRequest) {
        var xhr = new XMLHttpRequest();
    } else {
        var xhr = new ActiveXObject('Microsoft.XMLHTTP');
    }
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) {
            var status = xhr.status;
            if (status >= 200 && status < 300) {
                options.success && options.success(xhr.responseText, xhr.responseXML);
            } else {
                options.fail && options.fail(status);
            }
        }
    }
    if (options.type == "GET") {
        xhr.open("GET", options.url + "?" + params, true);
        xhr.send(null);
    } else if (options.type == "POST") {
        xhr.open("POST", options.url, true);
        xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
        xhr.send(params);
    }
}
// 添加节点与属性
f.addNode = function(parent_n, child_s, string, attr){
    var child_n = document.createElement(child_s);
    if (attr) {
        for (var i in attr) {
            child_n.setAttribute(i, attr[i]);
        }
    }
    if (string !== ''){
        child_n.innerText = string;
    }
    parent_n.appendChild(child_n);
    return child_n;
}