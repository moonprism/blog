<template>
    <div class="markdown" v-html="this.html" ref="article"></div>
</template>

<script>
    import markdown from 'moonprism-markdown'
    process.env.VUE_APP_FILE_ORIGIN
    import '@/style/markdown.css'

    import hljs from 'highlight.js/lib/highlight'
    import javascript from 'highlight.js/lib/languages/javascript'
    import go from 'highlight.js/lib/languages/go'
    import php from 'highlight.js/lib/languages/php'
    import shell from 'highlight.js/lib/languages/shell'
    import json from 'highlight.js/lib/languages/json'
    hljs.registerLanguage('javascript', javascript)
    hljs.registerLanguage('go', go)
    hljs.registerLanguage('php', php)
    hljs.registerLanguage('shell', shell)
    hljs.registerLanguage('json', json)

    import 'highlight.js/styles/gruvbox-dark.css'

    import {insertStr} from "@/utils/string"

    export default {
        data() {
            return {
                kira: '☆',
            }
        },
        methods: {
            scrollIndex() {
                this.$refs.article.innerHTML = this.$refs.article.innerHTML
                    .replace(eval('/'+this.kira+'/'), '<a class="cursor" id="i"></a>')
                this.$nextTick(() => {
                    // 滚动到 #i 所在位置
                    let cursor = document.getElementById('i')
                    if (cursor) {
                        this.$refs.article.scrollTop = cursor.offsetTop - 300
                    }
                })
            },
            highlightCode() {
                this.$refs['article'].getElementsByTagName('pre').forEach((element) => {
                    let code = element.getElementsByTagName('code')[0]
                    let html = ''
                    if (code.className !== "" && ["php", "javascript", "go", "shell", "json"].includes(code.className)) {
                        let lang = code.className.split('-')[1]
                        html = hljs.highlight(lang, code.innerText).value
                    } else {
                        html = hljs.highlightAuto(code.innerText).value
                    }
                    code.innerHTML = html
                })
            },
            // insertCursorIndex 插入自定义的锚点
            insertCursorIndexForMarkdown() {
                let lineArray = this.$props.articleShow.markdownText.split('\n')
                lineArray[this.$props.articleShow.cursor.line] = insertStr(lineArray[this.$props.articleShow.cursor.line], this.$props.articleShow.cursor.ch+1, this.kira)
                return lineArray.join('\n')
            }
        },
        computed: {
            html() {
                return markdown(this.insertCursorIndexForMarkdown(), this.$props.articleShow.config)
            }
        },
        props: {
            articleShow: {
                markdownText: '',
                // 编辑文本对应的cursor位置
                cursor: {
                    line: 0,
                    ch: 0
                },
                config: {}
            }
        },
        mounted() {
            this.highlightCode()
            this.scrollIndex()
        },
        updated() {
            this.highlightCode()
            this.scrollIndex()
        }
    }
</script>

<style scoped>
    .markdown {
        width: 800px;
        margin: 0 auto;
        padding: 0 20px;
        height: 1000px;
        overflow-y: auto;
    }
</style>
