<template>
   <footer id="appFooter" class="flex flex-column flex-center">
       <p class="main grey">© 2020 Delian Inc All rights reserved.</p>
       <p class="secondary grey">Burning Trips &reg; is Delian Inc product.</p>
   </footer>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';

@Component({
   name: 'AppFooter'
})
export default class  extends Vue {

   /* Props */

   /* Computed */

   /* Data */

   /* Methods */

   /* Lifecycle Hooks */

}
</script>

<style lang='scss'>
#appFooter {
    position: absolute;
    left: 0;
    bottom: 0;
    height: 3rem;
    width: 100%;
    background: rgba(238, 238, 238, 0.795);
    box-shadow: 0px 0px 9px 3px rgba(104, 103, 103, 0.25);
    padding: .5rem 1rem;

    p {
        &.main {
            font-size: .8rem;
        }
        &.grey {
            color: lighten($dark-grey, 30%);
        }
        &.secondary {
            font-size: .75rem;
            margin-top: .2rem;
        }
    }
}
</style>