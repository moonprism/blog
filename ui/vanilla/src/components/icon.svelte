<svelte:options customElement="mod-icon" />

<script>
  export let id = '';
</script>

{#if id === 'date'}
  <svg class="icon"><use href="#icon-date" /></svg>
{:else if id === 'tag'}
  <svg class="icon"><use href="#icon-tag" /></svg>
{/if}

<!-- License: PD. Made by Mary Akveo: https://maryakveo.com/ -->
<symbol
  id="icon-date"
  width="18px"
  height="18px"
  viewBox="0 0 24 24"
  xmlns="http://www.w3.org/2000/svg"
  ><path
    id="secondary"
    d="M3,10H21a0,0,0,0,1,0,0V20a1,1,0,0,1-1,1H4a1,1,0,0,1-1-1V10A0,0,0,0,1,3,10Z"
    style="fill: rgb(44, 169, 188); stroke-width: 2;"
  ></path><path
    id="primary"
    d="M20,21H4a1,1,0,0,1-1-1V10H21V20A1,1,0,0,1,20,21ZM21,6a1,1,0,0,0-1-1H4A1,1,0,0,0,3,6v4H21ZM16,3V7M8,3V7"
    style="fill: none; stroke: var(--foreground); stroke-linecap: round; stroke-linejoin: round; stroke-width: 2;"
  ></path></symbol
>

<!-- License: CC0. Made by SVG Repo: https://www.svgrepo.com/svg/97179/price-tag -->
<symbol
  version="1.1"
  id="icon-tag"
  xmlns="http://www.w3.org/2000/svg"
  xmlns:xlink="http://www.w3.org/1999/xlink"
  x="0px"
  y="0px"
  viewBox="0 0 297.5 297.5"
  style="enable-background:new 0 0 297.5 297.5;"
  xml:space="preserve"
>
  <g id="XMLID_48_">
    <g>
      <path
        style="fill:#BDC3C7;"
        d="M205.955,106.74v68.95H91.545v-68.95l46.99-35.24v31.3c-8.72,3.92-14.81,12.68-14.81,22.84
			c0,13.8,11.22,25.03,25.02,25.03c13.8,0,25.03-11.23,25.03-25.03c0-10.16-6.09-18.92-14.81-22.84V71.5L205.955,106.74z"
      />
      <rect x="91.545" y="196.12" style="fill:#D95459;" width="114.41" height="80.7" />
      <path
        style="fill:var(--foreground)"
        d="M226.385,101.63v185.41c0,5.64-4.58,10.21-10.22,10.21H81.325c-5.64,0-10.21-4.57-10.21-10.21V101.63
			c0-3.21,1.51-6.24,4.09-8.17l63.33-47.5v-35.5c0-5.64,4.57-10.21,10.21-10.21s10.22,4.57,10.22,10.21v35.5l63.33,47.5
			C224.865,95.39,226.385,98.42,226.385,101.63z M205.955,276.82v-80.7H91.545v80.7H205.955z M205.955,175.69v-68.95l-46.99-35.24
			v31.3c8.72,3.92,14.81,12.68,14.81,22.84c0,13.8-11.23,25.03-25.03,25.03c-13.8,0-25.02-11.23-25.02-25.03
			c0-10.16,6.09-18.92,14.81-22.84V71.5l-46.99,35.24v68.95H205.955z M153.345,125.64c0-2.54-2.06-4.6-4.6-4.6
			c-2.53,0-4.59,2.06-4.59,4.6c0,2.53,2.06,4.6,4.59,4.6C151.285,130.24,153.345,128.17,153.345,125.64z"
      />
    </g>
    <g> </g>
  </g>
</symbol>

<style>
  .icon {
    width: 1em;
    height: 1em;
    vertical-align: -0.15em;
    font-size: 17px;
  }
</style>
