<script lang="ts">
  import { page } from '$app/stores'
  import type { ArticleDetail } from '$src/types/stream'
  import { fet, isRequestIn } from '@/helpers/fetch'

  import { Carta, MarkdownEditor, type Icon, type Plugin } from 'carta-md'
  import './(styles)/editor_custom.css'

  // 斜杠命令
  import { slash } from '@cartamd/plugin-slash'
  import '@cartamd/plugin-slash/default.css'
  // 代码高亮
  import { code } from '@cartamd/plugin-code'
  import '@cartamd/plugin-code/default.css'
  // 冒号命令
  import { emoji } from '@cartamd/plugin-emoji'
  import '@cartamd/plugin-emoji/default.css'
  // 引用markdown-body样式
  import 'moonprism-blog-frontend/src/css/markdown.css'

  // @ts-ignore
  import imgLinks from '@pondorasti/remark-img-links'

  import SaveIcon from './(components)/save-icon.svelte'
  import toast from '$lib/helpers/toast'
  import { onMount } from 'svelte'

  const id = Number($page.params.slug)
  let article = {} as ArticleDetail

  const icon: Icon = {
    id: 'save',
    action: async () => {
      if ($isRequestIn) {
        return
      }
      const html = await carta.render(value)
      const res = await fet.put(`article/${id}`, { text: value, html })
      if (res.ok) {
        toast.success('保存成功，使用 <ESC> 退出编辑')
        originValue = value
        esc()
      }
    },
    component: SaveIcon
  }

  // issue: 这个编辑器 Esc 不退出编辑模式
  function esc() {
    // @ts-ignore
    document.activeElement?.blur()
  }

  function handleGlobalKeyDown(event: KeyboardEvent): void {
    if (event.key === 'Escape') {
      if (document.activeElement?.tagName === 'TEXTAREA') {
        esc()
      } else if (value === originValue) {
        window.history.back()
      } else {
        toast.success('请使用浏览器按钮强制返回')
      }
    }
  }

  onMount(() => {
    window.addEventListener('keydown', handleGlobalKeyDown)

    return () => {
      window.removeEventListener('keydown', handleGlobalKeyDown)
    }
  })

  const ext: Plugin = {
    icons: [icon],
    transformers: [
      {
        execution: 'async',
        type: 'remark',
        transform({ processor }) {
          processor.use(imgLinks, {
            absolutePath: 'https://kicoe-blog.oss-cn-shanghai.aliyuncs.com/'
          })
        }
      }
    ]
  }

  const carta = new Carta({
    sanitizer: false,
    extensions: [slash(), code({ theme: 'carta-dark' }), emoji(), ext]
  })

  let value = '',
    originValue = ''

  fet.get(`article/${id}`).then((respoi) => {
    if (respoi.ok) {
      article = <ArticleDetail>respoi.data
      originValue = value = article.content.text
    }
  })
</script>

<MarkdownEditor {carta} bind:value theme="custom" />
