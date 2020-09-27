<template>
    <el-carousel v-if="isThumbCarousel" trigger="click" height="150px">
        <el-carousel-item v-for="stop in stops" :key="stop.id">
            <h3 class="small">
                {{ stop.name }}
            </h3>
            <img
                :src="stop.imageUrl"
                :alt="`Image of ${stop.name}`"
            />
        </el-carousel-item>
    </el-carousel>
      <el-carousel v-else trigger="click" :interval="4000" type="card" height="200px">
        <el-carousel-item v-for="stop in stops" :key="stop.id">
            <h3 class="small">
                {{ stop.name }}
            </h3>
            <img
                :src="stop.imageUrl"
                :alt="`Image of ${stop.name}`"
            />
        </el-carousel-item>
    </el-carousel>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { IStop } from '../../../types/Stop';

@Component({
   name: 'LandmarkCarousel'
})
export default class LandmarkCarousel extends Vue {

   /* Props */
   @Prop({ type: Array as () => IStop[] })
   readonly stops: IStop[];

   @Prop({ type: String, default: 'thumb-carousel' })
   readonly type: string;

   /* Computed */
   get isThumbCarousel(): boolean {
       return this.type === 'thumb-carousel';
   }

   /* Data */

   /* Methods */

   /* Lifecycle Hooks */

}
</script>

<style lang='scss'>

</style>