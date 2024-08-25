<script lang="ts">
  import type { TableViewModel } from 'svelte-headless-table'
  import Button from '@/components/ui/button/button.svelte'
  import * as Select from '@/components/ui/select'
  import { ChevronLeft, ChevronRight, ChevronsLeft, ChevronsRight } from 'lucide-svelte'

  export let tableModel: TableViewModel<any>

  const { pluginStates } = tableModel

  const { hasNextPage, hasPreviousPage, pageIndex, pageCount, pageSize } = pluginStates.page
</script>

<div class="flex items-center justify-between px-2">
  <div class="flex-1 font-mono text-sm text-muted-foreground">
  <!-- issue: 服务器模式下该插件不保存总数 -->
    <!-- {$rows.length} row(s) -->
  </div>
  <div class="flex items-center">
    <div class="flex items-center space-x-1">
      <p class="text-sm font-medium">Rows per page</p>
      <Select.Root
        onSelectedChange={(selected) => pageSize.set(Number(selected?.value))}
        selected={{ value: 10, label: '10' }}
      >
        <Select.Trigger class="h-8 w-16">
          <Select.Value />
        </Select.Trigger>
        <Select.Content>
          <Select.Item value="10">10</Select.Item>
          <Select.Item value="20">20</Select.Item>
          <Select.Item value="30">30</Select.Item>
          <Select.Item value="50">50</Select.Item>
        </Select.Content>
      </Select.Root>
    </div>
    <div class="flex w-[100px] items-center justify-center text-sm font-medium">
      Page {$pageIndex + 1} of {$pageCount}
    </div>
    <div class="flex items-center space-x-2">
      <Button
        variant="outline"
        class="hidden h-8 w-8 p-0 lg:flex"
        on:click={() => ($pageIndex = 0)}
        disabled={!$hasPreviousPage}
      >
        <span class="sr-only">Go to first page</span>
        <ChevronsLeft size={15} />
      </Button>
      <Button
        variant="outline"
        class="h-8 w-8 p-0"
        on:click={() => ($pageIndex = $pageIndex - 1)}
        disabled={!$hasPreviousPage}
      >
        <span class="sr-only">Go to previous page</span>
        <ChevronLeft size={15} />
      </Button>
      <Button
        variant="outline"
        class="h-8 w-8 p-0"
        disabled={!$hasNextPage}
        on:click={() => ($pageIndex = $pageIndex + 1)}
      >
        <span class="sr-only">Go to next page</span>
        <ChevronRight size={15} />
      </Button>
      <Button
        variant="outline"
        class="hidden h-8 w-8 p-0 lg:flex"
        disabled={!$hasNextPage}
        on:click={() => ($pageIndex = $pageCount - 1)}
      >
        <span class="sr-only">Go to last page</span>
        <ChevronsRight size={15} />
      </Button>
    </div>
  </div>
</div>
