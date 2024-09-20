import type { JwtInfo } from '$src/types/stream'

// manage jwt from localStorage
const jwtKey = 'jwt_v'

export const getJwt = (): string => {
  return localStorage.getItem(jwtKey) || ''
}

export const setJwt = (token: string): void => {
  localStorage.setItem(jwtKey, token)
}

export const removeJwt = (): void => {
  localStorage.removeItem(jwtKey)
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
