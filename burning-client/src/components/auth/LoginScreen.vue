<template>
	<div>
		<el-form 
			:model="userToSign" 
			:rules="rules" 
			ref="new_stop_form" 
			label-width="120px" 
			class="new-stop-form"
		>
			<el-form-item label="email" prop="email">
				<el-input v-model="userToSign.email"></el-input>
			</el-form-item>
            <el-form-item label="Password" prop="password">
				<el-input v-model="userToSign.password"></el-input>
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="signIn">
                    Login
                </el-button>
				<el-button @click="$router.push('/register')">
                    Create Account
                </el-button>
			</el-form-item>
		</el-form>
	</div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { IUserToSign } from '@/types/User';
import { Get } from 'vuex-pathify';

@Component({
   name: 'LoginScreen'
})
export default class LoginScreen extends Vue {

    /* Computed */

    @Get('auth/dummy')
    dummy: string;

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
			{ min: 8, max: 36, message: 'Password should 3 to 32 characters', trigger: 'blur' }
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
        const result = await this.$store.dispatch('auth/signIn', this.userToSign);
        // send request to authenticate
    }

}
</script>

<style lang='scss' scoped>

</style>
