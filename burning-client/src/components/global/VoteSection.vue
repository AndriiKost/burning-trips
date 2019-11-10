<template>
    <div class="vote-section flex-row space-between">
        <div class="icon-wrapper relative">
            <div 
                v-if="trending" 
                class="icon-trending icon-btn inline-block" 
                @click="handleVote"
            ></div>
            <div 
                v-else 
                class="icon-camera icon-btn inline-block" 
                @click="handleVote"
            ></div>
            <transition name="el-fade-in">
                <span v-show="showUserVotes" class="user-votes bg">
                    +{{ curVotes }}
                </span>
            </transition>
        </div>
        <span class="vote-amount bg">
            {{ totalVotes }}
        </span>
    </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';

@Component({
   name: 'VoteSection'
})
export default class VoteSection extends Vue {

    /* Props */
    @Prop({ type: Number, default: 25 })
    readonly voteLimit: Number;

    @Prop({ type: Number, required: true })
    readonly totalVotes: Number;

    @Prop({ type: Boolean, default: false })
    readonly trending: Boolean;

    /* Computed */

    /* Data */
    timer: any = null;
    curVotes: number = 0;
    showUserVotes: boolean = false;

    /* Methods */
    handleVote(): void {
		clearTimeout(this.timer);
		this.showUserVotes = true;
		if (this.curVotes < this.voteLimit) this.curVotes++;
		this.timer = setTimeout(() => {
			this.showUserVotes = false;
			this.$emit('update-votes', this.curVotes);
		}, 2000);
	}

    /* Lifecycle Hooks */

}
</script>

<style lang='scss' scoped>
@import 'src/scss/variables.scss';
.vote-section {
    .icon-btn {
        color: darken($red, 10%);
        font-size: 1.75rem;
        margin-right: .25rem;
        cursor: pointer;
    }
    .vote-amount {
        color: $blue;
    }
    .user-votes {
        position: absolute;
        top: -2rem;
        left: 0;
        color: $red;
        padding: .25rem;
        border-radius: 50%;
        font-size: .85rem;
    }
}
</style>