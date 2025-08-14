const initChart = () => {
  const ctx = document.getElementById('ohlcChart').getContext('2d')

	let plugin_chart = {
		title: {
			display: true,
			text: `Grafik Historikal - ${stock_name}`,
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
				color: '#a0a0a0',
				callback: function(value) {
					return 'Rp ' + value.toLocaleString('id-ID');
				}
			},
			title: {
				display: true,
				text: 'Harga (Rp)',
				color: '#ffffff'
			}
		},
		y1: {
			type: 'linear',
			display: false,
			position: 'right',
			grid: {
				drawOnChartArea: false,
			},
			ticks: {
				color: '#a0a0a0',
				callback: function(value) {
					return value;
				}
			},
			title: {
				display: true,
				text: 'Volume',
				color: '#ffffff'
			}
		}
	}
  

  chart = new Chart(ctx, {
    type: 'line',
      data: {
        labels: ohlcData.labels,
        datasets: [
          ohlcData.datasets.open,
          ohlcData.datasets.high,
          ohlcData.datasets.low,
          ohlcData.datasets.close
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