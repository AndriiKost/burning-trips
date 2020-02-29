<template>
   <section id="explore_screen">
        <location-search 
            @search="search" 
            @click:location="onLocationClick" 
        />
   </section>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator'
import LocationSearch from '@/components/search/LocationSearch.vue';
import { ISearchQuery } from '@/types/Explore';

@Component({
   name: 'ExploreScreen',
   components: {
       LocationSearch
   }
})
export default class ExploreScreen extends Vue {

   /* Props */

   /* Computed */

   /* Data */
   searchPlace = null;

   /* Methods */
   onLocationClick(place) {
       this.searchPlace = place;
   }

    search() {
        if (this.searchPlace != null) {
            let lat = this.searchPlace.geometry.location.lat();
            let lng = this.searchPlace.geometry.location.lng();
            let loc = this.searchPlace.formatted_address;
            if (!lng || !lat) {
                // handle error
                return;
            }
            this.$router.push(`/explore?lng=${lng}&lat=${lat}&loc=${loc}`);
        }
    }

   /* Lifecycle Hooks */

}
</script>

<style lang='scss'>
#explore_screen {

}
</style>