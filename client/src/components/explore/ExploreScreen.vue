<template>
   <section id="explore_screen">
       	<el-form 
			ref="search_form" 
			label-width="120px" 
			class="search-form"
            @submit.prevent="search"
		>
            <address-autocomplete
                label="Stop Address"
                @location-click="onLocationClick"
            />
            <el-form-item>
				<el-button type="primary" @click.prevent="search">
				    Explore
			    </el-button>
            </el-form-item>
       	</el-form>
   </section>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import AddressAutocomplete from '@/components/global/AddressAutocomplete.vue';
import { ISearchQuery } from '@/types/Explore';

@Component({
   name: 'ExploreScreen',
   components: {
       AddressAutocomplete
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
            if (!lng || !lat) {
                // handle error
                return;
            }
            this.$router.push(`/explore?lng=${lng}&lat=${lat}`);
        }
    }

   /* Lifecycle Hooks */

}
</script>

<style lang='scss'>
#explore_screen {

}
</style>