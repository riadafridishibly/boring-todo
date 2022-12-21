const baseAPIURI = "http://localhost:8989/api";

export const fetchAllTodos = async () => {
  return await (await fetch(baseAPIURI + "/todos")).json();
};

/**
 * @param {string} id
 */
export async function fetchOneTodo(id) {
  return await (await fetch(`${baseAPIURI}/todos/${id}`)).json();
}

/**
 * @param {string} id
 */
export async function deleteOneTodo(id) {
  return await (
    await fetch(`${baseAPIURI}/todos/${id}`, { method: "DELETE" })
  ).json();
}

/**
 * @param {string} title
 */
export async function createTodo(title) {
  return await (
    await fetch(`${baseAPIURI}/todos`, {
      method: "POST",
      body: JSON.stringify({
        title: title,
      }),
    })
  ).json();
}

/**
 * @param {string} id
 */
export async function updateTodo(id, values) {
  return await (
    await fetch(`${baseAPIURI}/todos/${id}`, {
      method: "PUT",
      body: JSON.stringify({
        ...values,
      }),
    })
  ).json();
}

/**
 * @param {string} id
 * @param {boolean} done
 */
export async function setDone(id, done) {
  return await (
    await fetch(`${baseAPIURI}/todos/${id}`, {
      method: "PUT",
      body: JSON.stringify({
        done: done,
      }),
    })
  ).json();
}
