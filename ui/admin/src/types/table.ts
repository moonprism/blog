import Circle from 'svelte-radix/Circle.svelte'
import QuestionMarkCircled from 'svelte-radix/QuestionMarkCircled.svelte'
import Stopwatch from 'svelte-radix/Stopwatch.svelte'

export const statuses = [
  {
    value: '0',
    label: 'Draft',
    icon: QuestionMarkCircled
  },
  {
    value: '1',
    label: 'Published',
    icon: Circle
  },
  {
    value: '2',
    label: 'Hidden',
    icon: Stopwatch
  }
]
