<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="leftDrawerOpen = !leftDrawerOpen"
        />

        <q-toolbar-title>
          Lista de tarefas
        </q-toolbar-title>

        <div>TodoList v{{ $q.version }}</div>
      </q-toolbar>
    </q-header>

    <q-drawer
      v-model="leftDrawerOpen"
      show-if-above
      bordered
      content-class="bg-grey-1"
    >
      <q-list>
        <q-item-label
          header
          class="text-grey-8"
        >
          Ações
        </q-item-label>
        <EssentialLink
          v-for="link in essentialLinks"
          :key="link.title"
          v-bind="link"
        />
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
import EssentialLink from 'components/EssentialLink.vue'

export default {
  name: 'MainLayout',

  components: {
    EssentialLink
  },

  data () {
    return {
      leftDrawerOpen: false,
      essentialLinks: [
        {
          title: 'Listar Atividades',
          caption: 'Verificar tarefas não concluídas.',
          icon: 'assignment_turned_in',
          link: '/'
        },
        {
          title: 'Criar Nova Atividade',
          caption: 'Criar uma nova atividade pendente.',
          icon: 'create',
          link: '/nova-atividade'
        },
        {
          title: 'Editar Atividade',
          caption: 'Editar uma tividade criada',
          icon: 'list',
          link: '/editar-atividade'
        },
        {
          title: 'Excluir Atividade',
          caption: 'Remover uma atividade.',
          icon: 'delete',
          link: '/remover-atividade'
        }
      ]
    }
  }
}
</script>
