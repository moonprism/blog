/**
 * Stream Types
 * Defines the data types transmitted in network streams.
 * 此文件定义了穿行于流式网络中灵魂的具象
 */

export interface LoginRequest {
  username: string
  password: string
}

/**
 * 服务器 status == 200 直接返回 { data }
 * status != 200 则返回 BadResponse
 */
export interface BadResponse {
  code: number
  message: string
}

export interface JwtInfo {
  username: string
}
