<template>
    <div>
        <div v-if="activeRoute">
            <route-details :route="activeRoute" />
            <stop-map :stops="activeRoute.stops" />
        </div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import StopMap from '@/components/maps/StopMap.vue';
import RouteDetails from './RouteDetails.vue';
import { IRoute } from '../../types/Route';

@Component({
   name: 'RouteDetailsPage',
   components: {
       StopMap,
       RouteDetails
   }
})
export default class RouteDetailsPage extends Vue {

   /* Props */

   /* Computed */

   /* Data */
   activeRoute: IRoute = null;

   /* Methods */
   async init() {
       const { id } = this.$route.params;
	   const route = await this.$store.dispatch('route/getRoute', id);
	   if (!route) {
		// add error to the notif stack
		this.$router.push('/routes');
	   }
	   this.activeRoute = route;
   }

   /* Lifecycle Hooks */
   beforeMount() {
       this.init();
   }

}
</script>

<style lang='scss'>

</style>