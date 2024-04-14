<script lang="ts">
  import { createTable, Render, Subscribe, createRender } from 'svelte-headless-table'
  import {
    addColumnFilters,
    addHiddenColumns,
    addPagination,
    addSelectedRows,
    addSortBy,
    addTableFilter
  } from 'svelte-headless-table/plugins'

  import { get, writable } from 'svelte/store'
  import * as Table from '@/components/ui/table'

  import DataTableActions from './data-table-actions.svelte'

  import DataTablePagination from './data-table-pagination.svelte'
  import { fet } from '@/helpers/fetch'
  import TableColumnImage from './table-column-image.svelte'
  import type { Article } from '$src/types/stream'

  import DataTableToolbar from './data-table-toolbar.svelte'
  import TableColumnTitle from './table-column-title.svelte'
  import TableHeaderSort from './table-header-sort.svelte'
  import { BookLock } from 'lucide-svelte'

  let data: Article[] = []

  let storeData = writable(data)

  function getArticles() {
    fet.get('article').then((respoi) => {
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
        }
      },
      cell: ({ value }) => {
        return createRender(BookLock, {})
      }
    }),
    table.column({
      accessor: 'title',
      header: 'Title',
      cell: ({ value }) => {
        return createRender(TableColumnTitle, {
          text: value
        })
      }
    }),
    table.column({
      accessor: 'image',
      header: 'Image',
      cell: ({ value }) => {
        return createRender(TableColumnImage, {
          src: `https://kicoe-blog.oss-cn-shanghai.aliyuncs.com/${value}`
        })
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
      }
    }),
    table.column({
      accessor: 'updated',
      header: 'Updated',
      cell: ({ value }) => {
        const time = new Date(value * 1000)
        return time.toLocaleDateString('en')
      }
    }),
    table.column({
      accessor: ({ id }) => id,
      header: '',
      cell: ({ value }) => {
        return createRender(DataTableActions, { id: value })
      }
    })
  ])

  const tableModel = table.createViewModel(columns)

  const { headerRows, pageRows, tableAttrs, tableBodyAttrs, pluginStates } = tableModel
</script>

<div class="space-y-3">
  <DataTableToolbar {tableModel} />

  <div class="rounded-md border px-3">
    <Table.Root {...$tableAttrs}>
      <Table.Header>
        {#each $headerRows as headerRow}
          <Subscribe rowAttrs={headerRow.attrs()}>
            <Table.Row>
              {#each headerRow.cells as cell (cell.id)}
                <Subscribe attrs={cell.attrs()} let:attrs props={cell.props()} let:props>
                  <Table.Head {...attrs}>
                    {#if cell.id !== '' && cell.id !== 'actions'}
                      <TableHeaderSort {props}>
                        <Render of={cell.render()} />
                      </TableHeaderSort>
                    {:else}
                      <Render of={cell.render()} />
                    {/if}
                  </Table.Head>
                </Subscribe>
              {/each}
            </Table.Row>
          </Subscribe>
        {/each}
      </Table.Header>
      <Table.Body {...$tableBodyAttrs}>
        {#each $pageRows as row (row.id)}
          <Subscribe rowAttrs={row.attrs()} let:rowAttrs>
            <Table.Row {...rowAttrs}>
              {#each row.cells as cell (cell.id)}
                <Subscribe attrs={cell.attrs()} let:attrs>
                  <Table.Cell {...attrs}>
                    <Render of={cell.render()} />
                  </Table.Cell>
                </Subscribe>
              {/each}
            </Table.Row>
          </Subscribe>
        {/each}
      </Table.Body>
    </Table.Root>
  </div>

  <DataTablePagination {tableModel} />
</div>
