<template>
  <v-snackbar
    :timeout="3000"
    bottom
    right
    :auto-height="true"
    :color="snackbar.color"
    v-model="snackbarShow"
  >
    {{ snackbar.text }}
    <v-btn dark flat @click.native="snackbarShow = false" icon>
      <v-icon>close</v-icon>
    </v-btn>
  </v-snackbar>
</template>
<script>

  export default {

    computed: {
      snackbar: function() {
        return this.$store.state.app.snackbar
      },

      snackbarShow: {
        get: function()  {
          return this.$store.state.app.snackbar.show
        },
        set:function (newVal)  {
          this.$store.dispatch('app/showSnackbar', newVal)
        }
      },

    },
  }
</script>