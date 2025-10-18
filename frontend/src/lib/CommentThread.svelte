<script lang="ts">
    import { addComment, comments } from "../store/post";
    import { user } from "../store/auth";
    import { get } from "svelte/store";
    import timeDiff from "../utils/timeDiff";
    import { ArrowUp, Plus, Minus } from "lucide-svelte";

    export let comment: any;
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

    let upvotes_reply_id = 0;
    let upvotes = 0;
    async function handleUpvotes() {
        const updates = {
            upvotes: upvotes + 1,
        };
        if (upvotes_reply_id != 0) {
            const token = localStorage.getItem("token");
            const res = await fetch(
                `${import.meta.env.VITE_BACKEND_URL}/post/${upvotes_reply_id}`,
                {
                    method: "PATCH",
                    headers: {
                        "Content-Type": "application/json",
                        Authorization: `Bearer ${token}`,
                    },
                    body: JSON.stringify(updates),
                },
            );
            const data = await res.json();

            let tmp = get(comments);
            if (tmp == null) tmp = [];
            for (let i = 0; i < tmp.length; i++) {
                if (tmp[i].id == upvotes_reply_id) {
                    tmp[i].upvotes = upvotes + 1;
                }
            }
            comments.set(tmp);
        }
    }

    let showConvo = true;
</script>

<div class="comment" style="margin-left: {depth * 1}rem;">
    <div class="comment-user">
        {#if showConvo}
            <button
                on:click={() => {
                    showConvo = false;
                }}
            >
                <Minus />
            </button>
        {:else}
            <button
                on:click={() => {
                    showConvo = true;
                }}
            >
                <Plus />
            </button>
        {/if}
        <img class="profile-img" src={comment.avatar} alt="profile" />
        <p>{comment.username}</p>
        <span>{timeDiff(comment.updated_at)}</span>
    </div>
    <p>{comment.text}</p>
    <div class="comment-upvotes">
        <span class="upvotes">{comment.upvotes}</span>
        <button
            on:click={() => {
                upvotes = comment.upvotes;
                upvotes_reply_id = comment.id;
                comment.upvotes++;
                handleUpvotes();
            }}
        >
            <ArrowUp size={18} />
        </button>
    </div>

    {#if showReplyBox}
        <div class="reply-box">
            <input
                type="text"
                placeholder="Write your reply..."
                bind:value={replyText}
            />
            <button on:click={handleAddReply}>Post</button>
            <button
                on:click={() => {
                    showReplyBox = false;
                    replyText = "";
                }}>Cancel</button
            >
        </div>
    {:else}
        <button class="reply-btn" on:click={() => (showReplyBox = true)}>
            Reply
        </button>
    {/if}

    {#if comment.replies && comment.replies.length > 0}
        {#if showConvo}
            <div class="replies">
                {#each comment.replies as reply}
                    <svelte:self comment={reply} depth={depth + 1} />
                {/each}
            </div>
        {/if}
    {/if}
</div>

<style>
    button {
        border: none;
        border-radius: 1000rem;
        background-color: transparent;
        color: white;
        cursor: pointer;
    }
    .comment {
        border-left: 2px solid #646cff;
        padding-left: 1rem;
        margin: 0.75rem 0;
        background: #1e1e1e;
        border-radius: 8px;
        padding: 10px 12px;
        color: #f0f0f0;
        font-family: sans-serif;
    }

    .comment-user {
        display: flex;
        align-items: center;
        gap: 8px;
        margin-bottom: 4px;
    }

    .profile-img {
        width: 36px;
        height: 36px;
        border-radius: 50%;
        object-fit: cover;
        border: 1px solid #646cff;
    }

    .comment-user p {
        font-weight: 600;
        margin: 0;
    }

    .comment-user span {
        font-size: 0.75rem;
        color: #aaa;
    }

    .comment-upvotes {
        display: flex;
        align-items: center;
        gap: 4px;
        margin-top: 4px;
    }

    .comment-upvotes span {
        font-size: 0.85rem;
    }

    .comment-upvotes button {
        background-color: #646cff;
        border: none;
        color: #fff;
        padding: 2px 6px;
        border-radius: 4px;
        cursor: pointer;
        font-weight: bold;
        transition:
            background-color 0.2s ease,
            transform 0.1s ease;
    }

    .comment-upvotes button:hover {
        background-color: #5058d4;
        transform: scale(1.1);
    }

    .reply-btn {
        background: none;
        border: none;
        color: #646cff;
        cursor: pointer;
        margin-top: 4px;
        font-size: 0.85rem;
        transition: color 0.2s ease;
    }

    .reply-btn:hover {
        color: #5058d4;
        text-decoration: underline;
    }

    .reply-box {
        display: flex;
        gap: 0.5rem;
        margin-top: 0.5rem;
    }

    .reply-box input {
        flex: 1;
        padding: 8px;
        border-radius: 6px;
        border: 1px solid #555;
        background-color: #2c2c2c;
        color: #f0f0f0;
        font-size: 0.9rem;
    }

    .reply-box input:focus {
        outline: none;
        border-color: #646cff;
        box-shadow: 0 0 5px #646cff;
    }

    .reply-box button {
        padding: 8px 12px;
        border: none;
        border-radius: 6px;
        background-color: #646cff;
        color: #fff;
        cursor: pointer;
        transition: all 0.2s ease;
    }

    .reply-box button:hover {
        background-color: #5058d4;
        transform: scale(1.05);
    }

    .replies {
        margin-top: 0.5rem;
        margin-left: 1rem;
    }

    .upvotes {
        font-weight: 600;
        margin-right: 5px;
    }
</style>
