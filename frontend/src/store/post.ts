import { get, writable, type Writable } from "svelte/store";
import { user } from "./auth";

export interface Comment {
  id: number;
  parent_id: number;
  text: string;
  upvotes: number;
  user_id: number;
  created_at: string;
  updated_at: string;
  username: string;
}

export const comments: Writable<Comment[] | null> = writable(null);

export async function getPosts() {
  try {
    const token = localStorage.getItem("token");
    const resp = await fetch(`/v1/post/`, {
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
    const tmp = [];

    for (let i = 0; i < data.length; i++) {
      const usres = await fetch(`/v1/user/${data[i].user_id}`);
      const usrdata = await usres.json();
      if (usrdata) {
        tmp.push({
          id: data[i].id,
          parent_id: data[i].parent_id,
          text: data[i].text,
          upvotes: data[i].upvotes,
          user_id: data[i].user_id,
          created_at: data[i].created_at,
          updated_at: data[i].updated_at,
          username: usrdata.username,
          avatar: usrdata.avatar,
        });
      }
    }

    comments.set(tmp);
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

    const res = await fetch(`/v1/post/`, {
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
      console.log(get(user));
      data["avatar"] = get(user)?.avatar;
      data["username"] = get(user)?.username;
      temp?.push(data);
      comments.set(temp);
    } else {
      alert("Unable to post");
    }
  } catch (err) {
    console.error("Failed to post: ", err);
  }
}
