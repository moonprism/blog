<script lang="ts">
  import { createRender, createTable } from 'svelte-headless-table'
  import {
    addColumnFilters,
    addHiddenColumns,
    addPagination,
    addSelectedRows,
    addSortBy,
    addTableFilter
  } from 'svelte-headless-table/plugins'

  import { get, writable } from 'svelte/store'
  import { DataTable } from '@/components/blocks/table/index'

  import { fet } from '@/helpers/fetch'
  import type { Article } from '$src/types/stream'
  import type { filter } from '$src/types/table'
  import { statuses } from '../(data)/data'
  import TableActions from './table-actions.svelte'
  import TableRowStatus from './table-row-status.svelte'
  import TableRowTitle from './table-row-title.svelte'

  let data: Article[] = []

  let storeData = writable(data)

  function getArticles() {
    fet.get('article').then(respoi => {
      if (respoi.ok) {
        data = <Article[]>respoi.data
        storeData.set(data)
      }
    })
  }
  
  getArticles()

  const table = createTable(storeData, {
    select: addSelectedRows(),
    page: addPagination(),
    filter: addTableFilter({
      fn: ({ filterValue, value }) => {
        return value.toLowerCase().includes(filterValue.toLowerCase())
      }
    }),
    colFilter: addColumnFilters(),
    hide: addHiddenColumns(),
    sort: addSortBy({
      toggleOrder: ['asc', 'desc']
    })
  })

  const columns = table.createColumns([
    table.column({
      accessor: 'id',
      header: 'ID',
      plugins: {
        colFilter: {
          fn: ({ filterValue, value }) => {
            if (filterValue.length === 0) return true
            if (!Array.isArray(filterValue)) return true
            return filterValue.includes(value)
          },
          initialFilterValue: [],
          render: ({ filterValue }) => {
            return get(filterValue)
          }
        }
      }
    }),
    table.column({
      accessor: 'status',
      header: 'Status',
      plugins: {
        colFilter: {
          fn: ({ filterValue, value }) => {
            if (filterValue.length === 0) return true
            if (!Array.isArray(filterValue)) return true
            return filterValue.includes(value)
          },
          initialFilterValue: [],
          render: ({ filterValue }) => {
            return get(filterValue)
          }
        },
        filter: { exclude: true }
      },
      cell: ({ value }) => {
        return createRender(TableRowStatus, { text: statuses.find((op) => op.id == value)?.label })
      }
    }),
    table.column({
      accessor: 'title',
      header: 'Title',
      cell: ({ value }) => {
        return createRender(TableRowTitle, { text: value })
      }
    }),
    table.column({
      accessor: 'image',
      header: 'Image'
    }),
    table.column({
      accessor: 'summary',
      header: 'Summary'
    }),
    table.column({
      accessor: 'created',
      header: 'Created',
      cell: ({ value }) => {
        const time = new Date(value * 1000)
        return time.toLocaleDateString('en')
      },
      plugins: {
        filter: { exclude: true }
      }
    }),
    table.column({
      accessor: 'updated',
      header: 'Updated',
      cell: ({ value }) => {
        const time = new Date(value * 1000)
        return time.toLocaleDateString('en')
      },
      plugins: {
        filter: { exclude: true }
      }
    }),
    table.display({
      id: 'actions',
      header: () => {
        return ''
      },
      cell: ({ row }) => {
        if (row.isData() && row.original) {
          return createRender(TableActions, {
            row: row.original
          })
        }
        return ''
      }
    })
  ])

  const filters: filter[] = [
    {
      name: 'status',
      options: statuses
    },
    {
      name: 'id',
      options: statuses
    }
  ]

  const tableModel = table.createViewModel(columns)
</script>

<DataTable {tableModel} {filters} />
