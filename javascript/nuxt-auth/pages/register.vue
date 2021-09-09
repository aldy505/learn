<template>
  <div class="bg-gray-200 h-screen">
      <div class="container px-56 py-16">
          <div class="border rounded-lg shadow-lg bg-white py-4 px-8">
              <h1 class="text-4xl text-purple-700">Register yourself!</h1>
              <form @submit.prevent="formSubmit">
                  <div class="flex flex-row items-center py-2">
                      <div class="flex-1">Full name: </div>
                      <div class="flex-2"><input type="text" v-model="register.name" required name="nameInput" id="nameInput" class="shadow bg-gray-100  appearance-none border border-gray-500 focus:border-purple-600 rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"></div>
                  </div>
                  <div class="flex flex-row items-center py-2">
                      <div class="flex-1">Email:</div>
                      <div class="flex-2"><input type="email" v-model="register.email" required name="emailInput" id="emailInput" class="shadow bg-gray-100  appearance-none border border-gray-500 focus:border-purple-600 rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"></div>
                  </div>
                  <div class="flex flex-row items-center py-2">
                      <div class="flex-1">Password:</div>
                      <div class="flex-2">
                        <input type="password" v-model="register.password" required name="passwordInput" id="passwordInput" class="shadow bg-gray-100  appearance-none border border-gray-500 focus:border-purple-600 rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline">
                        <div class="py-1 text-xs">{{ checkPassStrength(register.password) }}</div>
                      </div>
                  </div>
                  
                  <div class="bg-red-200 text-black text-base font-bold rounded-md my-2 p-4" v-show="isThereAnError" :key="isThereAnError">
                    Oops! {{ errorMessage }}
                  </div>
                  <input type="hidden" name="_csrf" v-model="register._csrf">
                  <div class="flex py-2">
                      <div class="flex-1">
                          <button type="submit" value="Register" class="w-100 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">Register</button>
                      </div>
                      <div class="flex-2">
                        <p class="text-right"></p>
                      </div>
                  </div>
              </form>
          </div>
      </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  layout: 'default',
  data() {
    return {
      isThereAnError: false,
      errorMessage: '',
      register: {
        _csrf: '',
        password: '',
        email: '',
        name: ''
      }
    }
  },
  async created() {
    if (this.authorized) {
      console.log('Index.vue logic line 74: You are authorized')
      this.$router.push({
        name: 'dashboard'
      })
    } else {
      console.log('Index.vue logic line 77: You are not authorized (which is good if you first came in here)');
      try {
        /**
         * Fetching CSRF data from the API.
         * For more documentation on CSRF: https://github.com/expressjs/csurf 
         */
        await axios.get('http://localhost:3000/api/preform')
          .then(res => {
            this.register._csrf = res.data.csrfToken
          })
          .catch(err => {
            this.isThereAnError = true
            console.log('Index.vue logic line 88: Catch from CSRF fetch, see error details below:')
            this.errorMessage = 'Index.vue logic line 88: Catch from CSRF fetch, see error details in console.'
            console.log(err)
          })
      } catch (error) {
        this.isThereAnError = true
        this.errorMessage = error
        console.log(error)
      }
    }
  },
  methods: {
    async formSubmit() {
      try {
        /**
         * This means calling the login action from ../store/auth.js
         * The action fetch data from the API, store it globally in Vuex State, and return a Promise.
         * For more documentation on Vuex: https://vuex.vuejs.org/ 
         */
        await this.$store.dispatch('auth/register', {
            name: this.register.name,
            email: this.register.email,
            password: this.register.password,
            csrf: this.register._csrf
          }).then(res => {
            this.$router.push({
              name: 'index'
            })
          })
          .catch(err => {
            this.isThereAnError = true
            this.errorMessage = err
            console.log('Register.vue logic line 109: error on store dispatch. See error notes below.');
            console.log(err)
          })
      } catch (err) {
        this.isThereAnError = true
        this.errorMessage = err
        console.log('Register.vue logic line 115: error on trycatch. See error notes below.')
        console.log(err)
      }
    },
    scorePassword(pass) {
      var score = 0;
      if (!pass)
        return score;

      // award every unique letter until 5 repetitions
      var letters = new Object();
      for (var i = 0; i < pass.length; i++) {
        letters[pass[i]] = (letters[pass[i]] || 0) + 1;
        score += 5.0 / letters[pass[i]];
      }

      // bonus points for mixing it up
      var variations = {
        digits: /\d/.test(pass),
        lower: /[a-z]/.test(pass),
        upper: /[A-Z]/.test(pass),
        nonWords: /\W/.test(pass),
      }

      var variationCount = 0;
      for (var check in variations) {
        variationCount += (variations[check] == true) ? 1 : 0;
      }
      score += (variationCount - 1) * 10;

      return parseInt(score);
    },
    checkPassStrength(pass) {
      var score = this.scorePassword(pass);
      if (score > 95)
        return "Woah, that's one heck of a password!";
      if (score > 80)
        return "Nice password you got there."
      if (score > 45)
        return "Can you make it a little bit more secure?";
      if (score >= 20)
        return "Oh come on, you're better than this.";

      return "";
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