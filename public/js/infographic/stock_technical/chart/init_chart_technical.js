const initTechnicalChart = () => {
  const ctx = document.getElementById('technicalChart').getContext('2d')

  // plugin chart
  let plugin_chart = {
    title: {
      display: true,
      text: `Grafik Teknikal - ${stock_name}`,
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
      borderWidth: 1
    }
  }


  // scales chart
  let scales_chart = {
    x: {
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
      type: 'linear',
      display: true,
      position: 'left',
      grid: {
        color: 'rgba(255, 255, 255, 0.1)'
      },
      ticks: {
        color: '#a0a0a0'
      },
      title: {
        display: true,
        text: 'MFI / RSI',
        color: '#ffffff'
      },
      min: 0,
      max: 100
    },
    y1: {
      type: 'linear',
      display: true,
      position: 'right',
      grid: {
        drawOnChartArea: false,
      },
      ticks: {
        color: '#a0a0a0'
      },
      title: {
        display: true,
        text: 'MACD',
        color: '#ffffff'
      }
    }
  }
 
  // technical chart
  technicalChart = new Chart(ctx, {
    type: 'line',
    data: {
      labels: technicalData.labels,
      datasets: [
        technicalData.datasets.mfi,
        technicalData.datasets.rsi,
        technicalData.datasets.macd
      ]
    },
    options: {
      responsive: true,
      maintainAspectRatio: false,
      plugins: plugin_chart,
      scales:  scales_chart,
      interaction: {
        mode: 'nearest',
        axis: 'x',
        intersect: false
      }
    }
  })
}
