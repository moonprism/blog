import type { Tag } from '$src/types/stream'
import { writable } from 'svelte/store'

export const tableData = writable([] as Tag[])

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
