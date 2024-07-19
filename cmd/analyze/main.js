const {ChartJSNodeCanvas} = require('chartjs-node-canvas');
//
// const width = 400; //px
// const height = 400; //px
//
// const backgroundColour = 'white'; // Uses https://www.w3schools.com/tags/canvas_fillstyle.asp
// const chartJSNodeCanvas = new ChartJSNodeCanvas({ width, height, backgroundColour});
//
//
// (async () => {
//     const configuration = {
//         type: 'line',
//         data: {},
//         options: {},
//         plugins: []
//     };
//     const image = await chartJSNodeCanvas.renderToBuffer(configuration);
//     await fs.writeFile('./example.png', image, 'base64');
//     // const dataUrl = await chartJSNodeCanvas.renderToDataURL(configuration);
//     // const stream = chartJSNodeCanvas.renderToStream(configuration);
// })();
//
//
// const csv = require('csv-parser')
// const fs = require('fs')
// const results = [];
//
// fs.createReadStream('results/raw/redis-readheavy-img-239-1721347791.csv')
//     .pipe(csv())
//     .on('data', (data) => results.push(data))
//     .on('end', () => {
//         console.log(results);
//         // [
//         //   { NAME: 'Daffy Duck', AGE: '24' },
//         //   { NAME: 'Bugs Bunny', AGE: '22' }
//         // ]
//     });


// import { ChartJSNodeCanvas, ChartCallback } from './';
// import { ChartConfiguration } from 'chart.js';

const { BoxPlotChart } = require('@sgratzl/chartjs-chart-boxplot');

function randomValues(count, min, max, extra) {

    if (!extra) {
        extra = []
    }

    const delta = max - min;
    return [...Array.from({ length: count }).map(() => Math.random() * delta + min), ...extra];
}

const data = {
    labels: ['A', 'B', 'C', 'D'],
    datasets: [
        {
            label: 'Dataset 1',
            data: [
                randomValues(100, 0, 100),
                randomValues(100, 0, 20, [110]),
                randomValues(100, 20, 70),
                // empty data
                [],
            ],
        },
        {
            label: 'Dataset 2',
            data: [
                randomValues(100, 60, 100, [5, 10]),
                randomValues(100, 0, 100),
                randomValues(100, 0, 20),
                randomValues(100, 20, 40),
            ],
        },
    ],
};

const config = {
    type: 'boxplot',
    data,
};


const fs = require('fs')

async function main() {

    const width = 400;
    const height = 400;
    const configuration = {
        type: 'bar',
        data: {
            labels: ['Red', 'Blue', 'Yellow', 'Green', 'Purple', 'Orange'],
            datasets: [{
                label: '# of Votes',
                data: [12, 19, 3, 5, 2, 3],
                backgroundColor: [
                    'rgba(255, 99, 132, 0.2)',
                    'rgba(54, 162, 235, 0.2)',
                    'rgba(255, 206, 86, 0.2)',
                    'rgba(75, 192, 192, 0.2)',
                    'rgba(153, 102, 255, 0.2)',
                    'rgba(255, 159, 64, 0.2)'
                ],
                borderColor: [
                    'rgba(255,99,132,1)',
                    'rgba(54, 162, 235, 1)',
                    'rgba(255, 206, 86, 1)',
                    'rgba(75, 192, 192, 1)',
                    'rgba(153, 102, 255, 1)',
                    'rgba(255, 159, 64, 1)'
                ],
                borderWidth: 1
            }]
        },
        options: {},
        plugins: [{
            id: 'background-colour',
            beforeDraw: (chart) => {
                const ctx = chart.ctx;
                ctx.save();
                ctx.fillStyle = 'white';
                ctx.fillRect(0, 0, width, height);
                ctx.restore();
            }
        }]
    };
    const chartCallback = (ChartJS) => {
        ChartJS.defaults.responsive = true;
        ChartJS.defaults.maintainAspectRatio = false;
    };
    const chartJSNodeCanvas = new ChartJSNodeCanvas({width, height, chartCallback});
    const buffer = await chartJSNodeCanvas.renderToBuffer(config);
    await fs.writeFileSync('./example.png', buffer, 'base64');
}

main();

// console.log("hello world");
