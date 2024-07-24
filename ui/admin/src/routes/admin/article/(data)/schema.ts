import { z } from 'zod'

export const formSchema = z.object({
  title: z.string().min(1).max(20),
  status: z.number().int(),
  image: z.string().min(1).max(20),
  summary: z.string().min(1).max(20),
  tags: z.array(z.number()),
})

export type FormSchema = typeof formSchema
