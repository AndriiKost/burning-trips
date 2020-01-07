import LoginScreen from '@/components/auth/LoginScreen.vue';
import RouteScreen from '@/components/routes/RouteScreen.vue';
import CreateStopSection from '@/components/stops/create/CreateStopSection.vue';
import StopDetailsPage from '@/components/stops/StopDetailsPage.vue';
import StopScreen from '@/components/stops/StopScreen.vue';
import Vue from "vue";
import VueRouter from "vue-router";
import { loggedInGuard } from './middleware';

Vue.use(VueRouter);

const routes = [
  {
    path: "/stops",
    component: StopScreen,
    name: 'Stop Screen',
    meta: {
      allowGuest: true
    }
  },
  {
    path: "/stops/create",
    component: CreateStopSection,
    name: 'Create Stop'
  },
  {
    path: "/stops/:id",
    component: StopDetailsPage,
    name: 'Stop Details',
    meta: {
      allowGuest: true
    }
  },
  {
    path: "/routes",
    component: RouteScreen,
    name: 'Routes Screen',
    meta: {
      allowGuest: true
    }
  },
  {
    path: "/login",
    component: LoginScreen,
    name: 'Login'
  },
  {
    path: "*",
    component: RouteScreen,
    name: 'Routes Screen'
  }
  // {
    // path: "/about",
    // name: "about",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    // component: () => import(/* webpackChunkName: "about" */ "../views/About.vue")
  // }
];

const router = new VueRouter({
  mode: "history",
  // base: process.env.VUE_APP_BASE_URL,
  routes
});

router.beforeEach(loggedInGuard);

export default router;
