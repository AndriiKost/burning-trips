import { Route } from 'vue-router';
import store from '@/store';

export async function loggedInGuard(to: Route, from: Route, next: any) {
    const isLoggedIn = await store.dispatch('auth/initSession')
    if (!isLoggedIn && to.path !== '/login' && !to.matched.some(r => r.meta.allowGuest))  {
        next('/login');
    } else if(to.path === '/login' && isLoggedIn) {
        next('/stops');
    } else {
        next();
    }
    // TODO: redirect after logging ing
}