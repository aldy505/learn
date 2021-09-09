<template>
  <div class="container px-24 py-6 mx-auto">
    <div v-if="error.statusCode === 404">
      <h1 class="text-4xl">Page not found</h1>
      <p>{{ error.message }}</p>
    </div>
    <div v-else>
      <h1 class="text-4xl">An error occurred</h1>
    </div>
    <nuxt-link to="/" class="text-blue-600 hover:text-blue-800 hover:underline">Go back to home page</nuxt-link>
  </div>
</template>

<script>
export default {
  props: ['error'],
  layout: 'default' // If you prefers you can set a custom layout for the error page
}
</script>