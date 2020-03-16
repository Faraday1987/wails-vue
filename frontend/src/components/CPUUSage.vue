<template>
  <div class="container">
    <div class="logo"></div>
    <apexchart class="chart" width="200" type="radialBar" :options="options" :series="series"></apexchart>
  </div>
</template>

<script>
export default {
  data() {
    return {
      message: " ",
      series: [0],
      options: {
        labels: ['CPU Usage'],
      }
    };
  },
  mounted: function() {
    console.log(wails);
    wails.Events.On("cpu_usage", cpu_usage => {
      if (cpu_usage) {
        this.series = [cpu_usage.avg];
        this.message = cpu_usage.avg;
      }
    });
  },
  methods: {
    getMessage: function() {
      var self = this;
      window.backend.Stats.GetCPUUsage().then(cpu_usage => {
        self.message = cpu_usage.avg;
      });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.chart {
  position: absolute;
  right: 0;
}
.logo {
  position: absolute;
  width: 100px;
  height: 100px;
  top: 0;
  left: 0;  
  background-image: url('../assets/images/logo.png');
  background-attachment: inherit;
  background-size: cover;
}
</style>
