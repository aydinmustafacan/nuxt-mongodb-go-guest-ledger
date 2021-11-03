<template>

  <div class="card">
          <div class="bg-white max-w-xs shadow-lg   mx-auto border-b-4 border-indigo-500 rounded-2xl overflow-hidden  hover:shadow-2xl transition duration-500 transform hover:scale-105 cursor-pointer" >
            <div class="bg-blue-600 pt-1 px-2 bg-gradient-to-b from-indigo-100 to-indigo-200 rounded-2xl shadow-lg w-52">
              <a href="#" class="font-display max-w-sm text-1xl font-bold leading-tight">
                <span class="link link-underline link-underline-black text-black"> {{ post.email}} </span>
              </a>
            </div>

            <div class="bg-white-100 pt-1 px-2 bg-gradient-to-b from-white-100 to-white-50 rounded-xs shadow-lg w-52">
              <a href="#" class="font-display max-w-sm text-1xs font-bold leading-tight">
                <span class="link link-underline link-underline-black text-black"> {{post.message}} </span>
              </a>
            </div>
                      <button class="btn" @click.prevent="$router.push('tobeadded')" type="button">
                        expand
                      </button>
          </div>

    </div>


</template>
<script>
    export default {
        props: {
            isAdmin: {
                required: true,
                type: Boolean
            },
            post: {
                required: true,
                type : Object
            }
        }
    }

</script>
<style lang="scss" scoped>
.card {
  @apply p-10 rounded-lg text-emerald-500 bg-emerald-100 shadow-lg;
  .title {
    @apply text-6xl font-bold;
  }
}
.link-underline {
  border-bottom-width: 0;
  background-image: linear-gradient(transparent, transparent), linear-gradient(#fff, #fff);
  background-size: 0 3px;
  background-position: 0 100%;
  background-repeat: no-repeat;
  transition: background-size .5s ease-in-out;
}

.link-underline-black {
  background-image: linear-gradient(transparent, transparent), linear-gradient(#E2C, #B2C)
}


.link-underline-purp {
  background-image: linear-gradient(transparent, transparent), linear-gradient(#F11, #F30)
}

.link-underline:hover {
  background-size: 100% 3px;
  background-position: 0 100%
}
.btn {
  @apply px-4 py-2 text-gray-500 rounded-lg hover:bg-emerald-100 hover:text-emerald-700;

  &.nuxt-link-exact-active {
    @apply bg-emerald-100 text-emerald-700;
  }
}
</style>
