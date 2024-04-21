import type { Tag } from "$src/types/stream";
import { writable } from "svelte/store";

export const tableData = writable([] as Tag[])

export const defaultFormData = {id: 0, name:'', color: '#000000'} as Tag
export const formData = writable(defaultFormData)
export const formOpen = writable(false)