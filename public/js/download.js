const downloadReport = (informationType, stockCode) => {
  const button = event.target
  const originalText = button.innerHTML
  
  // show loading state
  button.innerHTML = '<i class="fas fa-spinner fa-spin"></i> Mengunduh...'
  button.disabled  = true
  
  setTimeout(() => {
    // reset button state
    button.innerHTML = originalText
    button.disabled  = false

    if (informationType === "fundamental") {
      pdfUrl = `https://huggingface.co/datasets/qywok/indonesia_stocks/resolve/main/fundamentals/${stockCode}.pdf`
    
    } else if (informationType === "historicals"){
      pdfUrl = `https://huggingface.co/datasets/qywok/indonesia_stocks/resolve/main/historicals/${stockCode}.pdf`

    } else if (informationType === "technicals") {
      pdfUrl = `https://huggingface.co/datasets/qywok/indonesia_stocks/resolve/main/indicators/${stockCode}.pdf`

    } else {
      pdfUrl = 'https://huggingface.co/datasets/qywok/indonesia_stocks/resolve/main/emiten_saham.pdf'
    }

    fetch(pdfUrl)
    .then(res => res.blob())
    .then(blob => {
      const blobUrl = URL.createObjectURL(blob)
      const printWindow = window.open(blobUrl)
      printWindow.addEventListener('load', () => {
        printWindow.print()
      })
    })
  }, 1000)
}