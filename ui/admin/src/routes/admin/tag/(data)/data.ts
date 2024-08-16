import type { Tag } from '$src/types/stream'
import { fet } from '@/helpers/fetch'
import type { SvelteComponent } from 'svelte'
import { writable } from 'svelte/store'
import Form from '../(components)/form.svelte'

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

let formComponent:SvelteComponent

const formOpen = writable(false)

export function openForm(t?: Tag) {
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

export const selectedViewOption = writable([] as string[])