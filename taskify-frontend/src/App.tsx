import React, { useEffect } from 'react';
import './App.css';
import { useTodoStore } from './app/TodoStore';
import { TodoInputField } from './components/TodoInputField';
import { TodoList } from './components/TodoList';

function App() {
  const fetchTodos = useTodoStore(state => state.fetchTodos)

  useEffect(() => {
    fetchTodos()
  }, [fetchTodos])

  return (
    <div className="App">
      <span className="heading">Taskify</span>
      <TodoInputField />
      <TodoList />
    </div>
  );
}

export default App;
