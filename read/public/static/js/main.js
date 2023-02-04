// markdown
var main_markdown_config = {
    linkTargetBlank: true,
    debug: false,
    imageCDN: 'https://kicoe-blog.oss-cn-shanghai.aliyuncs.com/'
}
var avatar_cdn = 'https://dn-qiniu-avatar.qbox.me/avatar/'
function markd(md) {
    let text = md.replace(/\nxxx\n([\s\S]*?)\nxxx/g, '') // 自定义隐藏
    return markdown(text, main_markdown_config).replace(/\:bread\:/g, '🍞')
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
        .replace(/\:cherries\:/g, '🍒')
}

document.querySelectorAll('.markdown > .md').forEach((md) => {
    md.parentNode.innerHTML = markd(md.innerText)
})

document.querySelectorAll('.markdown > .fav').forEach((fav) => {
    let favP = fav.firstChild
    favP.dataset.content = "▸"
    let favL = fav.lastChild
    favP.onclick = () => {
        if (favP.dataset.content == "▸") {
            fav.style.maxHeight = favL.clientHeight+40+'px'
            favP.setAttribute('data-content', "▾");
        } else {
            fav.style.maxHeight = '30px'
            favP.setAttribute('data-content', "▸");
        }
    }
})

var $ = (id) => document.getElementById(id)

// 回到顶部 #up
var upButton = $('up')
if (upButton) {
    upButton.innerHTML = '<a><b>^</b></a>';
    upButton.style.display = 'none';
    var timer = null;
    window.onscroll = () => {
        if (timer) {
            clearTimeout(timer)
        }
        timer = setTimeout(function(){
            if (document.documentElement.scrollTop||document.body.scrollTop > 500 && upButton.style.display == 'none') {
                upButton.style.display='block';
            } else if(document.documentElement.scrollTop||document.body.scrollTop < 500 && upButton.style.display == 'block'){
                upButton.style.display='none';
            }
        },80);
    }
    upButton.onclick = () => {
        window.scrollTo(0, 0)
    }
}

// 预览图片 #curtain .preview
var curtain = $('curtain')
if (curtain) {
    curtain.onclick = () => {
        curtain.style.display = 'none'
    }
    document.querySelectorAll('img').forEach((p) => {
        p.style.cursor = 'pointer'
        p.onclick = (i) => {
            let wh = document.documentElement.clientHeight || document.body.clientHeight
            let h = (wh - i.srcElement.naturalHeight) / 2.1
            let paddingHeigh = h > 0 ? h : 15
            $('curtain').style.display = 'block'
            $('curtain').innerHTML = '<img style="margin: '+paddingHeigh+'px auto" src="'+ i.srcElement.currentSrc +'">'
        }
    })
}

// six years ago
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

// 请求(重写ajax)
f.fetch = async (url, options) => {
    let response
    if (options.method == 'GET') {
        response = await fetch(url)
    } else {
        let params = []
        for (var i in options.data) {
            params.push(i+"="+encodeURIComponent(options.data[i]))
        }
        response = await fetch(url, {
            method: 'POST',
            headers: {
                "Content-type": "application/x-www-form-urlencoded; charset=UTF-8",
            },
            body: params.join('&')
        })
    }
    if (response.ok) {
        options.success(await response.json())
    } else {
        options.fail(response.status)
    }
}
