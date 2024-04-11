<script lang="ts">
  import { toggleMode } from 'mode-watcher'
  import { Sun, Moon } from 'lucide-svelte'

  import { Button } from '$lib/components/ui/button/index.js'
  import * as Card from '$lib/components/ui/card/index.js'
  import { Input } from '$lib/components/ui/input/index.js'
  import { Label } from '$lib/components/ui/label/index.js'
  import { Crown } from 'lucide-svelte';
  import { LoaderCircle } from 'lucide-svelte';

  import type { LoginRequest } from '$src/types/stream'
  import { login } from '@/helpers/fetch'


  let data: LoginRequest = {
    username: '',
    password: '',
  }
  let isLoading = false

  /**
   * 推荐使用sveltekit-superforms进行表单验证，该插件官网有SPA模式使用说明
   * ssr参照: https://shadcn-svelte.com/docs/components/form
   * 个人项目就不用了随便写写
   */
  let checkOn = false
  $: isShowError = checkOn && data.username == ''

  function handleLogin() {
    if (!data.username) {
      checkOn = true
      return
    }
    isLoading = true
    login(data).catch((err) => {
      isLoading = false
    })
  }
  
</script>

<div class="mx-auto flex justify-between px-6 py-4">
  <div></div>
  <div>
    <Button on:click={toggleMode} variant="outline" size="icon">
      <Sun
        class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
      />
      <Moon
        class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
      />
    </Button>
  </div>
</div>

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
      {#if isLoading}
        <Button class="w-full">
          <LoaderCircle class="mr-2 h-4 w-4 animate-spin" />
        </Button>
      {:else}
        <Button class="w-full" on:click={handleLogin}><span class="font-bold">Log in</span></Button>
      {/if}
    </Card.Footer>
  </Card.Root>
</div>
