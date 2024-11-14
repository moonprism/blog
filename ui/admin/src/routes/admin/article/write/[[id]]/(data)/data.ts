import { Carta } from 'carta-md'
import type { InputEnhancer, UnifiedTransformer } from 'carta-md'
import remarkAdmonitions from 'remark-github-beta-blockquote-admonitions'
import { emoji } from '@cartamd/plugin-emoji'
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-expect-error
import remarkImgLinks from '@pondorasti/remark-img-links'
import { code } from '@cartamd/plugin-code'
import { get } from 'svelte/store'
import { appInfo } from '$src/routes/admin/(data)/data'

const alertTypes = ['NOTE', 'IMPORTANT', 'WARNING', 'TIP', 'CAUTION']

// https://github.com/myl7/remark-github-beta-blockquote-admonitions
const remarkAdConfig = {
  classNameMaps: {
    block: (title: string) => `admonition ad-${title.toLowerCase()}`,
    title: 'admonition-title'
  },
  titleFilter: alertTypes.map((v) => `[!${v}]`)
}

export const getMiddlewareTransformers = (
  absolutePath: string
): UnifiedTransformer<'sync' | 'async'>[] => {
  return [
    {
      execution: 'async',
      type: 'remark',
      transform({ processor }) {
        // remark plugins
        processor
          .use(remarkImgLinks, {
            absolutePath
          })
          .use(remarkAdmonitions, remarkAdConfig)
      }
    }
  ]
}

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

function newCarta(isDark = false) {
  let codeExt = code()
  if (isDark) {
    codeExt = code({ theme: 'carta-dark' })
  }
  return new Carta({
    sanitizer: false,
    extensions: [
      {
        transformers: getMiddlewareTransformers(get(appInfo).attachmentCDN)
      },
      codeExt,
      emoji()
    ]
  })
}

export async function renderMD(value: string, isDark = false): Promise<string> {
  return await newCarta(isDark).render(value)
}
