import { writable, type Unsubscriber } from "svelte/store"
import Alert from "./alert.svelte"

const open = writable(false)
const title = writable('')
const description = writable('')
const confirm = writable(false)
let unsubscribe: Unsubscriber

new Alert({
  target: document.body,
  props: {
    open,
    description,
    confirm,
    title,
  }
})

export const alertDialog = (desc = '', t = '') => {
  description.set(desc)
  title.set(t)
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