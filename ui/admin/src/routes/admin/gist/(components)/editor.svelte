<script lang="ts">
  import { Carta, MarkdownEditor, type Icon, type InputEnhancer, type Plugin } from 'carta-md'
  import '../(styles)/editor.css'

  // 代码高亮
  import { code } from '@cartamd/plugin-code'
  import '@cartamd/plugin-code/default.css'
  // 引用markdown-body样式
  import 'moonprism-blog-frontend/src/styles/gist.md.css'

  import FormImageFlow from '@/components/blocks/cell/form-image-flow.svelte'
  import { writable } from 'svelte/store'

  import AttachmentIcon from '../../article/write/[[id]]/(components)/attachment-icon.svelte'
  import { middlewareTransformers } from '../../article/write/[[id]]/(data)/data'

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

  const ext: Plugin = {
    icons: [attachmentIcon],
    transformers: middlewareTransformers
  }

  const carta = new Carta({
    disableIcons: true,
    sanitizer: false,
    extensions: [code(), ext]
  })

  export let value = ''

  export const render = async () => {
    return await carta.render(value)
  }
</script>

<div class="gist">
  <MarkdownEditor {carta} bind:value theme="gist" mode="tabs" />

  <FormImageFlow
    open={isOpenImageFlow}
    callback={(v) => {
      currentInput.insertAt(currentInput.getSelection().start, `![](${v.key})`)
      currentInput.update()
    }}
  ></FormImageFlow>
</div>
