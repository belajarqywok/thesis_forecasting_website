const toggleTechnicalDataset = (datasetName) => {
  const checkbox = document.getElementById(datasetName + '-check')
  const checkboxContainer = checkbox.parentElement
  
  checkbox.checked = !checkbox.checked
  
  if (checkbox.checked) {
    checkboxContainer.classList.add('checked')
  } else {
    checkboxContainer.classList.remove('checked')
  }
  
  updateTechnicalChart()
}