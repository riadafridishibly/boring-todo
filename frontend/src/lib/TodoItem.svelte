<script lang="ts">
  import { link } from "svelte-spa-router";
  import Time from "svelte-time";
  import { refetch, toggleDone, todos, watchDone } from "../stores/store";
  import { setDone } from "./api";
  import Icon from "./DoneStatusIcon.svelte";
  export let id: string;
  export let done: boolean = false;
  export let title: string = "";
  export let createdAt: string = "";
  const toggle = () => {
    done = !done;
    toggleDone(id);
    setDone(id, done)
      .then((v) => console.log(v))
      .catch((err) => console.error(err));
  };
</script>

<!-- Only render if messege is present -->
{#if title}
  <div class="relative p-4 flex flex-col border-b">
    <div class="flex items-center">
      <button on:click={() => toggle()}>
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
      <Time live relative timestamp={createdAt} />
    </div>
  </div>
{/if}
