<script lang="ts">
  import DataActionButton from '@/components/blocks/cell/data-action-button.svelte'
  import Table from './(components)/table.svelte'
  import { openForm } from './(data)/data'
</script>

<div class="container mx-auto">
  <DataActionButton text="New Article" onClick={() => openForm()} />
  <Table />
</div>
