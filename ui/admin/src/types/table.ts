import type { ComponentType } from 'svelte'

export interface option {
  id: number
  label: string
  icon: ComponentType
}

export interface filter {
  name: string
  options: option[]
}