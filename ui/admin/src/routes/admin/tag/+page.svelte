<script lang="ts">
  import Table from './(components)/table.svelte'
  import DataActionButton from '@/components/blocks/cell/data-action-button.svelte'
  import { openForm } from './(data)/data'
  import Button from '@/components/ui/button/button.svelte'
  import { Undo } from 'lucide-svelte'
  import { goto } from '$app/navigation'
  import { base } from '$app/paths'
</script>

<div class="container mx-auto">
  <Button
    variant="ghost"
    size="icon"
    on:click={() => {
      goto(`${base}/admin/article`)
    }}
  >
    <Undo class="h-4 w-4" />
  </Button>
  <DataActionButton text="New Tag" onClick={() => openForm()} />
  <Table />
</div>
