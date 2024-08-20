import type { Article, Tag } from '$src/types/stream'
import type { Option } from '$src/types/table'
import { fet } from '@/helpers/fetch'
import { Radio, VenetianMask } from 'lucide-svelte'
import { derived, readable, writable } from 'svelte/store'
import Form from '../(components)/form.svelte'
import type { SvelteComponent } from 'svelte'
import type { ReadOrWritable } from '$src/types/store'
import { createRender } from 'svelte-headless-table'
import TableRowTagColors from '../(components)/table-row-tag-colors.svelte'
import { tableData as tagTableData } from '../../tag/(data)/data'

export const tableData = writable([] as Article[])

export function initTableData() {
  fet.get('article').then((respoi) => {
    if (respoi.ok) {
      tableData.set(<Article[]>respoi.data)
    }
  })
}

export const statuses: ReadOrWritable<Option[]> = readable([
  {
    id: 0,
    label: 'Draft',
    icon: createRender(VenetianMask, { class: 'h-4 w-4' })
  },
  {
    id: 1,
    label: 'Published',
    icon: createRender(Radio, { class: 'h-4 w-4' })
    //icon: Radio
  }
])

export const tags = derived(tagTableData, (t) =>
  t.map((v) => {
    return {
      id: v.id,
      label: v.name,
      icon: createRender(TableRowTagColors, { colors: [v.color] })
    }
  })
)

export function getDefaultFormData() {
  return {
    id: 0,
    title: '',
    status: 0,
    rune: 0,
    image: '',
    summary: '',
    tags: [] as Tag[],
  } as Article
}

let formComponent: SvelteComponent

// 为了动画
const formOpen = writable(false)

export function openForm(t?: Article) {
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
      formOpen: formOpen,
    }
  })
  formOpen.set(true)
}

export function closeForm() {
  formOpen.set(false)
}

export const selectedViewOption = writable(['image'])