const updateChart = () => {
  const datasets = []
  
  if (document.getElementById('open-check').checked) {
    datasets.push(ohlcData.datasets.open)
  }
  if (document.getElementById('high-check').checked) {
    datasets.push(ohlcData.datasets.high)
  }
  if (document.getElementById('low-check').checked) {
    datasets.push(ohlcData.datasets.low)
  }
  if (document.getElementById('close-check').checked) {
    datasets.push(ohlcData.datasets.close)
  }
  if (document.getElementById('volume-check').checked) {
    datasets.push(ohlcData.datasets.volume)
    chart.options.scales.y1.display = true
  } else {
    chart.options.scales.y1.display = false
  }
  
  chart.data.datasets = datasets
  chart.update('active')
}