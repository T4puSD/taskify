import { useTodoStore } from '../app/TodoStore';
import { SingleTodo } from './SingleTodo';
import './styles.css'

export const TodoList = () => {
    const todos = useTodoStore((state) => (state.todos))

    return (
        <div className="todos">
            {todos.map(t=><SingleTodo key={t.id} todo={t} />)}
        </div>
    )
}