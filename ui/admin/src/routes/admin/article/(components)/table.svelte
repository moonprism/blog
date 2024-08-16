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

  import { derived, get } from 'svelte/store'
  import { DataTable } from '@/components/blocks/table/index'

  import type { Filter } from '$src/types/table'
  import { tableData, statuses, initTableData, selectedViewOption } from '../(data)/data'
  import TableActions from './table-actions.svelte'
  import TableRowStatus from './table-row-status.svelte'
  import TableRowTitle from './table-row-title.svelte'

  import {
    tableData as tagTableData,
    initTableData as initTagTableData
  } from '../../tag/(data)/data'
  import { BookLock } from 'lucide-svelte'
  import type { Tag } from '$src/types/stream'
  import type { ViewOption } from '$src/types/table'

  if ($tagTableData.length === 0) {
    initTagTableData()
  }

  if ($tableData.length === 0) {
    initTableData()
  }

  const table = createTable(tableData, {
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
      header: 'ID'
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
        return createRender(TableRowStatus, { text: $statuses.find((op) => op.id == value)?.label })
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
      accessor: 'tags',
      header: 'Tags',
      cell: ({ value }) => {
        return value.map((e) => e.name).join(',')
      },
      plugins: {
        colFilter: {
          fn: ({ filterValue, value }) => {
            if (filterValue.length === 0) return true
            if (!Array.isArray(filterValue)) return true
            const vIndex = value.map((e: Tag) => e.id)
            return filterValue.every((e) => vIndex.includes(e))
          },
          initialFilterValue: [],
          render: ({ filterValue }) => {
            return get(filterValue)
          }
        }
      }
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

  const tags = derived(tagTableData, (t) =>
    t.map((v) => {
      return { id: v.id, label: v.name, icon: BookLock }
    })
  )

  // filter.name 对应该字段配置的colFilter
  const filters: Filter[] = [
    {
      name: 'status',
      options: statuses
    },
    {
      name: 'tags',
      options: tags
    }
  ]

  const tableModel = table.createViewModel(columns)
  
  const viewOption: ViewOption = {
    type: 'hideColumn',
    selected: selectedViewOption,
    options: []
  }
</script>

<DataTable {tableModel} {filters} {viewOption} />
