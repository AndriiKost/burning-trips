import LoginScreen from '@/components/auth/LoginScreen.vue';
import CreateStoryScreen from '@/components/stories/CreateStoryScreen.vue';
import EditStoryScreen from '@/components/stories/EditStoryScreen.vue';
import StoryDetails from '@/components/stories/StoryDetails.vue';
import StoryScreen from '@/components/stories/StoryScreen.vue';
import Vue from "vue";
import VueRouter from "vue-router";
import { loggedInGuard } from './middleware';

Vue.use(VueRouter);

const routes = [
  // {
  //   path: "/stops",
  //   component: StopScreen,
  //   name: 'Stop Screen',
  //   meta: {
  //     allowGuest: true
  //   }
  // },
  // {
  //   path: "/stops/create",
  //   component: CreateStopSection,
  //   name: 'Create Stop'
  // },
  // {
  //   path: "/stops/:id",
  //   component: StopDetailsPage,
  //   name: 'Stop Details',
  //   meta: {
  //     allowGuest: true
  //   }
  // },
  // {
  //   path: '/account',
  //   component: AccountScreen,
  //   name: 'Account Screen',
  // },
  {
    path: '/',
    component: StoryScreen,
    name: 'Story Screen',
    meta: {
      allowGuest: true
    }
  },
  {
    path: '/story/details/:storyId',
    component: StoryDetails,
    name: 'story-details',
    meta: {
      allowGuest: true
    }
  },
  {
    path: '/story/create',
    component: CreateStoryScreen,
    name: 'story-create'
  },
  {
    path: '/story/edit/:storyId',
    component: EditStoryScreen,
    name: 'story-edit',
    meta: {
      allowGuest: true
    }
  },
  {
    path: '*',
    component: StoryScreen,
    name: 'Story Screen',
    meta: {
      allowGuest: true
    }
  },
  // {
  //   path: "/routes",
  //   component: RouteScreen,
  //   name: 'Routes Screen',
  //   meta: {
  //     allowGuest: true
  //   }
  // },
  {
    path: "/login",
    component: LoginScreen,
    name: 'Login'
  },
  // {
  //   path: "*",
  //   component: RouteScreen,
  //   name: 'Routes Screen',
  //   meta: {
  //     allowGuest: true
  //   }
  // }
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
