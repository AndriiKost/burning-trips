<template>
   <div class="route-list-wrapper relative">
		<div 
			v-for="route in routes"
			:key="route.id"
         class="route-summary-wrapper"
		>
         <route-summary-card
            :route="route"
            @update-votes="updateUserVote"
         />
		</div>
      <!-- <add-button @click.native="addNewRoute" /> -->
   </div>
</template>

<script lang="ts">
import Vue from 'vue';
import RouteSummaryCard from './RouteSummaryCard.vue';
import AddButton from '../global/AddButton.vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { Get, Sync } from 'vuex-pathify';
import { IStop } from '../../types/Stop';
import { IStopVote, IRouteVote } from '../../types/Vote';
import { IRoute } from '../../types/Route';

@Component({
   name: 'RouteScreen',
   components: {
      RouteSummaryCard,
      AddButton
   }
})
export default class RouteScreen extends Vue {

   /* Props */

   /* Computed */
   @Sync('route/routes')
   routes: IRoute[];

   @Get('stop/stops')
   stops: IStop[];

   /* Data */

   /* Methods */
   addNewRoute() {
      this.$router.push('/stops/create')
   }

   updateUserVote(userVote: IRouteVote) {
      this.routes.forEach(route => {
         route.votes.some(v => v.id === userVote.id)
            ? route.votes.forEach(v => v.id === userVote.id ? v.count = userVote.count : null)
            : route.votes.push(userVote);
      });
   }

   async init() {
      await this.$store.dispatch('route/getAllRoutes');
   }

   /* Lifecycle Hooks */

   async beforeMount() {
      await this.init();
   }

}
</script>

<style lang='scss' scoped>
.route-list-wrapper {
   .route-summary {
      margin: 1.2rem auto;
   }
}
</style>