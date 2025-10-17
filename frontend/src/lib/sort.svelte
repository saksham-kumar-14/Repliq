<script lang="ts">
    import { get } from "svelte/store";
    import { comments } from "../store/post";

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
    }

    function sortByUpvotes() {
        let tmp = get(comments);
        if (tmp == null) tmp = [];
        tmp.sort((a, b) => b.upvotes - a.upvotes);
        comments.set(tmp);
    }
</script>

<div>
    <div>
        Sory by:
        <button on:click={sortByNewest}>Newest</button>
        <button on:click={sortByOldest}>Oldest</button>
        <button on:click={sortByUpvotes}>Upvotes</button>
    </div>
</div>

<style></style>
