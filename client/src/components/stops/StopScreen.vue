<template>
   <div class="stop-list-wrapper relative">
		<div 
			v-for="stop in stops"
			:key="stop.id"
         class="stop-summary-wrapper"
		>
         <stop-summary-card :stop="stop" @update-votes="updateUserVote" />
		</div>
      <!-- <add-button @click.native="addNewStop" /> -->
   </div>
</template>

<script lang="ts">
import Vue from 'vue';
import StopSummaryCard from './StopSummaryCard.vue';
import AddButton from '../global/AddButton.vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { Get, Sync } from 'vuex-pathify';
import { IStop } from '../../types/Stop';
import { IStopVote } from '../../types/Vote';

@Component({
   name: 'StopScreen',
   components: {
      StopSummaryCard,
      AddButton
   }
})
export default class StopScreen extends Vue {

   /* Props */

   /* Computed */
   @Sync('stop/stops')
   readonly stops: IStop[];

   /* Data */

   /* Methods */
   addNewStop() {
      this.$router.push('/stops/create')
   }

   updateUserVote(userVote: IStopVote) {
      this.stops.forEach(stop => {
         stop.votes.some(v => v.id === userVote.id)
            ? stop.votes.forEach(v => v.id === userVote.id ? v.count = userVote.count : null)
            : stop.votes.push(userVote);
      });
   }

   async init() {
      await this.$store.dispatch('stop/getAllStops');
      // console.log(this.stops[0].votes)
   }

   /* Lifecycle Hooks */

   created() {
      this.init();
   }

}
</script>

<style lang='scss'>
.stop-summary-wrapper {
   .stop-summary {
      margin: 1.2rem auto;
   }
}
</style>