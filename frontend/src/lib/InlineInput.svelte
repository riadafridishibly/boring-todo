<script lang="ts">
  import { create } from "../stores/store";
  import { link } from "svelte-spa-router";

  let value: string =
    "Lorem ipsum dolor, sit amet consectetur adipisicing elit.";
  let submittedValue = null;
  const handleSubmit = () => {
    submittedValue = value;
  };
</script>

<div class="flex items-center">
  <a href="/" class="block flex-none px-4" use:link>
    <div class="flex items-center">
      <svg
        class="w-6 h-6 stroke-1"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
        xmlns="http://www.w3.org/2000/svg"
        ><path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M15 19l-7-7 7-7"
        /></svg
      >
      <span class="text-2xl font-light px-1">Go Back</span>
    </div>
  </a>
  <form class="relative flex-grow" on:submit|preventDefault={handleSubmit}>
    <label class="inline">
      <input
        bind:value
        class="py-3 px-5 text-2xl font-light outline-none focus:border-b-2 border-b-2 border-b-gray-400 focus:border-b-black w-full"
        placeholder="Add new todo"
      />
    </label>

    <button
      class="absolute right-0 top-1/2 transform -translate-x-1/2 -translate-y-1/2"
      on:click={() => console.log("button clicked")}
    >
      <svg
        class="w-6 h-6 cursor-pointer stroke-gray-500 stroke-1 "
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
        xmlns="http://www.w3.org/2000/svg"
        ><path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z"
        /></svg
      >
    </button>
  </form>
  <button class="flex-none px-4">
    <div class="flex items-center">
      <svg
        class="w-6 h-6 stroke-1 stroke-rose-400"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
        xmlns="http://www.w3.org/2000/svg"
        ><path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
        /></svg
      >
      <!-- <span class="text-2xl font-light px-1">Delete</span> -->
    </div>
  </button>
</div>
