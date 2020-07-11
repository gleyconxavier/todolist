
const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/Index.vue') },
      { path: '/nova-atividade', component: () => import('pages/NovaAtividade.vue') },
      { path: '/editar-atividade', component: () => import('pages/EditarAtividade.vue') },
      { path: '/remover-atividade', component: () => import('pages/RemoverAtividade.vue') }
    ]
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '*',
    component: () => import('pages/Error404.vue')
  }
]

export default routes
