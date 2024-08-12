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
 * if response.ok == true
 *   badRespo
 * else
 *   { data }
 */
export interface BadRespo {
  code: number
  message: string
}

export interface Respoi extends BadRespo {
  data: any
  ok: boolean
}

export interface JwtInfo {
  username: string
}

// 泛用数据模型
interface DataModel {
  id: number
  created: number
  updated: number
}

export interface ArticleBody {
  title: string
  status: number
  rune: number
  image: string
  summary: string
  content: string
  tags: Tag[]
}

export interface Article extends DataModel, ArticleBody {}

export interface TagBody {
  name: string
  color: string
}

export interface Tag extends DataModel, TagBody {}
