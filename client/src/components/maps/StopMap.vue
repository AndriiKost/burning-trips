<template>
    <div id="stopMap"></div>
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
            style: 'mapbox://styles/mapbox/light-v10'
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
#stopMap {
    position: absolute;
    top:0;
    left:0;
    right:0;
    bottom:0;
}
</style>