import { z } from 'zod'

export const formSchema = z.object({
  link: z.string().min(1),
  summary: z.string().max(20),
})

export type FormSchema = typeof formSchema
