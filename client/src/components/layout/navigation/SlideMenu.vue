<template>
    <slider 
        :width="300" 
        format="overlay" 
        direction="right" 
        :opacity="0.15" 
    >
        <router-link to="/profile" v-if="isLoggedIn">
            <div class="relative link-wrapper">
                <i class="el-icon el-icon-sm el-icon-user" />
                <span class="link-text">Profile</span>
            </div>
        </router-link>
        <router-link to="/login" v-if="!isLoggedIn">
            <div class="relative link-wrapper">
                <i class="el-icon el-icon-sm el-icon-user" />
                <span class="link-text">Sign In</span>
            </div>
        </router-link>
        <router-link to="/routes">
            <div class="relative link-wrapper">
                <i class="el-icon el-icon-sm el-icon-map-location" />
                <span class="link-text">Routes</span>
            </div>
        </router-link>
        <router-link to="/stories">
            <div class="relative link-wrapper">
                <i class="el-icon el-icon-sm el-icon-notebook-2" />
                <span class="link-text">Stories</span>
            </div>
        </router-link>
        <router-link to="/">
            <div class="relative link-wrapper">
                <i class="el-icon el-icon-sm el-icon-search" />
                <span class="link-text">Search</span>
            </div>
        </router-link>
    </slider>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { Get } from 'vuex-pathify';
import Slider from './Slider.vue';

@Component({
   name: 'SlideMenu',
   components: {
       'slider': Slider
   }
})
export default class SlideMenu extends Vue {

   /* Props */

   /* Computed */
   @Get('auth/isLoggedIn')
   readonly isLoggedIn: boolean;

   /* Data */

   /* Methods */

   /* Lifecycle Hooks */

}
</script>

<style lang='scss'>
.link-wrapper, .el-icon {
    color: #000;
}
</style>