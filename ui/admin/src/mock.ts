import articles from '$src/routes/admin/article/(data)/mock.json'
import articleDetail from '$src/routes/admin/article/write/[[id]]/(data)/mock.json'
import tags from '$src/routes/admin/tag/(data)/mock.json'
import attachmentMock from '$src/routes/admin/attachment/(data)/mock.json'
import gistMock from '$src/routes/admin/gist/(data)/mock.json'
import comments from '$src/routes/admin/comment/(data)/mock.json'
export const apiData = {
  article: {
    list: articles,
    detail: articleDetail
  },
  attachment: {
    list: attachmentMock.attachments
  },
  tag: {
    list: tags
  },
  group: {
    year: attachmentMock['group-year'],
    'gist-lang': gistMock['gist-lang']
  },
  gist: {
    list: gistMock.gists
  },
  comment: {
    list: comments
  }
}
