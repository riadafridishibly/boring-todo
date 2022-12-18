import { derived, writable } from "svelte/store";

function randomDate() {
    const start = new Date('2015-12-17T03:24:00');
    const end = new Date()
    return new Date(start.getTime() + Math.random() * (end.getTime() - start.getTime()));
}

let dummyData = []
function populate() {
    if (dummyData.length > 0) {
        return
    }
    for (let i = 0; i < 10; i++) {
        dummyData.push(
            {
                id: `${i}`,
                done: false,
                title: `Item ${i}`,
                date: randomDate().toISOString(),
                content: 'Some more title'
            },
        )
    }
}

populate()

export let todos = writable(dummyData)


export const create = async (title) => {
    console.log("CREATED CALLED:", title)
    todos.update(curr => [{
        id: `${curr.length}`,
        done: false,
        date: new Date().toISOString(),
        title: title,
        content: ""
    }, ...curr])
}

export const allReadable = derived(todos, ($todo) => {
    return $todo
})

export const update = async (id, values) => {
    todos.update(curr => {
        let index = curr.findIndex(item => item.id === id)
        if (index < 0) {
            return curr
        }

        // TODO: Check if modifying object triggers update
        // TODO: Only update specific fields, not all, like created date
        curr[index] = values
        return [...curr]
    })
}

export const toggleDone = async (id) => {
    if (!id) {
        return null
    }
    todos.update(curr => {
        let index = curr.findIndex(item => item.id === id)
        if (index < 0) {
            return null
        }

        curr[index].done = !curr[index].done
        return [...curr]
    })
}

export const deleteItem = async (id) => {
    todos.update(curr => {
        let index = curr.findIndex(item => item.id === id)
        if (index < 0) {
            return curr
        }

        // TODO: update this method
        curr.splice(index, 1);
        return [...curr]
    })
}

export const readOne = (id) => {
    if (!id) {
        return null
    }

    const { subscribe } = derived(todos, ($todo) => {
        const value = $todo.filter(item => item.id === id)
        if (value.length === 0) {
            return null
        }
        return value[0]
    })

    return { subscribe }
}

export const read = (id) => {
    console.log(typeof id)
    if (id) {
        return derived(todos, ($todo) => {
            return $todo.filter(item => item.id === id)
        })
    }
    return allReadable
}

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
        content: 'Lorem ipsum dolor sit, amet consectetur adipisicing elit.Ratione, consequuntur.Perferendis expedita saepe voluptatum est placeat ullam accusantium quidem magnam!'
    })

    setTimeout(() => {
        set(
            {
                id: id,
                done: false,
                title: `Item ${id}`,
                date: new Date("10-Dec-2022").toUTCString(),
                content: 'Lorem ipsum dolor sit, amet consectetur adipisicing elit.Ratione, consequuntur.Perferendis expedita saepe voluptatum est placeat ullam accusantium quidem magnam!'
            }
        )
    }, 2000)

    return {
        subscribe,
        set,
        update
    }
}

