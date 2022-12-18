<script lang="ts">
  import { link } from "svelte-spa-router";
  import { todos } from "../stores/store";
  import TodoInput from "./TodoInput.svelte";
  import TodoItem from "./TodoItem.svelte";
</script>

<a href="/" use:link class="block">
  <h1
    class="text-6xl w-full flex-shrink-0 text-center font-thin p-8 text-orange-600"
  >
    TODO APP
  </h1>
</a>
<div class="overflow-auto w-full h-full flex flex-col lg:w-[800px]">
  <TodoInput />
  <div
    class="py-5 scrollbar-thin scrollbar-thumb-gray-400 scrollbar-thumb-rounded-full scrollbar-track-rounded-full scrollbar-track-gray-100 overflow-y-auto"
  >
    <ul class="list-inside ">
      {#each $todos as item}
        <li>
          <TodoItem id={item.id} title={item.title} date={item.date} />
        </li>
      {/each}
    </ul>
  </div>
</div>
