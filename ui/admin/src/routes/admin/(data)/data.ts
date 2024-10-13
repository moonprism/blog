import type { AppConfig } from '$src/types/stream'
import { fet, isExternalLink } from '@/helpers/fetch'

const res = await fet.get('system')
export const appInfo = <AppConfig>res.data

export const fileCDN = appInfo.attachmentCDN

export const getRealSrc = (key: string) => {
  if (key === '' || isExternalLink(key)) {
    return key
  }
  return `${fileCDN}${key}`
}