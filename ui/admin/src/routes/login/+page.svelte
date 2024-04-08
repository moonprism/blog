<script lang="ts">
  import { Button } from '$lib/components/ui/button/index.js';
  import * as Card from '$lib/components/ui/card/index.js';
  import { Input } from '$lib/components/ui/input/index.js';
  import { Label } from '$lib/components/ui/label/index.js';

  import { Avatar } from 'svelte-radix';
</script>

<div class="mt-16 flex w-full items-center justify-center">
  <Card.Root class="w-full max-w-sm">
    <Card.Header>
      <Card.Title class="text-2xl">Admin</Card.Title>
    </Card.Header>
    <Card.Content class="grid gap-4">
      <div class="grid gap-2">
        <Label for="email">Name</Label>
        <Input id="email" type="text" required />
      </div>
      <div class="grid gap-2">
        <Label for="password">Password</Label>
        <Input id="password" type="password" required />
      </div>
    </Card.Content>
    <Card.Footer>
      <Button class="w-full">Log in</Button>
    </Card.Footer>
  </Card.Root>
</div>
