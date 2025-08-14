function setTechnicalPreset(preset) {
  const checkboxes = ['mfi-check', 'rsi-check', 'macd-check']

  checkboxes.forEach(id => {
    const checkbox = document.getElementById(id)
    const container = checkbox.parentElement
    checkbox.checked = false
    container.classList.remove('checked')
  })
  
  switch(preset) {
    case 'mfi':
      document.getElementById('mfi-check').checked = true
      document.getElementById('mfi-check').parentElement.classList.add('checked')
      break
    case 'rsi':
      document.getElementById('rsi-check').checked = true
      document.getElementById('rsi-check').parentElement.classList.add('checked')
      break
    case 'macd':
      document.getElementById('macd-check').checked = true
      document.getElementById('macd-check').parentElement.classList.add('checked')
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
  
  updateTechnicalChart()
}