<template>
    <div class="circle-menu-bottom">
        <div class="wrapper">
            <circle-menu
                type="top"
                :number="4"
                animate="animated jello"
                mask='white'
                circle
                @toggle-menu="toggleMenu"
            >
                <button type="button" slot="item_btn" class="transparent">
                    <i class="el-icon el-icon-sm el-icon-menu" v-show="!isOpen" />
                    <i class="el-icon el-icon-sm el-icon-close" v-show="isOpen" />
                </button>
                <router-link to="/" slot="item_1">
                    <i class="el-icon el-icon-sm el-icon-search" />
                </router-link>
                <router-link to="/" slot="item_2">
                    <i class="el-icon el-icon-sm el-icon-place" />
                </router-link>
                <router-link to="/" slot="item_3">
                    <i class="el-icon el-icon-sm el-icon-place" />
                </router-link>
                <router-link to="/" slot="item_4">
                    <i class="el-icon el-icon-sm el-icon-place" />
                </router-link>
            </circle-menu>
        </div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import CircleMenu from './CircleMenu.vue';

@Component({
    name: 'BottomCircleMenu',
    components: {
        CircleMenu
    }
})
export default class BottomCircleMenu extends Vue {

   /* Props */

   /* Computed */

   /* Data */
   isOpen: boolean = false;

   /* Methods */
   toggleMenu(open: boolean) {
       this.isOpen = open;
   }

   /* Lifecycle Hooks */

}
</script>

<style lang='scss'>
.circle-menu-bottom {
    position: fixed;
    align-items: flex-end;
    justify-content: center;
    display: flex;
    top: 0;
    bottom: 4.5rem;
    right: 1rem;
}
</style>