<script lang="ts">
  import { link } from "svelte-spa-router";
  import { getAll } from "../stores/store";
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
      {#each $getAll as item}
        <li>
          <TodoItem message={item.title} done={item.done} date={item.date} />
        </li>
      {/each}
      <li>
        <TodoItem
          message={"Lorem ipsum dolor sit amet consectetur adipisicing elit."}
          done
          date="12-Nov-2022"
        />
      </li>
    </ul>
  </div>
</div>
