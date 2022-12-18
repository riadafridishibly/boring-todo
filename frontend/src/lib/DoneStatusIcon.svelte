<script>
  export let done;
</script>

{#if done}
  <svg
    xmlns="http://www.w3.org/2000/svg"
    class="w-8 h-8 mr-2 stroke-1 stroke-gray-400 cursor-pointer"
    width="24"
    height="24"
    viewBox="0 0 24 24"
    stroke-width="2"
    stroke="currentColor"
    fill="none"
    stroke-linecap="round"
    stroke-linejoin="round"
  >
    <path stroke="none" d="M0 0h24v24H0z" fill="none" />
    <circle cx="12" cy="12" r="9" />
    <path d="M9 12l2 2l4 -4" />
  </svg>
{:else}
  <svg
    xmlns="http://www.w3.org/2000/svg"
    class="w-8 h-8 mr-2 stroke-1 cursor-pointer"
    width="24"
    height="24"
    viewBox="0 0 24 24"
    stroke-width="2"
    stroke="currentColor"
    fill="none"
    stroke-linecap="round"
    stroke-linejoin="round"
  >
    <path stroke="none" d="M0 0h24v24H0z" fill="none" />
    <circle cx="12" cy="12" r="9" />
  </svg>
{/if}
