<template>
   <div class="stop-list-wrapper relative">
      <div class="stop-search">
         <search-field
            v-model="query"
         />
      </div>
		<div 
			v-for="stop in filteredStops"
			:key="stop.id"
         class="stop-summary-wrapper"
		>
         <stop-summary-card :stop="stop" @update-votes="updateUserVote" />
		</div>

      <add-button @click="showCreateOptions" />
      
      <modal-wrapper v-if="activeModal" />
   </div>
</template>

<script lang="ts">
import Vue from 'vue';
import StopSummaryCard from './StopSummaryCard.vue';
import AddButton from '../global/AddButton.vue';
import ModalWrapper from './modals/ModalWrapper.vue';
import SearchField from '../global/SearchField.vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { Get, Sync } from 'vuex-pathify';
import { IStop } from '../../types/Stop';
import { IStopVote } from '../../types/Vote';
import { ModalType } from '@/types/components/Modal';

@Component({
   name: 'StopScreen',
   components: {
      StopSummaryCard,
      ModalWrapper,
      SearchField,
      AddButton
   }
})
export default class StopScreen extends Vue {

   /* Props */

   /* Computed */
   @Sync('stop/stops')
   readonly stops: IStop[];

   @Sync('stop/activeModal')
   activeModal: ModalType;

   get filteredStops(): IStop[] {
      return this.stops.filter(stop => stop.name.indexOf(this.query) > 0);
   }

   /* Data */
   query: string = null;

   /* Methods */

   updateUserVote(userVote: IStopVote) {
      this.stops.forEach(stop => {
         stop.votes.some(v => v.id === userVote.id)
            ? stop.votes.forEach(v => v.id === userVote.id ? v.count = userVote.count : null)
            : stop.votes.push(userVote);
      });
   }

   showCreateOptions() {
      this.activeModal = ModalType.CreateContentFromStops;
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