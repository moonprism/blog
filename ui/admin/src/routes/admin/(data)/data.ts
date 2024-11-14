import type { AppConfig } from '$src/types/stream'
import { fet, isExternalLink } from '@/helpers/fetch'
import { get, writable } from 'svelte/store'

export const appInfo = writable({} as AppConfig)
export async function initAppInfo() {
  if (Object.keys(get(appInfo)).length === 0) {
    const res = await fet.get('settings')
    appInfo.set(res.data)
  }
}

export const getRealSrc = (key: string) => {
  if (key === '' || isExternalLink(key)) {
    return key
  }
  return `${get(appInfo).attachmentCDN}${key}`
}
