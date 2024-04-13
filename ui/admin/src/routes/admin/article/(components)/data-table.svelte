<script lang="ts">
  import { createTable, Render, Subscribe, createRender } from 'svelte-headless-table'
  import { addColumnFilters, addPagination, addTableFilter } from 'svelte-headless-table/plugins'

  import { get, writable } from 'svelte/store'
  import * as Table from '@/components/ui/table'

  import DataTableActions from './data-table-actions.svelte'

  import DataTablePagination from './data-table-pagination.svelte'
  import { fet } from '@/helpers/fetch'
  import DataTableImage from './data-table-image.svelte'
  import type { Article } from '$src/types/stream'

  import DataTableToolbar from './data-table-toolbar.svelte'

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
    page: addPagination(),
    filter: addTableFilter({
      fn: ({ filterValue, value }) => {
        return value.toLowerCase().includes(filterValue.toLowerCase())
      }
    }),
    colFilter: addColumnFilters()
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
            if (typeof value !== 'string') {
              let v = String(value)
              return filterValue.some((filter) => {
                return v.includes(filter)
              })
            }
            return filterValue.some((filter) => {
              return value.includes(filter)
            })
          },
          initialFilterValue: [],
          render: ({ filterValue }) => {
            return get(filterValue)
          }
        }
      }
    }),
    table.column({
      accessor: 'title',
      header: 'Title'
    }),
    table.column({
      accessor: 'image',
      header: 'Image',
      cell: ({ value }) => {
        return createRender(DataTableImage, {
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
      header: 'Created'
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

<div>
  <DataTableToolbar {tableModel} {data} />
  <div class="rounded-md border">
    <Table.Root {...$tableAttrs}>
      <Table.Header>
        {#each $headerRows as headerRow}
          <Subscribe rowAttrs={headerRow.attrs()}>
            <Table.Row>
              {#each headerRow.cells as cell (cell.id)}
                <Subscribe attrs={cell.attrs()} let:attrs props={cell.props()}>
                  <Table.Head {...attrs}>
                    <Render of={cell.render()} />
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
                    {#if cell.id === 'image'}
                      <Render of={cell.render()} />
                    {:else}
                      <Render of={cell.render()} />
                    {/if}
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
