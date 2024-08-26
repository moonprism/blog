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
 *   badRespoi
 * else
 *   { data }
 */
export interface BadRespoi {
  code: number
  message: string
}

export interface Respoi extends BadRespoi {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  data: any
  ok: boolean
}

export interface JwtInfo {
  username: string
}

// 泛用数据模型
export interface DataModel {
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
  tags: Tag[]
}

export interface Article extends DataModel, ArticleBody {}

interface ArticleContent {
  content: {
    text: string
  }
}

export interface ArticleDetail extends Article, ArticleContent {}

export interface TagBody {
  name: string
  color: string
}

export interface Tag extends DataModel, TagBody {}

export interface AttachmentBody {
  key: string
  summary: string
  year: number
}

export interface Attachment extends DataModel, AttachmentBody {}

// todo 还应细分
export interface GroupInfo {
  year: number
  count: number
}