const foreseeStartChar = '☾🔮☽';
const foreseeEndChar = '☾†🔮☽';
/**
 * 转换查询标记
 * @param {string} str
 * @returns {string}
 */
export function foresee(str) {
  return str
    .replaceAll(foreseeStartChar, '<span class="em">')
    .replaceAll(foreseeEndChar, '</span>');
}

/**
 * 检测 .md
 * @param {string} lang
 */
export function isMdLang(lang) {
  return lang.replace(new RegExp(`^${foreseeStartChar}|${foreseeEndChar}$`, 'g'), '') === 'md';
}

/**
 *从url中获取保存的查询结果id列表
 * @returns {Array<number>}
 */
export function getUrlIds() {
  const urlSlices = window.location.href.split('#(:');
  if (urlSlices.length > 1 && urlSlices[1].length > 0) {
    return decodeURI(urlSlices[1])
      .split(',')
      .map((s) => Number(s));
  }
  return [];
}

/**
 * @param {Array<number>} ids
 * 保存查询结果id列表到url
 */
export function setUrlIds(ids = []) {
  if (ids.length !== 0) {
    window.history.pushState({}, 0, '#(:' + ids.join(','));
  } else {
    window.history.pushState({}, 0, ' ');
  }
}
