import { derived, writable } from "svelte/store";
import { createTodo, fetchAllTodos } from "../lib/api";

export let todos = writable([]);
export let refetch = writable(false);

refetch.subscribe(async (currentValue) => {
  if (currentValue) {
    todos.set(await fetchAllTodos());
    refetch.set(false);
  }
});

export const create = async (title) => {
  const res = await createTodo(title);
  todos.update((curr) => [{ ...res }, ...curr]);
};

export const toggleDone = async (id) => {
  todos.update((curr) => {
    let index = curr.findIndex((item) => item.id === id);
    if (index < 0) {
      return null;
    }

    curr[index].done = !curr[index].done;
    let newArray = [...curr];
    return newArray;
  });
};
