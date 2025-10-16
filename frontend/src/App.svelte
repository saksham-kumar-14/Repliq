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
</style>
