<script lang="ts">
    import { addComment } from "../store/post";
    import { user } from "../store/auth";
    import { get } from "svelte/store";

    export let comment;
    export let depth = 0;

    let showReplyBox = false;
    let replyText = "";

    function handleAddReply() {
        const currentUser = get(user);
        if (currentUser && currentUser.user_id && replyText.trim() !== "") {
            addComment(currentUser.user_id, replyText, comment.id);
            replyText = "";
            showReplyBox = false;
        }
    }
</script>

<div class="comment" style="margin-left: {depth * 1}rem;">
    <p>{comment.text}</p>

    {#if showReplyBox}
        <div class="reply-box">
            <input
                type="text"
                placeholder="Write your reply..."
                bind:value={replyText}
            />
            <button on:click={handleAddReply}>Post</button>
        </div>
    {:else}
        <button class="reply-btn" on:click={() => (showReplyBox = true)}>
            Reply
        </button>
    {/if}

    {#if comment.replies && comment.replies.length > 0}
        <div class="replies">
            {#each comment.replies as reply}
                <svelte:self comment={reply} depth={depth + 1} />
            {/each}
        </div>
    {/if}
</div>

<style>
    .comment {
        border-left: 2px solid #ddd;
        padding-left: 1rem;
        margin: 0.75rem 0;
    }

    .replies {
        margin-top: 0.5rem;
    }

    .reply-box {
        display: flex;
        gap: 0.5rem;
        margin-top: 0.25rem;
    }

    .reply-btn {
        background: none;
        border: none;
        color: #007bff;
        cursor: pointer;
        margin-top: 0.25rem;
    }

    .reply-btn:hover {
        text-decoration: underline;
    }
</style>
