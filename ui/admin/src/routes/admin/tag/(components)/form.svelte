<script lang="ts">
  import { Button } from '@/components/ui/button/index.js'
  import * as Dialog from '@/components/ui/dialog/index.js'
  import { Input } from '@/components/ui/input/index.js'
  import { Label } from '@/components/ui/label/index.js'
  import type { Tag, TagBody } from '$src/types/stream'
  import { tableData, closeForm } from '../(data)/data'
  import { fet, isReuqestIn } from '@/helpers/fetch'
  import Badge from '@/components/ui/badge/badge.svelte'

  import * as Form from '@/components/ui/form'
  import { formSchema, type FormSchema } from '../(data)/schema'
  import { superForm, defaults } from 'sveltekit-superforms'
  import { zod, zodClient } from 'sveltekit-superforms/adapters'
  import { LoaderCircle } from 'lucide-svelte'
  import type { ReadOrWritable } from 'svelte-headless-table'

  const form = superForm(defaults(zod(formSchema)), {
    validators: zodClient(formSchema),
    SPA: true,
    onUpdate({ form }) {
      if (form.valid) {
        save()
      }
    },
    resetForm: false
  })

  const { form: formValidData, enhance } = form

  export let formData: Tag
  export let formOpen: ReadOrWritable<boolean>

  const isCreate = formData.id === 0

  $: $formValidData = formData

  function save() {
    const body: TagBody = {
      name: formData.name,
      color: formData.color
    }
    if (isCreate) {
      fet.post('tag', body).then((res) => {
        if (res.ok) {
          $tableData = [<Tag>res.data, ...$tableData]
          closeForm()
        }
      })
    } else {
      fet.put(`tag/${formData.id}`, body).then((res) => {
        if (res.ok) {
          formData.updated = Date.parse(new Date().toString()) / 1000
          $tableData[$tableData.findIndex((v) => v.id === formData.id)] = formData
          closeForm()
        }
      })
    }
  }
</script>

<Dialog.Root bind:open={$formOpen}>
  <Dialog.Content class="sm:max-w-[330px]">
    <!-- https://github.com/huntabyte/bits-ui/issues/427#issuecomment-2025696636-->
    <!-- svelte-ignore a11y-autofocus -->
    <input class="fixed left-0 top-0 h-0 w-0" type="checkbox" autofocus={true} />

    <Dialog.Header>
      <Dialog.Title>{isCreate ? 'New' : 'Edit'} Tag</Dialog.Title>
    </Dialog.Header>
    <form method="POST" use:enhance class="space-y-2">
      <Form.Field {form} name="name">
        <Form.Control let:attrs>
          <Form.Label>Name</Form.Label>
          <Input {...attrs} bind:value={formData.name} autocomplete="off" />
        </Form.Control>
        <Form.FieldErrors />
      </Form.Field>
      <Form.Field {form} name="color">
        <Form.Control let:attrs>
          <Form.Label>Color</Form.Label>
          <div class="flex items-center space-x-2">
            <Input {...attrs} bind:value={formData.color} autocomplete="off" />
            <input type="color" bind:value={formData.color} class="border-0" />
          </div>
        </Form.Control>
        <Form.FieldErrors />
      </Form.Field>
      <Label>Preview</Label>
      <div>
        <Badge class="mb-1" style="background-color:{formData.color};color:white"
          >{formData.name}</Badge
        >
      </div>
      {#if $isReuqestIn}
        <Button class="w-full">
          <LoaderCircle class="mr-2 h-4 w-4 animate-spin" />
        </Button>
      {:else}
        <Form.Button class="w-full">Save</Form.Button>
      {/if}
    </form>
  </Dialog.Content>
</Dialog.Root>
