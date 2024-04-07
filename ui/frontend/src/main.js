import './css/main.css'
import './css/markdown.css'

// markdown
var main_markdown_config = {
    linkTargetBlank: true,
    debug: false,
    imageCDN: 'https://kicoe-blog.oss-cn-shanghai.aliyuncs.com/'
}
function markd(md) {
    let text = md.replace(/\nxxx\n([\s\S]*?)\nxxx/g, '') // 自定义隐藏
    return markdown(text, main_markdown_config)
}
