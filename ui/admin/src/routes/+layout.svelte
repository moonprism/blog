<script lang="ts">
  import '../app.pcss'

  // 夜间模式
  import { ModeWatcher } from 'mode-watcher'

  // Toast组件
  import { Toaster } from '$lib/components/ui/sonner'
</script>

<ModeWatcher />

<Toaster position="top-center" />

<slot />
