<script lang="ts">
  import { createRender, createTable } from 'svelte-headless-table'
  import {
    addColumnFilters,
    addHiddenColumns,
    addPagination,
    addSortBy,
    addTableFilter
  } from 'svelte-headless-table/plugins'

  import { DataTable } from '@/components/blocks/table/index'

  import type { Filter, SearchParams } from '$src/types/table'
  import { tableData, serverItemCount, selectedViewOption, searchUrlQuery, initTableData } from '../(data)/data'

  import TableActions from './table-actions.svelte'

  import type { ViewOption } from '$src/types/table'
  import { isMockMode } from '@/helpers/fetch'
  import { onMount } from 'svelte'
  import { debounce } from '@/helpers/system'
  import { DateFormat } from '@/helpers/date'
  
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
      serverSide,
      fn: ({ filterValue, value }) => {
        return value.toLowerCase().includes(filterValue.toLowerCase())
      }
    }),
    colFilter: addColumnFilters({ serverSide }),
    hide: addHiddenColumns(),
    sort: addSortBy({
      serverSide,
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
      accessor: 'email',
      header: 'Email'
    }),
    table.column({
      accessor: 'link',
      header: 'Link'
    }),
    table.column({
      accessor: 'content',
      header: 'Content'
    }),
    table.column({
      accessor: 'article_id',
      header: 'Article ID'
    }),
    table.column({
      accessor: 'reply_comment_id',
      header: 'Reply Comment ID'
    }),
    table.column({
      accessor: 'root_comment_id',
      header: 'Root Comment ID'
    }),
    table.column({
      accessor: 'created',
      header: 'Created',
      cell: ({ value }) => {
        return DateFormat(value)
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

  // filter.name 对应该字段配置的colFilter
  const filters: Filter[] = [
  ]

  const tableModel = table.createViewModel(columns)

  const viewOption: ViewOption = {
    type: 'hideColumn',
    selected: selectedViewOption,
    options: []
  }

  const { pluginStates } = tableModel

  const { pageIndex, pageSize } = pluginStates.page
  const filterText = pluginStates.filter.filterValue
  const { filterValues } = pluginStates.colFilter
  const { sortKeys } = pluginStates.sort

  let mounted = false
  onMount(() => {
    mounted = true
  })
  
  let params: SearchParams

  const fetchData = debounce(() => {
    const q = encodeURIComponent(JSON.stringify(params))
    if ($searchUrlQuery !== q) {
      initTableData(q)
    }
  }, 200)

  $: {
    if (mounted) {
      params = {
        page: $pageIndex,
        page_size: $pageSize,
        filter_text: $filterText,
        filter_values: <{ [index: string]: number[] }>$filterValues,
        sort_key: $sortKeys[0]
      }
      fetchData()
    }
  }
</script>

<DataTable {tableModel} {filters} {viewOption} />
