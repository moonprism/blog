<script lang="ts">
  import Form from './(components)/form.svelte'
  import DataActionButton from '@/components/blocks/buttons/data-action-button.svelte'
  import Table from './(components)/table.svelte'
  import { openForm } from './(data)/data'
  function add() {
    openForm()
  }
</script>

<div class="container mx-auto">
  <DataActionButton text="New Article" onClick={add} />
  <Form />
  <Table />
</div>
