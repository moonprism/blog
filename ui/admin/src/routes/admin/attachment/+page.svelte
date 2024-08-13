<script lang="ts">
  import Table from './(components)/table.svelte'
  import DataActionButton from '@/components/blocks/buttons/data-action-button.svelte'
  import { openForm } from './(data)/data'
  function add() {
    openForm()
  }
</script>

<div class="container mx-auto">
  <DataActionButton text="Upload" onClick={add} />
  <Table />
</div>
