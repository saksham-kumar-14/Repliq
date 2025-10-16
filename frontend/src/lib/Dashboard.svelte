<script lang="ts">
    import { onMount } from "svelte";
    import { getPosts, comments } from "../store/post";
    import { addComment } from "../store/post";
    import { user } from "../store/auth";
    import { get } from "svelte/store";

    let comment = "";
    let parentID = 0;
    function handleAddComment(e: Event) {
        e.preventDefault();
        const currentUser = get(user);
        if (currentUser && currentUser.user_id) {
            addComment(currentUser.user_id, comment, parentID);
        }
    }

    onMount(() => {
        getPosts();
    });
</script>

<div>
    <form on:submit={handleAddComment} class="give-comment">
        <input type="text" bind:value={comment} />
        <button type="submit">Post</button>
    </form>

    <div class="comments">
        {#if $comments === null}
            <p>Loading comments...</p>
        {:else if $comments.length === 0}
            <p>No comments yet.</p>
        {:else}
            {#each $comments as comment}
                <div class="comment">
                    <p class="text">{comment.text}</p>
                    <small>Upvotes: {comment.upvotes ?? 0}</small>
                </div>
            {/each}
        {/if}
    </div>
</div>

<style>
</style>
