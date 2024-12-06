<svelte:options customElement="mod-tag" />

<script>
  import { onMount } from 'svelte';

  /** @type {HTMLElement} */
  let element;
  export let i = '';
  export let color = '';

  onMount(() => {
    if (i === '') {
      element.removeAttribute('href');
      element.style.backgroundColor = 'var(--background)'
      element.style.color = element.style.borderColor = color
    }
  });
</script>

<a bind:this={element} href="/posts/tag/{i}" style="--color: {color}">
  <slot></slot>
</a>

<style>
  a {
    padding: 0.2rem 0.4rem 0.1rem;
    min-width: 2rem;
    text-align: center;
    text-decoration: none;
    line-height: 13px;
    background-color: var(--color);
    display: inline-block;
    color: white;
    border-radius: 3px;
    border: 1px solid var(--background);
    font-size: 12px;
    transition: color 0.3s ease;
  }
  a:hover {
    color: var(--color);
    border-color: var(--color);
    background-color: var(--background);
  }
</style>
