<template>
  <div id="pageDashboard">
    <v-container grid-list-xl fluid>
      <v-layout row wrap>

        <!-- mini statistic  end -->
        <v-flex lg7 sm12 xs12>
          <v-widget title="Annual expenses report" content-bg="white">
            <v-btn icon slot="widget-header-action">
              <v-icon class="text--secondary">refresh</v-icon>
            </v-btn>
            <div slot="widget-content">
              <e-chart
                :path-option="[
                  ['dataset.source', expensesData],
                  ['color', [color.lightBlue.base,]],
                  ['legend.show', true],
                  ['xAxis.axisLabel.show', true],
                  ['yAxis.axisLabel.show', true],
                  ['grid.left', '2%'],
                  ['grid.bottom', '5%'],
                  ['grid.right', '3%'],
                  ['series[0].type', 'bar'],
                  ['series[0].areaStyle', {}],
                  ['series[0].smooth', true],
                ]"
                height="400px"
                width="100%"
              >
              </e-chart>
            </div>
          </v-widget>
        </v-flex>

        <v-flex lg5 sm12 xs12>
          <v-widget title="Expenses by category all the time" content-bg="white">
            <div slot="widget-content">
              <e-chart
                :path-option="[
                  ['dataset.source', categoriesData],
                  ['legend.bottom', '0'],
                  ['color', [color.lightBlue.base, color.indigo.base, color.pink.base, color.green.base, color.cyan.base, color.teal.base]],
                  ['xAxis.show', false],
                  ['yAxis.show', false],
                  ['series[0].type', 'pie'],
                  ['series[0].avoidLabelOverlap', true],
                  ['series[0].radius', ['50%', '70%']],
                ]"
                height="400px"
                width="100%"
              >
              </e-chart>
            </div>
          </v-widget>
        </v-flex>

      </v-layout>
      <v-layout row wrap>
        <v-flex lg4 sm6 xs12 v-for="(item,index) in categories" :key="'c-trending'+index">
          <circle-statistic
            :title="item.subheading"
            :sub-title="item.headline"
            :caption="item.caption"
            :icon="item.icon.label"
            :color="item.linear.color"
            :value="item.linear.value"
          >
          </circle-statistic>
        </v-flex>
      </v-layout>
    </v-container>
  </div>
</template>

<script>
  import EChart from '../components/chart/echart';
  import VWidget from '../components/VWidget';
  import Material from 'vuetify/es5/util/colors';
  import CircleStatistic from '../components/widgets/statistic/CircleStatistic'

  export default {
    components: {
      VWidget,
      EChart,
      CircleStatistic
    },
    data: () => ({
      color: Material,
      categories: [
        {
          subheading: 'Daily',
          headline: '20,000 / 40,000',
          caption: '(Remain 20,000)',
          percent: 50,
          icon: {
            color: 'info'
          },
          linear: {
            value: 50,
            color: 'info'
          }
        },
        {
          subheading: 'Hang out',
          headline: '5,000/10,000',
          caption: '(Remain 5,000)',
          percent: 50,
          icon: {
            label: 'list',
            color: 'primary'
          },
          linear: {
            value: 50,
            color: 'success'
          }
        },
        {
          subheading: 'Study',
          headline: '10,000/10,000',
          caption: '(Remain 0)',
          percent: 100,
          icon: {
            label: 'bug_report',
            color: 'primary'
          },
          linear: {
            value: 100,
            color: 'error'
          }
        },
      ],


    }),
    computed: {

      expensesData() {

        const shortMonth = [
          'Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'
        ];
        const monthVisitData = shortMonth.map(m => {
          return {
            'month': m,
            'Total': Math.floor(Math.random() * 100000) + 2000,
          };
        });

        return monthVisitData
      },

      categoriesData() {
        const data = [
          {
            value: 50,
            name: 'Ăn uông'
          },
          {
            value: 35,
            name: 'Nhà cửa'
          },
          {
            value: 25,
            name: 'Du lịch'
          },
          {
            value: 10,
            name: 'Sinh hoạt phí'
          },
          {
            value: 10,
            name: 'Other'
          }
        ];

        return data
      }
    },

  };
</script>
