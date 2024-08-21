<script lang="ts">
  import { Button } from '$lib/components/ui/button/index.js'
  import * as Card from '$lib/components/ui/card/index.js'
  import { Input } from '$lib/components/ui/input/index.js'
  import { Label } from '$lib/components/ui/label/index.js'
  import { Crown, LoaderCircle } from 'lucide-svelte'
  import Solar from './solar.svelte'

  import type { LoginRequest } from '$src/types/stream'
  import { fet, isRequestIn } from '@/helpers/fetch'
  import { getJwtInfo, setJwt } from '@/helpers/jwt'
  import toast from '$lib/helpers/toast'
  import { goto } from '$app/navigation'

  let data: LoginRequest = {
    username: '',
    password: ''
  }

  /**
   * 推荐使用sveltekit-superforms进行表单验证，该插件官网有SPA模式使用说明
   * ssr参照: https://shadcn-svelte.com/docs/components/form
   */
  let checkOn = false
  $: isShowError = checkOn && data.username == ''

  function handleLogin() {
    if (!data.username) {
      checkOn = true
      return
    }
    fet.post('login', data).then((respo) => {
      if (respo.ok) {
        setJwt(respo.data.token)
        toast.success(`登陆成功，${getJwtInfo()?.username}`)
        goto('/admin')
      }
    })
  }
</script>

<Solar />

<div class="mt-16 flex w-full items-center justify-center">
  <Card.Root class="w-full max-w-sm">
    <Card.Header>
      <Card.Title class="text-2xl"><Crown class="mb-2 mr-1 inline-block" />.admin</Card.Title>
    </Card.Header>
    <Card.Content class="grid gap-4">
      <div class="grid gap-2">
        <Label for="name">Name</Label>
        <Input id="name" type="text" bind:value={data.username} />
        {#if isShowError}
          <p class="border-l-2 border-l-red-400 pl-1 text-xs text-muted-foreground">
            用户名不能为空
          </p>
        {/if}
      </div>
      <div class="grid gap-2">
        <Label for="password">Password</Label>
        <Input id="password" type="password" bind:value={data.password} />
      </div>
    </Card.Content>
    <Card.Footer>
      {#if $isRequestIn}
        <Button class="w-full">
          <LoaderCircle class="mr-2 h-4 w-4 animate-spin" />
        </Button>
      {:else}
        <Button class="w-full" on:click={handleLogin}><span class="font-bold">Log in</span></Button>
      {/if}
    </Card.Footer>
  </Card.Root>
</div>
