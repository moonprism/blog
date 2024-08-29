import { z } from 'zod'

export const formSchema = z.object({
  lang: z.string().min(1).max(10),
  title: z.string().min(1).max(100),
  content: z.string(),
})

export type FormSchema = typeof formSchema
