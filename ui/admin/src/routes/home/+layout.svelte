<script lang="ts">
  import { Button } from '$lib/components/ui/button/index.js'
  import * as Sheet from '$lib/components/ui/sheet/index.js'
  import * as Tooltip from '$lib/components/ui/tooltip/index.js'

  import { Castle } from 'lucide-svelte'
  import { BookA } from 'lucide-svelte'
  import { Scroll } from 'lucide-svelte'
  import { Award } from 'lucide-svelte'
  import { Mountain } from 'lucide-svelte'
  import { WandSparkles } from 'lucide-svelte'
  import { toggleMode } from 'mode-watcher'
  import { ChevronsLeft } from 'lucide-svelte'

  const menus = [
    {
      title: 'Article',
      href: '/article',
      icon: BookA,
      isActive: false
    },
    {
      title: 'Code',
      href: '/code',
      icon: Scroll,
      isActive: true
    },
    {
      title: 'Tag',
      href: '/tag',
      icon: Award,
      isActive: false
    },
    {
      title: 'Attachment',
      href: '/attachment',
      icon: Mountain,
      isActive: false
    }
  ]
</script>

<div class="flex min-h-screen w-full flex-col bg-muted/40">
  <aside class="fixed inset-y-0 left-0 z-10 hidden w-14 flex-col border-r bg-background sm:flex">
    <nav class="flex flex-col items-center gap-4 px-2 py-4">
      <a
        href="##"
        class="group flex h-9 w-9 shrink-0 items-center justify-center gap-2 rounded-full bg-primary text-lg font-semibold text-primary-foreground md:h-8 md:w-8 md:text-base"
      >
        <Castle class="h-5 w-5 transition-all group-hover:scale-110" />
      </a>
      {#each menus as menu}
        <Tooltip.Root>
          <Tooltip.Trigger asChild let:builder>
            <a
              href="##"
              class="flex h-9 w-9 items-center justify-center rounded-lg transition-colors hover:text-foreground md:h-8 md:w-8 {menu.isActive
                ? 'bg-accent text-accent-foreground transition-colors hover:text-foreground md:h-8 md:w-8'
                : ''}"
              use:builder.action
              {...builder}
            >
              <svelte:component this={menu.icon} class="h-5 w-5" />
              <span class="sr-only">{menu.title}</span>
            </a>
          </Tooltip.Trigger>
          <Tooltip.Content side="right">{menu.title}</Tooltip.Content>
        </Tooltip.Root>
      {/each}
    </nav>
    <nav class="mt-auto flex flex-col items-center gap-4 px-2 py-4">
      <Tooltip.Root>
        <Tooltip.Trigger asChild let:builder>
          <a
            on:click={toggleMode}
            href="##"
            class="flex h-9 w-9 items-center justify-center rounded-lg text-muted-foreground transition-colors hover:text-foreground md:h-8 md:w-8"
            use:builder.action
            {...builder}
          >
            <WandSparkles class="h-5 w-5" />
            <span class="sr-only">Settings</span>
          </a>
        </Tooltip.Trigger>
        <Tooltip.Content side="right">Settings</Tooltip.Content>
      </Tooltip.Root>
    </nav>
  </aside>
  <div class="flex flex-col sm:gap-4 sm:py-4 sm:pl-14">
    <header
      class="sticky top-0 z-30 flex h-12 items-center gap-4 border-b bg-background px-4 sm:static sm:h-auto sm:border-0 sm:bg-transparent sm:px-6"
    >
      <Sheet.Root>
        <Sheet.Trigger asChild let:builder>
          <Button
            builders={[builder]}
            size="icon"
            variant="outline"
            class="border-0 shadow-transparent sm:hidden"
          >
            <ChevronsLeft class="h-5 w-5" />
            <span class="sr-only">Toggle Menu</span>
          </Button>
        </Sheet.Trigger>
        <Sheet.Content side="left" class="sm:max-w-xs">
          <nav class="grid gap-6 text-lg font-medium">
            <a
              href="##"
              class="group flex h-10 w-10 shrink-0 items-center justify-center gap-2 rounded-full bg-primary text-lg font-semibold text-primary-foreground md:text-base"
            >
              <Castle class="h-5 w-5 transition-all group-hover:scale-110" />
            </a>
            {#each menus as menu}
              <a
                href="##"
                class="flex items-center gap-4 px-2.5 {menu.isActive
                  ? ''
                  : 'text-muted-foreground'}"
              >
                <svelte:component this={menu.icon} class="h-5 w-5" />
                {menu.title}
              </a>
            {/each}
          </nav>
        </Sheet.Content>
      </Sheet.Root>
    </header>
    <main class="grid flex-1 items-start gap-4 p-4 sm:px-6 sm:py-0 md:gap-8"><slot /></main>
  </div>
</div>