<template>
  <div>

    <v-navigation-drawer
      id="appDrawer"
      :mini-variant.sync="mini"
      fixed
      app
      v-model="drawer"
      width="260"
    >
      <v-toolbar color="primary darken-1" dark>
        <img v-bind:src="logo" height="36" alt="My Wallet">
        <v-toolbar-title class="ml0 pl-3">
          <span class="hidden-sm-and-down">My Wallet</span>
        </v-toolbar-title>
      </v-toolbar>

    <app-drawer></app-drawer>
    </v-navigation-drawer>

    <v-toolbar
      color="primary"
      fixed
      dark
      app
    >
      <v-toolbar-title class="ml-0 pl-3">
        <v-toolbar-side-icon @click.stop="handleDrawerToggle"></v-toolbar-side-icon>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn icon href="https://github.com/thanhchungbtc">
        <v-icon medium>fa fa-github</v-icon>
      </v-btn>

      <v-menu offset-y origin="center center" :nudge-bottom="10" transition="scale-transition">
        <v-btn icon large flat slot="activator">
          <v-avatar size="30px">
            <img :src="avatar" alt="BTC"/>
          </v-avatar>
        </v-btn>
        <v-list class="pa-0">
          <v-list-tile v-for="(item,index) in items" :to="!item.href ? { name: item.name } : null" :href="item.href" @click="item.click" ripple="ripple" :disabled="item.disabled" :target="item.target" rel="noopener" :key="index">
            <v-list-tile-action v-if="item.icon">
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-tile-action>
            <v-list-tile-content>
              <v-list-tile-title>{{ item.title }}</v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-menu>
    </v-toolbar>

  </div>
</template>

<script>

  import AppDrawer from './AppDrawer'

  export default {
    components: {AppDrawer,},

    data: () => ({
      mini: false,
      drawer: false,
      items: [
        {
          icon: 'account_circle',
          href: '#',
          title: 'Profile',
          click: (e) => {
            console.log(e);
          }
        },
        {
          icon: 'settings',
          href: '#',
          title: 'Settings',
          click: (e) => {
            console.log(e);
          }
        },
        {
          icon: 'fullscreen_exit',
          href: '#',
          title: 'Logout',
          click: (e) => {
          }
        }
      ],
    }),
    computed: {

      logo() {
        // return '/images/logo.png'
        return '/images/wallet.png'
      },
      avatar() {
        return '/images/avatar/btc.jpeg'
      }
    },

    methods: {
      handleDrawerToggle() {
        this.drawer = !this.drawer
      }
    }
  }

</script>