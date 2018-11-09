<template>
  <div id="app">
    <div id="nav">
      <v-app id="inspire">

        <app-toolbar class="app--drawer"></app-toolbar>

        <v-content>
          <router-view></router-view>
        </v-content>

        <!-- Go to top -->
        <app-fab></app-fab>

      </v-app>
    </div>
  </div>
</template>

<script>

  import AppToolbar from './components/AppToolbar'
  import AppFab from './components/AppFab'

  export default {
    name: 'app',
    components: {
      AppToolbar,
      AppFab,
    }
  }
</script>

<style>
  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    color: #2c3e50;
  }
</style>
