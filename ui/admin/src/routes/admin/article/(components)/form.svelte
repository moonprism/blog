<script lang="ts">
  import { Button } from '@/components/ui/button/index.js'
  import * as Dialog from '@/components/ui/dialog/index.js'
  import { Input } from '@/components/ui/input/index.js'
  import type { Article } from '$src/types/stream'
  import { tableData, closeForm, statuses } from '../(data)/data'
  import { fet, isReuqestIn } from '@/helpers/fetch'

  import * as Form from '@/components/ui/form'
  import { formSchema, type FormSchema } from '../(data)/schema'
  import { superForm, defaults } from 'sveltekit-superforms'
  import { zod, zodClient } from 'sveltekit-superforms/adapters'
  import { LoaderCircle } from 'lucide-svelte'

  import * as Select from '@/components/ui/select'
  import Textarea from '@/components/ui/textarea/textarea.svelte'
  import Checkbox from '@/components/ui/checkbox/checkbox.svelte'

  import { tableData as tagTableData } from '../../tag/(data)/data'
  import type { Writable } from 'svelte/store'

  const form = superForm(defaults(zod(formSchema)), {
    validators: zodClient(formSchema),
    SPA: true,
    onUpdated({ form }) {
      if (form.valid) {
        save()
      }
    },
    resetForm: false
  })

  const { form: vform, enhance } = form

  export let formData: Article
  export let formOpen: Writable<boolean>

  const isCreate = formData.id === 0

  $: $vform = formData

  function save() {
    const body = {
      title: $vform.title,
      status: $vform.status,
      image: $vform.image,
      summary: $vform.summary,
      tags: $vform.tags
    }
    if (isCreate) {
      fet.post('article', body).then((res) => {
        if (res.ok) {
          $tableData = [<Article>res.data, ...$tableData]
          closeForm()
        }
      })
    } else {
      fet.put(`article/${formData.id}`, body).then((res) => {
        if (res.ok) {
          formData = <Article>$vform
          formData.updated = Date.parse(new Date().toString()) / 1000
          $tableData[$tableData.findIndex((v) => v.id === formData.id)] = formData
          closeForm()
        }
      })
    }
  }
</script>

<Dialog.Root bind:open={$formOpen}>
  <Dialog.Content class="sm:max-w-[450px]">
    <!-- https://github.com/huntabyte/bits-ui/issues/427#issuecomment-2025696636-->
    <!-- svelte-ignore a11y-autofocus -->
    <input class="fixed left-0 top-0 h-0 w-0" type="checkbox" autofocus={true} />

    <Dialog.Header>
      <Dialog.Title>{isCreate ? 'New' : 'Edit'} Article</Dialog.Title>
    </Dialog.Header>
    <form method="POST" use:enhance class="space-y-2">
      <Form.Field {form} name="title">
        <Form.Control let:attrs>
          <Form.Label>Title</Form.Label>
          <Input {...attrs} bind:value={$vform.title} autocomplete="off" />
        </Form.Control>
        <Form.FieldErrors />
      </Form.Field>

      <Form.Field {form} name="status">
        <Form.Control let:attrs>
          <Form.Label>Status</Form.Label>
          <Select.Root
            selected={{
              label: $statuses.find((e) => e.id === $vform.status)?.label,
              value: $vform.status
            }}
            onSelectedChange={(v) => {
              v && ($vform.status = Number(v.value))
            }}
          >
            <Select.Trigger>
              <Select.Value placeholder="Draft" />
            </Select.Trigger>
            <Select.Content>
              {#each $statuses as status}
                <Select.Item value={status.id} label={status.label}>{status.label}</Select.Item>
              {/each}
            </Select.Content>
          </Select.Root>
        </Form.Control>
        <Form.FieldErrors />
      </Form.Field>

      <Form.Field {form} name="tags">
        <Form.Control let:attrs>
          <Form.Label>Tags</Form.Label>
          <div>
            {#each $tagTableData as tag}
              {@const checked = $vform.tags.some((e) => e.id === tag.id)}
              <div class="mr-2 inline-block">
                <Form.Control let:attrs>
                  <Checkbox
                    {...attrs}
                    {checked}
                    onCheckedChange={(v) => {
                      if (v) {
                        // add
                        $vform.tags = [...$vform.tags, tag]
                      } else {
                        // remove
                        $vform.tags = $vform.tags.filter((e) => e.id !== tag.id)
                      }
                    }}
                  />
                  <Form.Label class="text-sm font-normal">
                    {tag.name}
                  </Form.Label>
                </Form.Control>
              </div>
            {/each}
          </div>
        </Form.Control>
        <Form.FieldErrors />
      </Form.Field>

      <Form.Field {form} name="image">
        <Form.Control let:attrs>
          <Form.Label>Image</Form.Label>
          <Input {...attrs} bind:value={$vform.image} autocomplete="off" />
        </Form.Control>
        <Form.FieldErrors />
      </Form.Field>

      <Form.Field {form} name="summary">
        <Form.Control let:attrs>
          <Form.Label>Summary</Form.Label>
          <Textarea {...attrs} bind:value={$vform.summary} rows={3} />
        </Form.Control>
        <Form.FieldErrors />
      </Form.Field>

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
