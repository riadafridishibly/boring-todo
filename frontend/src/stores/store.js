import { derived, writable } from "svelte/store";
import { fetchAllTodos } from "../lib/api";

function randomDate() {
  const start = new Date("2015-12-17T03:24:00");
  const end = new Date();
  return new Date(
    start.getTime() + Math.random() * (end.getTime() - start.getTime())
  );
}

let dummyData = [];
function populate() {
  if (dummyData.length > 0) {
    return;
  }
  for (let i = 0; i < 3; i++) {
    dummyData.push({
      id: `${i + 1}`,
      done: false,
      title: `Item ${i}`,
      date: randomDate().toISOString(),
      content: "Some more title",
    });
  }
}

populate();

export let todos = writable([]);
export let refetch = writable(false);

refetch.subscribe(async (currentValue) => {
  if (currentValue) {
    todos.set(await fetchAllTodos());
    refetch.set(false);
  }
});

const createApi = async (title) => {
  return await fetch("http://localhost:8989/api/todos", {
    method: "POST",
    body: JSON.stringify({
      title: title,
    }),
  });
};

const listAllApi = async () => {
  return await fetch("http://localhost:8989/api/todos", {
    method: "GET",
  });
};

const listAll = async () => {
  const allItems = await listAllApi();
  const jsons = await allItems.json();
  todos.set(jsons);
};

// listAll();

export const create = async (title) => {
  const res = await createApi(title);
  const jsonData = await res.json();
  jsonData.id = `${jsonData.id}`;
  todos.update((curr) => [{ ...jsonData }, ...curr]);
};

export const allReadable = derived(todos, ($todo) => {
  return $todo;
});

export const update = async (id, values) => {
  todos.update((curr) => {
    let index = curr.findIndex((item) => item.id === id);
    if (index < 0) {
      return curr;
    }

    // TODO: Check if modifying object triggers update
    // TODO: Only update specific fields, not all, like created date
    curr[index] = values;
    return [...curr];
  });
};

export const watchDone = (id) => {
  return {
    found: derived(todos, ($todos) => {
      return $todos.filter((item) => item.id === id).length > 0;
    }),
    done: derived(todos, ($todo) => {
      let item = $todo.filter((item) => item.id === id);
      if (item.length > 0) {
        return item[0].done;
      }
      return undefined;
    }),
    toggle: () => toggleDone(id),
  };
};

export const toggleDone = async (id) => {
  todos.update((curr) => {
    let index = curr.findIndex((item) => item.id === id);
    console.log("toggling:", id, index, curr[index]);
    if (index < 0) {
      return null;
    }

    curr[index].done = !curr[index].done;
    let newArray = [...curr];
    return newArray;
  });
};

export const deleteItem = async (id) => {
  todos.update((curr) => {
    let index = curr.findIndex((item) => item.id === id);
    if (index < 0) {
      return curr;
    }

    // TODO: update this method
    curr.splice(index, 1);
    return [...curr];
  });
};

export const readOne = (id) => {
  if (!id) {
    return null;
  }

  const { subscribe } = derived(todos, ($todo) => {
    const value = $todo.filter((item) => item.id === id);
    if (value.length === 0) {
      return null;
    }
    return value[0];
  });

  return { subscribe };
};

export const read = (id) => {
  console.log(typeof id);
  if (id) {
    return derived(todos, ($todo) => {
      return $todo.filter((item) => item.id === id);
    });
  }
  return allReadable;
};

// TODO: POC with every single item!!
/**
 * @param {string} id string
 */
export function createTodoStore(id) {
  // TODO: get data from some api call?
  const { subscribe, set, update } = writable({
    id: id,
    done: false,
    title: undefined, // `Item ${id}`,
    date: new Date("10-Dec-2022").toUTCString(),
    content:
      "Lorem ipsum dolor sit, amet consectetur adipisicing elit.Ratione, consequuntur.Perferendis expedita saepe voluptatum est placeat ullam accusantium quidem magnam!",
  });

  setTimeout(() => {
    set({
      id: id,
      done: false,
      title: `Item ${id}`,
      date: new Date("10-Dec-2022").toUTCString(),
      content:
        "Lorem ipsum dolor sit, amet consectetur adipisicing elit.Ratione, consequuntur.Perferendis expedita saepe voluptatum est placeat ullam accusantium quidem magnam!",
    });
  }, 2000);

  return {
    subscribe,
    set,
    update,
  };
}
