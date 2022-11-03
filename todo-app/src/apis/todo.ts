interface Response {
  message: string;
  data?: any;
}

const API_ENDPOINT = "https://api.aws-demo.authorizer.dev/todo";

export const getTodo = async (token: string): Promise<Response> => {
  try {
    const res = await fetch(API_ENDPOINT, {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    });

    const json = await res.json();

    if (res.status >= 400) {
      throw new Error(json);
    }

    return json;
  } catch (err) {
    throw err;
  }
};

export const addTodo = async (
  token: string,
  title: string
): Promise<Response> => {
  try {
    const res = await fetch(API_ENDPOINT, {
      method: "POST",
      body: JSON.stringify({
        title: title,
      }),
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    });

    const json = await res.json();

    if (res.status >= 400) {
      throw new Error(json);
    }

    return json;
  } catch (err) {
    throw err;
  }
};

export const updateTodo = async (
  token: string,
  todoId: string,
  isCompleted: boolean
): Promise<Response> => {
  try {
    const res = await fetch(`${API_ENDPOINT}/${todoId}`, {
      method: "PUT",
      body: JSON.stringify({
        is_completed: isCompleted,
      }),
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    });

    const json = await res.json();

    if (res.status >= 400) {
      throw new Error(json);
    }

    return json;
  } catch (err) {
    throw err;
  }
};

export const deleteTodo = async (
  token: string,
  todoId: string
): Promise<Response> => {
  try {
    const res = await fetch(`${API_ENDPOINT}/${todoId}`, {
      method: "DELETE",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    });

    const json = await res.json();

    if (res.status >= 400) {
      throw new Error(json);
    }

    return json;
  } catch (err) {
    throw err;
  }
};
