<script>
  // This snipped is copied and modified from here
  // https://svelte.dev/examples/modal
  import { createEventDispatcher, onDestroy } from "svelte";

  const dispatch = createEventDispatcher();
  const close = () => dispatch("close");

  let modal;

  const handle_keydown = (e) => {
    if (e.key === "Escape") {
      close();
      return;
    }

    if (e.key === "Tab") {
      // trap focus
      const nodes = modal.querySelectorAll("*");
      const tabbable = Array.from(nodes).filter((n) => n.tabIndex >= 0);

      let index = tabbable.indexOf(document.activeElement);
      if (index === -1 && e.shiftKey) index = 0;

      index += tabbable.length + (e.shiftKey ? -1 : 1);
      index %= tabbable.length;

      tabbable[index].focus();
      e.preventDefault();
    }
  };

  const previously_focused =
    typeof document !== "undefined" && document.activeElement;

  if (previously_focused) {
    onDestroy(() => {
      // @ts-ignore
      previously_focused.focus();
    });
  }
</script>

<svelte:window on:keydown={handle_keydown} />

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div
  class="fixed inset-0 h-full w-full z-20 bg-gray-400 opacity-40"
  on:click={close}
/>

<div class="modal" role="dialog" aria-modal="true" bind:this={modal}>
  <slot />
  <!-- svelte-ignore a11y-autofocus -->
  <!-- <button autofocus on:click={close}>close modal</button> -->
</div>

<style>
  .modal-background {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 999;
    background: rgba(0, 0, 0, 0.5);
  }

  .modal {
    z-index: 10000;
    position: absolute;
    left: 50%;
    top: 50%;
    width: calc(100vw - 4em);
    max-width: 32em;
    max-height: calc(100vh - 4em);
    overflow: auto;
    transform: translate(-50%, -50%);
    padding: 1em;
    border-radius: 0.2em;
    background: white;
  }

  button {
    display: block;
  }
</style>
