<script lang="ts">
  import { createTable, Render, Subscribe, createRender } from 'svelte-headless-table'
  import { addPagination } from 'svelte-headless-table/plugins'

  import { readable } from 'svelte/store'
  import * as Table from '@/components/ui/table'

  import DataTableActions from './data-table-actions.svelte'

  import { Button } from '$lib/components/ui/button'

  import DataTablePagination from './data-table-pagination.svelte'

  type Article = {
    id: number
    title: string
    status: number
    rune: number
    image: string
    summary: string
    content: string
  }

  const data: Article[] = [
    {
      id: 1,
      title: 'xxx',
      status: 0,
      rune: 100,
      image: 'xxx',
      summary: 'xxxx',
      content: 'xxxxxxxxx'
    },
    {
      id: 1,
      title: 'xxx',
      status: 0,
      rune: 100,
      image: 'xxx',
      summary: 'xxxx',
      content: 'xxxxxxxxx'
    },
    {
      id: 1,
      title: 'xxx',
      status: 0,
      rune: 100,
      image: 'xxx',
      summary: 'xxxx',
      content: 'xxxxxxxxx'
    },
    {
      id: 1,
      title: 'xxx',
      status: 0,
      rune: 100,
      image: 'xxx',
      summary: 'xxxx',
      content: 'xxxxxxxxx'
    }
  ]

  const table = createTable(readable(data), {
    page: addPagination()
  })

  const columns = table.createColumns([
    table.column({
      accessor: 'id',
      header: 'ID'
    }),
    table.column({
      accessor: 'status',
      header: 'Status'
    }),
    table.column({
      accessor: 'title',
      header: 'Title'
    }),
    table.column({
      accessor: 'image',
      header: 'Image',
      cell: ({ value }) => {
        return value
      }
    }),
    table.column({
      accessor: 'summary',
      header: 'Summary'
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
