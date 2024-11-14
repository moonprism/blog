import './styles/main.css';
import './styles/art.md.css';
import './styles/gist.md.css';
import './styles/link.md.css';
import './styles/about.md.css';
import './styles/wind.css';

// 检查用户的系统主题偏好
const userPrefersDark =
  window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches;
// 尝试从 localStorage 中加载用户设置的主题，如果没有则根据系统偏好设置
const savedTheme = localStorage.getItem('theme');

// 如果用户之前设置了主题，则使用该主题，否则使用系统偏好
if (savedTheme) {
  if (savedTheme === 'dark') {
    document.body.classList.add('dark');
  } else {
    document.body.classList.remove('dark');
  }
} else {
  if (userPrefersDark) {
    document.body.classList.add('dark');
  }
}

// 标签:hover动画，(用js实现会有一种不流畅的美感
document.querySelectorAll('.art > .art-item > .anno').forEach((anno) => {
  anno.querySelectorAll('.art-tag').forEach((a) => {
    a.addEventListener('mouseover', () => {
      a.style.color = a.style.borderColor = a.style.backgroundColor;
      a.style.backgroundColor = 'var(--background)';
    });
    a.addEventListener('mouseout', () => {
      a.style.backgroundColor = a.style.color;
      a.style.color = 'white';
      a.style.borderColor = 'var(--background)';
    });
  });
});

// Links 页面随机排序等处理
/**
 * 获取两元素之间的所有元素
 * @param {Element} el1
 * @param {Element} el2
 * @param {string} tagName 筛选标签名
 * @returns
 */
function getElementsBetween(el1, el2, tagName = '') {
  const elements = [];
  let el = el1.nextElementSibling;

  while (el && el !== el2) {
    if (tagName === '' || el.tagName.toUpperCase() === tagName.toUpperCase()) {
      elements.push(el);
    }
    el = el.nextElementSibling;
  }

  return elements;
}

const heartSvg = `
<?xml version="1.0" encoding="utf-8"?>
<!-- License: MIT. Made by halfmage: https://github.com/halfmage/majesticons -->
<svg width="21px" height="21px" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" fill="none"><path fill="#e5404f" stroke="#e5404f" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 4c-3.2 0-5 2.667-5 4 0-1.333-1.8-4-5-4S3 6.667 3 8c0 7 9 12 9 12s9-5 9-12c0-1.333-.8-4-4-4z"/></svg>
`;

const starSvg = `
<?xml version="1.0" encoding="utf-8"?>
<!-- License: PD. Made by CFPB: https://github.com/cfpb/design-system -->
<svg fill="#FFAC33" width="18px" height="18px" viewBox="-2 0 19 19" xmlns="http://www.w3.org/2000/svg" class="cf-icon-svg"><path d="m12.673 10.779.798 4.02c.221 1.11-.407 1.566-1.395 1.013L8.5 13.81l-3.576 2.002c-.988.553-1.616.097-1.395-1.013l.397-2.001.401-2.02-1.51-1.397-1.498-1.385c-.832-.769-.592-1.507.532-1.64l2.026-.24 2.044-.242 1.717-3.722c.474-1.028 1.25-1.028 1.724 0l1.717 3.722 2.044.242 2.026.24c1.124.133 1.364.871.533 1.64L14.184 9.38z"/></svg>
`;

const fishcakeSvg = `
<!-- License: CC Attribution. Made by Muhammad Ridlo: mailto:m.ridlo@gmail.com -->
<svg
  width="29px"
  height="29px"
  viewBox="0 0 64 64"
  style="fill-rule:evenodd;clip-rule:evenodd;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:1.5;"
  version="1.1"
  xml:space="preserve"
  xmlns="http://www.w3.org/2000/svg"
  xmlns:serif="http://www.serif.com/"
  xmlns:xlink="http://www.w3.org/1999/xlink"
  ><g transform="matrix(1,0,0,1,-84,0)"
    ><g id="Naruto" transform="matrix(1,0,0,1,-188.333,0)"
      ><rect height="64" style="fill:none;" width="64" x="272.333" y="0" /><g
        transform="matrix(1,0,0,1,180.333,0)"
        ><path
          d="M121.475,5.673C121.904,4.659 122.899,4 124,4C125.101,4 126.096,4.659 126.525,5.673L128.37,10.03L131.742,6.711C132.526,5.938 133.697,5.71 134.715,6.131C135.733,6.553 136.399,7.542 136.408,8.644L136.445,13.375L140.83,11.598C141.851,11.185 143.02,11.422 143.799,12.201C144.578,12.98 144.815,14.149 144.402,15.17L142.625,19.555L147.356,19.592C148.458,19.601 149.447,20.267 149.869,21.285C150.29,22.303 150.062,23.474 149.289,24.258L145.97,27.63L150.327,29.475C151.341,29.904 152,30.899 152,32C152,33.101 151.341,34.096 150.327,34.525L145.97,36.37L149.289,39.742C150.062,40.526 150.29,41.697 149.869,42.715C149.447,43.733 148.458,44.399 147.356,44.408L142.625,44.445L144.402,48.83C144.815,49.851 144.578,51.02 143.799,51.799C143.02,52.578 141.851,52.815 140.83,52.402L136.445,50.625L136.408,55.356C136.399,56.458 135.733,57.447 134.715,57.869C133.697,58.29 132.526,58.062 131.742,57.289L128.37,53.97L126.525,58.327C126.096,59.341 125.101,60 124,60C122.899,60 121.904,59.341 121.475,58.327L119.63,53.97L116.258,57.289C115.474,58.062 114.303,58.29 113.285,57.869C112.267,57.447 111.601,56.458 111.592,55.356L111.555,50.625L107.17,52.402C106.149,52.815 104.98,52.578 104.201,51.799C103.422,51.02 103.185,49.851 103.598,48.83L105.375,44.445L100.644,44.408C99.542,44.399 98.553,43.733 98.131,42.715C97.71,41.697 97.938,40.526 98.711,39.742L102.03,36.37L97.673,34.525C96.659,34.096 96,33.101 96,32C96,30.899 96.659,29.904 97.673,29.475L102.03,27.63L98.711,24.258C97.938,23.474 97.71,22.303 98.131,21.285C98.553,20.267 99.542,19.601 100.644,19.592L105.375,19.555L103.598,15.17C103.185,14.149 103.422,12.98 104.201,12.201C104.98,11.422 106.149,11.185 107.17,11.598L111.555,13.375L111.592,8.644C111.601,7.542 112.267,6.553 113.285,6.131C114.303,5.71 115.474,5.938 116.258,6.711L119.63,10.03L121.475,5.673Z"
          style="fill:white;stroke:rgb(34,32,77);stroke-width:2px;"
        /></g
      ><g transform="matrix(1.01035,0,0,1.05877,185.386,2.13651)"
        ><path
          d="M116,31C116,31 114.349,24.48 122,25C129.651,25.52 126.254,32.473 123,35C119.746,37.527 107.184,38.659 109,28C110.816,17.341 120,20 120,20"
          style="fill:none;stroke:rgb(205,90,70);stroke-width:1.93px;"
        /></g
      ></g
    ></g
  ></svg
>
`;

/**
 * 使字符串hash到一个正整数
 * @param {string} str
 * @returns {number}
 */
function hashStringToInt(str) {
  let hash = 0;
  for (let i = 0; i < str.length; i++) {
    // 使用字符的 UTF-16 编码值进行哈希
    hash = (hash << 5) - hash + str.charCodeAt(i); // hash * 31 + str.charCodeAt(i)
    hash |= 0; // 强制转换为 32 位整数
  }
  return Math.abs(hash);
}

// 适应字体大小
document.querySelectorAll('.link .markdown-body blockquote').forEach((e) => {
  const bio = e.querySelector('code');
  if (bio.innerText.length >= 21) {
    bio.style.fontSize = '13px';
  }
});

const linkHeadingEl = document.querySelector('.link .markdown-body h2');
const collectionHEadingEl = document.querySelector('.link .markdown-body h3');
// 设置图标
if (linkHeadingEl) {
  linkHeadingEl.insertAdjacentHTML('afterbegin', heartSvg);
  linkHeadingEl.style.marginBottom = '15px';
  if (collectionHEadingEl) {
    collectionHEadingEl.insertAdjacentHTML('afterbegin', starSvg);
  }
  const boardHeadingEl = document.querySelector('.link .markdown-body h2:nth-of-type(2)');
  boardHeadingEl.insertAdjacentHTML('afterbegin', fishcakeSvg);
  boardHeadingEl.style.fontSize = '17px';
}
const seedEl = document.querySelector('.link .markdown-body > p > code');
if (seedEl) {
  // 随机排序
  const blockquoteEls = getElementsBetween(
    linkHeadingEl,
    collectionHEadingEl ? collectionHEadingEl : boardHeadingEl,
    'blockquote'
  );
  const url = window.location.href;
  const seed =
    url.indexOf('?') != -1
      ? parseInt(url.split('?')[1])
      : Math.round(Math.random() * 899999) + 100000;

  blockquoteEls.map((v, i) => {
    const j = (seed * hashStringToInt(v.innerText)) % blockquoteEls.length;
    [blockquoteEls[j].innerHTML, v.innerHTML] = [v.innerHTML, blockquoteEls[j].innerHTML];
  });
  window.history.pushState({}, 0, '?' + seed);
  seedEl.innerHTML += `<span style="font-size:.95em;">${seed}</span>`;
}
