<script lang="ts">
  import { link } from "svelte-spa-router";
  import { toggleDone, todos } from "../stores/store";
  import { timeTo } from "../utils/time";
  import Icon from "./DoneStatusIcon.svelte";
  export let id: string;
  export let done: boolean = false;
  export let title: string = "";
  export let date: string = "";
  todos.subscribe((values) => {
    let item = values.filter((item) => item.id === id);
    if (item.length > 0) {
      console.log("Value", item[0]);
      done = item[0].done;
    }
  });
</script>

<!-- Only render if messege is present -->
{#if title}
  <div class="relative p-4 flex flex-col border-b">
    <div class="flex items-center">
      <button on:click={() => toggleDone(id)}>
        <Icon {done} />
      </button>
      <a href="/todo/{id}" use:link>
        <p
          class="text-2xl text-left font-light hover:cursor-pointer {done
            ? 'line-through text-gray-400'
            : ''} hover:underline"
        >
          {title}
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
