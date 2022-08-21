import axios from 'axios'
import create from 'zustand'
import {devtools, persist} from 'zustand/middleware'
import { Todo } from '../model'

interface TodoStore {
    todos: Todo[]
    fetchTodos: () => void
    addTodo: (todo: Todo) => void
    editTodoContent: (todoId: string, content: string) => void
    deleteTodo: (todoId: string) => void
    completeTodo: (todo: Todo) => void
}

export const useTodoStore = create<TodoStore>()(
    devtools(
        persist(
            (set) => ({
                todos: [],
                fetchTodos: async () => {
                    const response = await axios.get('http://localhost:3001/todos')
                    if (response.status === 200) {
                        set({todos: response.data.data})
                    } else {
                        alert("Unable to fetch todo from backend! Please check if backend is running!!!")
                    }
                },
                addTodo: async (todo) => {
                    const response = await axios.post('http://localhost:3001/todos', {
                        content: todo.content
                    });

                    if (response.status === 201) {
                        const todoId = response.data.id
                        if (todoId) {
                            set ( state => ({todos: [{id: todoId, content: todo.content, isDone: false}, ...state.todos]}))
                            return
                        } 
                    } 
                    alert("Error during saving todo!")
                },
                editTodoContent: async (todoId, content) => {
                    const response = await axios.put(`http://localhost:3001/todos/${todoId}`, {
                        content: content
                    });

                    if (response.status === 200) {
                        set(state => ({todos: state.todos.map(t => t.id === todoId ? {...t, content: content} : t)}))
                    } else {
                        alert("Error during todo content update!")
                    }
                },
                deleteTodo: async (todoId) => {
                    const response = await axios.delete(`http://localhost:3001/todos/${todoId}`);

                    if (response.status === 200) {
                        set(state => ({todos: state.todos.filter(t => t.id !== todoId)}))
                    } else {
                        alert("Error during deleteing todo!")
                    }
                },
                completeTodo: async (todo) => {
                    const response = await axios.put(`http://localhost:3001/todos/${todo.id}`, {
                        is_done: "" + !todo.isDone
                    });
                    
                    if (response.status === 200) {
                        set(state => ({todos: state.todos.map(t => t.id === todo.id ? {...t, isDone: !t.isDone} : t)}))
                    } else {
                        alert("Error during maring todo as complete!")
                    }
                }
            })
        )
    )
)