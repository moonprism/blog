import './iconfont.js';

import './styles/main.css';
import './styles/art.md.css';
import './styles/gist.md.css';
import './styles/link.md.css';
import './styles/wind.css';

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

// 代码copy按钮
document.querySelectorAll('.art > .markdown-body > pre').forEach((pre) => {
  pre.addEventListener('click', (event) => {
    const offsetX = event.offsetX;
    const offsetY = event.offsetY;
    if (offsetX > 597 && offsetY < 17) {
      navigator.clipboard.writeText(pre.innerText);
      pre.style.setProperty('--copy-btn-text', "'copied'");
      pre.style.setProperty('--copy-btn-color', '#d2e3c8');
    }
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

const appleSvg = `
<?xml version="1.0" encoding="iso-8859-1"?>
<!-- Generator: Adobe Illustrator 16.0.0, SVG Export Plug-In . SVG Version: 6.00 Build 0)  -->
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<!-- License: CC0. Made by SVG Repo: https://www.svgrepo.com/svg/42619/apple-fruit -->
<svg version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
	 fill="#8FD14F" width="19.5px" height="19.5px" viewBox="0 0 45.721 45.721" style="enable-background:new 0 0 45.721 45.721;"
	 xml:space="preserve">
<g>
	<path d="M31.686,11.161c-2.615-0.057-5.047,1.063-7.073,2.882c-0.344-2.371-1.271-4.273-1.271-4.273
		c10.967-1.485,7.901-9.753,7.901-9.753c-7.738-0.328-9.239,4.304-9.399,7.344c-2.385-2.55-5.077-4.009-6.202-4.518
		c-1.071-0.484-2.329,0.074-2.775,1.159c-0.446,1.088,0.094,2.284,1.161,2.774c1.411,0.649,5.841,2.251,7.181,7.409
		c-2.058-1.907-4.548-3.024-7.231-3.024c-7.066,0-13.598,5.845-11.318,17.281c1.866,9.357,8.678,17.279,15.744,17.279
		c1.842,0,3.291-0.527,4.427-1.475c1.138,0.946,2.586,1.484,4.428,1.475c7.307-0.045,13.542-7.505,15.744-17.279
		C45.232,18.552,40.071,11.343,31.686,11.161z"/>
</g>
</svg>
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

document.addEventListener('DOMContentLoaded', function () {
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
  linkHeadingEl.insertAdjacentHTML('afterbegin', heartSvg);
  linkHeadingEl.style.marginBottom = '15px';
  if (collectionHEadingEl) {
    collectionHEadingEl.insertAdjacentHTML('afterbegin', starSvg);
  }

  const boardHeadingEl = document.querySelector('.link .markdown-body h2:nth-of-type(2)');
  boardHeadingEl.insertAdjacentHTML('afterbegin', appleSvg);

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
});
