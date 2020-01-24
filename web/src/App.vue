<template>
  <div id="app">

    <router-view name="headertop"></router-view>
    <router-view> </router-view>

  </div>
</template>

<script>
  export default {
    name: 'app'
  }
</script>

<style lang="scss">
   @import './assets/scss/main.scss';
</style>
