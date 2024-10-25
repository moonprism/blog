<script lang="ts">
  import { goto } from '$app/navigation'
  import Button from '@/components/ui/button/button.svelte'
  import { getJwtInfo, removeJwt } from '@/helpers/jwt'
  import { base } from '$app/paths'
  import { appInfo, getRealSrc } from './(data)/data'

  import * as Form from '@/components/ui/form'
  import { formSchema, type FormSchema } from './(data)/schema'
  import { superForm, defaults } from 'sveltekit-superforms'
  import { zod, zodClient } from 'sveltekit-superforms/adapters'
  import { Input } from '@/components/ui/input/index.js'
  import { fet, isRequestIn } from '@/helpers/fetch'
  import { LoaderCircle, Eraser } from 'lucide-svelte'
  import FormImageFlow from '@/components/blocks/cell/form-image-flow.svelte'
  import { writable } from 'svelte/store'
  import { Skeleton } from '$lib/components/ui/skeleton'
  import toast from '$lib/helpers/toast'
  import Textarea from '@/components/ui/textarea/textarea.svelte'
  import { Slider } from '$lib/components/ui/slider'

  const form = superForm(defaults(zod(formSchema)), {
    validators: zodClient(formSchema),
    SPA: true,
    onUpdated({ form }) {
      if (form.valid) {
        const body = {
          title: $vform.title,
          background: $vform.background,
          marginBottom: $vform.marginBottom
        }
        fet.post('settings', body).then((res) => {
          if (res.ok) {
            toast.success('更新成功')
          }
        })
      }
    },
    resetForm: false
  })
  const { form: vform, enhance } = form

  $vform = appInfo

  let isOpenImageFlow = writable(false)

  function setBgDefault() {
    $vform.background = `background: linear-gradient( 180deg, rgba(238, 174, 202, 1) 0%, rgba(148, 187, 233, 1) 100%);`
  }

  function setBg(selectImage: string) {
    $vform.background = `background-image: url(${getRealSrc(selectImage)});background-size: cover;`
  }

  let marginBottoms = [$vform.marginBottom]

  $: $vform.marginBottom = marginBottoms[0]
</script>

<FormImageFlow
  open={isOpenImageFlow}
  callback={(v) => {
    setBg(v.key)
  }}
></FormImageFlow>

<div class="container mx-auto mt-4">
  <div class="my-2 text-sm text-muted-foreground">我从来没有觉得写代码开心过。</div>
  <div class="my-4">
    <form method="POST" use:enhance class="space-y-3">
      <Form.Field {form} name="title">
        <Form.Control let:attrs>
          <div class="flex w-[330px] flex-row items-center">
            <Form.Label class="min-w-[80px]">站点标题</Form.Label>
            <Input {...attrs} bind:value={$vform.title} autocomplete="off" />
          </div>
        </Form.Control>
        <Form.FieldErrors />
      </Form.Field>

      <Form.Field {form} name="background">
        <Form.Control let:attrs>
          <div class="flex w-[480px] flex-row items-center">
            <Form.Label class="min-w-[80px]">背景</Form.Label>
            <Textarea
              {...attrs}
              bind:value={$vform.background}
              autocomplete="off"
              rows={1}
              class="min-h-[48px] px-2 py-1"
            />
            <Button variant="link" class="group" on:click={setBgDefault}>
              <Eraser
                class="h-4 w-4 cursor-pointer text-muted-foreground/70 group-hover:text-muted-foreground"
              ></Eraser>
            </Button>
            <Button
              class="mx-1"
              on:click={() => ($isOpenImageFlow = true)}
              variant="outline"
              size="sm">选择图片</Button
            >
          </div>
          <div
            style={$vform.background}
            class="preview flex h-[270px] w-[480px] flex-col items-center overflow-auto"
          >
            <div
              class="mt-4 flex h-[25px] w-1/2 items-center justify-center space-x-3 rounded-[2px] bg-background p-1 text-[11px]"
            >
              <Skeleton class="h-[10px] w-[27px]" />
              <Skeleton class="h-[10px] w-[27px]" />
              <Skeleton class="h-[10px] w-[27px]" />
              <Skeleton class="h-[10px] w-[27px]" />
            </div>
            <div
              class="mt-3 w-1/2 space-y-1 rounded-[2px] bg-background p-3"
              style="margin-bottom: {marginBottoms[0] / 2.5}px;"
            >
              <Skeleton class="h-[10px] w-[60px]" />
              <Skeleton class="h-[5px] w-[90px]" />
              <Skeleton class="h-[105px] w-full" />
              <Skeleton class="h-[6px] w-[66px]" />
              {#each Array(3) as _, index}
                <div class="h-1"></div>
                <Skeleton class="h-[10px] w-[40px]" />
                <Skeleton class="h-[5px] w-[80px]" />
                <Skeleton class="h-[6px] w-full" />
                <Skeleton class="h-[6px] w-[170px]" />
              {/each}
            </div>
          </div>
        </Form.Control>
        <Form.FieldErrors />
      </Form.Field>

      <Form.Field {form} name="marginBottom">
        <Form.Control let:attrs>
          <div class="flex w-[480px] flex-row items-center">
            <Form.Label class="min-w-[80px]">下边距</Form.Label>
            <Slider bind:value={marginBottoms} max={1000} step={1} />
            <span class="min-w-[80px] text-right text-sm text-muted-foreground"
              >{marginBottoms[0]}px</span
            >
          </div>
        </Form.Control>
        <Form.FieldErrors />
      </Form.Field>

      <Form.Button size="sm" variant="outline">
        {#if $isRequestIn}
          <LoaderCircle class="h-4 w-4 animate-spin" />
        {:else}
          更新
        {/if}
      </Form.Button>
    </form>
  </div>
  <div class="mt-7 text-sm text-secondary-foreground">
    <!--todo-->
    <p>上次登陆时间：{appInfo.lastLoginTime}</p>
    <Button
      size="sm"
      class="mt-1 h-7 rounded-sm"
      on:click={() => {
        removeJwt()
        goto(`${base}/login`)
      }}>退出登陆</Button
    >
  </div>
</div>

<style>
  .preview::-webkit-scrollbar {
    width: 3px;
  }
</style>
