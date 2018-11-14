<template>

  <v-card>

    <v-card-title>
      <span class="headline">{{ formTitle }}</span>
    </v-card-title>

    <v-divider></v-divider>

    <v-card-text>
      <v-container grid-list-md>
        <v-layout wrap>
          <v-flex xs12 sm12 md12>
            <!--<date-picker></date-picker>-->
          </v-flex>
          <v-flex xs12 sm12 md12>
            <v-text-field label="Budget" v-model="item.budget" type="number"></v-text-field>
          </v-flex>
          <v-flex xs12 sm12 md12>
            <v-text-field label="Memo" v-model="item.memo"></v-text-field>
          </v-flex>

        </v-layout>
      </v-container>

      <small>*indicates required field</small>
    </v-card-text>

    <v-card-actions>
      <v-spacer></v-spacer>
      <!--<v-btn color="blue darken-1" flat @click.native="close">Close</v-btn>-->
      <!--<v-btn color="blue darken-1" flat @click.native="dialog = false">Save</v-btn>-->
    </v-card-actions>

  </v-card>

</template>

<script>
  export default {

    data: () => ({
      item: {

      }
    }),
    computed: {
      formTitle() {
        return this.isUpdate === true ? 'Update category' : 'Add category'
      },
    }
  }

</script>