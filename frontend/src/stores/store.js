import { derived, writable } from "svelte/store";

let a = []
function populate() {
    if (a.length > 0) {
        return
    }
    for (let i = 0; i < 10; i++) {
        a.push(
            {
                id: i,
                done: false,
                title: `Item ${i}`,
                date: new Date("10-Dec-2022").toUTCString(),
                content: 'Some more title'
            },
        )
    }
}

populate()

export const todo = writable(a)

todo.subscribe((v) => {
    console.log('Subscribe hit', v)
})

todo.update(v => {
    return v
})

export const create = async (title) => {
    todo.update(curr => [{
        id: curr.length,
        done: false,
        date: new Date().toISOString(),
        title: title,
        content: ""
    }, ...curr])
}


export const getAll = derived(todo, ($todo) => {
    return $todo
})