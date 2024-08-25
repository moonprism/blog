<script lang="ts">
  import { createRender, createTable } from 'svelte-headless-table'
  import {
    addColumnFilters,
    addHiddenColumns,
    addPagination,
    addSortBy,
    addTableFilter
  } from 'svelte-headless-table/plugins'

  import { get } from 'svelte/store'
  import { DataTable } from '@/components/blocks/table/index'

  import type { Filter } from '$src/types/table'
  import {
    tableData,
    statuses,
    initTableData,
    selectedViewOption,
    tags,
    serverItemCount,
    searchParams
  } from '../(data)/data'
  import TableActions from './table-actions.svelte'
  import TableRowTitle from './table-row-title.svelte'

  import {
    tableData as tagTableData,
    initTableData as initTagTableData
  } from '../../tag/(data)/data'
  import type { Tag } from '$src/types/stream'
  import type { ViewOption } from '$src/types/table'
  import TableRowTagColors from './table-row-tag-colors.svelte'
  import TableRowImage from './table-row-image.svelte'
  import { getRealSrc } from '@/helpers/fetch'
  import { onMount } from 'svelte'

  if ($tagTableData.length === 0) {
    initTagTableData()
  }

  const serverSide = true

  const table = createTable(tableData, {
    page: addPagination({
      serverSide,
      serverItemCount
    }),
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
        const p = $statuses.find((op) => op.id == value)
        if (!p) return ''
        return p.icon
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
      header: 'Image',
      cell: ({ value }) => {
        return createRender(TableRowImage, { src: getRealSrc(value) })
      }
    }),
    table.column({
      accessor: 'summary',
      header: 'Summary'
    }),
    table.column({
      accessor: 'tags',
      header: 'Tags',
      cell: ({ value }) => {
        return createRender(TableRowTagColors, {
          colors: value.map((e) => e.color),
          class: 'h-2 w-2'
        })
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

  const { pluginStates } = tableModel

  const { pageIndex, pageSize } = pluginStates.page
  const filterText = pluginStates.filter.filterValue
  const { filterValues } = pluginStates.colFilter
  const { sortKeys } = pluginStates.sort

  let mounted = false
  onMount(() => {
    mounted = true
  })

  $: {
    if (mounted) {
      $searchParams = {
        page: $pageIndex,
        page_size: $pageSize,
        filter_text: $filterText,
        filter_values: <{ [index: string]: number[] }>$filterValues,
        sort_key: $sortKeys[0]
      }
      initTableData($searchParams)
    } else {
      // 编辑文章涉及到路由的改变，只好缓存这些查询参数了
      // not undefined
      if ($searchParams.page) {
        $pageIndex = $searchParams.page
      }
      if ($searchParams.page_size) {
        $pageSize = $searchParams.page_size
      }
      if ($searchParams.filter_text) {
        $filterText = $searchParams.filter_text
      }
      if ($searchParams.filter_values) {
        $filterValues = $searchParams.filter_values
      }
      if ($searchParams.sort_key) {
        $sortKeys = [$searchParams.sort_key]
      }
    }
  }
</script>

<DataTable {tableModel} {filters} {viewOption} />
