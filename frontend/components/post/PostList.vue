<template>
  <li class="card">
    <article
    >
        <Post
                :post="post"
                v-for="post in posts"
                :is-admin="isAdmin"
                :key="post.id"
        />
    </article>
    </li>
</template>
<script>
    import Post from "@/components/post/Post"

    export default {
        components: {
            Post
        },
        props: {
            isAdmin: {
                type: Boolean,
                required: false,
                default: false
            },
            posts: {
                required: true,
                type: Array,
            }
        }
    }
</script>
