<template>
  <div id="app" v-show="!booting">
    <div id="nav">
      <v-app id="inspire">


        <app-toolbar v-if="!$route.meta.public" class="app--drawer"></app-toolbar>

        <v-content>
          <router-view></router-view>
        </v-content>

        <!-- Go to top -->
        <app-fab></app-fab>

        <app-snackbar></app-snackbar>

      </v-app>
    </div>
  </div>
</template>

<script>

  import AppToolbar from './components/AppToolbar'
  import AppFab from './components/AppFab'
  import AppSnackbar from './components/AppSnackbar'

  export default {
    name: 'app',
    components: {
      AppToolbar,
      AppFab,
      AppSnackbar,
    },

    data: () => ({
      booting: true,
    }),

    created() {
      console.log('hit')
      this.$store.dispatch('auth/verify')
        .then(() => {
          this.$router.push({name: 'dashboard'})
          this.booting = false
        })
        .catch(err => {
          console.log(err)
          this.$store.dispatch('auth/logout')
          this.booting = false
        })

    },
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
