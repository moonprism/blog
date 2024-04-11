<script lang="ts">
  import { Button } from '$lib/components/ui/button/index.js'
  import * as Card from '$lib/components/ui/card/index.js'
  import { Input } from '$lib/components/ui/input/index.js'
  import { Label } from '$lib/components/ui/label/index.js'
  import { IdCard } from 'svelte-radix'
  import Reload from 'svelte-radix/Reload.svelte'

  import type { LoginRequest } from '$src/types/stream'
  import { login } from '@/helpers/fetch'

  let data = <LoginRequest>{}
  let isLoading = false

  /**
   * 推荐使用sveltekit-superforms进行表单验证，该插件官网有SPA模式使用说明
   * ssr参照: https://shadcn-svelte.com/docs/components/form
   */
  let checkOn = false

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

<div class="mt-16 flex w-full items-center justify-center">
  <Card.Root class="w-full max-w-sm">
    <Card.Header>
      <Card.Title class="text-2xl"><IdCard class="mb-1 mr-2 inline-block" />Admin</Card.Title>
    </Card.Header>
    <Card.Content class="grid gap-4">
      <div class="grid gap-2">
        <Label for="name">Name</Label>
        <Input id="name" type="text" bind:value={data.username} />
        {#if checkOn && data.username == ''}
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
          <Reload class="mr-2 h-4 w-4 animate-spin" />
        </Button>
      {:else}
        <Button class="w-full" on:click={handleLogin}><span class="font-bold">Log in</span></Button>
      {/if}
    </Card.Footer>
  </Card.Root>
</div>
