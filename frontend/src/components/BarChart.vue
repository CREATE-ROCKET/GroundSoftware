<template>
    <div>
        <canvas ref="canvas" id="chart"></canvas>
    </div>
</template>

<script>
import { Chart, LinearScale, PointElement, LineController, LineElement } from 'chart.js';

// Register the necessary components
Chart.register(LinearScale, PointElement, LineController, LineElement);

export default {
    methods: {
        renderChart() {
            let ctx = this.$refs.canvas.getContext('2d');
            this.chart = new Chart(ctx, {
                type: 'line',
                data: {
                    datasets: [{
                        label: 'データポイント',
                        data: [],
                        backgroundColor: 'rgba(75, 192, 192, 0.5)',
                        pointRadius: 2
                    }]
                },
                options: {
                    scales: {
                        x: {
                            type: 'linear',
                            position: 'bottom'
                        },
                        y: {
                            // min: 0,
                            // max: 10
                        }
                    },
                    animation: false
                }
            });
        },
        addDataPoint(x, y) {
            this.chart.data.datasets[0].data.push({ x: x, y: y });
            this.chart.update();
        },
    },
    mounted() {
        this.renderChart();
    }
};
</script>