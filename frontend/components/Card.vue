<template>
  <li class="card">
      <article
        @click="flipCard(card, $event)"
      >
        <div class="bg-white max-w-xs shadow-lg   mx-auto border-b-4 border-indigo-500 rounded-2xl overflow-hidden  hover:shadow-2xl transition duration-500 transform hover:scale-105 cursor-pointer" >

          <div class="bg-blue-600 pt-1 px-2 bg-gradient-to-b from-indigo-100 to-indigo-200 rounded-2xl shadow-lg w-52">
            <a href="#" class="font-display max-w-sm text-1xl font-bold leading-tight">
              <span class="link link-underline link-underline-black text-black"> {{ card.front }} </span>
            </a>
          </div>
          </br>
          <div class="bg-white-100 pt-1 px-2 bg-gradient-to-b from-white-100 to-white-50 rounded-xs shadow-lg w-52">
            <a href="#" class="font-display max-w-sm text-1xs font-bold leading-tight">
              <span class="link link-underline link-underline-black text-black"> {{card.back}} </span>
            </a>
          </div>
        </div>

        <br/>
<!--        <div class="bg-blue-600 pt-1 px-2 bg-gradient-to-b from-indigo-100 to-indigo-200 rounded-xl shadow-lg w-52">-->
<!--        <a href="#" class="font-display max-w-sm text-2xl font-bold leading-tight">-->
<!--          <span class="link link-underline link-underline-black text-black"> {{ card.flipped ? card.back : card.front }} </span>-->
<!--        </a>-->
<!--        </div>-->

        <br/>
        <div class="p-4">

          <div class="flex justify-between mt-2">

<!--            <a @click="likeCard(card)" class="border-2 rounded-lg border-pink-500 font-bold text-pink-500 px-4 py-3 transition duration-300 ease-in-out hover:bg-pink-500 hover:text-white mr-6">-->
<!--              ❤️-->
<!--            </a>-->
<!--            <a @click="likeCard(card)" class="border-2 rounded-full border-pink-500 font-bold text-pink-500 px-4 py-3 transition duration-300 ease-in-out hover:bg-pink-500 hover:text-white mr-6">-->
<!--             {{ card.liked ?  card.dislike : card.symbollike }}-->
<!--            </a>-->
            <a href="" @click="deleteCard(card.id)" class="border-2 rounded-full border-green-500 font-bold text-green-500 px-4 py-3 transition duration-300 ease-in-out hover:bg-green-500 hover:text-white mr-6">
              🗑️
            </a>

          </div>
        </div>

      </article>
  </li>
</template>

<script>
export default {
  name: "Card",
  props: ["card"],
  methods: {
    flipCard(card) {
        card.flipped = !card.flipped;
    },
    deleteCard(id) {
      this.$emit("deleteTrigger", id);
    },
    likeCard(card) {
      card.liked = !card.liked;
      this.$emit("likeTrigger", !!card.liked);
    }
  }
};
</script>
<style>
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
</style>


