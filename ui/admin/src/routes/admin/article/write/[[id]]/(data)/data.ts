import { Carta } from 'carta-md'
import type { InputEnhancer, UnifiedTransformer } from 'carta-md'
import { fileCDN } from '$src/routes/admin/(data)/data'
import remarkAdmonitions from 'remark-github-beta-blockquote-admonitions'
import { emoji } from '@cartamd/plugin-emoji'
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-expect-error
import remarkImgLinks from '@pondorasti/remark-img-links'
import { code } from '@cartamd/plugin-code'

const alertTypes = ['NOTE', 'IMPORTANT', 'WARNING', 'TIP', 'CAUTION']

// https://github.com/myl7/remark-github-beta-blockquote-admonitions
const remarkAdConfig = {
  classNameMaps: {
    block: (title: string) => `admonition ad-${title.toLowerCase()}`,
    title: 'admonition-title'
  },
  titleFilter: alertTypes.map((v) => `[!${v}]`)
}

export const middlewareTransformers: UnifiedTransformer<'sync' | 'async'>[] = [
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

export const slashSnippets = alertTypes.map((v) => {
  return {
    id: `alert${v}`,
    title: `gi${v}`,
    description: `Create a Github-style ${v} alert blockquote`,
    group: 'Basic',
    action: (input: InputEnhancer) => {
      const s = `> [!${v}]\n> `
      const line = input.getLine()
      input.insertAt(line.start, s)
      const newPos = line.end + s.length
      input.textarea.selectionStart = newPos
      input.textarea.selectionEnd = newPos
    }
  }
})

// TODO 冗余
const carta = new Carta({
  sanitizer: false,
  extensions: [
    {
      transformers: middlewareTransformers
    },
    code(),
    emoji()
  ]
})

const cartaDark = new Carta({
  sanitizer: false,
  extensions: [
    {
      transformers: middlewareTransformers
    },
    code({ theme: 'carta-dark' }),
    emoji()
  ]
})

export async function renderMD(value: string, isDark = false): Promise<string> {
  if (isDark) {
    return await cartaDark.render(value)
  } else {
    return await carta.render(value)
  }
}
