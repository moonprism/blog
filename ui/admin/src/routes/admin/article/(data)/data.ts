import type { Article, Tag } from '$src/types/stream'
import type { option } from '$src/types/table'
import { fet } from '@/helpers/fetch'
import { BookLock, BookOpen, VenetianMask } from 'lucide-svelte'
import type { ReadOrWritable } from 'svelte-headless-table'
import { readable, writable } from 'svelte/store'

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

export const formData = writable(getDefaultFormData())
export const formOpen = writable(false)

export function openForm(t?: Article) {
  if (!t) {
    t = getDefaultFormData()
  }
  formData.set(t)
  formOpen.set(true)
}

export function closeForm() {
  formData.set(getDefaultFormData())
  formOpen.set(false)
}