<template>
    <el-card :body-style="{ padding: '0px' }" shadow="always">
        <div id="stopMap"></div>
   </el-card>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { IStop } from '@/types/Stop';

import mapboxgl from 'mapbox-gl/dist/mapbox-gl.js';
import config from '../../config';

@Component({
   name: 'StopMap'
})
export default class StopMap extends Vue {

   /* Props */
   @Prop({ type: Array as () => IStop[], required: true })
   readonly stops: IStop[];

   /* Computed */

   /* Data */

   /* Methods */
   init() {
       this.initMap();
   }

   initMap() {
        mapboxgl.accessToken = config.MAPBOX_API;
        var map = new mapboxgl.Map({
            container: 'stopMap',
            style: 'mapbox://styles/mapbox/streets-v11'
        });

        this.stops.forEach(stop => {
           new mapboxgl.Marker()
            .setLngLat([stop.latitude, stop.longitude])
            .setPopup(
                new mapboxgl.Popup()
                    .setHTML(`<div class="popup-wrapper"><h2>${stop.name}</h2><p>${stop.address}</p></div>`)
            )
            .addTo(map)
       })
   }

   /* Lifecycle Hooks */
   mounted() {
       this.init();
   }

}
</script>

<style lang='scss'>

</style>