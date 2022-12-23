<script lang="ts">
  import { onMount } from "svelte";
  import { link } from "svelte-spa-router";
  import Time from "svelte-time";
  import { toggleDone } from "../stores/store";
  import { fetchOneTodo, setDone } from "./api";
  import Icon from "./DoneStatusIcon.svelte";
  export let id: string;
  export let done: boolean = false;
  export let title: string = "";
  export let createdAt: string = "";
  let doneAt: string = "";

  const checkDone = async (done: boolean) => {
    if (done) {
      const value = await fetchOneTodo(id);
      doneAt = value?.done_at;
    }
  };

  onMount(() => checkDone(done));

  const toggle = () => {
    done = !done;
    toggleDone(id);
    setDone(id, done)
      .then((v) => (doneAt = v?.done_at))
      .catch((err) => console.error(err));
  };
</script>

<!-- Only render if messege is present -->
{#if title}
  <div class="relative py-5 flex flex-col border-b">
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
      {#if done && doneAt}
        Done <Time live relative timestamp={doneAt} />
      {:else}
        Created <Time live relative timestamp={createdAt} />
      {/if}
    </div>
  </div>
{/if}
