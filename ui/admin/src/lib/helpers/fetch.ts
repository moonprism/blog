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
let requestCount = 0
let fakeGlobalID = 1001

// custom wrapper for "fetch" function
const request = async (path: string, method: string, data?: unknown): Promise<Respoi> => {
  requestCount++
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
    let fakeData: DataModel | DataModel[] | { token: string } = <DataModel>data
    const ro = path.split('?')[0]
    switch (method) {
      case 'DELETE':
        break
      case 'GET':
        // article/{id}, group/year, group/gist-lang
        if (ro.includes('/')) {
          const modules = ro.split('/')
          const action = modules[0] === 'article' ? 'detail' : modules[1]
          // eslint-disable-next-line @typescript-eslint/ban-ts-comment
          // @ts-expect-error
          fakeData = apiData[modules[0]][action]
        } else {
          // eslint-disable-next-line @typescript-eslint/ban-ts-comment
          // @ts-expect-error
          fakeData = apiData[ro].list
        }
        break
      case 'POST':
        if (ro === 'login') {
          fakeData = { token: '' }
          break
        }
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
      }, 500)
    })
  }

  try {
    const response = await fetchWithTimeout(`${host}${path}`, options)
    const data = await response.json()

    if (--requestCount === 0) {
      isRequestIn.set(false)
    }
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
    if (--requestCount === 0) {
      isRequestIn.set(false)
    }
    return result
  }
}

function fetchWithTimeout(url: string, options: RequestInit, timeout = 5000) {
  return Promise.race([
    fetch(url, options),
    new Promise((_, reject) => setTimeout(() => reject(new Error('Request timed out')), timeout))
  ]) as Promise<Response>
}
