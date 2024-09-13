<script lang="ts">
  import DataActionButton from '@/components/blocks/cell/data-action-button.svelte'
  import Table from './(components)/table.svelte'
  import { openForm } from './(data)/data'
  import { Award } from 'lucide-svelte'
  import Button from '@/components/ui/button/button.svelte'
  import { goto } from '$app/navigation'
  import * as Tooltip from '$lib/components/ui/tooltip/index.js'
</script>

<div class="container mx-auto">
  <div class="flex items-center">
    <DataActionButton text="New Article" onClick={() => openForm()} />
    <Tooltip.Root>
      <Tooltip.Trigger asChild let:builder>
        <Button
          builders={[builder]}
          variant="ghost"
          size="icon"
          on:click={() => {
            goto('/admin/tag')
          }}
        >
          <Award class="h-5 w-5"></Award>
        </Button>
      </Tooltip.Trigger>
      <Tooltip.Content>标签管理</Tooltip.Content>
    </Tooltip.Root>
  </div>
  <Table />
</div>
