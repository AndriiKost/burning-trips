<template>
   <div class="search-field-wrapper">
       <el-input 
            :placeholder="placeholder" 
            v-model="query"
            clearable
        >
        </el-input>
   </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';

@Component({
   name: 'SearchField'
})
export default class SearchField extends Vue {

   /* Props */
   @Prop({ type: String, required: true })
   readonly value: string;

   @Prop({ type: String, default: 'Search' })
   readonly placeholder: string;

   /* Computed */

   /* Data */
   query: string = null;

   /* Methods */

   /* Lifecycle Hooks */
   beforeMount() {
       this.query = this.value ? this.value : '';
   }

   @Watch('query')
   onQueryChange(nextVal: string, prevVal: string) {
       if (nextVal != prevVal) {
           this.$emit('input', nextVal);
       }
   }

}
</script>

<style lang='scss' scoped>

</style>