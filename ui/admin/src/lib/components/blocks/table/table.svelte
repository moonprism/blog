<script lang="ts">
  import type { TableViewModel } from 'svelte-headless-table'
  import { Subscribe, Render } from 'svelte-headless-table'

  import * as Table from '@/components/ui/table'
  import type { Filter, ViewOption } from '$src/types/table'

  import TableToolbar from './table-toolbar.svelte'
  import TableHeadSort from './table-head-sort.svelte'
  import TablePagination from './table-pagination.svelte'
  import type { AnyPlugins } from 'svelte-headless-table/plugins'

  export let tableModel: TableViewModel<any, AnyPlugins>
  export let filters: Filter[] = []
  export let viewOption: ViewOption
  
  const { headerRows, pageRows, tableAttrs, tableBodyAttrs } = tableModel
</script>

<div class="space-y-3">
  <TableToolbar {tableModel} {filters} {viewOption} />
  <div class="rounded-md border px-3">
    <Table.Root {...$tableAttrs}>
      <Table.Header>
        {#each $headerRows as headerRow}
          <Subscribe rowAttrs={headerRow.attrs()}>
            <Table.Row>
              {#each headerRow.cells as cell (cell.id)}
                <Subscribe attrs={cell.attrs()} let:attrs props={cell.props()} let:props>
                  <Table.Head {...attrs}>
                    {#if cell.id !== 'actions'}
                      <TableHeadSort sort={props.sort}>
                        <Render of={cell.render()} />
                      </TableHeadSort>
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
  <TablePagination {tableModel} />
</div>
