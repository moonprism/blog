<script lang="ts">
  import { Button } from '$lib/components/ui/button/index.js';
  import * as Card from '$lib/components/ui/card/index.js';
  import { Input } from '$lib/components/ui/input/index.js';
  import { Label } from '$lib/components/ui/label/index.js';
  import { IdCard } from 'svelte-radix';

  import type { LoginRequest } from '$src/types/stream';
  import { fet, setJwt } from '@/helpers/fetch';

  let loginData: LoginRequest = {
    username: '',
    password: ''
  };
  
  function handleLogin() {
    fet('login', 'post', loginData).then(response => {
      if(response.ok) {
        response.json().then(data => {
          setJwt(data.token)
        })
      }
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
        <Label for="email">Name</Label>
        <Input id="email" type="text" bind:value={loginData.username} required />
      </div>
      <div class="grid gap-2">
        <Label for="password">Password</Label>
        <Input id="password" type="password" bind:value={loginData.password} required />
      </div>
    </Card.Content>
    <Card.Footer>
      <Button class="w-full" on:click={handleLogin}>Log in</Button>
    </Card.Footer>
  </Card.Root>
</div>
