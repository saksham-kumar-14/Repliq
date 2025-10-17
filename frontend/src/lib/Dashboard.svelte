<script lang="ts">
    import { onMount } from "svelte";
    import { getPosts, comments, addComment } from "../store/post";
    import { user } from "../store/auth";
    import { get } from "svelte/store";
    import buildCommentTree from "../utils/nestedComments";
    import CommentThread from "./CommentThread.svelte";
    import Sort from "./sort.svelte";

    let comment_txt = "";
    let parentID = 0;

    function handleAddComment(e: Event) {
        e.preventDefault();
        const currentUser = get(user);
        if (currentUser && currentUser.user_id && comment_txt.trim() !== "") {
            addComment(currentUser.user_id, comment_txt, 0);
            comment_txt = "";
        }
    }

    onMount(() => {
        getPosts();
    });

    $: nestedComments = $comments ? buildCommentTree($comments) : [];
</script>

<div>
    <Sort />
    <form on:submit={handleAddComment} class="give-comment">
        <input
            type="text"
            bind:value={comment_txt}
            placeholder="Write a comment..."
        />
        <button type="submit">Post</button>
    </form>

    <!-- Comments list -->
    <div class="comments">
        {#if !$comments}
            <p>Loading comments...</p>
        {:else if $comments.length === 0}
            <p>No comments yet.</p>
        {:else}
            {#each nestedComments as comment}
                <CommentThread {comment} />
            {/each}
        {/if}
    </div>
</div>

<style>
    .give-comment {
        display: flex;
        gap: 0.5rem;
        margin-bottom: 1rem;
    }

    .comments {
        margin-top: 1rem;
    }
</style>
