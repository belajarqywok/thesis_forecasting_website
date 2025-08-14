const updateTechnicalChart = () => {
  const datasets = []
  
  if (document.getElementById('mfi-check').checked) {
    datasets.push(technicalData.datasets.mfi)
  }
  if (document.getElementById('rsi-check').checked) {
    datasets.push(technicalData.datasets.rsi)
  }
  if (document.getElementById('macd-check').checked) {
    datasets.push(technicalData.datasets.macd)
    technicalChart.options.scales.y1.display = true
  } else {
    technicalChart.options.scales.y1.display = false
  }
  
  technicalChart.data.datasets = datasets
  technicalChart.update('active')
}