import { z } from 'zod'

export const formSchema = z.object({
  title: z.string().min(1).max(30),
  background: z.string().max(255),
})

export type FormSchema = typeof formSchema
