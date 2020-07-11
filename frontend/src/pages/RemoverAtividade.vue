<template>
  <q-page class="flex flex-center">

    <q-img
      class="q-mb-md"
      spinner-color="white"
      style="height: 168px; max-width: 230px"
      alt="TodoList Ilustration"
      src="~assets/to-do-list.svg"
    />
    <div class="row">
      <div class="col-12" v-if="info">
        <q-field v-for="(item, index) in info" :key="item.id" filled readonly>
            <div class="self-center">{{ item.title }}</div>

          <template v-slot:after>
            <q-btn
              color="primary"
              icon="delete_forever"
              label="Remover"
              @click="removerAtividade(item, index, info)"
            />
          </template>
        </q-field>
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
      filteredInfo: '',
      item: ''
    }
  },
  mounted () {
    this.$axios
      .get('/listar-atividades')
      .then(response => (this.info = response.data))
      .catch(error => console.log(error))
  },
  methods: {
    removerAtividade: function (item, index, array) {
      const bar = this.$refs.bar
      bar.start()

      this.$axios
        .delete('/remover-atividade/' + item.id)
        .then(response => (alert('Removido com sucesso!')))
        .then(response => {
          array.splice(index, 1)
        })
        .catch(error => console.log(error))
      bar.stop()
    }
  }
}
</script>
