import type { option } from '$src/types/table'
import { BookLock, BookOpen, VenetianMask } from 'lucide-svelte'

export const statuses: option[] = [
  {
    id: 0,
    label: 'Draft',
    icon: BookLock
  },
  {
    id: 1,
    label: 'Published',
    icon: BookOpen
  },
  {
    id: 2,
    label: 'Hidden',
    icon: VenetianMask
  }
]