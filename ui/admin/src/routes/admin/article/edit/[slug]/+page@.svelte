<script lang="ts">
  import { page } from '$app/stores'
  import type { ArticleDetail } from '$src/types/stream'
  import { fet } from '@/helpers/fetch'

  const id = Number($page.params.slug)

  let article = {} as ArticleDetail

  import { Carta, MarkdownEditor } from 'carta-md'
  import './(styles)/editor_custom.css'

  import { component } from '@cartamd/plugin-component'
  import { svelte, initializeComponents } from '@cartamd/plugin-component/svelte'
  import PreviewImage from './(components)/preview-image.svelte'
  import { slash } from '@cartamd/plugin-slash'
  import '@cartamd/plugin-slash/default.css'
  import { code } from '@cartamd/plugin-code'
  import '@cartamd/plugin-code/default.css'
  import { emoji } from '@cartamd/plugin-emoji'
  import '@cartamd/plugin-emoji/default.css'
  const mapped = [svelte('img', PreviewImage) /* other components ... */]

  const carta = new Carta({
    sanitizer: false,
    extensions: [component(mapped, initializeComponents), slash(), code(), emoji()]
  })

  let value = ''

  fet.get(`article/${id}`).then((respoi) => {
    if (respoi.ok) {
      article = <ArticleDetail>respoi.data
      value = article.content.text
    }
  })
</script>

<MarkdownEditor {carta} bind:value theme="custom" />