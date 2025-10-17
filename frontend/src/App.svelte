<script lang="ts">
    import Navbar from "./lib/Navbar.svelte";
    import Dashboard from "./lib/Dashboard.svelte";
    import Auth from "./lib/Auth.svelte";
    import { onMount, onDestroy } from "svelte";
    import { checkAuth, isLoggedIn } from "./store/auth";

    onMount(() => {
        checkAuth();
    });
</script>

<main>
    <Navbar />

    {#if $isLoggedIn}
        <Dashboard />
    {:else}
        <Auth />
    {/if}
</main>

<style>
    main {
        min-height: 100vh;
        background: linear-gradient(
            135deg,
            #1e1e1e,
            #2c2c2c,
            #3a3a3a,
            #4a4a4a,
            #5a5a5a
        );
        color: #f0f0f0;
        display: flex;
        flex-direction: column;
        align-items: center;
        padding: 0;
        margin: 0;
        font-family: sans-serif;
    }
</style>
