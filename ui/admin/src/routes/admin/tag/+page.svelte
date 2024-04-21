<script lang="ts">
  import Form from './(components)/form.svelte'
  import Table from './(components)/table.svelte'
  import DataActionButton from '@/components/blocks/buttons/data-action-button.svelte'
  import { defaultFormData, formData, formOpen } from './(data)/data'

  function add() {
    $formOpen = true
    $formData = defaultFormData
  }
</script>

<div class="container mx-auto">
  <DataActionButton text="New Tag" onClick={add} />
  <Form />
  <Table />
</div>
