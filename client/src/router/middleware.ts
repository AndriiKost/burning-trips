import store from '@/store';
import { Route } from 'vue-router';

export async function loggedInGuard(to: Route, from: Route, next: any) {
    const isLoggedIn = await store.dispatch('auth/initSession')
    if (!isLoggedIn && to.path !== '/login' && !to.matched.some(r => r.meta.allowGuest))  {
        next('/login');
    } else if(to.path === '/login' && isLoggedIn) {
        next('/stops');
    } else {
        next();
    }
    if (localStorage.redirectAfterLogin) {
        next(localStorage.redirectAfterLogin);
    }
    next(to);
}

export function preventDuplicate(to: Route, from: Route, next: any) {
    if (to.fullPath === from.fullPath) return next();
    else next();
}