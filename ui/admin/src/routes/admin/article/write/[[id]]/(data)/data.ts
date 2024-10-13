import { Carta } from 'carta-md'
import type { UnifiedTransformer } from 'carta-md'
import { fileCDN } from '$src/routes/admin/(data)/data'
import remarkAdmonitions from 'remark-github-beta-blockquote-admonitions'
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-expect-error
import remarkImgLinks from '@pondorasti/remark-img-links'
import { code } from '@cartamd/plugin-code'

// https://github.com/myl7/remark-github-beta-blockquote-admonitions
const remarkAdConfig = {
  classNameMaps: {
    block: (title: string) => `admonition ad-${title.toLowerCase()}`,
    title: 'admonition-title'
  },
  titleFilter: ['[!NOTE]', '[!IMPORTANT]', '[!WARNING]', '[!TIP]', '[!CAUTION]']
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

// TODO 冗余
const carta = new Carta({
  sanitizer: false,
  extensions: [
    {
      transformers: middlewareTransformers
    },
    code()
  ]
})

const cartaDark = new Carta({
  sanitizer: false,
  extensions: [
    {
      transformers: middlewareTransformers
    },
    code({ theme: 'carta-dark' })
  ]
})

export async function renderMD(value: string, isDark = false): Promise<string> {
  if (isDark) {
    return await cartaDark.render(value)
  } else {
    return await carta.render(value)
  }
}
