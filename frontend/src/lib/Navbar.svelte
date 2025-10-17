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
    {#if isLoggedIn}
        <button on:click={handleLogout}>Logout</button>
    {/if}
</nav>

<style>
    .nav {
        position: sticky;
        display: flex;
        align-items: center;
        justify-content: center;
        width: 100vw;
        border-bottom: solid #646cff 1px;
        padding: 5px;
    }
</style>
