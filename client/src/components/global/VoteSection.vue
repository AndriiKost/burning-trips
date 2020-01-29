<template>
    <div class="vote-section flex flex-row space-between">
        <div class="icon-wrapper relative">
            <el-button 
                circle 
                @click="handleVote" 
                icon="el-icon-camera-solid" 
                style="font-size: 22px;padding: 8px;"
            ></el-button>
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
    readonly voteLimit: number;

    @Prop({ type: Number, required: true })
    readonly totalVotes: number;

    @Prop({ type: Boolean, default: false })
    readonly trending: boolean;

    @Prop({ type: Number, default: 0 })
    readonly curUserVotes: number;

    /* Computed */

    /* Data */
    timer: any = null;
    curVotes: number = this.curUserVotes;
    showUserVotes: boolean = false;

    /* Methods */
    handleVote(): void {
		clearTimeout(this.timer);
		this.showUserVotes = true;
		if (this.curVotes < this.voteLimit) this.curVotes++;
		this.timer = setTimeout(() => {
			this.showUserVotes = false;
			this.$emit('save-votes', this.curVotes);
		}, 2000);
	}

    /* Lifecycle Hooks */

}
</script>

<style lang='scss' scoped>
@import 'src/scss/variables.scss';
.vote-section {
    width: 65px;

    .icon-btn {
        // color: darken($red, 10%);
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