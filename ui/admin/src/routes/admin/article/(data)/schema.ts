import { z } from 'zod'

export const formSchema = z.object({
  title: z.string().min(1).max(30),
  status: z.number().int(),
  summary: z.string().max(255),
  image: z.string(),
  tags: z.array(z.any()),
})

export type FormSchema = typeof formSchema
