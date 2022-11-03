import React, { useState, useEffect } from "react";
import { useAuthorizer } from "@authorizerdev/authorizer-react";
import { addTodo, deleteTodo, getTodo, updateTodo } from "../apis/todo";

interface Todo {
  id: string;
  title: string;
  is_completed: boolean;
}

export default function Dashboard() {
  const { token } = useAuthorizer();
  const [fetchingTodos, setFetchingTodos] = useState(true);
  const [todos, setTodos] = useState<Todo[]>([]);
  const [savingTodo, setSavingTodo] = useState(false);
  const [updatingTodo, setUpdatingTodo] = useState(""); // set the updating id
  const [deletingTodo, setDeletingTodo] = useState(""); // set the deleting id
  const [err, setErr] = useState("");

  useEffect(() => {
    let isMounted = true;
    async function fetchTodo() {
      try {
        const res = await getTodo(token?.id_token || "");
        if (isMounted) {
          setTodos(res.data || []);
        }
      } catch (err) {
        if (err instanceof Error && isMounted) {
          setErr(err.message);
        }
      } finally {
        if (isMounted) {
          setFetchingTodos(false);
          setErr("");
        }
      }
    }

    fetchTodo();

    return () => {
      isMounted = false;
    };
    // eslint-disable-next-line
  }, []);

  const handleAddTask = async (e: any) => {
    e.preventDefault();
    const title = e.target.elements["task"].value || "";

    if (title.trim()) {
      try {
        setErr("");
        setSavingTodo(true);
        const res = await addTodo(token?.id_token || "", title);
        console.log({ res });
        if (res.data) {
          setTodos([res.data, ...todos]);
          e.target.reset();
        }
      } catch (err) {
        if (err instanceof Error) {
          setErr(err.message);
        }
      } finally {
        setSavingTodo(false);
      }
    } else {
      setErr("title is required");
    }
  };

  const handleUpdateTask = async (todoId: string, isCompleted: boolean) => {
    try {
      setErr("");
      setUpdatingTodo(todoId);
      const res = await updateTodo(token?.id_token || "", todoId, isCompleted);
      if (res.data) {
        setTodos(
          todos.map((i) => {
            if (i.id === todoId) {
              return res.data;
            }
            return i;
          })
        );
      }
    } catch (err) {
      if (err instanceof Error) {
        setErr(err.message);
      }
    } finally {
      setUpdatingTodo("");
    }
  };

  const handleDeleteTodo = async (todoId: string) => {
    try {
      setErr("");
      setDeletingTodo(todoId);
      await deleteTodo(token?.id_token || "", todoId);
      setTodos(todos.filter((i) => i.id !== todoId));
    } catch (err) {
      if (err instanceof Error) {
        setErr(err.message);
      }
    } finally {
      setDeletingTodo("");
    }
  };

  return (
    <div className="dashboard inner-container">
      <form onSubmit={handleAddTask}>
        <input
          name="task"
          id="task"
          type="text"
          placeholder="Enter your task"
          required
        />
        <button type="submit" disabled={savingTodo} className="btn">
          {savingTodo ? "Saving..." : "Save Task"}
        </button>
      </form>
      <>
        {Boolean(err) && <div className="err">{err}</div>}
        {fetchingTodos ? (
          <div>Fetching tasks..</div>
        ) : (
          <>
            {Boolean(todos.length) ? (
              <ul>
                {todos.map((todo) => (
                  <li key={todo.id}>
                    <div
                      className={`title ${
                        todo.is_completed ? "completed" : ""
                      }`}
                    >
                      {todo.title}
                    </div>
                    <div>
                      <button
                        disabled={todo.id === updatingTodo}
                        onClick={() =>
                          handleUpdateTask(todo.id, !todo.is_completed)
                        }
                        className="btn green"
                        style={{
                          marginRight: 10,
                        }}
                      >
                        {todo.id === updatingTodo
                          ? "Updating..."
                          : `Mark As ${
                              todo.is_completed ? "Remaining" : "Completed"
                            }`}
                      </button>
                      <button
                        className="btn red"
                        disabled={todo.id === deletingTodo}
                        onClick={() => handleDeleteTodo(todo.id)}
                      >
                        {deletingTodo === todo.id ? "Deleting..." : "Delete"}
                      </button>
                    </div>
                  </li>
                ))}
              </ul>
            ) : (
              <div>No tasks found</div>
            )}
          </>
        )}
      </>
    </div>
  );
}
