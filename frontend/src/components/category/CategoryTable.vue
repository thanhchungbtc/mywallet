<template>
  <!--Category list-->
  <v-card>

    <!--Category list-->
    <v-toolbar card dense color="transparent">
      <v-toolbar-title><h4>Category list</h4></v-toolbar-title>
      <v-spacer></v-spacer>

      <!--Dialog-->
      <v-dialog v-model="dialog" persistent max-width="500px">

        <v-btn slot="activator" color="primary" flat icon>
          <v-icon>add</v-icon>
        </v-btn>

        <v-card>

          <v-card-title>
            <span class="headline">{{ formTitle }}</span>
          </v-card-title>

          <v-divider></v-divider>

          <v-card-text>
            <v-container grid-list-md>
              <v-layout wrap>
                <v-flex xs12 sm12 md12>
                  <v-text-field label="*Name" v-model="category.name" required></v-text-field>
                </v-flex>
                <v-flex xs12 sm12 md12>
                  <v-text-field label="Balance" v-model="category.balance" type="number"></v-text-field>
                </v-flex>
                <v-flex xs12 sm12 md12>
                  <v-text-field label="Memo" v-model="category.memo"></v-text-field>
                </v-flex>

              </v-layout>
            </v-container>

            <small>*indicates required field</small>
          </v-card-text>

          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="blue darken-1" flat @click.native="close">Close</v-btn>
            <v-btn color="blue darken-1" flat @click.native="dialog = false">Save</v-btn>
          </v-card-actions>

        </v-card>
      </v-dialog>

    </v-toolbar>

    <v-divider></v-divider>

    <v-card-text class="pa-0">
      <template>
        <v-data-table
          :headers="headers"
          :items="categories"
          hide-actions
          class="elevation-0"
        >
          <template slot="items" slot-scope="props">
            <td>
              <v-avatar size="36px">
                <img :src="props.item.avatar" :alt="props.item.username"/>
              </v-avatar>
            </td>
            <td>{{ props.item.name }}</td>
            <td>{{ props.item.memo }}</td>
            <td class="text-xs-left">
              <v-progress-linear :value="props.item.progress" height="5" :color="props.item.color"></v-progress-linear>
            </td>
            <td class="text-xs-right">
              <v-btn flat icon color="grey" @click.stop="edit(props.item)">
                <v-icon>edit</v-icon>
              </v-btn>
              <v-btn flat icon color="grey">
                <v-icon>delete</v-icon>
              </v-btn>
            </td>
          </template>

        </v-data-table>

      </template>

      <v-divider></v-divider>

    </v-card-text>


  </v-card>
</template>

<script>
  export default {
    data() {
      return {
        isUpdate: false,
        defaultCategory: {
          name: '',
          memo: '',
          balance: '',
        },
        category: {
          name: '',
          memo: '',
          balance: '',
        },

        dialog: false,

        headers: [
          {
            text: '',
            align: 'center',
            sortable: false,
            value: 'avatar'
          },
          {
            text: 'Name',
            align: 'left',
            sortable: true,
            value: 'name'
          },
          {
            text: 'Memo',
            align: 'left',
            sortable: true,
            value: 'memo'
          },
          {text: 'Progress', value: 'progress', sortable: true,},
          {text: 'Actions', value: 'action', align: 'right', sortable: false,},

        ],
      };
    },
    computed: {
      formTitle() {
        return this.isUpdate === true ? 'Update category' : 'Add category'
      },

      categories() {
        return [
          {
            avatar: '/images/category/food.png',
            name: 'Food',
            memo: 'Food',
            progress: 90,
            color: 'pink',
          },
          {
            avatar: '/images/category/family.png',
            name: 'Family',
            memo: 'Family',
            progress: 70,
            color: 'success'
          },
          {
            avatar: '/images/category/baby.JPG',
            name: 'Baby',
            memo: 'Baby',
            progress: 50,
            color: 'info'
          },
          {
            avatar: '/images/category/house.png',
            name: 'House',
            memo: 'House',
            progress: 30,
            color: 'teal'
          },
          {
            avatar: '/images/category/travel.png',
            name: 'Travel',
            memo: 'Travel',
            progress: 15,
            color: 'grey'
          },

        ];
        return Projects;
      }
    },

    methods: {
      edit(item) {
        this.category = item
        this.dialog = true
        this.isUpdate = true
      },
      close() {
        this.dialog = false
        this.category = {...this.defaultCategory}
      }
    }
  };
</script>
