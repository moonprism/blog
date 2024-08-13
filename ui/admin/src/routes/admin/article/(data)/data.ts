import type { Article, Tag } from '$src/types/stream'
import type { option } from '$src/types/table'
import { fet } from '@/helpers/fetch'
import { BookLock, BookOpen, VenetianMask } from 'lucide-svelte'
import type { ReadOrWritable } from 'svelte-headless-table'
import { readable, writable } from 'svelte/store'
import Form from '../(components)/form.svelte'
import type { SvelteComponent } from 'svelte'

export const tableData = writable([] as Article[])

export function initTableData() {
  fet.get('article').then((respoi) => {
    if (respoi.ok) {
      tableData.set(<Article[]>respoi.data)
    }
  })
}

export const statuses: ReadOrWritable<option[]> = readable([
  {
    id: 0,
    label: 'Draft',
    icon: BookLock
  },
  {
    id: 1,
    label: 'Published',
    icon: BookOpen
  },
  {
    id: 2,
    label: 'Hidden',
    icon: VenetianMask
  }
])

export function getDefaultFormData() {
  return {
    id: 0,
    title: '',
    status: 0,
    rune: 0,
    image: '',
    summary: '',
    content: '',
    tags: [] as Tag[],
  } as Article
}

let formComponent:SvelteComponent

// 为了动画
const formOpen = writable(false)

export function openForm(t?: Article) {
  if (!t) {
    t = getDefaultFormData()
  }
  if (formComponent) {
    formComponent.$destroy()
  }
  formComponent =  new Form({
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