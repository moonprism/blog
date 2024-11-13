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

// 切换主题函数
export function toggleTheme(isDark = false) {
  //const isDark = document.body.classList.toggle('dark');
  if (isDark) {
    document.body.classList.add('dark');
  } else {
    document.body.classList.remove('dark');
  }
  // 将用户选择的主题保存到 localStorage
  localStorage.setItem('theme', isDark ? 'dark' : 'light');
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
