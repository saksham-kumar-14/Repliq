<script lang="ts">
    import { isLoggedIn } from "../store/auth";

    function handleLogout(e: Event) {
        e.preventDefault();
        localStorage.clear();
        window.location.reload();
    }
</script>

<nav class="nav">
    <a class="nav-logo" href="/">Repliq</a>
    {#if $isLoggedIn}
        <button on:click={handleLogout}>Logout</button>
    {/if}
</nav>

<style>
    .nav {
        position: sticky;
        top: 0;
        display: flex;
        align-items: center;
        justify-content: space-between;
        width: 96%;
        padding: 10px 20px;
        background: linear-gradient(90deg, #2c2c2c, #1e1e1e);
        border-bottom: 1px solid #646cff;
        box-shadow: 0 2px 5px rgba(0, 0, 0, 0.5);
        z-index: 100;
        font-family: sans-serif;
    }

    .nav-logo {
        font-size: 1.5rem;
        font-weight: bold;
        color: #f0f0f0;
        text-decoration: none;
        margin-left: 2rem;
    }

    button {
        background-color: #646cff;
        color: #fff;
        border: none;
        padding: 6px 12px;
        border-radius: 5px;
        cursor: pointer;
        transition: background-color 0.2s ease;
        margin-left: 2rem;
    }

    button:hover {
        background-color: #5058d4;
    }
</style>
