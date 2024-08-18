import { z } from 'zod'

export const formSchema = z.object({
  key: z.string().min(1),
  summary: z.string().max(20),
})

export type FormSchema = typeof formSchema
