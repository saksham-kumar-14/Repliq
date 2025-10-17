import { get, writable, type Writable } from "svelte/store";

export interface Comment {
  id: number;
  parent_id: number;
  text: string;
  upvotes: number;
  user_id: number;
  created_at: string;
  updated_at: string;
}

export const comments: Writable<Comment[] | null> = writable(null);

export async function getPosts() {
  try {
    const token = localStorage.getItem("token");
    const resp = await fetch(`${import.meta.env.VITE_BACKEND_URL}/post/`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    });

    if (!resp.ok) {
      throw new Error(`HTTP error! status: ${resp.status}`);
    }

    const data: Comment[] = await resp.json();
    comments.set(data);
  } catch (err) {
    console.error("Failed to fetch posts:", err);
    comments.set([]);
  }
}

export async function addComment(
  userID: number,
  text: string,
  parentID: number,
) {
  try {
    if (text == "") {
      throw new Error("Text can't be empty!");
    }
    const token = localStorage.getItem("token");

    const postData = {
      text: text,
      user_id: userID,
      parent_id: parentID,
    };

    const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/post/`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(postData),
    });

    if (!res.ok) {
      throw new Error("Posting failed.");
    }
    const data = await res.json();

    if (data) {
      let temp = get(comments);
      temp?.push(data);
      comments.set(temp);
    } else {
      alert("Unable to post");
    }
  } catch (err) {
    console.log("Failed to post: ", err);
  }
}
