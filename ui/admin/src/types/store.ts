import { type Readable, type Writable } from 'svelte/store';

export type ReadOrWritable<T> = Readable<T> | Writable<T>;