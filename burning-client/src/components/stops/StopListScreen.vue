<template>
   <div class="stop-list-wrapper">
		<div 
			v-for="stop in stops"
			:key="stop.id"
         class="stop-summary-wrapper"
		>
         <stop-summary :stop="stop" />
		</div>
   </div>
</template>

<script lang="ts">
import Vue from 'vue';
import StopSummary from './StopSummary.vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { Get } from 'vuex-pathify';
import { Stop } from '@/types/stops.type';

@Component({
   name: 'StopListScreen',
   components: {
      StopSummary
   }
})
export default class  extends Vue {

   /* Props */

   /* Computed */
   @Get('stop/stops')
   stops: Array<Stop>

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