<template>
   <div class="stop-list-wrapper">
		<div 
			v-for="stop in stops"
			:key="stop.id"
         class="stop-summary-wrapper"
		>
         <stop-summary-card :stop="stop" />
		</div>
   </div>
</template>

<script lang="ts">
import Vue from 'vue';
import StopSummaryCard from './StopSummaryCard.vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { Get } from 'vuex-pathify';
import { IStop } from '../../types/Stop';

@Component({
   name: 'StopListScreen',
   components: {
      StopSummaryCard
   }
})
export default class StopListScreen extends Vue {

   /* Props */

   /* Computed */
   @Get('stop/stops')
   stops: Array<IStop>

   /* Data */

   /* Methods */

   /* Lifecycle Hooks */

}
</script>

<style lang='scss' scoped>
.stop-summary-wrapper {
   .stop-summary {
      margin: 1.2rem auto;
   }
}
</style>