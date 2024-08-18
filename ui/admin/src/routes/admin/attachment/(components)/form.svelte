<script lang="ts">
  import { Button } from '@/components/ui/button/index.js'
  import * as Dialog from '@/components/ui/dialog/index.js'
  import { Input } from '@/components/ui/input/index.js'
  import { tableData, closeForm } from '../(data)/data'
  import * as Popover from '@/components/ui/popover'
  import { fet, getRealSrc, isMockMode, isReuqestIn } from '@/helpers/fetch'

  import * as Form from '@/components/ui/form'
  import { formSchema, type FormSchema } from '../(data)/schema'
  import { superForm, defaults } from 'sveltekit-superforms'
  import { zod, zodClient } from 'sveltekit-superforms/adapters'
  import { LoaderCircle, Maximize, Trash2 } from 'lucide-svelte'
  import type { Attachment } from '$src/types/stream'
  import Textarea from '@/components/ui/textarea/textarea.svelte'
  import type { Writable } from 'svelte/store'

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

  export let formData: Attachment
  export let formOpen: Writable<boolean>

  const isCreate = formData.id === 0

  $: $vform = formData

  function save() {
    if (isMockMode) {
      formData.link = previewUrl
    }
    if (isCreate) {
      fet.post('attachment', formData).then((res) => {
        if (res.ok) {
          $tableData = [<Attachment>res.data, ...$tableData]
          closeForm()
        }
      })
    } else {
      fet.put(`attachment/${formData.id}`, { summary: $vform.summary }).then((res) => {
        if (res.ok) {
          formData = <Attachment>$vform
          formData.updated = Date.parse(new Date().toString()) / 1000
          $tableData[$tableData.findIndex((v) => v.id === formData.id)] = formData
          closeForm()
        }
      })
    }
  }

  let previewUrl: string = getRealSrc(formData.link)
  // 图片预览
  function handleFileChange(event: Event) {
    const input = event.target as HTMLInputElement
    const file = input.files?.[0]

    if (file) {
      $vform.link = file.name
      const reader = new FileReader()
      reader.onload = (e) => {
        previewUrl = e.target?.result as string
      }
      reader.readAsDataURL(file)
    } else {
      $vform.link = previewUrl = ''
    }
  }

  function del() {
    fet.delete(`attachment/${formData.id}`).then((res) => {
      if (res.ok) {
        $tableData = $tableData.filter((v) => v.id !== formData.id)
        closeForm()
      }
    })
  }

  let isDel = false

  let isPreview = false
</script>

<Dialog.Root bind:open={isPreview}>
  <Dialog.Content class="max-w-screen z-[51] h-full cursor-pointer">
    <!-- https://github.com/huntabyte/bits-ui/issues/427#issuecomment-2025696636-->
    <!-- svelte-ignore a11y-autofocus -->
    <input class="fixed left-0 top-0 h-0 w-0" type="checkbox" autofocus={true} />
    <!-- svelte-ignore a11y-click-events-have-key-events a11y-no-static-element-interactions -->
    <div class="flex h-full w-full justify-center p-3" on:click={() => (isPreview = false)}>
      <img
        src={previewUrl}
        alt="preview"
        class="scale-down h-auto max-h-[92vh] w-auto object-contain"
      />
    </div>
  </Dialog.Content>
</Dialog.Root>

<Dialog.Root bind:open={$formOpen}>
  <Dialog.Content class="sm:max-w-[460px]">
    <!-- https://github.com/huntabyte/bits-ui/issues/427#issuecomment-2025696636-->
    <!-- svelte-ignore a11y-autofocus -->
    <input class="fixed left-0 top-0 h-0 w-0" type="checkbox" autofocus={true} />

    <Dialog.Header>
      <Dialog.Title>{isCreate ? '上传' : '更新'}</Dialog.Title>
    </Dialog.Header>
    <form method="POST" use:enhance class="space-y-2" enctype="multipart/form-data">
      <Form.Field {form} name="link">
        <Form.Control let:attrs>
          <Form.Label>
            <div>
              文件
              {#if !isCreate}
                <span class="mx-1 text-xs text-muted-foreground">{$vform.link}</span>
              {/if}
            </div>
          </Form.Label>
          {#if previewUrl !== ''}
            <!-- svelte-ignore a11y-click-events-have-key-events a11y-no-static-element-interactions -->
            <div
              on:click={() => (isPreview = true)}
              class="group relative inline-block cursor-pointer hover:shadow"
            >
              <img src={previewUrl} alt="preview" class="block h-auto max-h-[300px] w-full" />
              <div
                class="absolute right-0 top-0 hidden bg-white/30 backdrop-blur-sm group-hover:block"
              >
                <Maximize class="m-1 h-4 w-4"></Maximize>
              </div>
            </div>
          {/if}
          <div class="mr-1 flex items-center justify-between">
            <Input
              disabled={!isCreate}
              {...attrs}
              type="file"
              on:change={handleFileChange}
              class="w-[240px]"
              accept="image/*"
            />
            {#if !isCreate}
              <Popover.Root bind:open={isDel}>
                <Popover.Trigger>
                  <Trash2 class="h-4 w-4 duration-100 hover:text-red-500"></Trash2>
                </Popover.Trigger>
                <Popover.Content class="w-56">
                  <div class="grid gap-3">
                    <h4 class="font-medium leading-none">完全删除该数据？</h4>
                    <div class="flex justify-end gap-1">
                      <Button variant="outline" size="sm" on:click={() => (isDel = false)}
                        >再想想</Button
                      >
                      <Button size="sm" on:click={del}>确认</Button>
                    </div>
                  </div>
                </Popover.Content>
              </Popover.Root>
            {/if}
          </div>
        </Form.Control>
        <Form.FieldErrors />
      </Form.Field>
      <Form.Field {form} name="summary">
        <Form.Control let:attrs>
          <Form.Label>描述</Form.Label>
          <Textarea {...attrs} bind:value={$vform.summary} rows={2} />
        </Form.Control>
        <Form.FieldErrors />
      </Form.Field>
      {#if $isReuqestIn}
        <Button class="w-full">
          <LoaderCircle class="mr-2 h-4 w-4 animate-spin" />
        </Button>
      {:else}
        <Form.Button class="w-full">保存</Form.Button>
      {/if}
    </form>
  </Dialog.Content>
</Dialog.Root>
