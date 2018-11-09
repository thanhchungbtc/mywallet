<template>
  <!--Category list-->
  <v-card>

    <!--Dialog-->
    <v-dialog v-model="dialog" persistent max-width="500px">
      <v-card>
        <v-card-title>
          <span class="headline">Add category</span>
        </v-card-title>
        <v-divider></v-divider>
        <v-card-text>

          <v-container grid-list-md>
            <v-layout wrap>
              <v-flex xs12 sm12 md12>
                <v-text-field label="*Name" required></v-text-field>
              </v-flex>
              <v-flex xs12 sm12 md12>
                <v-text-field label="Balance" type="number"></v-text-field>
              </v-flex>
              <v-flex xs12 sm12 md12>
                <v-text-field label="Memo"></v-text-field>
              </v-flex>

            </v-layout>
          </v-container>

          <small>*indicates required field</small>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" flat @click.native="dialog = false">Close</v-btn>
          <v-btn color="blue darken-1" flat @click.native="dialog = false">Save</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!--Category list-->
    <v-toolbar card dense color="transparent">
      <v-toolbar-title><h4>Category list</h4></v-toolbar-title>
      <v-spacer></v-spacer>

      <v-btn flat icon color="grey" @click.native="dialog=true">
        <v-icon>add</v-icon>
      </v-btn>

    </v-toolbar>

    <v-divider></v-divider>

    <v-card-text class="pa-0">
      <template>
        <v-data-table
          :headers="headers"
          :items="projects"
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
              <v-btn flat icon color="grey">
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
            value: 'name'
          },
          {
            text: 'Memo',
            align: 'left',
            value: 'memo'
          },
          {text: 'Progress', value: 'progress'},
          {text: 'Action', value: 'action', align: 'right'},

        ],
      };
    },
    computed: {
      projects() {
        const Projects = [
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
    }
  };
</script>
