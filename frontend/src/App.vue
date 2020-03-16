<template>
  <div id="app">
    <CPUUSage/>
  </div>
</template>

<script>
import CPUUSage from "./components/CPUUSage.vue";
import "./assets/css/main.css";

export default {
  name: "app",
  components: {
    CPUUSage
  }
};
</script>
