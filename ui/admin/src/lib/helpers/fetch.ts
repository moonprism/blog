import type { JwtInfo, LoginRequest } from '$src/types/stream'
import toast from '$lib/helpers/toast'

import { PUBLIC_API_ADDR } from '$env/static/public'
import { goto } from '$app/navigation'

// ff custom wrapper for "fetch" funtion
export const fet = async (path: string, method: string, data: any): Promise<Response> => {
  const address = PUBLIC_API_ADDR
  let headers: HeadersInit = {
    'Content-Type': 'application/json'
  }
  const token = getJwt()
  if (token) {
    headers['Authorization'] = `Bearer ${token}`
  }
  const response = await fetchWithTimeout(`${address}${path}`, {
    method,
    headers,
    body: JSON.stringify(data)
  })
  if (!response.ok) {
    toast.error(await response.text())
    // 好吧，咱根本不需要错误处理
    // let data = <BadResponse> await response.json()
    // toast.error(data.code.toString(), {
    //   description: data.message
    // })
  }
  return response
}

function fetchWithTimeout(url: string, options: RequestInit, timeout = 5000) {
  return Promise.race([
    fetch(url, options),
    new Promise((_, reject) => setTimeout(() => reject(new Error('Request timed out')), timeout))
  ]) as Promise<Response>
}

// manage jwt from localStorage
const jwtKey = 'jwt_v'

const getJwt = (): string => {
  return localStorage.getItem(jwtKey) || ''
}

const setJwt = (token: string): void => {
  localStorage.setItem(jwtKey, token)
}

export const getJwtInfo = (): JwtInfo | null => {
  const jwt = getJwt()
  if (jwt == '') {
    return null
  }
  return <JwtInfo>(
    JSON.parse(
      decodeURIComponent(window.atob(jwt.split('.')[1].replace(/-/g, '+').replace(/_/g, '/')))
    )
  )
}

export const login = (data: LoginRequest) => {
  return fet('login', 'post', data)
    .then((response) => {
      if (response.ok) {
        response.json().then((data) => {
          setJwt(data.token)
          toast.success(`登陆成功，${getJwtInfo()?.username}`)
          goto('/home')
        })
      } else {
        throw Error('login failed')
      }
    })
    .catch((err) => {
      throw err
    })
}
