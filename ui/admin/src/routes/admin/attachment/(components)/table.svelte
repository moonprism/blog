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
  import { formOpen, initTableData, selectedViewOption, tableData } from '../(data)/data'
  import TableRowImage from './table-row-image.svelte'
  import { PUBLIC_ATTACHMENT_CDN } from '$env/static/public'
  import TableRowDate from './table-row-date.svelte'
  import TableRowSummary from './table-row-summary.svelte'
  import type { ViewOption } from '$src/types/table'
  import { resetMode } from 'mode-watcher'

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
        return createRender(TableRowImage, {
          src: `${value.startsWith('data') ? '' : PUBLIC_ATTACHMENT_CDN}${value}`
        })
      }
    }),
    table.column({
      accessor: 'summary',
      header: 'Summary',
      cell: ({ value }) => {
        return createRender(TableRowSummary, { text: value })
      }
    }),
    table.column({
      accessor: 'created',
      header: 'Created',
      cell: ({ value }) => {
        const time = new Date(value * 1000)
        return createRender(TableRowDate, { date: time.toLocaleDateString('en') })
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

  const viewOption: ViewOption = {
    type: 'cardSize',
    selected: selectedViewOption,
    options: ['max-w-xs | 320', 'max-w-sm | 384', 'max-w-md | 448', 'max-w-xl | 576']
  }

  let flowTableComponent: FlowDataTable

  let initFlag = false
  $: {
    if ($formOpen) {
      initFlag = true
    }
    if (!$formOpen && initFlag && flowTableComponent) {
      setTimeout(() => {
        // 意外发现这样能正常运行
        flowTableComponent.reset()
      }, 300)
    }
  }

  $: {
    if ($selectedViewOption.length != 0) {
      let size = parseInt($selectedViewOption[0].split('|')[1])
      if (flowTableComponent) {
        flowTableComponent.reset(size)
      }
    }
  }
</script>

<FlowDataTable bind:this={flowTableComponent} {tableModel} {viewOption} />
