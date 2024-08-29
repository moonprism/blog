<script lang="ts">
  import Table from './(components)/table.svelte'
  import DataActionButton from '@/components/blocks/cell/data-action-button.svelte'
  import { openForm } from './(data)/data'
</script>

<div class="container mx-auto">
  <DataActionButton text="Upload" onClick={() => openForm()} />
  <Table />
</div>
