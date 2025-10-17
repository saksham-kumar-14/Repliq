<script lang="ts">
    import { login, register } from "../store/auth";

    let username = "";
    let email = "";
    let password = "";

    const handleSubmit = (event: Event) => {
        event.preventDefault();
        register(username, email, password);
    };
</script>

<form on:submit={handleSubmit} class="form">
    <h2>Register</h2>
    <input type="text" placeholder="Username" bind:value={username} />
    <input type="email" placeholder="Email" bind:value={email} />
    <input type="password" placeholder="Password" bind:value={password} />
    <button type="submit">Register</button>
</form>

<style>
    .form {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        background: linear-gradient(135deg, #2c2c2c, #1e1e1e);
        padding: 40px 30px;
        border-radius: 12px;
        box-shadow: 0 4px 15px rgba(0, 0, 0, 0.5);
        color: #f0f0f0;
        width: 320px;
        gap: 15px;
        font-family: sans-serif;
    }

    h2 {
        margin-bottom: 10px;
        color: #646cff;
    }

    input {
        width: 100%;
        padding: 10px;
        border-radius: 6px;
        border: 1px solid #555;
        background-color: #1e1e1e;
        color: #f0f0f0;
        font-size: 1rem;
    }

    input:focus {
        outline: none;
        border-color: #646cff;
        box-shadow: 0 0 5px #646cff;
    }

    button {
        width: 100%;
        padding: 10px;
        border: none;
        border-radius: 6px;
        background-color: #646cff;
        color: #fff;
        font-size: 1rem;
        cursor: pointer;
        transition:
            background-color 0.2s ease,
            transform 0.1s ease;
    }

    button:hover {
        background-color: #5058d4;
        transform: scale(1.02);
    }

    button:active {
        transform: scale(0.98);
    }
</style>
