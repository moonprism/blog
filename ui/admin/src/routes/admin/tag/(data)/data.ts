import type { Tag } from '$src/types/stream'
import { fet } from '@/helpers/fetch'
import { writable } from 'svelte/store'

export const tableData = writable([] as Tag[])

export function initTableData() {
  fet.get('tag').then((respoi) => {
    if (respoi.ok) {
      tableData.set(<Tag[]>respoi.data)
    }
  })
}

export function getDefaultFormData() {
  return { id: 0, name: '', color: '#000000' } as Tag
}

export const formData = writable(getDefaultFormData())
export const formOpen = writable(false)

export function openForm(t?: Tag) {
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
