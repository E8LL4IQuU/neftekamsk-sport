import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: '',
      components: {
        // route level code-splitting
        // this generates a separate chunk (news.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        default: () => import('@/views/Page.vue'),
        navbar: () => import('@/components/NavBar.vue'),
        footer: () => import('@/components/Footer.vue')
      }
    },
    {
      path: '/events',
      name: 'events',
      components: {
        default: () => import('@/views/EventsView.vue'),
        navbar: () => import('@/components/NavBar.vue'),
        footer: () => import('@/components/Footer.vue')
      }
    },
    {
      path: '/events/:id',
      components: {
        default: () => import('@/views/EventView.vue'),
        navbar: () => import('@/components/NavBar.vue'),
        footer: () => import('@/components/Footer.vue')
      }
    },
    {
      path: '/news',
      name: 'news',
      components: {
        default: () => import('@/views/NewsBulkView.vue'),
        navbar: () => import('@/components/NavBar.vue'),
        footer: () => import('@/components/Footer.vue')
      }
    },
    {
      path: '/news/:id',
      name: 'news',
      components: {
        default: () => import('@/views/NewsView.vue'),
        navbar: () => import('@/components/NavBar.vue'),
        footer: () => import('@/components/Footer.vue')
      }
    },
    {
      path: '/schedule',
      name: 'schedule',
      components: {
        default: () => import('@/views/ScheduleView.vue'),
        navbar: () => import('@/components/NavBar.vue'),
        footer: () => import('@/components/Footer.vue')
      }
    },
    {
      path: '/sections',
      name: 'sections',
      components: {
        default: () => import('@/views/SectionsView.vue'),
        navbar: () => import('@/components/NavBar.vue'),
        footer: () => import('@/components/Footer.vue')
      }
    },
    {
      path: '/gallery',
      name: 'gallery',
      components: {
        default: () => import('@/views/GalleryView.vue'),
        navbar: () => import('@/components/NavBar.vue'),
        footer: () => import('@/components/Footer.vue')
      }
    },
    {
      path: '/documents',
      name: 'documents',
      components: {
        default: () => import('@/views/DocumentsView.vue'),
        navbar: () => import('@/components/NavBar.vue'),
        footer: () => import('@/components/Footer.vue')
      }
    },
    {
      path: '/people',
      name: 'people',
      components: {
        default: () => import('@/views/PeopleView.vue'),
        navbar: () => import('@/components/NavBar.vue'),
        footer: () => import('@/components/Footer.vue')
      }
    },

    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/LoginView.vue'),
    },
    {
      path: '/logout',
      name: 'logout',
      component: () => import('@/views/LogoutView.vue'),
    },
    {
      path: '/manage',
      redirect: '/manage/events'
    },
    {
      path: '/manage/events',
      name: 'manage.events',
      components: {
        default: () => import('@/views/Manage/ManageEvents.vue'),
        manageSidebar: () => import('@/components/Manage/Sidebar.vue')
      }
    },
    {
      path: '/manage/events/create',
      name: 'manage.events.create',
      component: () => import('@/views/Manage/CreateEvent.vue'),
    },
    {
      path: '/manage/news',
      name: 'manage.news',
      components: {
        default: () => import('@/views/Manage/ManageNews.vue'),
        manageSidebar: () => import('@/components/Manage/Sidebar.vue')
      }
    },
    {
      path: '/manage/news/create',
      name: 'manage.news.create',
      component: () => import('@/views/Manage/CreateNews.vue'),
    },
    {
      path: '/manage/news/:id',
      name: 'newsEdit',
      component: () => import('@/views/Manage/EditNews.vue'),
    },
    {
      path: '/manage/signups',
      name: 'manage.signups',
      components: {
        default: () => import('@/views/Manage/ManageSignups.vue'),
        manageSidebar: () => import('@/components/Manage/Sidebar.vue')
      }
    },
    {
      path: '/:catchAll(.*)',
      redirect: '/',
    }

  ]
})

export default router
