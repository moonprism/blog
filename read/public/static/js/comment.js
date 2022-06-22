var comment = $('comment')
if (comment) {
    let emoticons = ['("▔□▔)/']; //'(｡・`ω´･)', '(*゜ロ゜)'];
    comment.innerHTML = `
    <div class="comment">
        <span id="to_name"></span>
        <div id="com_up">
            <input type="hidden" name="art_id" value="" >
            <input type="hidden" name="to_id" value="0" >
            <label for="name">
                <svg class="icon" aria-hidden="true">
                    <use xlink:href="#icon-shiliangzhinengduixiang"></use>
                </svg>
                Name:
            </label>
            <input id="name" name="name" autocomplete="off" placeholder="your name." type="text" >
            <label for="email">
                <svg class="icon" aria-hidden="true">
                    <use xlink:href="#icon-Email"></use>
                </svg> 
                Email:
            </label>
            <input id="email" name="email" autocomplete="off" placeholder="Email" type="text" >
            <div class="exp-p" id="exp_p" style="display: none;"></div>
            <textarea name="content"></textarea>
            <button id="exp" class="exp">`+emoticons[Math.floor((Math.random()*emoticons.length))]+`</button>
            <input type="submit" class="submit" value="留言" id="re">
            <div id="repl_line" class='repl_line'>
                <a onclick="cancel_repl()">
                    <svg class="icon" aria-hidden="true">
                        <use xlink:href="#icon-close"></use>
                    </svg>
                </a>
            </div>
        </div>
        <div class="com-list" id="com_list"></div>
        <a class="more_a" id="more_a">all</a>
    </div>
    `;

    var commentInputs = $('com_up').getElementsByTagName('input');
    commentInputs[0].value = comment.dataset.id; //??
    var commentText = $('com_up').getElementsByTagName('textarea')[0];
    var commentToName = $('to_name');
    var commentList = $('com_list');

    //表情
    var emoticonList = ["(泣)", "(汗)", "(笑)", "∑(O_O；)", "(T_T)", "(๑❛ᴗ❛๑)", "^_^", "(๑＞ ＜)☆", "(*゜ロ゜)ノ ", "(・ω< )★", "_(:3」∠)_", "(｡･ω･｡)", "(^・ω・^ )", "( ˃᷄˶˶̫˶˂᷅ )", "(๑Ő௰Ő๑)", "(ﾟДﾟ≡ﾟдﾟ)!?", "(\"▔□▔)/", "(๑˙ー˙๑)", "（︶︿︶）", "✧( ु•⌄• )◞◟( •⌄• ू )✧", "ฅ^•ﻌ•^ฅ", "╮(｡>口<｡)╭", "<(｀^´)>", "╮(￣▽￣)╭", "๑乛◡乛๑", "(⁄ ⁄•⁄ω⁄•⁄ ⁄)", "Ծ‸Ծ", "(๑¯◡¯๑)", "(｀・ω・´)", "(=・ω・=)", "（￣▽￣）", "(｡・`ω´･)"];
    var emoticonNode = $('exp_p');
    emoticonNode.onclick = $('exp').onclick = () => {
        if ( emoticonNode.style.display == 'none' ){
            emoticonNode.style.display = 'block';
        } else {
            emoticonNode.style.display = 'none';
        }
    }

    for (var i = emoticonList.length - 1; i >= 0; i--) {
        f.addEve(f.addNode(exp_p, 'a', emoticonList[i]), 'click',(function(i){
            return function(){
                commentText.value += emoticonList[i];
                commentText.focus();
            };
        })(i) );
    }
    // 显示评论列表
    var currentBeferId = 0;
    var currentPageSize = 3;
    var more_a = $('more_a');
    function comment_list(){
        // 啊这...
        commentInputs[1].value = '0';
        commentInputs[4].value = '留言';
        cancel_repl();

        more_a.style.display='none';
        let loading = f.addNode(commentList, 'div', '', {class: 'loading'});
        // 加载缓存
        var m_url = '/comment/'+commentInputs[0].value+'?before_id='+currentBeferId+'&page_size='+currentPageSize;
        var m_type = "GET";
        var m_data = {};
        f.fetch(m_url, {
            method: m_type,
            data: m_data,
            success: function (response) {
                response.data.comments.forEach(co => {
                    var com = document.createElement('div');
                    com.className = 'com';
                    commentList.insertBefore(com, null);
                    f.addNode(com, 'img', '', {src:'https://gravatar.cat.net/avatar/'+co['email']});
                    f.addNode(com, 'a', co['name'], {
                        'class': 'repl_name',
                        onclick: 'repl(this)',
                        'data-id': co['id'],
                    });
                    f.addNode(com, 'span', getDateDiff(co['created_time']));
                    f.addNode(com, 'p', co['text']);
                    let h = f.addNode(com, 'a', '', {
                        onclick: 'repl(this)',
                        'data-id': co['id'],
                        'data-name': co['name'],
                        'class': 'repl',
                    })
                    h.innerHTML = '<svg class="icon" aria-hidden="true"><use xlink:href="#icon-icon-test"></use></svg>';
                    var indexX = {}
                    indexX[co['id']] = co
                    co.sub_comments.forEach(subCo => {
                        indexX[subCo['id']] = subCo
                        let subCom = f.addNode(com, 'div', '', {'class':'com'});
                        f.addNode(subCom, 'img', '', {src:'https://gravatar.cat.net/avatar/'+subCo['email']});
                        f.addNode(subCom, 'a', subCo['name']+'@'+indexX[subCo['to_id']]['name'], {
                            'class': 'repl_name',
                            onclick: 'repl(this)',
                            'data-id': subCo['id'],
                        });
                        f.addNode(subCom, 'span', getDateDiff(subCo['created_time']));
                        f.addNode(subCom, 'p', subCo['text']);
                        let h = f.addNode(subCom, 'a', '', {
                            onclick: 'repl(this)',
                            'data-id': subCo['id'],
                            'data-name': subCo['name'],
                            'class': 'repl'
                        })
                        h.innerHTML = '<svg class="icon" aria-hidden="true"><use xlink:href="#icon-icon-test"></use></svg>';
                    })
                    currentBeferId = co['id']
                })
                if (response.data.comments.length == currentPageSize) {
                    more_a.style.display='block';
                } else {
                    if (currentPageSize != 2) {
                        commentList.style.paddingBottom = '25px'
                    }
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
            return rtf.format(-parseInt(weekC), 'week');
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
    comment_list();
    more_a.onclick = () => {
        currentPageSize += 1
        comment_list()
    }

    // 回复某人
    function repl(el){
        let data_id = el.getAttribute('data-id');

        $('repl_line').style.visibility = 'visible';

        let names = document.getElementsByClassName('repl_name');
        let name_el
        [].slice.call(names).forEach(nel => {
            if (nel.getAttribute('data-id') == data_id) {
                // 用留言的用户名位置来定位准确点
                name_el = nel;
                name_el.style.color = '#ff2071';
            } else {
                nel.style.color = '';
            }
        });
        // let index = [].slice.call(repls).findIndex(repl => repl.getAttribute('data-id') == data_id);

        let re_el = $('re');
        re_el.style['border-color'] = '#e2e2e2';

        let pos = re_el.offsetLeft + re_el.clientWidth / 2;
        $('repl_line').style.left = pos+'px';

        let height = name_el.offsetTop - re_el.offsetTop - 28;
        $('repl_line').style.height = height+'px';

        window.scrollTo(0, document.querySelector('.markdown').offsetHeight+100);
        commentInputs[1].value = data_id;
        // commentToName.innerHTML = '';
        // f.addNode(commentToName, 'span', '@', {style: 'font-family:none;margin-right:3px'});
        // f.addNode(commentToName, 'span', name, {class: 'c-name'});
        commentInputs[4].value = '回复';
        // commentToName.innerHTML += '<a href="javascript:de_repl()"><svg class="icon" aria-hidden="true"><use xlink:href="#icon-close"></use></svg></a>';
    }
    function cancel_repl(){
        $('repl_line').style.visibility = 'hidden';
        $('repl_line').style.height = '0px';
        commentInputs[1].value = "0";
        commentInputs[4].value = '留言';
        let names = document.getElementsByClassName('repl_name');
        [].slice.call(names).forEach(nel => {
            nel.style.color = '';
        });
        let re_el = $('re');
        re_el.style['border-color'] = '';
        // commentToName.innerHTML = '';
    }
    f.addEve(commentInputs[4], 'click', function(){
        let comment_text = commentText.value;
        commentList.innerHTML = '';
        let loading = f.addNode(commentList, 'div', '', {class: 'loading'});
        $('more_a').style.display = 'none';
        let to_id = commentInputs[1].value;
        cancel_repl();
        f.fetch('/comment/'+commentInputs[0].value, {
            method: "POST",
            data: { art_id: commentInputs[0].value, to_id, name: commentInputs[2].value, email: commentInputs[3].value, text: comment_text},
            success: function (response) {
                if (response.code !== 200) {
                    alert(response.message)
                } else {
                    commentInputs[1].value = '0';
                    commentInputs[2].value = '';
                    commentInputs[3].value = '';
                    commentText.value = '';
                    // commentToName.innerHTML = '';
                }
                commentList.innerHTML = '';
                currentPageSize = 2;
                currentBeferId = 0;
                comment_list(1);
            },
            fail: function (status) {
                loading.parentNode.removeChild(loading);
                alert(status);
            }
        });
    });
}
