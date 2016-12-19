Vue.component('app-view', {
  props: ['title'],
  template: '<h2>{{ title }}</h2>'
})

var vm = new Vue({
  el: '#app',
  data: {
    title: 'Red!'
  }
})