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
                <router-link to="/profile" slot="item_1">
                    <div class="relative link-wrapper">
                        <i class="el-icon el-icon-sm el-icon-user" />
                        <span class="link-text">profile</span>
                    </div>
                </router-link>
                <router-link to="/routes" slot="item_2">
                    <div class="relative link-wrapper">
                        <i class="el-icon el-icon-sm el-icon-map-location" />
                        <span class="link-text">routes</span>
                    </div>
                </router-link>
                <router-link to="/stories" slot="item_3">
                    <div class="relative link-wrapper">
                        <i class="el-icon el-icon-sm el-icon-notebook-2" />
                        <span class="link-text">stories</span>
                    </div>
                </router-link>
                <router-link to="/" slot="item_4">
                    <div class="relative link-wrapper">
                        <i class="el-icon el-icon-sm el-icon-search" />
                        <span class="link-text">search</span>
                    </div>
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

    a {
        text-decoration: none;
    }
    
    .link-wrapper {
        // display: flex;
        // flex-direction: column;
        // text-decoration: none;

        .link-text, .el-icon {
            color: #000;
            text-decoration: none;
        }
        
        .link-text {
            position: absolute;
            bottom: -1.25rem;
            left: 0;
            right: 0;
        }
    }
}
</style>