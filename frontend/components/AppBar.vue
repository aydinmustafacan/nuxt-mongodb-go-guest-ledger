<template>
  <div>
    <nav>
      <div class="logo">
        <NuxtLink to="/">Mustafa Can Aydin</NuxtLink>
      </div>
      <div>

        <NuxtLink to="/messages" class="btn">Messages</NuxtLink>
        <NuxtLink to="/about" class="btn">About</NuxtLink>
      </div>
    </nav>
  </div>
</template>

<script>
export default {};
</script>

<style lang="scss" scoped>
nav {
  @apply flex justify-between items-center py-2 px-10 fixed w-full bg-white shadow-lg;

  .logo {
    @apply p-4 font-bold;
  }
  .btn {
    @apply px-4 py-2 text-gray-500 rounded-lg hover:bg-emerald-100 hover:text-emerald-700;

    &.nuxt-link-exact-active {
      @apply bg-emerald-100 text-emerald-700;
    }
  }
}
</style>
