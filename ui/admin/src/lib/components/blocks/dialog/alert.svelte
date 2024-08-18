<script lang="ts">
  import * as AlertDialog from '$lib/components/ui/alert-dialog/index.js'
  import type { ReadOrWritable } from '$src/types/store'

  export let open: ReadOrWritable<boolean>
  export let title: ReadOrWritable<string>
  export let description: ReadOrWritable<string>
  export let confirm: ReadOrWritable<boolean>
</script>

<AlertDialog.Root bind:open={$open}>
  <AlertDialog.Content class="md:w-[400px]">
    <AlertDialog.Header>
      <AlertDialog.Title>{$title ? $title : 'Are you absolutely sure?'}</AlertDialog.Title>
      <AlertDialog.Description>
        {$description}
      </AlertDialog.Description>
    </AlertDialog.Header>
    <AlertDialog.Footer>
      <AlertDialog.Cancel>Cancel</AlertDialog.Cancel>
      <AlertDialog.Action on:click={() => ($confirm = true)}>Continue</AlertDialog.Action>
    </AlertDialog.Footer>
  </AlertDialog.Content>
</AlertDialog.Root>
