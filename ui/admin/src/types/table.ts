import type { ComponentType } from 'svelte'

export interface option {
  id: number
  label: string
  icon: ComponentType
}

export interface filter {
  // todo keyof ...
  name: string
  options: option[]
}
