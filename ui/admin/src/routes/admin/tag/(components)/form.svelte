<script lang="ts">
  import { Button } from '@/components/ui/button/index.js'
  import * as Dialog from '@/components/ui/dialog/index.js'
  import { Input } from '@/components/ui/input/index.js'
  import { Label } from '@/components/ui/label/index.js'
  import { cn } from '@/utils'
  import type { Tag, TagBody } from '$src/types/stream'
  import { tableData, formOpen, formData, getDefaultFormData } from '../(data)/data'
  import { fet } from '@/helpers/fetch'
  import Badge from '@/components/ui/badge/badge.svelte'

  let isCheck = false

  $: inputClass = isCheck && $formData.name === '' ? 'text-red-500' : undefined

  function save() {
    if ($formData.name === '') {
      isCheck = true
      return
    }

    if ($formData.id === 0) {
      fet.post('tag', $formData).then((res) => {
        if (res.ok) {
          $tableData = [<Tag>res.data, ...$tableData]
          $formData = getDefaultFormData()
          $formOpen = false
        }
      })
    } else {
      let form: TagBody = {
        name: $formData.name,
        color: $formData.color
      }
      fet.put(`tag/${$formData.id}`, form).then((res) => {
        if (res.ok) {
          $tableData = $tableData.map((v) => {
            if (v.id === $formData.id) {
              return $formData
            }
            return v
          })
          $formData = getDefaultFormData()
          $formOpen = false
        }
      })
    }
  }
</script>

<Dialog.Root bind:open={$formOpen}>
  <Dialog.Content class="sm:max-w-[425px]">
    <!-- https://github.com/huntabyte/bits-ui/issues/427#issuecomment-2025696636-->
    <!-- svelte-ignore a11y-autofocus -->
    <input class="fixed left-0 top-0 h-0 w-0" type="checkbox" autofocus={true} />
    <Dialog.Header>
      <Dialog.Title>{$formData.id === 0 ? 'New' : 'Edit'} Tag</Dialog.Title>
    </Dialog.Header>
    <div class="">
      <Label for="name" class={cn('text-right', inputClass)}>Name:</Label>
      <Input id="name" bind:value={$formData.name} autocomplete="off" />
      <Label for="color">Color:</Label>
      <div class="flex items-center space-x-2">
        <Input id="color" bind:value={$formData.color} autocomplete="off" />
        <input type="color" bind:value={$formData.color} class="border-0" />
      </div>
      <Label for="color">Preview:</Label>
      <div>
        <Badge class="ml-1" style="background-color:{$formData.color};color:white"
          >{$formData.name}</Badge
        >
      </div>
    </div>
    <Button type="submit" on:click={save}>Save</Button>
  </Dialog.Content>
</Dialog.Root>
