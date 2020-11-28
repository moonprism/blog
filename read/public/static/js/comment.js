var com_inputs = $('com_up').getElementsByTagName('input');
var com_text = $('com_up').getElementsByTagName('textarea')[0];
var to_name = $('to_name');
var com_list = $('com_list');

//表情
var exp_i = ["(泣)", "(汗)", "(笑)", "_(:3」∠)_", "（#-_-)┯━┯", "(╯°口°)╯(┴—┴", "ヽ(`Д´)ﾉ", "(￣ε(#￣) Σ", "（￣へ￣）", "(-_-#)", "(´･_･`)", "ε=ε=(ノ≧∇≦)ノ", "(●￣(ｴ)￣●)", "(｡･ω･｡)", "(^・ω・^ )", "（/TДT)/", "(´；ω；`)", "Σ( ￣□￣||)", "Σ(ﾟдﾟ;)", "(ﾟДﾟ≡ﾟдﾟ)!?", "(\"▔□▔)/", "(;¬_¬)", "(>_>)", "(<_<)", "→_→", "←_←", "( ´_ゝ｀)", "╮(￣▽￣)╭", "(￣3￣)", "(°∀°)ﾉ", "(･∀･)", "(〜￣△￣)〜", "(｀・ω・´)", "(=・ω・=)", "（￣▽￣）", "(｡・`ω´･)"];
// 表情绑定
var exp = $('exp');
var exp_p = $('exp_p');
var exx = function(){
    if ( exp_p.style.display == 'none' ){
        exp_p.style.display = 'block';
    } else {
        exp_p.style.display = 'none';
    }
};
// 按钮
f.addEve(exp,'click',exx);
// 表情生成
for (var i = exp_i.length - 1; i >= 0; i--) {
    f.addEve(f.addNode(exp_p, 'a', exp_i[i]), 'click',(function(i){
        return function(){
            com_text.value += exp_i[i];
            exx();
            com_text.focus();
        };
    })(i) );
}
// 显示评论列表
var more_a = $('more_a');
function comment_list(total){
    com_list.innerHTML = '';
    let loading = f.addNode(com_list, 'div', '', {class: 'loading'});
    // 加载缓存
    if (total === 10) {
        var m_url = '/json/comment/'+com_inputs[0].value;
        var m_type = "GET";
        var m_data = {};
    } else {
        more_a.style.display = 'none';
        var m_url = "/comment/list/"+com_inputs[0].value;
        var m_type = "GET";
        var m_data = {};
    }
    f.ajax({
        url: m_url,
        type: m_type,
        data: m_data,
        dataType: "json",
        success: function (response, xml) {
            if (response!=='[]') {
                var responsejson = JSON.parse(response);
                // 评论树最上层
                var index_tree = {};
                var last_node = null;
                // 保存下标到id的映射
                var index_id = {};
                for (var i = responsejson.length-1; i >= 0; i--) {
                    responsejson[i]['name'] = replace_sym(responsejson[i]['name']);
                    responsejson[i]['text'] = replace_sym(responsejson[i]['text']);
                    var src = '';
                    // 邮箱后缀与头像获取
                    src = 'https://cdn.v2ex.com/gravatar/'+responsejson[i]['email'];
                    var dateDiff = getDateDiff(responsejson[i]['created_time']);
                    let h = null;
                    if(responsejson[i]['to_id'] === 0){
                        // 一层
                        var com = document.createElement('div');
                        com.className = 'com';
                        com_list.insertBefore(com, last_node);
                        last_node = com;
                        f.addNode(com, 'img', '', {src:src});
                        f.addNode(com, 'a', responsejson[i]['name'], {href:'javascript:repl('+responsejson[i]['id']+',"'+responsejson[i]['name']+'")'});
                        f.addNode(com, 'span', dateDiff);
                        f.addNode(com, 'p', responsejson[i]['text']);
                        h = f.addNode(com, 'a', '', {href:'javascript:repl('+responsejson[i]['id']+',"'+responsejson[i]['name']+'")','class':'repl'});
                        index_tree[responsejson[i]['id']] = com;
                    } else {
                        // 二层
                        var in_node = index_tree[responsejson[i]['to_id']];
                        if (index_tree[responsejson[i]['to_id']] === undefined) {
                            // 三层
                            in_node = index_tree[responsejson[index_id[responsejson[i]['to_id']]]['to_id']];
                        }  
                        var com = f.addNode(in_node, 'div', '', {'class':'com'});
                        f.addNode(com, 'img', '', {src:src});
                        f.addNode(com, 'a', responsejson[i]['name']+' @ '+responsejson[index_id[responsejson[i]['to_id']]]['name'], {href:'javascript:repl('+responsejson[i]['id']+',"'+responsejson[i]['name']+'")'});
                        f.addNode(com, 'span', dateDiff);
                        f.addNode(com, 'p', responsejson[i]['text']);
                        h = f.addNode(com, 'a', '', {href:'javascript:repl('+responsejson[i]['id']+',"'+responsejson[i]['name']+'")','class':'repl'});
                        index_tree[responsejson[i]['id']] = in_node;
                    }
                    h.innerHTML = '<svg class="icon" aria-hidden="true"><use xlink:href="#icon-icon-test"></use></svg>';
                    index_id[responsejson[i]['id']] = i;
                }
                if(total === 10){
                    if(responsejson.length>=5)
                        more_a.style.display='block';
                }
                else
                    window.scrollTo(0, $('markdown').offsetHeight+350);
            }
            loading.parentNode.removeChild(loading);
        },
        fail: function (status) {
            
        }
    });
}
function getDateDiff(date){
    const commentData = new Date(date);
    const dateTimeStamp = commentData.getTime();
    const minute = 1000 * 60;
    const hour = minute * 60;
    const day = hour * 24;
    const month = day * 30;
    const now = new Date().getTime();
    const diffValue = now - dateTimeStamp;
    if(diffValue < 0){return;}
    const monthC = diffValue/month;
    const weekC = diffValue/(7*day);
    const dayC = diffValue/day;
    const hourC = diffValue/hour;
    const minC = diffValue/minute;
    const rtf = new Intl.RelativeTimeFormat('zh', {
        numeric: 'auto'
    });
    if (monthC>12) {
        return rtf.format(-parseInt(monthC/12), 'year');
    }else if(monthC>=1 && monthC<=12){
        return rtf.format(-parseInt(monthC), 'month');
    } else if(weekC>=1){
        return rtf.format(-parseInt(monthC), 'week');
    } else if(dayC>=1){
        return rtf.format(-parseInt(dayC), 'day');
    } else if(hourC>=1){
        return rtf.format(-parseInt(hourC), 'hour');
    } else if(minC>=1){
        return rtf.format(-parseInt(minC), 'minute');
    } else {
        return '刚刚';
    }
}
comment_list(10);
// 回复某人
function repl(to_id, name){
    window.scrollTo(0, $('markdown').offsetHeight+100);
    com_inputs[1].value = to_id;
    to_name.innerHTML = '';
    f.addNode(to_name, 'span', '@', {style: 'font-family:none;margin-right:3px'});
    f.addNode(to_name, 'span', name, {style: 'color: #233;'});
    com_inputs[4].value = '回复';
    to_name.innerHTML += '<a href="javascript:de_repl()"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-close"></use></svg></a>';
}
function de_repl(){
    com_inputs[1].value = "0";
    com_inputs[4].value = '留言';
    to_name.innerHTML = '';
}
f.addEve(com_inputs[4], 'click', function(){
    if (com_inputs[2].value.length > 10 || com_inputs[2].value.length === 0) {
        alert('昵称长度  (´・ω・)ﾉ');
    } else if(com_inputs[3].value.match(/^([a-zA-Z0-9_\.\-])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})+$/ ) === null){
        alert('邮箱格式  ( ・◇・)？');
    }else if(com_text.value === ''){
        alert('没什么想说的吗 ,,Ծ‸Ծ,,');
    } else {
        let comment_text = com_text.value;
        com_text.value = '';
        com_list.innerHTML = '';
        let loading = f.addNode(com_list, 'div', '', {class: 'loading'});
        $('more_a').style.display = 'none';
        f.ajax({
            url: "/comment/up",
            type: "POST",
            data: { art_id: com_inputs[0].value, to_id: com_inputs[1].value, name: com_inputs[2].value, email: com_inputs[3].value, text: comment_text},
            dataType: "json",
            success: function (response, xml) {
                com_inputs[1].value = '0';
                com_inputs[4].value = '留言';
                com_inputs[2].value = '';
                com_inputs[3].value = '';
                com_text.value = '';
                to_name.innerHTML = '';
                com_list.innerHTML = '';
                comment_list(10);
            },
            fail: function (status) {
                loading.parentNode.removeChild(loading);
                alert(status);
            }
        });
    }
});