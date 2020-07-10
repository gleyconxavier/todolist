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
        <q-field v-for="item in info.data" :key="item.id" class="" filled readonly>
            <div class="self-center">{{ item.title }}</div>
          </q-field>
      </div>
    </div>

  </q-page>
</template>

<script>
export default {
  name: 'PageIndex',
  data () {
    return {
      info: ''
    }
  },
  mounted () {
    this.$axios
      .get('/listar-atividades')
      .then(response => (this.info = response))
      .catch(error => console.log(error))
  }
}
</script>
