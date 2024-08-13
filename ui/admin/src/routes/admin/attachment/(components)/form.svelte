<script lang="ts">
  import { Button } from '@/components/ui/button/index.js'
  import * as Dialog from '@/components/ui/dialog/index.js'
  import { Input } from '@/components/ui/input/index.js'
  import { tableData, closeForm } from '../(data)/data'
  import { fet, isReuqestIn } from '@/helpers/fetch'

  import * as Form from '@/components/ui/form'
  import { formSchema, type FormSchema } from '../(data)/schema'
  import { superForm, defaults } from 'sveltekit-superforms'
  import { zod, zodClient } from 'sveltekit-superforms/adapters'
  import { LoaderCircle } from 'lucide-svelte'
  import type { ReadOrWritable } from 'svelte-headless-table'
  import type { Attachment, AttachmentBody } from '$src/types/stream'
  import Textarea from '@/components/ui/textarea/textarea.svelte'
  import { PUBLIC_MOCK_MODE } from '$env/static/public'

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

  export let formData: Attachment
  export let formOpen: ReadOrWritable<boolean>

  $: $formValidData = formData

  function save() {
    if (PUBLIC_MOCK_MODE) {
      $formValidData.link = previewUrl
    }
    const body: AttachmentBody = {
      link: $formValidData.link,
      summary: $formValidData.summary
    }
    if (formData.id === 0) {
      fet.post('attachment', body).then((res) => {
        if (res.ok) {
          $tableData = [<Attachment>res.data, ...$tableData]
          closeForm()
        }
      })
    } else {
      fet.put(`attachment/${formData.id}`, body).then((res) => {
        if (res.ok) {
          if (PUBLIC_MOCK_MODE) {
            $formValidData.link = previewUrl
          }
          let newFormData = <Attachment>$formValidData
          newFormData.updated = Date.parse(new Date().toString()) / 1000
          $tableData[$tableData.findIndex((v) => v.id === formData.id)] = newFormData
          closeForm()
        }
      })
    }
  }

  export let previewUrl:string = ''
  // 图片预览
  function handleFileChange(event: Event) {
    const input = event.target as HTMLInputElement
    const file = input.files?.[0]

    if (file) {
      $formValidData.link = file.name
      const reader = new FileReader()
      reader.onload = (e) => {
        previewUrl = e.target?.result as string
      }
      reader.readAsDataURL(file)
    } else {
      previewUrl = ''
    }
  }
</script>

<Dialog.Root bind:open={$formOpen}>
  <Dialog.Content class="sm:max-w-[460px]">
    <!-- https://github.com/huntabyte/bits-ui/issues/427#issuecomment-2025696636-->
    <!-- svelte-ignore a11y-autofocus -->
    <input class="fixed left-0 top-0 h-0 w-0" type="checkbox" autofocus={true} />

    <Dialog.Header>
      <Dialog.Title>Upload</Dialog.Title>
    </Dialog.Header>
    <form method="POST" use:enhance class="space-y-2" enctype="multipart/form-data">
      <Form.Field {form} name="link">
        <Form.Control let:attrs>
          {#if previewUrl !== ''}
            <img src={previewUrl} alt="preview" class="max-h-[300px]" />
          {/if}
          <Input
            {...attrs}
            type="file"
            on:change={handleFileChange}
            class="w-[240px]"
          />
        </Form.Control>
        <Form.FieldErrors />
      </Form.Field>
      <Form.Field {form} name="summary">
        <Form.Control let:attrs>
          <Form.Label>Summary</Form.Label>
          <Textarea {...attrs} bind:value={$formValidData.summary} rows={2} />
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
