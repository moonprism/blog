<template>
    <div class="code-edit">
        <textarea
            style="display: none"
            ref="edit"
            v-model="$props.code.content">
        </textarea>
    </div>
</template>

<script>
    import CodeMirror from 'codemirror/lib/codemirror'
    import 'codemirror/lib/codemirror.css'
    import 'codemirror/mode/markdown/markdown'
    import 'codemirror/mode/go/go'
    import 'codemirror/mode/javascript/javascript'
    import 'codemirror/mode/shell/shell'
    import 'codemirror/mode/sql/sql'
    import 'codemirror/mode/css/css'
    import 'codemirror/mode/php/php'
    import 'codemirror/keymap/vim'
    import 'codemirror/addon/dialog/dialog.css'
    import 'codemirror/theme/monokai.css'

    export default {
        name: 'CodeEdit',
        data() {
            return {
                cmOptions: {
                    theme:  'monokai',
                    lineWrapping: true,
                    matchBrackets: true,
                    inputStyle: "contenteditable",
                    indentUnit: 4,
                    showCursorWhenSelecting: true,
                },
                editor: {}
            }
        },
        props: {
            code: {
                lang: '',
                content: ''
            },
            config: {
                useVim: false
            }
        },
        mounted() {
            this.init()
        },
        updated() {
            this.update()
        },
        methods: {
            init() {
                CodeMirror.Vim.map('jk', '<Esc>', 'insert')
                if (this.$props.config.useVim) {
                    this.cmOptions.keyMap = 'vim'
                }
                let editor = CodeMirror.fromTextArea(
                    this.$refs.edit, this.cmOptions
                )
                this.editor = editor
                if (this.$props.config.useVim) {
                    this.editor.focus()
                }
                this.updateLang()
            },
            update() {
                this.editor.setValue(this.$props.code.content)
                this.updateLang()
            },
            updateLang() {
                if (this.$props.code.lang === 'go') {
                    this.editor.setOption('indentWithTabs', true)
                } else {
                    this.editor.setOption('indentWithTabs', false)
                }
                let setLang = this.$props.code.lang
                switch (this.$props.code.lang) {
                    case 'md':
                        setLang = 'markdown'
                        break;
                    case 'js':
                        setLang = 'javascript'
                        break;
                }
                this.editor.setOption('mode', setLang)
            },
            getContent() {
                return this.editor.getValue()
            }
        }
    }
</script>

<style>
    .code-edit {
        line-height: 21px;
    }
    .code-edit .CodeMirror {
        padding: 5px 9px;
        height: auto;
    }
    .code-edit .CodeMirror {
        font-family: 'Space Mono', "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif !important;
    }
</style>
