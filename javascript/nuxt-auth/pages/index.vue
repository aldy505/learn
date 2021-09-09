<template>
  <div class="bg-gray-200">


    <div class="px-56 py-16">
      <div class="border rounded-lg shadow-lg bg-white mb-10 py-4 px-8">
        <h1 class="text-4xl text-purple-700 font-bold">Read me first!</h1>
        <p class="text-base whitespace-pre-line">
          For this form below, you can put any of these email: <code>reinaldy@mail.com / johndoe@mail.com /
            stevew@mail.com</code>
          With the same password: <code>password</code> / Yeah it's not that good but this is for learning/experimenting
          purpose so yeah.
          You shall be redirected to <code>/dashboard</code> once your login succeed.
        </p>

      </div>

      <div v-show="$store.state.auth.authorized" :key="$store.state.auth.authorized" class="border rounded-lg shadow-lg mb-10 bg-white py-4 px-8">
            <nuxt-link to="/dashboard" class="hover:font-bold text-lg text-purple-900">You seems like already logged in. Click here if you're not redirected!</nuxt-link>
      </div>

      <div class="border rounded-lg shadow-lg bg-white py-4 px-8">
        <h1 class="text-4xl font-bold text-purple-700">Please login</h1>
        <form @submit.prevent="formSubmit">
          <div class="flex py-2 flex-row items-center">
            <div class="flex-1">
              <p>Email</p>
            </div>
            <div class="flex-2"><input
                class="shadow bg-gray-100 appearance-none border border-gray-500 focus:border-purple-600 rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                type="email" v-model="login.email" name="emailInput" required id="emailInput"></div>
          </div>
          <div class="flex py-2 flex-row items-center">
            <div class="flex-1">
              <p>Password</p>
            </div>
            <div class="flex-2"><input
                class="shadow bg-gray-100  appearance-none border border-gray-500 focus:border-purple-600 rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                type="password" name="passwordInput" v-model="login.password" required id="passwordInput"></div>
          </div>
          <div class="bg-red-200 text-black text-base font-bold rounded-md my-2 p-4" v-show="isThereAnError" :key="isThereAnError">
            Oops! {{ errorMessage }}
          </div>
          <p class="text-base whitespace-pre-line my-2">
            Haven't registered yet? <span class="text-purple-700 hover:font-bold">
              <nuxt-link to="/register">Register now</nuxt-link>
            </span>
          </p>
          <div class="flex flex-col py-4" v-show="development">
            <div class="flex-1 text-base whitespace-pre-line">Form values: {{ login.email }} / {{ login.password }}
            </div>
            <div class="flex-1 text-base whitespace-pre-line">Error message: <span
                class="text-red-600">{{ errorMessage }}</span></div>
            <div class="flex-1 text-base whitespace-pre-line">Login status: {{ loginStatus }}</div>
            <input type="hidden" name="_csrf" v-model="login._csrf">
          </div>
          <div class="flex flex-row py-2">
            <div class="flex-1 mr-4">
              <button type="submit"
                class="w-100 bg-blue-500 hover:border-blue-700 text-white py-2 px-4 rounded focus:outline-none font-bold focus:shadow-outline">Login</button>
            </div>
          </div>
        </form>
        <div class="w-100">
          <button @click="development = !development"
            class="w-100 border-2 border-blue-800 hover:bg-blue-100 bg-white text-black py-2 px-4 rounded focus:outline-none focus:shadow-outline">See
            user input and results</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import { mapState } from 'vuex'

export default {
  layout: 'default',
  data() {
    return {
      development: false,
      isThereAnError: false,
      errorMessage: '',
      login: {
        _csrf: '',
        password: '',
        email: ''
      }
    }
  },
  async created() {
    if (this.authorized) {
      console.log('Index.vue logic line 74: You are authorized')
      this.$router.push({ name: 'dashboard' })
    } else {
      console.log('Index.vue logic line 77: You are not authorized (which is good if you first came in here)');
      try {
        /**
         * Fetching CSRF data from the API.
         * For more documentation on CSRF: https://github.com/expressjs/csurf 
         */
        await axios.get('http://localhost:3000/api/preform')
            .then(res => {
              this.login._csrf = res.data.csrfToken
            })
            .catch(err => {
              this.isThereAnError = true
              console.log('Index.vue logic line 109: Catch from CSRF fetch, see error details below:')
              this.errorMessage = 'Index.vue logic line 109: Catch from CSRF fetch, see error details in console.'
              console.log(err)
            })
      } catch (error) {
        console.log(error)
      }
    }
  },
  computed: {
    ...mapState({
      loginStatus: (state) => state.auth.loginStatus,
      authorized: (state) => state.auth.authorized,
      userEmail: (state) => state.auth.userEmail
    })    
  },
  methods: {
    async formSubmit() {
      try { 
        console.log('Index.vue logic line 128: Form submit received')
        /**
         * This means calling the login action from ../store/auth.js
         * The action fetch data from the API, store it globally in Vuex State, and return a Promise.
         * For more documentation on Vuex: https://vuex.vuejs.org/ 
         */
        await this.$store.dispatch('auth/login', {
          email: this.login.email,
          password: this.login.password,
          csrf: this.login._csrf
        }).then(res => {
            this.$router.push({name: 'dashboard'})
          })
          .catch(err => {
            this.isThereAnError = true
            this.errorMessage = err
            console.log('Index.vue logic line 144: error on store dispatch. See error notes below.');
            console.log(err)
          })
      } catch (err) {
        this.isThereAnError = true
        this.errorMessage = err
        console.log('Index.vue logic line 150: error on trycatch. See error notes below.')
        console.log(err)
      }
    }
  }
}
</script>

<style lang="sass">
.flex-2
  flex: '2 2 0%'

input:invalid
  @apply bg-red-100 border-red-600

input:valid
  @apply bg-green-100 border-green-600
</style>
