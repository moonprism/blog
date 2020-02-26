var backupHTML = $('searchResult').innerHTML;
var searchResult = {};
f.addEve($('searchInput'), 'input', () => {
    let searchText = $('searchInput').value;
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
        } else {
            setTimeout(function () {
                f.ajax({
                    url: "/code/search/" + searchText,
                    type: "GET",
                    dataType: "json",
                    success: function (response) {
                        if ($('searchInput').value == '') {
                            $('searchResult').innerHTML = backupHTML;
                            hoverMarkdown();
                            highlight();
                            return;
                        }
                        if ($('searchInput').value != searchText) {
                            return;
                        }
                        let result = JSON.parse(response);
                        searchResult[searchText] = result;
                        $('searchResult').innerHTML = '';
                        result.forEach(
                            (item) => {
                                render(item);
                            }
                        )
                        hoverMarkdown();
                        highlight();
                    },
                    fail: function (status) {
                        $('searchResult').innerHTML = '<div style="text-align:center;margin-top:35px">500</div>';
                    }
                });
            }, 50)
        }
    } else  {
        // 清空处理
        $('searchResult').innerHTML = backupHTML;
        hoverMarkdown()
        highlight();
    }
})
function render(searchItem) {
    $('searchResult').innerHTML += '<div class="search-item">\n'+
        '<blockquote>'+searchItem.description+'<span>.'+searchItem.lang+'</span></blockquote>\n'+
        '<div class="tags">'+searchItem.tags+'</div>\n'+
        '<pre class="c-'+searchItem.lang+'"><code class="'+searchItem.lang+'">'+searchItem.content+'</code><div class="markdown"></div></pre>\n'+
        '</div>';
}
function addLineNumber() {
    document.querySelectorAll('pre code').forEach((code) => {
        console.log(code.classList)
        if (!code.classList.contains('_c'))
            code.innerHTML = ('<ol><li><span class="cl">' + code.innerHTML.replace(/\n/g, '\n</span></li><li><span class="cl">') + '\n</span></li></ol>')
    })
}
function highlight() {
    document.querySelectorAll('pre code').forEach((block) => {
        hljs.highlightBlock(block);
    });
    addLineNumber();
}
function hoverMarkdown() {
    let preEl = document.getElementsByTagName('pre');
    let preList = Array.prototype.slice.call(preEl);
    preList.forEach((pre) => {
        if (pre.className != 'c-md') {
            return;
        }
        let code = pre.getElementsByTagName('code')[0];
        let show = pre.getElementsByTagName('div')[0];
        show.innerHTML = markdown(code.innerText);
        show.style.display = 'block';
        code.style.display = 'none';
        f.addEve(pre, 'click', (e) => {
            if (code.style.display != 'none') {
                show.style.display = 'block';
                code.style.display = 'none';
            } else {
                show.style.display = 'none';
                code.style.display = 'block';
            }
        })
    })
}
hoverMarkdown()
hljs.initHighlightingOnLoad();
addLineNumber();