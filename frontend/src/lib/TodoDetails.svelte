<script lang="ts">
  import { push } from "svelte-spa-router";
  import InlineInput from "./InlineInput.svelte";
  import "bytemd/dist/index.css";
  import { Editor } from "bytemd";
  import gfm from "@bytemd/plugin-gfm";
  import { onDestroy, onMount } from "svelte";
  import { refetch } from "../stores/store";
  import Modal from "./Modal.svelte";
  import { deleteOneTodo, fetchOneTodo, updateTodo } from "./api";

  export let params;
  console.log(params);

  let title = undefined;
  let value: string;

  onMount(async () => {
    const data = await fetchOneTodo(params.id);
    console.log(data);
    title = data?.title;
    value = data?.content;
  });

  const plugins = [gfm()];

  function handleChange(e) {
    value = e.detail.value;
  }

  async function uploadImage(f) {
    console.log(f);
    return [{ url: "https://picsum.photos/id/237/200/300" }];
  }

  let deleting = false;

  onDestroy(async () => {
    if (deleting) {
      return;
    }
    await updateTodo(params.id, {
      title: title,
      content: value,
    });
    refetch.set(true);
  });

  let showModal = false;
  const openModal = () => {
    showModal = true;
  };
  const closeModal = () => {
    showModal = false;
  };

  const onChangeInput = async (value) => {
    title = value;
  };

  const handleDelete = () => {
    deleting = true;
    refetch.set(true);
    deleteOneTodo(params.id);
    push("/");
  };
</script>

{#if showModal}
  <Modal on:close={() => (showModal = false)}>
    <div class="flex m-8 flex-col items-center">
      <!-- Header -->
      <div class="text-3xl font-light text-center select-none">
        Confirm delete?
      </div>

      <!-- Content -->
      <div class="flex mt-8 h-12 space-x-4 w-full justify-center">
        <button
          type="button"
          on:click={closeModal}
          class="w-32 border border-gray-600 bg-gray-500 h-full"
        >
          <span class="text-xl"> Close </span>
        </button>
        <button
          on:click={handleDelete}
          type="button"
          class="w-32 border  bg-red-500 border-red-600 h-full"
        >
          <span class="text-xl text-black"> Delete </span>
        </button>
      </div>
    </div>
  </Modal>
{/if}

<div class="w-full h-full flex flex-col">
  <InlineInput {onChangeInput} {openModal} {title} />
  <div class="h-full prose max-w-none">
    <Editor
      uploadImages={uploadImage}
      {value}
      {plugins}
      on:change={handleChange}
    />
  </div>
</div>

<style lang="postcss">
  :global(.bytemd) {
    @apply h-full p-2;
  }
</style>
