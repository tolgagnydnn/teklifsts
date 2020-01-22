<template>
  <div id="app">

    <!--Header-->
    <appHeader> </appHeader>

    <!-- Page Content -->
    <div class="container">
        <router-view></router-view>
    </div>

  </div>
</template>

<script>
import Header from './components/Header'
export default {
  name: 'app',
  components:{
    "appHeader": Header
  }
}
</script>


<style lang="scss">
   @import './assets/scss/main.scss';
</style>
