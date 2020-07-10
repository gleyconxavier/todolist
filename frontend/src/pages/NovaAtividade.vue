<template>
  <q-page class="flex flex-center">

    <div class="row">
      <div class="col">
        <h5>Digite uma nova atividade</h5>
        <q-input type="input" filled v-model="title">
        </q-input>
        <q-ajax-bar
          ref="bar"
          position="bottom"
          color="accent"
          size="10px"
          skip-hijack
        />
        <div v-if="status">
          <p class="text-positive">Sucesso!</p>
        </div>
        <div v-if="status && status !== 201">
          <p class="text-negative">Ocorreu um erro, tente novamente.</p>
        </div>
        <q-btn class="q-mt-md" color="primary" icon="check" label="OK" @click="enviarAtividade" />
      </div>
    </div>
    <q-img
        class=""
        spinner-color="white"
        style="height: 168px; max-width: 180px"
        alt="TodoList Ilustration"
        src="~assets/to-do-list.svg"
    />

  </q-page>
</template>

<script>
export default {
  name: 'PageIndex',
  data () {
    return {
      title: '',
      info: '',
      status: ''
    }
  },
  methods: {
    enviarAtividade: function () {
      const bar = this.$refs.bar
      bar.start()
      console.log(this.title)
      const data = {
        title: this.title
      }
      this.$axios
        .post('/criar-atividade', data)
        .then(response => (this.status = response.status))
        .catch(error => console.log(error))
      bar.stop()
    }
  }
}
</script>
