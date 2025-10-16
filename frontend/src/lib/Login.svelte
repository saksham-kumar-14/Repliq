<script lang="ts">
    import { login } from "../store/auth";

    let email = "";
    let password = "";

    const handleSubmit = (event: Event) => {
        event.preventDefault();
        login(email, password);
    };
</script>

<form on:submit={handleSubmit} class="form">
    <h2>Login</h2>
    <input type="email" placeholder="email" bind:value={email} />
    <input type="password" placeholder="password" bind:value={password} />
    <button type="submit">Login</button>
</form>

<style>
    .form {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }
</style>
