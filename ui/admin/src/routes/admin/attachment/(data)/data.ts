import type { Attachment, YearGroupInfo } from '$src/types/stream'
import { fet } from '@/helpers/fetch'
import type { SvelteComponent } from 'svelte'
import { writable } from 'svelte/store'
import Form from '../(components)/form.svelte'
import type { Option, SearchParams, ViewOption } from '$src/types/table'
import { createRender } from 'svelte-headless-table'
import CountFilter from '@/components/blocks/cell/count-filter.svelte'

export const tableData = writable([] as Attachment[])

export const serverItemCount = writable(0)

let data: Attachment[]

tableData.subscribe((v) => {
  data = v
})

export function initTableData(params: SearchParams, isAppend = true) {
  if (!isAppend) {
    tableData.set([])
  }
  params.page = params.page_size / 20 - 1
  params.page_size = 20
  const q = encodeURIComponent(JSON.stringify(params))
  fet.get(`attachment?q=${q}`).then((respoi) => {
    if (respoi.ok) {
      if (!isAppend) {
        tableData.set(<Attachment[]>respoi.data.data)
      } else {
        tableData.set([...data, ...(<Attachment[]>respoi.data.data)])
      }
      serverItemCount.set(respoi.data.pagination.count)
    }
  })
}

export const yearOptions = writable([] as Option[])

export function initGroupInfo() {
  fet.get('group/year?model=attachment').then((respoi) => {
    if (respoi.ok) {
      yearOptions.set(
        respoi.data.map((v: YearGroupInfo) => {
          return {
            id: v.year,
            label: String(v.year),
            icon: createRender(CountFilter, {text: String(v.count)})
          }
        })
      )
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
      formOpen: formOpen
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
