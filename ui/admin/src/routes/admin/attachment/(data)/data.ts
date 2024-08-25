import type { Attachment } from "$src/types/stream";
import { fet } from "@/helpers/fetch";
import type { SvelteComponent } from "svelte";
import { writable } from "svelte/store";
import Form from "../(components)/form.svelte";
import type { ViewOption } from "$src/types/table";

export const tableData = writable([] as Attachment[])

export function initTableData() {
  fet.get('attachment').then((respoi) => {
    if (respoi.ok) {
      tableData.set(<Attachment[]>respoi.data.data)
    }
  })
}

export function getDefaultFormData() {
  return { id: 0, key: '', summary: '' } as Attachment
}

let formComponent: SvelteComponent

export const formOpen = writable(false)

export function openForm(t?: Attachment) {
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

export const selectedViewOption = writable([] as string[])
export const viewOption: ViewOption = {
  type: 'cardSize',
  selected: selectedViewOption,
  options: ['max-w-xs | 320', 'max-w-sm | 384', 'max-w-md | 448']
}
selectedViewOption.set([viewOption.options[1]])