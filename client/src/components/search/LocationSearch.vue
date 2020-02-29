<template>
    <el-form 
        ref="search_form"
        :label-width="labelWidth"
        class="search-form"
        @submit.prevent="search"
    >
        <address-autocomplete
            :label="label"
            @location-click="onLocationClick"
        />
        <el-form-item>
            <el-button 
                type="primary" 
                @click.prevent="search"
            >
                {{ buttonText }}
            </el-button>
        </el-form-item>
    </el-form>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import AddressAutocomplete from '@/components/global/AddressAutocomplete.vue';

@Component({
   name: 'LocationSearch',
   components: {
       AddressAutocomplete,
   }
})
export default class LocationSearch extends Vue {

   /* Props */
   @Prop({ type: String, default: 'Search By location' })
   readonly label: string;

   @Prop({ type: String, default: '120px' })
   readonly labelWidth: string;

   @Prop({ type: String, default: 'Explore' })
   readonly buttonText: string;

   /* Computed */

   /* Data */

   /* Methods */
   search() {
       this.$emit('search');
   }

   onLocationClick(place) {
       this.$emit('click:location', place);
   }

   /* Lifecycle Hooks */

}
</script>

<style lang='scss' scoped>

</style>