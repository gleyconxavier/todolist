<template>
  <q-page class="flex flex-center">

    <div class="row">
      <q-img
        class="q-mb-md"
        spinner-color="white"
        style="height: 168px; max-width: 230px"
        alt="TodoList Ilustration"
        src="~assets/to-do-list.svg"
      />
      <div class="col-12" v-if="info">
        <q-input v-for="item in info.data" :key="item.id" v-model="item.title" filled>

          <template v-slot:after>
            <q-btn
              color="primary"
              icon="edit"
              label="Editar"
              @click="editarAtividade(item.id, item.title)"
            />
          </template>
        </q-input>
        <q-ajax-bar
          ref="bar"
          position="bottom"
          color="accent"
          size="10px"
          skip-hijack
        />
      </div>
    </div>

  </q-page>
</template>

<script>
export default {
  name: 'PageIndex',
  data () {
    return {
      info: '',
      newTitle: ''
    }
  },
  mounted () {
    this.$axios
      .get('/listar-atividades')
      .then(response => (this.info = response))
      .catch(error => console.log(error))
  },
  methods: {
    editarAtividade: function (id, title) {
      const bar = this.$refs.bar
      bar.start()
      const data = {
        title: title
      }
      this.$axios
        .put('/atualizar-atividade/' + id, data)
        .then(response => (alert('Editado com sucesso!')))
        .catch(error => console.log(error))
      bar.stop()
    }
  }
}
</script>
