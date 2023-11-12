<template>
    <canvas id="chart"></canvas>
    <button onclick="updateData()">Update</button>
</template>

<script>

//記事末尾で補足
import Chart from 'chart.js/auto';

export default {
    methods: {
        renderChart() {
            let ctx = document.getElementById("chart");
            new Chart(ctx, {
                type: 'line',
                data: {
                    datasets: [{
                        label: 'データポイント',
                        data: [
                            { x: 1, y: 5 },
                            { x: 2, y: 7 },
                            { x: 3, y: 3 },
                            { x: 4, y: 9 },
                            { x: 50, y: 6 }
                        ],
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
                            min: 0,
                            max: 10
                        }
                    }
                }
            });
        }
    },
    mounted() {
        this.renderChart();
    }
};

function updateData() {
    // 新しいデータ
    var newData = {
        datasets: [{
            label: 'データポイント',
            data: [
                { x: 1, y: 8 },
                { x: 2, y: 6 },
                { x: 3, y: 4 },
                { x: 4, y: 7 },
                { x: 5, y: 9 }
            ],
            backgroundColor: 'rgba(75, 192, 192, 0.5)',
            pointRadius: 8
        }]
    };

    // データの更新
    myScatterChart.data = newData;

    // グラフの再描画
    myScatterChart.update();
}
</script>
