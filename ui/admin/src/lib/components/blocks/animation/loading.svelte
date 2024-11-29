<!-- https://css-loaders.com/classic/ -->
<div class="loader"></div>

<style>
  .loader {
    --w: 10ch;
    --c: var(--loading);
    font-weight: bold;
    font-family: monospace;
    font-size: inherit;
    line-height: inherit;
    letter-spacing: var(--w);
    width: var(--w);
    overflow: hidden;
    white-space: nowrap;
    color: #0000;
    text-shadow:
      calc(0 * var(--w)) 0 var(--c),
      calc(-1 * var(--w)) 0 var(--c),
      calc(-2 * var(--w)) 0 var(--c),
      calc(-3 * var(--w)) 0 var(--c),
      calc(-4 * var(--w)) 0 var(--c),
      calc(-5 * var(--w)) 0 var(--c),
      calc(-6 * var(--w)) 0 var(--c),
      calc(-7 * var(--w)) 0 var(--c),
      calc(-8 * var(--w)) 0 var(--c),
      calc(-9 * var(--w)) 0 var(--c);
    animation: l20 2s infinite linear;
  }
  .loader:before {
    content: 'Loading...';
  }

  @keyframes l20 {
    9.09% {
      text-shadow:
        calc(0 * var(--w)) -10px var(--c),
        calc(-1 * var(--w)) 0 var(--c),
        calc(-2 * var(--w)) 0 var(--c),
        calc(-3 * var(--w)) 0 var(--c),
        calc(-4 * var(--w)) 0 var(--c),
        calc(-5 * var(--w)) 0 var(--c),
        calc(-6 * var(--w)) 0 var(--c),
        calc(-7 * var(--w)) 0 var(--c),
        calc(-8 * var(--w)) 0 var(--c),
        calc(-9 * var(--w)) 0 var(--c);
    }
    18.18% {
      text-shadow:
        calc(0 * var(--w)) 0 var(--c),
        calc(-1 * var(--w)) -10px var(--c),
        calc(-2 * var(--w)) 0 var(--c),
        calc(-3 * var(--w)) 0 var(--c),
        calc(-4 * var(--w)) 0 var(--c),
        calc(-5 * var(--w)) 0 var(--c),
        calc(-6 * var(--w)) 0 var(--c),
        calc(-7 * var(--w)) 0 var(--c),
        calc(-8 * var(--w)) 0 var(--c),
        calc(-9 * var(--w)) 0 var(--c);
    }
    27.27% {
      text-shadow:
        calc(0 * var(--w)) 0 var(--c),
        calc(-1 * var(--w)) 0 var(--c),
        calc(-2 * var(--w)) -10px var(--c),
        calc(-3 * var(--w)) 0 var(--c),
        calc(-4 * var(--w)) 0 var(--c),
        calc(-5 * var(--w)) 0 var(--c),
        calc(-6 * var(--w)) 0 var(--c),
        calc(-7 * var(--w)) 0 var(--c),
        calc(-8 * var(--w)) 0 var(--c),
        calc(-9 * var(--w)) 0 var(--c);
    }
    36.36% {
      text-shadow:
        calc(0 * var(--w)) 0 var(--c),
        calc(-1 * var(--w)) 0 var(--c),
        calc(-2 * var(--w)) 0 var(--c),
        calc(-3 * var(--w)) -10px var(--c),
        calc(-4 * var(--w)) 0 var(--c),
        calc(-5 * var(--w)) 0 var(--c),
        calc(-6 * var(--w)) 0 var(--c),
        calc(-7 * var(--w)) 0 var(--c),
        calc(-8 * var(--w)) 0 var(--c),
        calc(-9 * var(--w)) 0 var(--c);
    }
    45.45% {
      text-shadow:
        calc(0 * var(--w)) 0 var(--c),
        calc(-1 * var(--w)) 0 var(--c),
        calc(-2 * var(--w)) 0 var(--c),
        calc(-3 * var(--w)) 0 var(--c),
        calc(-4 * var(--w)) -10px var(--c),
        calc(-5 * var(--w)) 0 var(--c),
        calc(-6 * var(--w)) 0 var(--c),
        calc(-7 * var(--w)) 0 var(--c),
        calc(-8 * var(--w)) 0 var(--c),
        calc(-9 * var(--w)) 0 var(--c);
    }
    54.54% {
      text-shadow:
        calc(0 * var(--w)) 0 var(--c),
        calc(-1 * var(--w)) 0 var(--c),
        calc(-2 * var(--w)) 0 var(--c),
        calc(-3 * var(--w)) 0 var(--c),
        calc(-4 * var(--w)) 0 var(--c),
        calc(-5 * var(--w)) -10px var(--c),
        calc(-6 * var(--w)) 0 var(--c),
        calc(-7 * var(--w)) 0 var(--c),
        calc(-8 * var(--w)) 0 var(--c),
        calc(-9 * var(--w)) 0 var(--c);
    }
    63.63% {
      text-shadow:
        calc(0 * var(--w)) 0 var(--c),
        calc(-1 * var(--w)) 0 var(--c),
        calc(-2 * var(--w)) 0 var(--c),
        calc(-3 * var(--w)) 0 var(--c),
        calc(-4 * var(--w)) 0 var(--c),
        calc(-5 * var(--w)) 0 var(--c),
        calc(-6 * var(--w)) -10px var(--c),
        calc(-7 * var(--w)) 0 var(--c),
        calc(-8 * var(--w)) 0 var(--c),
        calc(-9 * var(--w)) 0 var(--c);
    }
    72.72% {
      text-shadow:
        calc(0 * var(--w)) 0 var(--c),
        calc(-1 * var(--w)) 0 var(--c),
        calc(-2 * var(--w)) 0 var(--c),
        calc(-3 * var(--w)) 0 var(--c),
        calc(-4 * var(--w)) 0 var(--c),
        calc(-5 * var(--w)) 0 var(--c),
        calc(-6 * var(--w)) 0 var(--c),
        calc(-7 * var(--w)) -10px var(--c),
        calc(-8 * var(--w)) 0 var(--c),
        calc(-9 * var(--w)) 0 var(--c);
    }
    81.81% {
      text-shadow:
        calc(0 * var(--w)) 0 var(--c),
        calc(-1 * var(--w)) 0 var(--c),
        calc(-2 * var(--w)) 0 var(--c),
        calc(-3 * var(--w)) 0 var(--c),
        calc(-4 * var(--w)) 0 var(--c),
        calc(-5 * var(--w)) 0 var(--c),
        calc(-6 * var(--w)) 0 var(--c),
        calc(-7 * var(--w)) 0 var(--c),
        calc(-8 * var(--w)) -10px var(--c),
        calc(-9 * var(--w)) 0 var(--c);
    }
    90.90% {
      text-shadow:
        calc(0 * var(--w)) 0 var(--c),
        calc(-1 * var(--w)) 0 var(--c),
        calc(-2 * var(--w)) 0 var(--c),
        calc(-3 * var(--w)) 0 var(--c),
        calc(-4 * var(--w)) 0 var(--c),
        calc(-5 * var(--w)) 0 var(--c),
        calc(-6 * var(--w)) 0 var(--c),
        calc(-7 * var(--w)) 0 var(--c),
        calc(-8 * var(--w)) 0 var(--c),
        calc(-9 * var(--w)) -10px var(--c);
    }
  }
</style>
