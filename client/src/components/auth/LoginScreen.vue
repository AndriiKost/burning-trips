<template>
	<div class="container-tiny login-screen-wrapper">
		<div style="margin: 1rem 2rem;">
			<h2>Please Sign In</h2>
		</div>
		<el-form 
			:model="userToSign" 
			:rules="rules" 
			ref="new_stop_form" 
			label-width="120px" 
			class="new-stop-form"
		>
			<el-form-item label="Email" prop="email">
				<el-input v-model="userToSign.email"></el-input>
			</el-form-item>
            <el-form-item label="Password" prop="password">
				<el-input v-model="userToSign.password" type="password" show-password></el-input>
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="signIn">
                    Login
                </el-button>
				<!-- <el-button @click="$router.push('/register')">
                    Create Account
                </el-button> -->
			</el-form-item>
		</el-form>
		<p>
			At this time, account creation is by invites only.
			If you are interested in participating in BETA, please send an inquiry to 
			<a href="mailto:info@burningtrips.com">info@burningtrips.com</a>.
		</p>
	</div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { IUserToSign } from '@/types/User';
import { Get } from 'vuex-pathify';
import { IUser } from '@/types/User';

@Component({
   name: 'LoginScreen'
})
export default class LoginScreen extends Vue {

    /* Data */
    userToSign: IUserToSign = {
       email: '',
       password: ''
    };
	valid: boolean = false;

	rules = {
		email: [
			{ required: true, message: 'Please enter your credentials', trigger: 'blur' },
			{ min: 3, max: 60, message: 'email should 3 to 32 characters', trigger: 'blur' }
		],
		password: [
			{ required: true, message: 'Please enter your credentials', trigger: 'blur' },
			{ min: 8, max: 36, message: 'Password should be 8 to 32 characters', trigger: 'blur' }
		],
	};

	/* Methods */

	validate() {
		// @ts-ignore
        this.$refs.new_stop_form.validate(valid => this.valid = valid);
    }
    
    async signIn() {
        this.validate();
        if (!this.valid) return;
		const status = await this.$store.dispatch('auth/signIn', this.userToSign);
		if (!status) return;
		let redirectTo = localStorage.getItem('redirectTo');
		if (!redirectTo) redirectTo = '/';
		this.$router.push(redirectTo);
    }

}
</script>

<style lang='scss'>
.login-screen-wrapper {
	margin-top: 3rem;
	padding-right: 1rem;
}
</style>
