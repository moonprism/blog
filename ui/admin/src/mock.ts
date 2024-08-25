import articles from '$src/routes/admin/article/(data)/mock.json'
import articleDetail from '$src/routes/admin/article/write/[slug]/(data)/mock.json'
import tags from '$src/routes/admin/tag/(data)/mock.json'
import attachments from '$src/routes/admin/attachment/(data)/mock.json'
export const apiData = {
  article: {
    list: articles,
    detail: articleDetail
  },
  attachment: {
    list: attachments
  },
  tag: {
    list: tags
  }
}
