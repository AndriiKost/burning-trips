<template>
   <div>
       <stop-summary-card v-if="activeStop" :stop="activeStop" />
   </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { IStop } from '../../types/Stop';
import StopSummaryCard from './StopSummaryCard.vue';

@Component({
   name: 'StopDetailsPage',
   components: {
       StopSummaryCard
   }
})
export default class StopDetailsPage extends Vue {

   /* Props */

   /* Computed */

   /* Data */
   activeStop: IStop = null;

   /* Methods */
   async init() {
       const { id } = this.$route.params;
       const stop = await this.$store.dispatch('stop/getStop', id);
       if (!stop) {
        // add error to the notif stack
        this.$router.push('/stops');
       }
       this.activeStop = stop;
   }

   /* Lifecycle Hooks */

   created() {
       this.init();
   }

}
</script>

<style lang='scss' scoped>

</style>