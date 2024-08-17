import type { SvelteComponent } from 'svelte'
import type { ReadOrWritable } from '$src/types/store';
import { type Writable } from 'svelte/store';
import type { ComponentRenderConfig } from 'svelte-headless-table';

export interface Option {
  id: number
  label: string
  icon: ComponentRenderConfig<SvelteComponent>
}

export interface Filter {
  // todo keyof ...
  name: string
  options: ReadOrWritable<Option[]>
}

export interface ViewOption {
  type: 'hideColumn' | 'cardSize',
  selected: Writable<string[]>,
  options: string[],
}