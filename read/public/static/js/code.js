var searchResult = $('searchResult');
if (searchResult) {
    var backupHTML = searchResult.innerHTML;
    var searchResult = {};
    f.addEve($('searchInput'), 'input', () => {
        let searchText = $('searchInput').value;
        window.history.pushState({}, 0, "#!/"+searchText);
        $('searchResult').innerHTML = '';
        if (searchText) {
            if (searchResult[searchText]) {
                $('searchResult').innerHTML = '';
                searchResult[searchText].forEach(
                    (item) => {
                        render(item);
                    }
                )
                hoverMarkdown();
                highlight();
                setTimeout(function () {$('load').innerHTML = '';}, 100)
            } else {
                if ($('load').innerHTML === '') {
                    f.addNode($('load'), 'div', '', {class: 'loading'});
                }
                setTimeout(function () {
                    if ($('searchInput').value != searchText) {
                        return;
                    }

                    f.fetch("/api/code?text=" + searchText, {
                        method: "GET",
                        success: function (response) {
                            if ($('searchInput').value == '') {
                                $('searchResult').innerHTML = backupHTML;
                                hoverMarkdown(true);
                                highlight();
                                return;
                            }
                            searchResult[searchText] = response;
                            $('searchResult').innerHTML = '';
                            setTimeout(function () {$('load').innerHTML = '';}, 100)
                            response.forEach(
                                (item) => {
                                    render(item);
                                }
                            )
                            hoverMarkdown();
                            highlight();
                        },
                        fail: function (status) {
                            $('searchResult').innerHTML = '<div style="text-align:center;margin-top:35px">500</div>';
                            $('load').innerHTML = '';
                            setTimeout(function () {$('load').innerHTML = '';}, 100)
                        }
                    });
                }, 500)
            }
        } else  {
            // 清空处理
            $('searchResult').innerHTML = backupHTML;
            window.history.pushState({}, 0, "/code/");
            hoverMarkdown(true)
            highlight();
            setTimeout(function () {$('load').innerHTML = '';}, 100)
        }
    })
    function render(searchItem) {
        let content = searchItem.content;
        let html = '';
        if (searchItem.lang === 'md') {
            content = searchItem.content.replaceAll('<em>', '').replaceAll('</em>', '');
            html = markd(searchItem.content);
        }
        var real_lang = searchItem.lang.replaceAll('<em>', '').replaceAll('</em>', '');
        $('searchResult').innerHTML += '<div class="search-item">\n'+
            '<blockquote><svg class="icon icon-code" aria-hidden="true"><use xlink:href="#icon-'+real_lang+'"></use></svg> '+
            searchItem.description+'<span>.'+searchItem.lang+'</span></blockquote>\n'+
            (searchItem.tags?('<div class="tags">'+searchItem.tags+'</div>\n'):'')+
            '<pre class="c-'+searchItem.lang+'"><code class="'+searchItem.lang+'">'+content+
            '</code><div class="m-cc">'+html+'</div><textarea style="display:none">'+content+'</textarea></pre>\n'+
            '</div>';
    }
    function addLineNumber() {
        document.querySelectorAll('.search-item>pre>code').forEach((code) => {
            code.innerHTML = ('<ol><li><span class="cl">' + 
                code.innerHTML.replace(/\n/g, '\n</span></li><li><span class="cl">') + '\n</span></li></ol>')
        })
    }
    function highlight() {
        document.querySelectorAll('pre code').forEach((block) => {
            if (typeof(hljs) != "undefined") {
                hljs.highlightBlock(block);
            }
            block.innerHTML = block.innerHTML.replace(/\<span class=\"hljs-comment\">(.+\s)(https?:\/\/[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|])\<\/span\>/, '<span class="hljs-comment">$1<a href="$2">$2</a></span>')
        });
            addLineNumber();
        }
    function hoverMarkdown(flash = false) {
        let preEl = document.getElementsByTagName('pre');
        let preList = Array.prototype.slice.call(preEl);
        preList.forEach((pre) => {
            if (pre.className != 'c-md') {
                return;
            }
            let code = pre.getElementsByTagName('code')[0];
            let show = pre.getElementsByTagName('div')[0];
            let text = pre.getElementsByTagName('textarea')[0];
            if (flash) show.innerHTML = markd(code.innerText);
            show.style.display = 'block';
            code.style.display = 'none';
            f.addEve(pre, 'click', (e) => {
                if (text.value) {
                    show.innerHTML = markd(text.value);
                    document.querySelectorAll('pre code').forEach((block) => {
                        hljs.highlightBlock(block);
                        block.innerHTML = block.innerHTML.replace(/\<span class=\"hljs-comment\">(.+\s)(https?:\/\/[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|])\<\/span\>/, '<span class="hljs-comment">$1<a href="$2">$2</a></span>')
                    });
                }
                // if (code.style.display != 'none') {
                //     show.style.display = 'block';
                //     code.style.display = 'none';
                // } else {
                //     show.style.display = 'none';
                //     code.style.display = 'block';
                // }
            })
        })
    }
    hoverMarkdown(true);
    highlight();

    let lo = window.location.href.split("#!/");
    if (lo.length > 1) {
        $('searchInput').value = decodeURI(lo[1]);
        let event = new InputEvent('input');
        $('searchInput').dispatchEvent(event);
    }
}
