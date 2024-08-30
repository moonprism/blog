import type { Gist, GistLangGroupInfo } from '$src/types/stream'
import { writable } from 'svelte/store'
import type { Option } from '$src/types/table'
import { createRender } from 'svelte-headless-table'
import { fet } from '@/helpers/fetch'
import type { SvelteComponent } from 'svelte'
import CountFilter from '@/components/blocks/cell/count-filter.svelte'
import Form from '../(components)/form.svelte'

export const tableData = writable([] as Gist[])

export const serverItemCount = writable(0)

export const selectedViewOption = writable(['content'])

export const langOptions = writable([] as Option[])
export function initGroupInfo() {
  fet.get('group/gist-lang').then((respoi) => {
    if (respoi.ok) {
      langOptions.set(
        respoi.data.map((v: GistLangGroupInfo) => {
          return {
            id: v.lang,
            label: getLangLabel(v.lang),
            icon: createRender(CountFilter, { text: String(v.count) })
          }
        })
      )
    }
  })
}
initGroupInfo()

export const searchUrlQuery = writable('')
export function initTableData(q = '') {
  fet.get(`gist?q=${q}`).then((respoi) => {
    if (respoi.ok) {
      tableData.set(<Gist[]>respoi.data.data)
      serverItemCount.set(respoi.data.pagination.count)
      searchUrlQuery.set(q)
    }
  })
}

export function getDefaultFormData() {
  return {
    id: 0,
    title: '',
    lang: '',
    content: ''
  } as Gist
}

let formComponent: SvelteComponent

const formOpen = writable(false)
export function openForm(t?: Gist) {
  if (!t) {
    t = getDefaultFormData()
  }
  if (formComponent) {
    formComponent.$destroy()
  }
  formComponent = new Form({
    target: document.body,
    props: {
      formData: t,
      formOpen: formOpen
    }
  })
  formOpen.set(true)
}

export function closeForm() {
  formOpen.set(false)
}

function getLangLabel(lang = '') {
  return langItems.find((i) => i.value === lang)?.label
}

export const langItems = [
  {
    value: 'md',
    label: 'Markdown'
  },
  {
    value: 'lua',
    label: 'Lua'
  },
  {
    value: 'go',
    label: 'Golang'
  },
  {
    value: 'js',
    label: 'JavaScript'
  },
  {
    value: 'sh',
    label: 'Shell'
  },
  {
    value: 'yml',
    label: 'YAML'
  },
  {
    value: 'css',
    label: 'CSS'
  },
  {
    value: 'php',
    label: 'PHP'
  },
  {
    value: 'sql',
    label: 'SQL'
  },
  {
    value: 'html',
    label: 'HTML'
  }
]
