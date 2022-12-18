<script lang="ts">
  import { link } from "svelte-spa-router";
  import { timeTo } from "../utils/time";
  import Icon from "./DoneStatusIcon.svelte";
  export let done: boolean;
  export let message: string;
  export let date: string;

  const toggle = () => {
    console.log(done);
    done = !done;
  };
</script>

<!-- Only render if messege is present -->
{#if message}
  <div class="relative p-4 flex flex-col border-b">
    <div class="flex items-center">
      <button on:click={toggle}>
        <Icon {done} />
      </button>
      <a href="/todo/123" use:link>
        <p
          class="text-3xl text-left font-light hover:cursor-pointer {done
            ? 'line-through text-gray-400'
            : ''} hover:underline"
        >
          {message}
        </p>
      </a>
    </div>
    <div
      class="absolute text-sm text-gray-400 text-right italic right-4 bottom-1"
    >
      Created {timeTo(date)}
    </div>
  </div>
{/if}
