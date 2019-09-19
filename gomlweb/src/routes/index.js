const routes = [
  { path: '/', component: () => import('@/pages/DashBoardPage') },
  { path: '/settings', component: () => import('@/pages/SettingsPage') },
  { path: '/dashboard', component: () => import('@/pages/DashBoardPage') },
  { path: '/login', component: () => import('@/pages/Login.vue') },
];

export default routes;
