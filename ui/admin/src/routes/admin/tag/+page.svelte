<script lang="ts">
  import Table from './(components)/table.svelte'
  import DataActionButton from '@/components/blocks/cell/data-action-button.svelte'
  import { openForm } from './(data)/data'
  import { Award } from 'lucide-svelte'
</script>

<div class="container mx-auto">
  <DataActionButton text="New Tag" onClick={() => openForm()} />
  <Table />
</div>
