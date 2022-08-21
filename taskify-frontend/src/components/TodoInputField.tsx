import React, { useRef, useState } from 'react'
import { useTodoStore } from '../app/TodoStore'
import './styles.css'

export const TodoInputField = () => {
  const [todo, setTodo] = useState<string>("")
  const { addTodo } = useTodoStore((state) => ({addTodo: state.addTodo}))

  const handleTodoSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (todo) {
      addTodo({id: Date.now().toString(), content: todo, isDone: false})
      setTodo("")
    }
  }

  const inputRef = useRef<HTMLInputElement>(null)
  return (
    <form className="input" onSubmit={(e) => {
        handleTodoSubmit(e);
        inputRef.current?.blur();
        }}>
        <input ref={inputRef}
        className="input__box" 
        type="input" 
        value={todo}
        onChange={
            e => {setTodo(e.target.value)}
        }
        placeholder="Enter a task"/>
        <button className="input__submit">Go</button>
    </form>
  )
}