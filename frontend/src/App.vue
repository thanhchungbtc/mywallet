<template>
  <div id="app">
    <div id="nav">
      <v-app id="inspire">

        <app-toolbar class="app--drawer"></app-toolbar>

        <v-content>
          <router-view></router-view>
        </v-content>

      </v-app>
    </div>
    <router-view/>
  </div>
</template>

<script>

  import AppToolbar from './components/AppToolbar'

  export default {
    name: 'app',
    components: {
      AppToolbar,
    }
  }
</script>

<style>
  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
  }
</style>
