<script lang="ts">
  import { createRender, createTable } from 'svelte-headless-table'
  import {
    addHiddenColumns,
    addPagination,
    addSelectedRows,
    addSortBy,
    addTableFilter
  } from 'svelte-headless-table/plugins'

  import { FlowDataTable } from '@/components/blocks/table/index'

  import TableActions from './table-actions.svelte'
  import { initTableData, tableData } from '../(data)/data'
  import TableRowImage from './table-row-image.svelte'
  import { PUBLIC_ATTACHMENT_CDN } from '$env/static/public'

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
      accessor: 'link',
      header: 'Link',
      cell: ({ value }) => {
        return createRender(TableRowImage, { src: `${PUBLIC_ATTACHMENT_CDN}${value}` })
      }
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

<FlowDataTable {tableModel} />