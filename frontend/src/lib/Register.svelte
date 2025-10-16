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
        gap: 1rem;
    }

    input,
    button {
        padding: 0.5rem;
        width: 200px;
    }
</style>
