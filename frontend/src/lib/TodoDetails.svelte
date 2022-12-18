<script>
  import { location, querystring } from "svelte-spa-router";
  import InlineInput from "./InlineInput.svelte";
  import "bytemd/dist/index.css";
  import { Editor, Viewer } from "bytemd";
  import gfm from "@bytemd/plugin-gfm";
  import { onDestroy } from "svelte";

  export let params;
  console.log(params);

  let value;
  const plugins = [
    gfm(),
    // Add more plugins here
  ];

  function handleChange(e) {
    value = e.detail.value;
  }

  let offHeight = 0;

  onDestroy(() => {
    console.log("I'm goin out");
  });

  console.log(offHeight);
</script>

<div class="w-full h-full flex flex-col">
  <InlineInput />
  <div bind:clientHeight={offHeight} class="h-full max-w-none">
    <Editor {value} {plugins} on:change={handleChange} />
  </div>
</div>

<style lang="postcss">
  :global(.bytemd) {
    @apply h-full p-2;
  }
</style>
