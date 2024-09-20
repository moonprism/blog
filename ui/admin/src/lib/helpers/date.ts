export const DateFormat = (ts: number) => {
  return new Date(ts * 1000).toLocaleDateString('zh-CN')
}
