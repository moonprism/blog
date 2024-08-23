import { writable, type Unsubscriber } from "svelte/store"
import type { SvelteComponent } from "svelte"
import Alert from "./alert.svelte"

let component: SvelteComponent
const open = writable(false)
const title = writable('')
const description = writable('')
const confirm = writable(false)
let unsubscribe: Unsubscriber
export const alertDialog = (desc = '', titl = '') => {
  description.set(desc)
  title.set(titl)
  if (!component) {
    component = new Alert({
      target: document.body,
      props: {
        open,
        description,
        confirm,
        title
      }
    })
  }
  open.set(true)
  return new Promise((resolve) => {
    unsubscribe = confirm.subscribe(value => {
      if (value === true) {
        resolve(value)
        confirm.set(false)
      }
      if (unsubscribe) {
        unsubscribe()
      }
    });
  });
}