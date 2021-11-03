<template>
  <div>
    <PostList :posts="fetchedPosts"/>

  </div>
</template>
<script>
import PostList from "@/components/post/PostList"

export default {
  components: {
    PostList,

  },
  created() {

  },
  computed : {
    fetchedPosts(){
      console.log("pages/index.vue")
      return this.$store.getters.getPosts
    }
  },
}
</script>
<style scoped>
.card {
  @apply p-10 rounded-lg text-emerald-500 bg-emerald-100 shadow-lg;
.title {
  @apply text-6xl font-bold;
}
}
</style>
