import type { ComponentType } from 'svelte'
import type { ReadOrWritable } from '$src/types/store';
import { type Writable } from 'svelte/store';

export interface Option {
  id: number
  label: string
  icon: ComponentType
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