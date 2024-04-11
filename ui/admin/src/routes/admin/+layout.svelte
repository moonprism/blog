<script lang="ts">
  import { Button } from '$lib/components/ui/button/index.js'
  import * as Sheet from '$lib/components/ui/sheet/index.js'
  import * as Tooltip from '$lib/components/ui/tooltip/index.js'

  import { Castle, BookA, Scroll, Award, Mountain } from 'lucide-svelte'
  import { WandSparkles } from 'lucide-svelte'
  import { toggleMode } from 'mode-watcher'
  import { ChevronsLeft } from 'lucide-svelte'

  import { page } from '$app/stores'
  import type { SvelteComponent } from 'svelte'

  export const adminPath = '/admin'

  class Menu {
    title: string
    href: string
    // 好麻烦
    // https://stackoverflow.com/questions/63551277/typing-sveltecomponents-this-property-in-typescript
    icon: typeof SvelteComponent
    constructor(name: string, icon: any) {
      this.title = name
      this.href = `${adminPath}/${name}`
      this.icon = icon
    }
  }

  export let menus: Menu[] = [
    new Menu('article', BookA),
    new Menu('code', Scroll),
    new Menu('tag', Award),
    new Menu('attachment', Mountain)
  ]
</script>

<div class="flex min-h-screen w-full flex-col bg-muted/40">
  <aside class="fixed inset-y-0 left-0 z-10 hidden w-14 flex-col border-r bg-background sm:flex">
    <nav class="flex flex-col items-center gap-4 px-2 py-4">
      <a
        href={adminPath}
        class="group flex h-9 w-9 shrink-0 items-center justify-center gap-2 rounded-full bg-primary text-lg font-semibold text-primary-foreground md:h-8 md:w-8 md:text-base"
      >
        <Castle class="h-5 w-5 transition-all group-hover:scale-110" />
      </a>
      {#each menus as menu}
        <Tooltip.Root>
          <Tooltip.Trigger asChild let:builder>
            <a
              href={menu.href}
              class="flex h-9 w-9 items-center justify-center rounded-lg transition-colors hover:text-foreground md:h-8 md:w-8 {$page
                .url.pathname == menu.href
                ? 'bg-accent text-accent-foreground transition-colors hover:text-foreground md:h-8 md:w-8'
                : ''}"
              use:builder.action
              {...builder}
            >
              <svelte:component this={menu.icon} class="h-5 w-5" />
            </a>
          </Tooltip.Trigger>
          <Tooltip.Content side="right">{menu.title}</Tooltip.Content>
        </Tooltip.Root>
      {/each}
    </nav>
    <nav class="mt-auto flex flex-col items-center gap-4 px-2 py-4">
      <Tooltip.Root>
        <Tooltip.Trigger asChild let:builder>
          <button
            on:click={toggleMode}
            class="flex h-9 w-9 items-center justify-center rounded-lg text-muted-foreground transition-colors hover:text-foreground md:h-8 md:w-8"
            use:builder.action
            {...builder}
          >
            <WandSparkles class="h-5 w-5" />
          </button>
        </Tooltip.Trigger>
        <Tooltip.Content side="right">Wand</Tooltip.Content>
      </Tooltip.Root>
    </nav>
  </aside>
  <div class="flex flex-col sm:gap-4 sm:py-4 sm:pl-14">
    <!--移动端nav-->
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
              href={adminPath}
              class="group flex h-10 w-10 shrink-0 items-center justify-center gap-2 rounded-full bg-primary text-lg font-semibold text-primary-foreground md:text-base"
            >
              <Castle class="h-5 w-5 transition-all group-hover:scale-110" />
            </a>
            {#each menus as menu}
              <a
                href={menu.href}
                class="flex items-center gap-4 px-2.5 {$page.url.pathname == menu.href
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
