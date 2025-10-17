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
    div {
        padding: 2rem;
    }
    .give-comment {
        display: flex;
        gap: 0.5rem;
        margin-bottom: 1rem;
        width: 100%;
    }

    .give-comment input {
        flex: 1;
        padding: 10px;
        border-radius: 6px;
        border: 1px solid #555;
        background-color: #1e1e1e;
        color: #f0f0f0;
        font-size: 1rem;
    }

    .give-comment input:focus {
        outline: none;
        border-color: #646cff;
        box-shadow: 0 0 5px #646cff;
    }

    .give-comment button {
        padding: 10px 16px;
        border: none;
        border-radius: 6px;
        background-color: #646cff;
        color: #fff;
        cursor: pointer;
        font-weight: 500;
        transition: all 0.2s ease;
    }

    .give-comment button:hover {
        background-color: #5058d4;
        transform: scale(1.05);
    }

    .comments {
        margin-top: 1rem;
        display: flex;
        flex-direction: column;
        gap: 1rem;
        width: 100%;
    }

    .comments p {
        color: #ccc;
    }
</style>
