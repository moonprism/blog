// ff custom wrapper for "fetch" funtion 
export const fet = async (
  path: string,
  method: string,
  data: any,
): Promise<Response> => {
  const address = 'http://localhost:2999/api/'
  let headers:HeadersInit = {
      "Content-Type": "application/json",
  }
  const token = getJwt()
  if (token) {
    headers["Authorization"] = `Bearer ${token}`
  }
  const response = await fetchWithTimeout(`${address}${path}`, {
    method,
    headers,
    body: JSON.stringify(data),
  })
  if (!response.ok) {
    // todo
  }
  return response
}

function fetchWithTimeout(url: string, options: RequestInit, timeout = 5000) {
  return Promise.race([
    fetch(url, options),
    new Promise((_, reject) =>
      setTimeout(() => reject(new Error('Request timed out')), timeout)
    )
  ]) as Promise<Response>;
}

// manage jwt from localStorage
const jwtKey = 'jwt_v'

export const getJwt = (): string => {
  return localStorage.getItem(jwtKey) || ''
}

export const setJwt = (token: string): void => {
  localStorage.setItem(jwtKey, token)
}