const predict = async () => {
  const days = document.getElementById('predictionDays').value
  const loadingSpinner = document.getElementById('loadingSpinner')
  
  loadingSpinner.classList.remove('d-none')
  loadingSpinner.classList.add('show')
  
  const apiUrl = 'https://qywok-cryptocurrency-prediction.hf.space/crypto/prediction'
  
  try {
    const response = await fetch(apiUrl, {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        days: parseInt(days),
        currency: "BBRI"
      })
    })
      
    if (!response.ok) throw new Error('Network response was not ok')
      
    const data = await response.json()
    updatePredictionChart(data)
      
  } catch (error) {
    console.error('Error fetching data:', error)
    alert('Terjadi kesalahan saat memproses prediksi. Silakan coba lagi.')

  } finally {
    loadingSpinner.classList.remove('show')
    loadingSpinner.classList.add('d-none')
  }
}
