<script lang="ts">
    import { get } from "svelte/store";
    import { comments } from "../store/post";

    let currSort = "oldest";

    function sortByNewest() {
        let tmp = get(comments);
        if (tmp == null) tmp = [];
        tmp.sort((a, b) => {
            return (
                new Date(b.updated_at).getTime() -
                new Date(a.updated_at).getTime()
            );
        });
        comments.set(tmp);
        currSort = "newest";
    }

    function sortByOldest() {
        let tmp = get(comments);
        if (tmp == null) tmp = [];
        tmp.sort((a, b) => {
            return (
                new Date(a.updated_at).getTime() -
                new Date(b.updated_at).getTime()
            );
        });
        comments.set(tmp);
        currSort = "oldest";
    }

    function sortByUpvotes() {
        let tmp = get(comments);
        if (tmp == null) tmp = [];
        tmp.sort((a, b) => b.upvotes - a.upvotes);
        comments.set(tmp);
        currSort = "upvotes";
    }
</script>

<div>
    <div>
        Sort by:
        <button
            class={currSort == "newest" && "highlight"}
            on:click={sortByNewest}>Newest</button
        >
        <button
            class={currSort == "oldest" && "highlight"}
            on:click={sortByOldest}>Oldest</button
        >
        <button
            class={currSort == "upvotes" && "highlight"}
            on:click={sortByUpvotes}>Upvotes</button
        >
    </div>
</div>

<style>
    div > div {
        display: flex;
        align-items: center;
        gap: 10px;
        background: #1e1e1e;
        padding: 10px 15px;
        border-radius: 8px;
        margin-bottom: 1rem;
        font-family: sans-serif;
        color: #f0f0f0;
        box-shadow: 0 2px 6px rgba(0, 0, 0, 0.5);
    }

    div > div button {
        padding: 6px 12px;
        border: none;
        border-radius: 6px;
        background-color: #646cff;
        color: #fff;
        cursor: pointer;
        font-weight: 500;
        transition: all 0.2s ease;
    }

    div > div button:hover {
        background-color: #5058d4;
        transform: scale(1.05);
    }

    div > div button:active {
        transform: scale(0.95);
    }

    .highlight {
        background-color: white;
        color: black;
    }
</style>
