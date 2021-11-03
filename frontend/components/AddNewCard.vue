<template>
  <div>
    <div class="card">
      <label>
        <input
          v-model="front"
          :class="{ '-error': error }"
          type="text"
          class="input"
          placeholder="email"
        />
      </label>
      <label>
        <input
          @keypress.enter="addCard"
          v-model="back"
          type="text"
          class="input"
          :class="{ '-error': error }"
          placeholder="message"
        />
      </label>
      <button class="btn" @click="asyncData(front, back)" type="button">
        Add an email with a message
      </button>
    </div>
    <!-- end of add card box -->

  </div>
</template>

<script>

import axios from 'axios'
const colors = ["-orange", "-red", "-purple", "-blue", "-green"];

export default {
  name: "AddNewCard",
  data() {
    return {
      front: "",
      back: "",
      dislike: "👎",
      error: false
    }
  },
  methods: {
    async asyncData(email, message){
      axios.post("http://localhost:8080/api/guestledger",
          {email, message}
      ).then(
          (response) => {
            console.log(response.data)
          }
      );
    },
    addCard() {
      if (!this.front || !this.back) {
        this.error = true;
      } else {

        const card = {
          id: 1,
          front: this.front,
          back: this.back,
          flipped: false,
          liked: false,
          // symbollike : "❤️",
          dislike : "👎",
          color: `${colors[Math.floor(Math.random() * colors.length)]}`
        };
        this.$emit("addCardTrigger", card);
        this.front = "";
        this.back = "";
        // this.symbollike = "❤️";
        this.dislike = "👎";
        this.error = false;
      }
    }
  }
};
</script>
<style lang="scss" scoped>
.btn {
  @apply px-4 py-2 text-gray-500 rounded-lg hover:bg-emerald-100 hover:text-emerald-700;

  &.nuxt-link-exact-active {
    @apply bg-emerald-100 text-emerald-700;
  }
}
</style>

