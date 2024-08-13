<script lang="ts">
  import { createRender, createTable } from 'svelte-headless-table'
  import {
    addHiddenColumns,
    addPagination,
    addSelectedRows,
    addSortBy,
    addTableFilter
  } from 'svelte-headless-table/plugins'

  import { DataTable } from '@/components/blocks/table/index'

  import TableActions from './table-actions.svelte'
  import { initTableData, tableData } from '../(data)/data'
  import TableRowColor from './table-row-color.svelte'

  if ($tableData.length === 0) {
    initTableData()
  }

  const table = createTable(tableData, {
    select: addSelectedRows(),
    page: addPagination(),
    filter: addTableFilter({
      fn: ({ filterValue, value }) => {
        return value.toLowerCase().includes(filterValue.toLowerCase())
      }
    }),
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
      accessor: 'name',
      header: 'Name'
    }),
    table.column({
      accessor: 'color',
      header: 'Color',
      cell: ({ value }) => {
        return createRender(TableRowColor, { color: value })
      },
      plugins: {
        filter: { exclude: true }
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

  const tableModel = table.createViewModel(columns)
</script>

<DataTable {tableModel} />
