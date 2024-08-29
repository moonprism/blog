<script lang="ts">
  import { fileCDN } from '@/helpers/fetch'

  import {
    Carta,
    MarkdownEditor,
    PreRendered,
    type Icon,
    type InputEnhancer,
    type Plugin
  } from 'carta-md'
  import './editor.css'

  // 代码高亮
  import { code } from '@cartamd/plugin-code'
  import '@cartamd/plugin-code/default.css'
  // 引用markdown-body样式
  //import 'moonprism-blog-frontend/src/css/markdown.css'

  // @ts-ignore
  import remarkImgLinks from '@pondorasti/remark-img-links'

  import FormImageFlow from '@/components/blocks/cell/form-image-flow.svelte'
  import { writable } from 'svelte/store'

  import remarkAdmonitions from 'remark-github-beta-blockquote-admonitions'
  import AttachmentIcon from '../../article/write/[slug]/(components)/attachment-icon.svelte'

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

  // https://github.com/myl7/remark-github-beta-blockquote-admonitions
  const remarkAdConfig = {
    classNameMaps: {
      block: (title: string) => `admonition ad-${title.toLowerCase()}`,
      title: 'admonition-title'
    },
    titleFilter: ['[!NOTE]', '[!IMPORTANT]', '[!WARNING]', '[!TIP]', '[!CAUTION]']
  }

  const ext: Plugin = {
    icons: [attachmentIcon],
    transformers: [
      {
        execution: 'async',
        type: 'remark',
        transform({ processor }) {
          // remark plugins
          processor
            .use(remarkImgLinks, {
              absolutePath: fileCDN
            })
            .use(remarkAdmonitions, remarkAdConfig)
        }
      }
    ]
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

<MarkdownEditor {carta} bind:value theme="gist" mode="tabs" />

<FormImageFlow
  open={isOpenImageFlow}
  callback={(v) => {
    currentInput.insertAt(currentInput.getSelection().start, `![](${v.key})`)
    currentInput.update()
  }}
></FormImageFlow>
