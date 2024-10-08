<script>
  import { debounce } from "../utils";

  // 回到顶部按钮
  let arrowBtn;
  function handleArrowBtn() {
    const scrollTop = window.scrollY || document.documentElement.scrollTop
    const maxHeight = 500
    if (scrollTop > maxHeight && arrowBtn.style.display === "") {
      arrowBtn.style.display = "block";
    } else if (scrollTop < maxHeight && arrowBtn.style.display === "block") {
      arrowBtn.style.display = "";
    }
  }
  window.addEventListener("scroll", debounce(handleArrowBtn, 200));
</script>

<button
  bind:this={arrowBtn}
  class="arrow"
  on:click={() => {
    window.scrollTo(0, 0);
  }}
>
</button>

<style>
  .arrow {
    position: fixed;
    right: 60px;
    bottom: 55px;
    display: none;
    animation: fade-in;
    animation-duration: 1.5s;
    -webkit-animation: fade-in 1.5s;
    transform: rotate(270deg);
    border: none;
    border-radius: 3px;
    padding: 11px;
    cursor: pointer;
    backdrop-filter: blur(4px);
    background-color: rgba(255, 255, 255, 0.3);
  }
  .arrow:before,
  .arrow:after {
    border-right: 2px solid;
    content: "";
    display: block;
    height: 8px;
    margin-top: -6px;
    position: absolute;
    transform: rotate(135deg);
    right: 10px;
    top: 50%;
    width: 0;
  }
  .arrow:after {
    margin-top: -1px;
    transform: rotate(45deg);
  }

  @keyframes fade-in {
    0% {
      opacity: 0;
    }

    40% {
      opacity: 0;
    }

    100% {
      opacity: 1;
    }
  }
</style>
