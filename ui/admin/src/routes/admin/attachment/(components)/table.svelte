<script lang="ts">
  import { createRender, createTable } from 'svelte-headless-table'
  import {
    addColumnFilters,
    addHiddenColumns,
    addPagination,
    addSortBy,
    addTableFilter
  } from 'svelte-headless-table/plugins'

  import { FlowDataTable } from '@/components/blocks/table/index'

  import TableActions from './table-actions.svelte'
  import {
    formOpen,
    initTableData,
    selectedViewOption,
    viewOption,
    tableData,
    serverItemCount,
    initGroupInfo,
    yearOptions
  } from '../(data)/data'
  import TableRowImage from './table-row-image.svelte'
  import TableRowDate from './table-row-date.svelte'
  import TableRowSummary from './table-row-summary.svelte'
  import type { Attachment } from '$src/types/stream'
  import { isMockMode } from '@/helpers/fetch'
  import { getRealSrc } from '$src/routes/admin/(data)/data'
  import { onMount } from 'svelte'
  import type { Filter, SearchParams } from '$src/types/table'
  import { get } from 'svelte/store'
  import { debounce } from '@/helpers/system'
  import { DateFormat } from '@/helpers/date'

  export let itemClick: ((row: Attachment) => void) | null = null

  const serverSide = !isMockMode
  let paginationConfig = {}
  if (serverSide) {
    paginationConfig = {
      serverSide,
      serverItemCount
    }
  }

  const table = createTable(tableData, {
    page: addPagination(paginationConfig),
    filter: addTableFilter({
      fn: ({ filterValue, value }) => {
        return value.toLowerCase().includes(filterValue.toLowerCase())
      },
      serverSide
    }),
    colFilter: addColumnFilters({ serverSide }),
    hide: addHiddenColumns(),
    sort: addSortBy({
      toggleOrder: ['asc', 'desc']
    })
  })

  const columns = table.createColumns([
    table.column({
      accessor: 'key',
      header: 'Key',
      cell: ({ value }) => {
        return createRender(TableRowImage, {
          src: getRealSrc(value)
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
      accessor: 'year',
      header: 'Year',
      cell: () => {
        return ''
      },
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
      }
    }),
    table.column({
      accessor: 'created',
      header: 'Created',
      cell: ({ value }) => {
        return createRender(TableRowDate, { date: DateFormat(value) })
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
            row: row.original,
            click: itemClick
          })
        }
        return ''
      }
    })
  ])

  const tableModel = table.createViewModel(columns)

  let flowTableComponent: FlowDataTable

  let mounted = false
  onMount(() => {
    mounted = true
    initGroupInfo()
  })

  const filters: Filter[] = [
    {
      name: 'year',
      options: yearOptions
    }
  ]

  $tableData = []

  $: {
    if (mounted && !$formOpen && flowTableComponent) {
      setTimeout(() => {
        // 编辑图片后触发重新布局
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

  const { pageSize } = tableModel.pluginStates.page
  const filterText = tableModel.pluginStates.filter.filterValue
  const { filterValues } = tableModel.pluginStates.colFilter

  const indexTableData = debounce(() => {
    const searchParams = {
      page_size: $pageSize,
      filter_text: $filterText,
      filter_values: <{ [index: string]: number[] }>$filterValues
    }
    initTableData(<SearchParams>searchParams)
  }, 100)
  $: {
    $pageSize
    if (mounted) {
      // 交给下面的监听器初始化
      if ($pageSize !== 20) {
        indexTableData()
      }
    }
  }

  const searchTableData = debounce(() => {
    const searchParams = {
      filter_text: $filterText,
      filter_values: <{ [index: string]: number[] }>$filterValues
    }
    $pageSize = 20
    initTableData(<SearchParams>searchParams, false)
  }, 300)
  $: {
    $filterText
    // todo values不需要防抖处理
    $filterValues
    if (mounted) {
      searchTableData()
    }
  }
</script>

<FlowDataTable bind:this={flowTableComponent} {filters} {tableModel} {viewOption} />
