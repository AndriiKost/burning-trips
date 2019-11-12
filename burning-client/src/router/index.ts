import Vue from "vue";
import VueRouter from "vue-router";
import StopListScreen from '@/components/stops/StopListScreen.vue';
import CreateStopSection from '@/components/stops/create/CreateStopSection.vue';
import LoginScreen from '@/components/auth/LoginScreen.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: "/stops",
    component: StopListScreen
  },
  {
    path: "/stops/create",
    component: CreateStopSection
  },
  {
    path: "/login",
    component: LoginScreen
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
  base: process.env.BASE_URL,
  routes
});

export default router;
