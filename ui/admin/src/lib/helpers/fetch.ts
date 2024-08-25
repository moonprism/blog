import type { BadRespoi, DataModel, Respoi } from '$src/types/stream'
import toast from '$lib/helpers/toast'

import { PUBLIC_API_ADDR, PUBLIC_ATTACHMENT_CDN, PUBLIC_MOCK_MODE } from '$env/static/public'
import { getJwt } from './jwt'
import { writable } from 'svelte/store'
import { apiData } from '$src/mock'

const host = PUBLIC_API_ADDR

export const isMockMode = PUBLIC_MOCK_MODE === 'true'
export const fileCDN = PUBLIC_ATTACHMENT_CDN

export const isExternalLink = (link: string) => {
  return link.startsWith('data') || link.startsWith('http')
}

export const getRealSrc = (key: string) => {
  if (key === '' || isExternalLink(key)) {
    return key
  }
  return `${fileCDN}${key}`
}

export const fet = {
  get: (path: string) => request(path, 'GET'),
  put: (path: string, data: unknown) => request(path, 'PUT', data),
  post: (path: string, data: unknown) => request(path, 'POST', data),
  delete: (path: string) => request(path, 'DELETE')
}

export const isRequestIn = writable(false)
let fakeGlobalID = 1001

// custom wrapper for "fetch" function
const request = async (path: string, method: string, data?: unknown): Promise<Respoi> => {
  isRequestIn.set(true)
  const headers: HeadersInit = {
    'Content-Type': 'application/json'
  }

  const token = getJwt()
  if (token) {
    headers['Authorization'] = `Bearer ${token}`
  }

  const options: RequestInit = {
    method,
    headers
  }
  if (data !== null) {
    options.body = JSON.stringify(data)
  }

  const result: Respoi = {
    ok: false,
    code: 0,
    message: '',
    data: null
  }

  if (isMockMode) {
    let fakeData: DataModel | DataModel[] = <DataModel>data
    let ro = ''
    switch (method) {
      case 'DELETE':
        break
      case 'GET':
        ro = path.split('?')[0]
        // article/{id}
        if (ro.includes('/')) {
          // eslint-disable-next-line @typescript-eslint/ban-ts-comment
          // @ts-expect-error
          fakeData = apiData[ro.split('/')[0]].detail
        } else {
          // eslint-disable-next-line @typescript-eslint/ban-ts-comment
          // @ts-expect-error
          fakeData = apiData[ro].list
        }
        break
      case 'POST':
        fakeData.created = Date.parse(new Date().toString()) / 1000
        fakeData.id = fakeGlobalID++
      // eslint-disable-next-line no-fallthrough
      default:
        fakeData.updated = Date.parse(new Date().toString()) / 1000
    }
    return new Promise((resolve) => {
      setTimeout(() => {
        result.ok = true
        result.data = fakeData
        // console.log(result)
        isRequestIn.set(false)
        resolve(result)
      }, 1000)
    })
  }

  try {
    const response = await fetchWithTimeout(`${host}${path}`, options)
    const data = await response.json()

    isRequestIn.set(false)
    if (!response.ok) {
      const badRespoi = <BadRespoi>data
      toast.error(badRespoi.code.toString(), {
        description: `${badRespoi.message}`
      })
      result.ok = false
      result.code = badRespoi.code
      result.message = badRespoi.message
      return result
    }

    result.ok = true
    result.data = data
    return result
  } catch (e) {
    result.message = (e as Error).message
    toast.error('system failure', {
      description: result.message
    })
    isRequestIn.set(false)
    return result
  }
}

function fetchWithTimeout(url: string, options: RequestInit, timeout = 5000) {
  return Promise.race([
    fetch(url, options),
    new Promise((_, reject) => setTimeout(() => reject(new Error('Request timed out')), timeout))
  ]) as Promise<Response>
}
