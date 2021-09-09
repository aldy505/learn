<template>
  <div class="p-8">
    <div v-if="authorized" :key="authorized">
      <h1 class="text-4xl">You're logged in!</h1>
      <p class="text-base">Well I guess this is still empty until my next free time. This is your data:</p>
      <pre class="bg-gray-400 text-black my-4 p-4 lg:w-2/5 md:w-1/2 w-100 text-sm">{{$store.state.auth.userData}}</pre>
      <p class="text-base">Navigate through here:</p>
      <div class="text-base">
        <nuxt-link class="text-blue-700 hover:text-blue-800 hover:underline" to="dashboard/users">Users page</nuxt-link>
        &nbsp;/&nbsp;
        <nuxt-link class="text-blue-700 hover:text-blue-800 hover:underline" to="dashboard/blog">Blog page</nuxt-link>
      </div>
      <div class="my-4">
        <!-- logout button -->
        <form @submit.prevent="logoutSubmit">
         Click here to&nbsp;&nbsp;<button class="w-100 border-2 border-blue-800 hover:bg-blue-100 bg-white text-black py-1 px-4 rounded focus:outline-none focus:shadow-outline">Logout</button>
        </form>
        
      </div>
    </div>
    <div v-else>
      <h1 class="text-4xl">You're not logged in</h1>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import axios from 'axios'

export default {
  layout: 'default',
  data() {
    return {
      logout: {
        _csrf: ''
      }
    }
  },
  computed: {
    ...mapState({
      authorized: (state) => state.auth.authorized
    })
  },
  async created() {
    if (!this.authorized) {
      console.log('Index.vue logic line 74: You are not authorized')
    } else {
      console.log('Index.vue logic line 77: You are authorized (which is good)');
      try {
        /**
         * Fetching CSRF data from the API.
         * For more documentation on CSRF: https://github.com/expressjs/csurf 
         */
        await axios.get('http://localhost:3000/api/preform')
          .then(res => {
            this.logout._csrf = res.data.csrfToken
          })
          .catch(err => {
            console.log('Index.vue logic line 88: Catch from CSRF fetch, see error details below:')
            console.log(err)
          })
      } catch (error) {
        console.log(error)
      }
    }
  },
  methods: {
    async logoutSubmit() {
      try {
        await this.$store.dispatch('auth/logout', {
            csrf: this.logout._csrf
          })
          .then(res => {
            this.$router.push({
              name: 'index'
            })
          })
          .catch(err => {
            console.log('Index.js logic line 122: error on store dispatch. See error notes below.');
            console.log(err)
          })
      } catch (error) {
        console.log(err)
      }
    }
  }
}
</script>

<style lang="sass" scoped></style>
