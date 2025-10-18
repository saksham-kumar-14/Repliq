import { writable, type Writable } from "svelte/store";

export interface User {
  user_id: number;
  email: string;
  username: string;
  avatar: string;
}

export interface AuthResp {
  valid: boolean;
  user: User;
}

export const isLoggedIn: Writable<boolean> = writable(false);
export const user: Writable<User | null> = writable(null);

export async function checkAuth(): Promise<void> {
  const token = localStorage.getItem("token");
  if (!token) {
    isLoggedIn.set(false);
    user.set(null);
    return;
  }

  try {
    const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/api/token`, {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    });

    if (!res.ok) {
      throw new Error("Authentication using token failed.");
    }

    const data: AuthResp = await res.json();

    if (data.valid) {
      const usres = await fetch(
        `${import.meta.env.VITE_BACKEND_URL}/user/${data.user?.user_id}`,
        {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
          },
        },
      );
      const usrdata = await usres.json();
      if (usrdata?.id) {
        isLoggedIn.set(true);
        data.user["avatar"] = usrdata["avatar"];
        data.user["username"] = usrdata["username"];
        user.set(data.user);
      } else {
        isLoggedIn.set(false);
        user.set(null);
      }
    } else {
      isLoggedIn.set(false);
      user.set(null);
    }
  } catch (err) {
    console.error("Auth check failed: ", err);
    isLoggedIn.set(false);
    user.set(null);
  }
}

export async function login(email: string, password: string): Promise<void> {
  try {
    const userData = {
      email: email,
      password: password,
    };

    const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/user/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(userData),
    });

    if (!res.ok) {
      throw new Error("Login failed.");
    }

    const data = await res.json();

    if (data.token) {
      localStorage.setItem("token", data.token);
      await checkAuth();
    } else {
      isLoggedIn.set(false);
      user.set(null);
    }
  } catch (err) {
    console.error("Login failed: ", err);
    isLoggedIn.set(false);
    user.set(null);
  }
}

export async function register(
  username: string,
  email: string,
  password: string,
): Promise<void> {
  try {
    const userData = {
      username: username,
      email: email,
      password: password,
    };

    const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/user`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(userData),
    });

    if (!res.ok) {
      throw new Error("Registration failed.");
    }

    const data = await res.json();

    if (data) {
      user.set(data ?? null);
      alert("Registered!");

      try {
        const loginres = await fetch(
          `${import.meta.env.VITE_BACKEND_URL}/user/login`,
          {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              email: email,
              password: password,
            }),
          },
        );

        if (!loginres.ok) {
          throw new Error("Login after registration failed.");
        }

        const logindata = await loginres.json();

        if (logindata.token) {
          isLoggedIn.set(true);
          localStorage.setItem("token", logindata.token);
          // Fetch complete user data after login
          await checkAuth();
        } else {
          isLoggedIn.set(false);
          user.set(null);
        }
      } catch (err) {
        console.error("Registered but Login failed: ", err);
        isLoggedIn.set(false);
        user.set(null);
      }
    } else {
      isLoggedIn.set(false);
      user.set(null);
    }
  } catch (err) {
    console.error("Registration failed: ", err);
    isLoggedIn.set(false);
    user.set(null);
  }
}
