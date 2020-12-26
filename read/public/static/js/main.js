var emoji_parse = function (str) {
    return str.replace(/\:bread\:/g, '🍞')
        .replace(/\:heart\:/g, '❤️')
        .replace(/\:sparkling_heart\:/g, '💖')
        .replace(/\:zap\:/g, '⚡️')
        .replace(/\:snowflake\:/g, '❄️')
        .replace(/\:books\:/g, '📚')
        .replace(/\:bookmark\:/g, '🔖')
        .replace(/\:dart\:/g, '🎯')
        .replace(/\:fish_cake\:/g, '🍥')
        .replace(/\:lollipop\:/g, '🍭')
        .replace(/\:ice_cream\:/g, '🍨')
        .replace(/\:star\:/g, '⭐️')
        .replace(/\:speech_balloon\:/g, '💬')
        .replace(/\:cloud\:/g, '☁️')
        .replace(/\:chestnut\:/g, '🌰')
        .replace(/\:jack_o_lantern\:/g, '🎃')
        .replace(/\:email\:/g, '✉️')
        .replace(/\:anchor\:/g, '⚓️')
        .replace(/\:triangular_flag_on_post\:/g, '🚩')
        .replace(/\:link\:/g, '🔗')
        .replace(/\:whale\:/g, '🐳')
        .replace(/\:tada\:/g, '🎉')
        .replace(/\:cake\:/g, '🍰')
        .replace(/\:art\:/g, '🎨')
        .replace(/\:book\:/g, '📖')
        .replace(/\:game\:/g, '🎮')
        .replace(/\:pushpin\:/g, '📌')
        .replace(/\:baby_chick\:/g, '🐤')
        .replace(/\:sparkles\:/g, '✨')
        .replace(/\:rocket\:/g, '🚀')
        .replace(/\:rabbit\:/g, '🐰')
        .replace(/\:cherry\:/g, '🌸')
        .replace(/\:pill\:/g, '💊')
        .replace(/\:watermelon\:/g, '🍉')
        .replace(/\:trophy\:/g, '🏆')
        .replace(/\:seedling\:/g, '🌱')
        .replace(/\:maple_leaf\:/g, '🍁')
        .replace(/\:evergreen_tree\:/g, '🌲')
        .replace(/\:octocat\:/g, '<img title=":octocat:" alt=":octocat:" src="https://github.githubassets.com/images/icons/emoji/octocat.png" height="20" width="20" align="absmiddle">')
        .replace(/\:cherries\:/g, '🍒');
}
var main_markdown_config = {
    inlineParse: emoji_parse,
    codeParse: emoji_parse,
    linkTargetBlank: true,
    debug: false,
    imageCDN: 'https://kicoe-blog.oss-cn-shanghai.aliyuncs.com/'
};
function replace_sym(re_str) {
    return re_str.replace(/\\"/g, '"').replace(/\\'/g, "'").replace(/\\\\/g, '\\');
}
var $ = function(id){return document.getElementById(id);};
var up_but = $('up');
var up_index = 50;
var main = $('content');
up_but.style.display = 'none';
function up(){
    window.scrollTo(0, 0);
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
        node.addEventListener(even,function(e){
            fun(e);
        }, false);
    } else { // IE
        node.attachEvent('on'+even, function(e){
            fun(e);
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
        xhr.open("GET", options.url, true);
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

function bg() {
    let imgEl = document.getElementsByTagName('img');
    let imgList = Array.prototype.slice.call(imgEl);
    f.addEve($('bg'), 'click', (i) => {
        $('bg').style.display = 'none';
    })
    imgList.forEach((img) => {
        img.className = 'cl-bg';
        f.addEve(img, 'click', (i) => {
            let wh = document.documentElement.clientHeight || document.body.clientHeight;
            let h = (wh - i.srcElement.naturalHeight) / 2.1;
            let paddingHeigh = h > 0 ? h : 15;
            $('bg').style.display = 'block';
            $('bg').innerHTML = '<img style="margin: '+paddingHeigh+'px auto" src="'+ i.srcElement.currentSrc +'">';
        })
    })
}
