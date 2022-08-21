import { Todo } from "../model"
import { AiTwotoneDelete, AiFillEdit} from 'react-icons/ai'
import { MdDone } from 'react-icons/md'
import './styles.css'
import React, { useState } from "react";
import { useTodoStore } from "../app/TodoStore";

interface SingleTodoProps {
    todo: Todo;
}

export const SingleTodo = ({ todo } : SingleTodoProps) => {
    const { editTodoContent, deleteTodo, completeTodo } = useTodoStore(
        (state) => ({
            editTodoContent: state.editTodoContent,
            deleteTodo: state.deleteTodo,
            completeTodo: state.completeTodo
        })
    )

    const [edit, setEdit] = useState<Boolean>(false)
    const [editedText, setEditedText] = useState<string>(todo.content)

    const handleEdit = () => {
        if (!todo.isDone) {
            setEdit(!edit)
        }
    }

    const handleEditSubmit = (e: React.FormEvent) => {
        e.preventDefault();

        if (editedText) {
            editTodoContent(todo.id, editedText)
            setEdit(!edit)
        }
    }

    const handleDelete = () => {
        deleteTodo(todo.id)
    }

    const handleComplete = () => {
        if (edit) {
            return
        }
        completeTodo(todo)
    }

    return (
        <form className="single-todo" onSubmit={e=> handleEditSubmit(e)}>
            {edit ? 
                (
                    <input type="text" 
                    value={editedText} 
                    onChange={e => setEditedText(e.target.value)}
                    autoFocus
                    />
                )
             : 
                (
                    todo.isDone ? 
                    (
                        <s>{todo.content}</s>
                    ) :
                    (
                        <span>{todo.content}</span>
                    )
                )
            }

            <div className="single-todo__icons">
                <span className="icon" onClick={() => handleEdit()}>
                    <AiFillEdit />
                </span>
                <span className="icon" onClick={() => handleDelete()}>
                    <AiTwotoneDelete />
                </span>
                <span className="icon" onClick={() => handleComplete()}>
                    <MdDone />
                </span>
            </div>
        </form>
    )
}