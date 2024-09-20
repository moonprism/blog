import type { Comment } from '$src/types/stream'
import { writable } from 'svelte/store'
import { fet } from '@/helpers/fetch'

export const tableData = writable([] as Comment[])

export const serverItemCount = writable(0)

export const selectedViewOption = writable(['reply_comment_id', 'root_comment_id', 'link'])

export const searchUrlQuery = writable('')
export function initTableData(q = '') {
  fet.get(`comment?q=${q}`).then((respoi) => {
    if (respoi.ok) {
      tableData.set(<Comment[]>respoi.data.data)
      serverItemCount.set(respoi.data.pagination.count)
      searchUrlQuery.set(q)
    }
  })
}
