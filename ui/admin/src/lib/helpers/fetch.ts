import type { BadRespo, Respoi } from '$src/types/stream'
import toast from '$lib/helpers/toast'

import { PUBLIC_API_ADDR } from '$env/static/public'
import { getJwt } from './jwt'

const host = PUBLIC_API_ADDR

export const fet = {
  get: (path: string) => request(path, 'GET'),
  put: (path: string, data: any) => request(path, 'PUT', data),
  post: (path: string, data: any) => request(path, 'POST', data),
  delete: (path: string) => request(path, 'DELETE')
}

// custom wrapper for "fetch" funtion
const request = async (path: string, method: string, data?: any): Promise<Respoi> => {
  const headers: HeadersInit = {
    'Content-Type': 'application/json'
  }

  const token = getJwt()
  if (token) {
    headers['Authorization'] = `Bearer ${token}`
  }

  let options: RequestInit = {
    method,
    headers
  }
  if (data !== null) {
    options.body = JSON.stringify(data)
  }

  let result: Respoi = {
    ok: false,
    code: 0,
    message: '',
    data: null
  }

  try {
    const response = await fetchWithTimeout(`${host}${path}`, options)
    const data = await response.json()

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
    return result
  }
}

function fetchWithTimeout(url: string, options: RequestInit, timeout = 5000) {
  return Promise.race([
    fetch(url, options),
    new Promise((_, reject) => setTimeout(() => reject(new Error('Request timed out')), timeout))
  ]) as Promise<Response>
}
