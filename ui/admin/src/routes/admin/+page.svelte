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

  const form = superForm(defaults(zod(formSchema)), {
    validators: zodClient(formSchema),
    SPA: true,
    onUpdated({ form }) {
      if (form.valid) {
        const body = {
          title: $vform.title,
          background: $vform.background
        }
        fet.post('settings', body).then((res) => {
          if (res.ok) {
            toast.success("更新成功")
          }
        })
      }
    },
    resetForm: false
  })
  const { form: vform, enhance } = form

  $vform = appInfo

  let isOpenImageFlow = writable(false)

  $: backgroundStyle =
    $vform.background !== ''
      ? `background-image: url(${getRealSrc($vform.background)});background-size: cover;`
      : `background: linear-gradient(
    180deg,
    rgba(238, 174, 202, 1) 0%,
    rgba(148, 187, 233, 1) 100%
  );`
</script>

<FormImageFlow
  open={isOpenImageFlow}
  callback={(v) => {
    $vform.background = v.key
  }}
></FormImageFlow>

<div class="container mx-auto mt-4">
  <div class="my-2">我从来没有觉得写代码开心过。</div>
  <div class="my-4">
    <form method="POST" use:enhance class="space-y-1">
      <Form.Field {form} name="title">
        <Form.Control let:attrs>
          <div class="flex w-[330px] flex-row items-center">
            <Form.Label class="min-w-[80px]">站点标题</Form.Label>
            <Input {...attrs} bind:value={$vform.title} autocomplete="off" />
          </div>
        </Form.Control>
        <Form.Control let:attrs>
          <div class="flex w-[480px] flex-row items-center">
            <Form.Label class="min-w-[80px]">背景</Form.Label>
            <Input
              {...attrs}
              disabled
              bind:value={$vform.background}
              autocomplete="off"
              placeholder="Default"
            />
            <Button variant="link" class="group" on:click={() => ($vform.background = '')}>
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
          <div style={backgroundStyle} class="flex h-[270px] w-[480px] flex-col items-center">
            <div
              class="mt-4 flex h-[25px] w-1/2 items-center justify-center space-x-3 rounded-[2px] bg-background text-[11px]"
            >
              <Skeleton class="h-[10px] w-[27px]" />
              <Skeleton class="h-[10px] w-[27px]" />
              <Skeleton class="h-[10px] w-[27px]" />
              <Skeleton class="h-[10px] w-[27px]" />
            </div>
            <div class="mt-3 h-full w-1/2 space-y-1 rounded-t-[2px] bg-background px-3 pt-3">
              <Skeleton class="h-[10px] w-[60px]" />
              <Skeleton class="h-[5px] w-[90px]" />
              <Skeleton class="h-[105px] w-full" />
              <Skeleton class="h-[6px] w-[66px]" />
              <div class="h-1"></div>
              <Skeleton class="h-[10px] w-[40px]" />
              <Skeleton class="h-[5px] w-[80px]" />
              <Skeleton class="h-[6px] w-full" />
              <Skeleton class="h-[6px] w-[170px]" />
            </div>
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
  <div class="text-sm text-secondary-foreground mt-5">
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
