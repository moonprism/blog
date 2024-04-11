<script lang="ts">
  import '../app.pcss'

  // 切换夜间模式
  import { ModeWatcher, toggleMode } from 'mode-watcher'
  import { Sun, Moon } from 'svelte-radix'
  import { Button } from '$lib/components/ui/button/index.js'

  // Toast组件
  import { Toaster } from '$lib/components/ui/sonner'
</script>

<ModeWatcher />

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

<Toaster />

<slot />
