import { z } from 'zod'

export const formSchema = z.object({
  name: z.string().min(1).max(20),
  color: z.string().regex(new RegExp('^#[0-9a-f]{6}$'))
})

export type FormSchema = typeof formSchema
