import "./iconfont.js";

import "./styles/main.css";
import "./styles/art.md.css";
import "./styles/wind.css";

// 标签:hover动画
document.querySelectorAll(".art > .art-item > .anno").forEach((anno) => {
  anno.querySelectorAll(".art-tag").forEach((a) => {
    a.addEventListener("mouseover", () => {
      a.style.color = a.style.borderColor = a.style.backgroundColor;
      a.style.backgroundColor = "var(--background)";
    });
    a.addEventListener("mouseout", () => {
      a.style.backgroundColor = a.style.color;
      a.style.color = "white";
      a.style.borderColor = "var(--background)";
    });
  });
});

// 代码copy按钮
document.querySelectorAll(".art > .markdown-body > pre").forEach((pre) => {
  pre.addEventListener("click", (event) => {
    const offsetX = event.offsetX;
    const offsetY = event.offsetY;
    if (offsetX > 597 && offsetY < 17) {
      navigator.clipboard.writeText(pre.innerText);
      pre.style.setProperty("--copy-btn-text", "'copied'");
      pre.style.setProperty("--copy-btn-color", "#d2e3c8");
    }
  });
});
