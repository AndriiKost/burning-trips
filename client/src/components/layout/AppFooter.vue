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
@import 'src/scss/mixins.scss';
@import 'src/scss/variables.scss';
#appFooter {
    position: absolute;
    left: 0;
    bottom: 0;
    height: 3rem;
    width: 100%;
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