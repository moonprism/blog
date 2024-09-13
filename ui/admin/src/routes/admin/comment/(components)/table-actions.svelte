<script lang="ts">
  import DotsHorizontal from 'svelte-radix/DotsHorizontal.svelte'
  import type { Comment } from '$src/types/stream.js'
  import * as DropdownMenu from '@/components/ui/dropdown-menu'
  import { Button } from '@/components/ui/button'
  import { tableData } from '../(data)/data'
  import { fet } from '@/helpers/fetch'
  import { alertDialog } from '@/components/blocks/dialog/alert'

  export let row: Comment

  function del() {
    fet.delete(`comment/${row.id}`).then((res) => {
      if (res.ok) {
        $tableData = $tableData.filter((v) => v.id !== row.id)
      }
    })
  }
</script>

<DropdownMenu.Root>
  <DropdownMenu.Trigger asChild let:builder>
    <Button
      variant="ghost"
      builders={[builder]}
      class="flex h-8 w-8 p-0 data-[state=open]:bg-muted"
    >
      <DotsHorizontal class="h-4 w-4" />
      <span class="sr-only">Open Menu</span>
    </Button>
  </DropdownMenu.Trigger>
  <DropdownMenu.Content class="w-[160px]" align="end">
    <DropdownMenu.Item
      on:click={() => {
        alertDialog(`删除评论 “ ${row.content} ”`).then(() => {
          del()
        })
      }}
    >
      删除
      <DropdownMenu.Shortcut>⌘⌫</DropdownMenu.Shortcut>
    </DropdownMenu.Item>
  </DropdownMenu.Content>
</DropdownMenu.Root>
