<template>
  <div class="p-8">
    <div v-if="authorized" :key="authorized">
    <h1 class="text-4xl">This is for seeing user list from the lowdb database, you might as well can delete user records
    </h1>
    <p class="text-base">Navigate through here:</p>
    <p class="text-base">
      <nuxt-link class="text-blue-700 hover:text-blue-800 hover:underline" to="/">Home page</nuxt-link>
      &nbsp;/&nbsp;
      <nuxt-link class="text-blue-700 hover:text-blue-800 hover:underline" to="blog">Blog page</nuxt-link>
    </p>
    </div>
    <div v-else>
      <h1 class="text-4xl">You're not logged in</h1>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
export default {
    computed: {
        ...mapState({
            authorized: (state) => state.auth.authorized
        })
    },
    created() {
    }
}
</script>

<style lang="sass" scoped></style>
