import type { BadRespo, Respoi } from '$src/types/stream'
import toast from '$lib/helpers/toast'

import { PUBLIC_API_ADDR, PUBLIC_ATTACHMENT_CDN, PUBLIC_MOCK_MODE } from '$env/static/public'
import { getJwt } from './jwt'
import { writable } from 'svelte/store'

const host = PUBLIC_API_ADDR

export const isMockMode = PUBLIC_MOCK_MODE === "true"

export const getRealSrc = (link: string) => {
  if (
    link !== '' &&
    !link.startsWith('data') &&
    !link.startsWith('http')
  ) {
    return `${PUBLIC_ATTACHMENT_CDN}${link}`
  }
  return ''
}

export const fet = {
  get: (path: string) => request(path, 'GET'),
  put: (path: string, data: any) => request(path, 'PUT', data),
  post: (path: string, data: any) => request(path, 'POST', data),
  delete: (path: string) => request(path, 'DELETE')
}

export const isReuqestIn = writable(false)
let fakeGlobalID = 1001

// custom wrapper for "fetch" funtion
const request = async (path: string, method: string, data?: any): Promise<Respoi> => {
  isReuqestIn.set(true)
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
    isReuqestIn.set(false)
    let fakeData: any
    switch (method) {
      case 'DELETE':
        break;
      case 'GET':
        fakeData = []
        break
      case 'POST':
        data.created = Date.parse(new Date().toString()) / 1000
        data.id = fakeGlobalID++
      default:
        data.updated = Date.parse(new Date().toString()) / 1000
        fakeData = data
    }
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        result.ok = true
        result.data = fakeData
        resolve(result)
      }, 1000)
    })
  }

  try {
    const response = await fetchWithTimeout(`${host}${path}`, options)
    const data = await response.json()

    isReuqestIn.set(false)
    if (!response.ok) {
      const badRespo = <BadRespo>data
      toast.error(badRespo.code.toString(), {
        description: `${badRespo.message}`
      })
      result.ok = false
      result.code = badRespo.code
      result.message = badRespo.message
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
    isReuqestIn.set(false)
    return result
  }
}

function fetchWithTimeout(url: string, options: RequestInit, timeout = 5000) {
  return Promise.race([
    fetch(url, options),
    new Promise((_, reject) => setTimeout(() => reject(new Error('Request timed out')), timeout))
  ]) as Promise<Response>
}
