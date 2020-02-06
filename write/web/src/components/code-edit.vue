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
    import 'codemirror/theme/3024-day.css'

    export default {
        name: 'CodeEdit',
        data() {
            return {
                cmOptions: {
                    theme:  '3024-day',
                    lineWrapping: true,
                    matchBrackets: true,
                    inputStyle: "contenteditable",
                    indentUnit: 4,
                    showCursorWhenSelecting: true,
                    lineNumbers: true,
                },
                editor: {}
            }
        },
        props: {
            code: {
                lang: '',
                content: ''
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
                let editor = CodeMirror.fromTextArea(
                    this.$refs.edit, this.cmOptions
                )
                this.editor = editor
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
                this.editor.setOption('mode', this.$props.code.lang)
            },
            getContent() {
                return this.editor.getValue()
            }
        }
    }
</script>

<style>
    .code-edit {
        line-height: 19px;
    }
    .code-edit .CodeMirror {
        padding: 8px 0;
    }
    .code-edit .CodeMirror {
        font-family: 'Roboto Mono', "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif !important;
    }

</style>
