<script>
  import Home from "./lib/Home.svelte";
  import Router from "svelte-spa-router";
  import TodoDetails from "./lib/TodoDetails.svelte";
  import { link } from "svelte-spa-router";

  const routes = {
    "/": Home,
    "/todo/:id": TodoDetails,
  };
</script>

<main class="w-screen h-screen overflow-x-hidden flex flex-col items-center">
  <div class="flex flex-col container items-center h-full">
    <Router {routes} />
  </div>
</main>
