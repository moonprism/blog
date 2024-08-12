import type { ComponentType } from 'svelte'
import type { ReadOrWritable } from 'svelte-headless-table'

export interface option {
  id: number
  label: string
  icon: ComponentType
}

export interface filter {
  // todo keyof ...
  name: string
  options: ReadOrWritable<option[]>
}
