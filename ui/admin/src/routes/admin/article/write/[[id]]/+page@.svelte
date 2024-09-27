<script lang="ts">
  import { page } from '$app/stores'
  import type { ArticleDetail } from '$src/types/stream'
  import { fet, isRequestIn } from '@/helpers/fetch'

  import { Carta, MarkdownEditor, type Icon, type InputEnhancer, type Plugin } from 'carta-md'
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

  import SaveIcon from './(components)/save-icon.svelte'
  import toast from '$lib/helpers/toast'
  import { onMount } from 'svelte'
  import AttachmentIcon from './(components)/attachment-icon.svelte'
  import FormImageFlow from '@/components/blocks/cell/form-image-flow.svelte'
  import { writable } from 'svelte/store'
  import { SquareX } from 'lucide-svelte'
  import { alertDialog } from '@/components/blocks/dialog/alert'

  import { goto } from '$app/navigation'
  import { middlewareTransformers } from './(data)/data'
  import { base } from '$app/paths'

  const id = Number($page.params.id)
  let article = {} as ArticleDetail

  const saveIcon: Icon = {
    id: 'save',
    action: () => {
      save()
    },
    component: SaveIcon
  }

  let isOpenImageFlow = writable(false)
  let currentInput: InputEnhancer

  const attachmentIcon: Icon = {
    id: 'attachment',
    action: (input) => {
      currentInput = input
      $isOpenImageFlow = true
    },
    component: AttachmentIcon
  }

  // issue: 这个编辑器 Esc 不退出编辑模式
  function esc() {
    // @ts-ignore
    document.activeElement?.blur()
  }

  async function save() {
    if ($isRequestIn) {
      return
    }
    const html = await carta.render(value)
    const res = await fet.put(`article/${id}`, { text: value, html })
    if (res.ok) {
      toast.success('保存成功')
      originValue = value
      //esc()
    }
  }

  function handlePopstate(): void {
    history.pushState(null, '', document.URL)
    if (originValue === value) {
      back()
      return
    }
    alertDialog('检测到未保存的文章内容！', '退出编辑').then(() => {
      back()
    })
  }

  function handleBeforeUnload(event: BeforeUnloadEvent) {
    if (originValue === value) {
      return
    }
    // Cancel the event as stated by the standard.
    event.preventDefault()
    // Chrome requires returnValue to be set.
    event.returnValue = ''
  }

  onMount(() => {
    // 禁止回退和刷新
    history.pushState(null, '', document.URL)
    window.addEventListener('popstate', handlePopstate)
    window.addEventListener('beforeunload', handleBeforeUnload)

    return () => {
      window.removeEventListener('popstate', handlePopstate)
      window.removeEventListener('beforeunload', handleBeforeUnload)
    }
  })

  function back() {
    goto(`${base}/admin/article`)
  }

  const ext: Plugin = {
    icons: [attachmentIcon, saveIcon],
    transformers: middlewareTransformers
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

<button
  class="group absolute z-10 flex h-8 w-8 items-center justify-center text-foreground/80 hover:text-foreground"
  on:click={handlePopstate}
>
  <SquareX class="h-5 w-5 group-hover:h-6 group-hover:w-6"></SquareX>
</button>

<div class="art">
  <MarkdownEditor {carta} bind:value theme="custom" />

  <FormImageFlow
    open={isOpenImageFlow}
    callback={(v) => {
      currentInput.insertAt(currentInput.getSelection().start, `![](${v.key})`)
      currentInput.update()
    }}
  ></FormImageFlow>
</div>
