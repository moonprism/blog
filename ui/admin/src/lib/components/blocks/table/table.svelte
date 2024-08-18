<script lang="ts">
  import type { TableViewModel } from 'svelte-headless-table'
  import { Subscribe, Render } from 'svelte-headless-table'

  import * as Table from '@/components/ui/table'
  import type { Filter, ViewOption } from '$src/types/table'

  import TableToolbar from './table-toolbar.svelte'
  import TableHeadSort from './table-head-sort.svelte'
  import TablePagination from './table-pagination.svelte'
  import type { AnyPlugins } from 'svelte-headless-table/plugins'
  import { Sparkle } from 'lucide-svelte'
  import { writable } from 'svelte/store'
  import { cn } from '@/utils'

  export let tableModel: TableViewModel<any, AnyPlugins>
  export let filters: Filter[] = []
  export let viewOption: ViewOption

  const { headerRows, pageRows, tableAttrs, tableBodyAttrs } = tableModel

  const showHeadSparkles = writable(false)
</script>

<div class="space-y-3">
  <TableToolbar {tableModel} {filters} {viewOption} {showHeadSparkles} />
  <div class="rounded-md border px-3">
    <Table.Root {...$tableAttrs}>
      <Table.Header>
        {#each $headerRows as headerRow}
          <Subscribe rowAttrs={headerRow.attrs()}>
            <Table.Row>
              {#each headerRow.cells as cell (cell.id)}
                {@const FP = cell.state?.flatColumns.find((e) => e.id === cell.id)?.plugins}
                <Subscribe attrs={cell.attrs()} let:attrs props={cell.props()} let:props>
                  <Table.Head {...attrs}>
                    <div class="relative">
                      {#if cell.id !== 'actions'}
                        <!--这行代码的数据查询可能有问题，但目前来说只要设置了filter: { exclude: true }都能正常运行-->
                        {#if !FP}
                          <Sparkle
                            class={cn(
                              'absolute left-[-8px] top-[-3px] h-3 w-3 transition-opacity duration-300',
                              $showHeadSparkles ? 'opacity-100' : 'opacity-0'
                            )}
                          ></Sparkle>
                        {/if}
                        <TableHeadSort sort={props.sort}>
                          <Render of={cell.render()} />
                        </TableHeadSort>
                      {:else}
                        <Render of={cell.render()} />
                      {/if}
                    </div>
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
