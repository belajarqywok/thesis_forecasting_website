const setPreset = (preset) => {
  const checkboxes = ['open-check', 'high-check', 'low-check', 'close-check', 'volume-check']

  checkboxes.forEach(id => {
    const checkbox   = document.getElementById(id)
    const container  = checkbox.parentElement
    checkbox.checked = false
    container.classList.remove('checked')
  })
  
  switch(preset) {
    case 'price':
    	document.getElementById('close-check').checked = true
      document.getElementById('close-check').parentElement.classList.add('checked')
      break
    case 'ohlc':
      ['open-check', 'high-check', 'low-check', 'close-check'].forEach(id => {
        document.getElementById(id).checked = true
        document.getElementById(id).parentElement.classList.add('checked')
      })
      break
    case 'volume':
      document.getElementById('volume-check').checked = true
      document.getElementById('volume-check').parentElement.classList.add('checked')
      break
    case 'all':
      checkboxes.forEach(id => {
        document.getElementById(id).checked = true
        document.getElementById(id).parentElement.classList.add('checked')
      })
      break
    case 'clear':
      break
  }
  
  updateChart()
}