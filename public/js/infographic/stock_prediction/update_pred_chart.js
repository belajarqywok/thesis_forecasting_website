const updatePredictionChart = (data) => {
  const actualDates  = data.data.predictions.actuals.map(entry => entry.date)
  const actualPrices = data.data.predictions.actuals.map(entry => entry.price)

  const predictionDates  = data.data.predictions.predictions.map(entry => entry.date)
  const predictionPrices = data.data.predictions.predictions.map(entry => entry.price)

  const labels         = [...actualDates, ...predictionDates]
  const actualData     = [...actualPrices, ...Array(predictionPrices.length).fill(null)]
  const predictionData = [...Array(actualPrices.length).fill(null), ...predictionPrices]

  if (predictionChart) {
      predictionChart.data.labels = labels
      predictionChart.data.datasets[0].data = actualData
      predictionChart.data.datasets[1].data = predictionData
      predictionChart.update()
      
  } else {
      const ctx = document.getElementById('predictionChart').getContext('2d')

      // chart data
      let chart_data = {
        labels: labels,
        datasets: [
          {
            label: 'Data Historis',
            data: actualData,
            borderColor: 'rgba(75, 192, 192, 1)',
            backgroundColor: 'rgba(75, 192, 192, 0.2)',
            borderWidth: 2,
            fill: false,
            tension: 0.1,
            pointRadius: 2,
            pointHoverRadius: 5
          },
          {
            label: 'Prediksi',
            data: predictionData,
            borderColor: 'rgba(255, 99, 132, 1)',
            backgroundColor: 'rgba(255, 99, 132, 0.2)',
            borderWidth: 2,
            borderDash: [5, 5],
            fill: false,
            tension: 0.1,
            pointRadius: 3,
            pointHoverRadius: 6
          }
        ]
      }

      // options plugins chart
      let options_plugins_chart = {
        title: {
          display: true,
          text: 'BBRI - Prediksi Harga Saham',
          color: '#ffffff',
          font: {
            size: 16,
            weight: 'bold'
          }
        },
        legend: {
          labels: {
            color: '#ffffff',
            usePointStyle: true,
            padding: 20
          }
        },
        tooltip: {
          mode: 'index',
          intersect: false,
          backgroundColor: 'rgba(26, 26, 46, 0.9)',
          titleColor: '#ffffff',
          bodyColor: '#ffffff',
          borderColor: '#00d4aa',
          borderWidth: 1,
          callbacks: {
            label: function (context) {
              return context.dataset.label + 
              ': Rp ' + context.parsed.y.toLocaleString('id-ID')
            }
          }
        }
      }

      // scales plugins chart
      let scales_plugins_chart = {
        x: {
          type: 'category',
          title: {
            display: true,
            text: 'Tanggal',
            color: '#ffffff'
          },
          grid: {
            color: 'rgba(255, 255, 255, 0.1)'
          },
          ticks: {
            color: '#a0a0a0'
          }
        },
        y: {
          beginAtZero: false,
          title: {
            display: true,
            text: 'Harga (Rp)',
            color: '#ffffff'
          },
          grid: {
            color: 'rgba(255, 255, 255, 0.1)'
          },
          ticks: {
            color: '#a0a0a0',
            callback: function (value) {
              return 'Rp ' + value.toLocaleString('id-ID')
            }
          }
        }
      }

      // prediction chart
      predictionChart = new Chart(ctx, {
        type: 'line',
        data: chart_data,
        options: {
          responsive: true,
          maintainAspectRatio: false,
          plugins: options_plugins_chart,
          scales:  scales_plugins_chart,
          interaction: {
            mode: 'nearest',
            axis: 'x',
            intersect: false
          },
          elements: {
            point: {
              hoverBorderWidth: 3
            }
          }
        }
      })

  }
}
