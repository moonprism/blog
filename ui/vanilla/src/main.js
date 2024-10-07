import "./iconfont.js";

import "./css/main.css";
import "./css/markdown.css";
import "./css/wind.css";

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
