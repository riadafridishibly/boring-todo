<script lang="ts">
  import { create } from "../stores/store";

  let value: string;
  let submittedValue = null;
  const handleSubmit = () => {
    submittedValue = value;
    create(submittedValue);
    value = "";
  };
</script>

<form
  class="relative flex-shrink-0 min-w-full border-2"
  on:submit|preventDefault={handleSubmit}
>
  <label class="inline">
    <input
      bind:value
      class="py-3 px-5 text-xl outline-none focus:border-2 border-2 border-transparent focus:border-black w-full"
      placeholder="Add new todo"
    />
  </label>

  <button
    class="absolute right-0 top-1/2 transform -translate-x-1/2 -translate-y-1/2"
  >
    <svg
      class="w-9 h-9 cursor-pointer stroke-orange-500 stroke-1 "
      fill="none"
      stroke="currentColor"
      viewBox="0 0 24 24"
      xmlns="http://www.w3.org/2000/svg"
      ><path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M12 6v6m0 0v6m0-6h6m-6 0H6"
      /></svg
    >
  </button>
</form>
