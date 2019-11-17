import { Route } from 'vue-router';
import store from '@/store';

export function loggedInGuard(to: Route, from: Route, next: any) {
    store.dispatch('auth/initSession').then(success => {
        if (!success && !to.matched.some(r => r.meta.allowGuest)) {
            next({
                path: '/login',
                params: { redirectTo: to.fullPath }
            });
        } else if (success && to.name === 'Login') {
            const path = to.params.redirectTo || '/';
            next({ path });
        } else {
            next();
        }
    })
}