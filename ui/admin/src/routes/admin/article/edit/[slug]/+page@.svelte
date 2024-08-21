<script lang="ts">
  import { page } from '$app/stores'
  import type { ArticleDetail } from '$src/types/stream'
  import { fet } from '@/helpers/fetch'

  import { Carta, MarkdownEditor, type Icon, type Plugin } from 'carta-md'
  import './(styles)/editor_custom.css'

  // 用来自定义解析的组件(img.src)
  import { component } from '@cartamd/plugin-component'
  import { svelte, initializeComponents } from '@cartamd/plugin-component/svelte'
  import PreviewImage from './(components)/preview-image.svelte'
  // 斜杠命令
  import { slash } from '@cartamd/plugin-slash'
  import '@cartamd/plugin-slash/default.css'
  // 代码高亮
  import { code } from '@cartamd/plugin-code'
  import '@cartamd/plugin-code/default.css'
  // 冒号命令
  import { emoji } from '@cartamd/plugin-emoji'
  import '@cartamd/plugin-emoji/default.css'
  // 引用样式
  import 'moonprism-blog-frontend/src/css/markdown.css'
  import { Heading } from 'svelte-radix'
  import { BookUp, Import, Save } from 'lucide-svelte'

  const id = Number($page.params.slug)
  let article = {} as ArticleDetail

  // 编辑器和预览的滚动条同步会有bug
  // 准备直接在remark修复这个问题 https://github.com/Pondorasti/remark-img-links
  const mapped = [svelte('img', PreviewImage)]

  const icon: Icon = {
    id: 'save',
    action: (input) => {
      console.log(
        carta.render(value).then((v) => {
          console.log(v)
        })
      )
    },
    component: BookUp
  }

  const ext: Plugin = {
    icons: [icon],
    transformers: [
      {
        execution: 'sync',
        type: 'rehype',
        transform({ processor }) {
        }
      }
    ]
  }

  const carta = new Carta({
    sanitizer: false,
    extensions: [
      //component(mapped, initializeComponents),
      slash(),
      code({ theme: 'carta-dark' }),
      emoji()
    ]
  })

  let value = '',
    html = ''

  fet.get(`article/${id}`).then((respoi) => {
    if (respoi.ok) {
      article = <ArticleDetail>respoi.data
      value = article.content.text
    }
  })
</script>

<MarkdownEditor {carta} bind:value theme="custom" />
