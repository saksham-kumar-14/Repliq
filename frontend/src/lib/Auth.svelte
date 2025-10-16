<script lang="ts">
    import Login from "./Login.svelte";
    import Register from "./Register.svelte";

    let login = true;
</script>

<div class="div">
    {#if login}
        <div>
            <Login />
            <p>
                Don't have an account?
                <span
                    on:click={() => {
                        login = false;
                    }}
                >
                    Register
                </span>
            </p>
        </div>
    {:else}
        <div>
            <Register />
            <p>
                Already have an account?
                <span
                    on:click={() => {
                        login = true; // switch to Login
                    }}
                >
                    Login
                </span>
            </p>
        </div>
    {/if}
</div>

<style>
    .div {
        max-width: 400px;
        margin: auto;
        padding: 2rem;
    }
    span {
        cursor: pointer;
        color: #646cff;
        text-decoration: underline;
    }
</style>
