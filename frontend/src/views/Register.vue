<template>
  <v-app id="login" class="primary">
    <v-content>
      <v-container fluid fill-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md4 lg4>
            <v-card class="elevation-1 pa-3">
              <v-card-text>
                <div class="layout column align-center">
                  <img :src="logo" alt="Vue Material Admin" width="120" height="120">
                  <h1 class="flex my-4 primary--text">Welcome to mywallet</h1>
                </div>
                <v-form>
                  <v-text-field append-icon="person" name="username" label="Username" type="text"
                                v-model="model.username"></v-text-field>
                  <v-text-field append-icon="email" name="email" label="Email" type="email"
                                v-model="model.email"></v-text-field>
                  <v-text-field append-icon="lock" name="password" label="Password" id="password" type="password"
                                v-model="model.password"></v-text-field>
                  <v-text-field append-icon="lock" name="password_confirmation" label="Password confirmation"
                                id="password_confirmation" type="password"
                                v-model="model.password_confirmation"></v-text-field>
                </v-form>
              </v-card-text>
              <v-card-actions>
                <v-btn block color="primary" @click="register" :loading="loading">Register</v-btn>
              </v-card-actions>
            </v-card>
          </v-flex>
        </v-layout>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
  export default {
    data: () => ({
      loading: false,
      model: {
        username: '',
        email: '',
        password: '',
        password_confirmation: ''
      }
    }),

    computed: {
      logo() {
        // return '/images/logo.png'
        return '/images/wallet.png'
      },
    },

    methods: {
      login() {
        this.loading = true
        this.$store.dispatch('register', this.model)
          .then(() => {
            this.loading = false
            this.$router.push({name: 'dashboard'})
          })
          .catch(err => {
            this.loading = false
            this.$store.dispatch('app/error', err.error)
          })
      }
    }
  }
</script>
<style scoped lang="css"> #login {
  height: 50%;
  width: 100%;
  position: absolute;
  top: 0;
  left: 0;
  content: "";
  z-index: 0;
}
</style>
