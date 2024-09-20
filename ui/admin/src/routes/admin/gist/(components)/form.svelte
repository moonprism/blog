<script lang="ts">
  import { Button } from '@/components/ui/button/index.js'
  import * as Dialog from '@/components/ui/dialog/index.js'
  import { Input } from '@/components/ui/input/index.js'
  import type { Gist, GistBody } from '$src/types/stream'
  import { tableData, closeForm, langItems } from '../(data)/data'
  import { fet, isRequestIn } from '@/helpers/fetch'

  import * as Form from '@/components/ui/form'
  import { formSchema, type FormSchema } from '../(data)/schema'
  import { superForm, defaults } from 'sveltekit-superforms'
  import { zod, zodClient } from 'sveltekit-superforms/adapters'
  import { LoaderCircle } from 'lucide-svelte'
  import type { Writable } from 'svelte/store'
  import Editor from './editor.svelte'
  import Combobox from '@/components/blocks/cell/combobox.svelte'

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

  const { form: vform, enhance } = form

  export let formData: Gist
  export let formOpen: Writable<boolean>

  const isCreate = formData.id === 0

  let editor: Editor

  $vform = formData

  async function save() {
    if ($vform.lang !== 'md') {
      const lines = $vform.content.trim().split(/\r?\n/)
      lines.pop()
      lines.shift()
      $vform.content = lines.join('\n')
    }
    const body: GistBody = {
      lang: $vform.lang,
      title: $vform.title,
      content: $vform.content,
      html: await editor.render()
    }
    if (isCreate) {
      fet.post('gist', body).then((res) => {
        if (res.ok) {
          $tableData = [<Gist>res.data, ...$tableData]
          closeForm()
        }
      })
    } else {
      fet.put(`gist/${formData.id}`, body).then((res) => {
        if (res.ok) {
          formData = <Gist>$vform
          formData.html = body.html
          formData.updated = Date.parse(new Date().toString()) / 1000
          $tableData[$tableData.findIndex((v) => v.id === formData.id)] = formData
          closeForm()
        }
      })
    }
  }

  let lang = isCreate ? 'md' : $vform.lang

  function applyLang() {
    const lines = $vform.content.trim().split(/\r?\n/)
    const header = `\`\`\`${lang}`
    if (lines[0].startsWith('```')) {
      if (lang !== 'md') {
        lines[0] = header
      }
    } else if (lang !== 'md') {
      lines.unshift(header)
    }
    if (lines[lines.length - 1] !== '```' && lang !== 'md') {
      lines.push('```')
    }
    $vform.lang = lang
    $vform.content = lines.join('\n')
  }

  $: {
    lang
    applyLang()
  }
</script>

<Dialog.Root bind:open={$formOpen} closeOnEscape={false}>
  <Dialog.Content class="sm:max-w-[540px]">
    <!-- https://github.com/huntabyte/bits-ui/issues/427#issuecomment-2025696636-->
    <!-- svelte-ignore a11y-autofocus -->
    <input class="fixed left-0 top-0 h-0 w-0" type="checkbox" autofocus={true} />

    <Dialog.Header>
      <Dialog.Title>{isCreate ? 'New' : 'Edit'} Gist</Dialog.Title>
    </Dialog.Header>
    <form method="POST" use:enhance class="space-y-2 w-full overflow-auto">
      <div class="flex flex-1 space-x-1">
        <Form.Field {form} name="title" class="w-full p-1">
          <Form.Control let:attrs>
            <Form.Label>标题</Form.Label>
            <Input {...attrs} bind:value={$vform.title} autocomplete="off" />
          </Form.Control>
          <Form.FieldErrors />
        </Form.Field>
        <Form.Field {form} name="lang" class="p-1">
          <Form.Control let:attrs>
            <Form.Label>语言</Form.Label>
            <Combobox items={langItems} bind:value={lang}></Combobox>
          </Form.Control>
          <Form.FieldErrors />
        </Form.Field>
      </div>
      <Form.Field {form} name="content">
        <Form.Control let:attrs>
          <Editor bind:value={$vform.content} bind:this={editor}></Editor>
        </Form.Control>
        <Form.FieldErrors />
      </Form.Field>
      {#if $isRequestIn}
        <Button class="w-full">
          <LoaderCircle class="mr-2 h-4 w-4 animate-spin" />
        </Button>
      {:else}
        <Form.Button class="w-full">保存</Form.Button>
      {/if}
    </form>
  </Dialog.Content>
</Dialog.Root>
