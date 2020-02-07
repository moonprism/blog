let backupHTML = $('searchResult').innerHTML;
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
                            highlight();
                        }
                        if ($('searchInput').value != searchText) {
                            return
                        }
                        let result = JSON.parse(response);
                        searchResult[searchText] = result;
                        $('searchResult').innerHTML = '';
                        result.forEach(
                            (item) => {
                                render(item);
                            }
                        )
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
        highlight();
    }
})
function render(searchItem) {
    $('searchResult').innerHTML += '<div class="search-item">\n'+
        '<blockquote>'+searchItem.description+'<span>.'+searchItem.lang+'</span></blockquote>\n'+
        '<div class="tags">'+searchItem.tags+'</div>\n'+
        '<pre><code class="'+searchItem.lang+'">'+searchItem.content+'</code></pre>\n'+
        '</div>';
}
function addLineNumber() {
    let codeEl = document.getElementsByTagName('code');
    let codeList = Array.prototype.slice.call(codeEl);
    codeList.forEach((code) => {
        code.innerHTML = ('<ol><li>' + code.innerHTML.replace(/\n/g, '\n</li><li>') + '\n</li></ol>')
    })
}
function highlight() {
    document.querySelectorAll('pre code').forEach((block) => {
        hljs.highlightBlock(block);
    });
    addLineNumber();
}
hljs.initHighlightingOnLoad();
addLineNumber();