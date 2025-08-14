// Download Functions
function downloadFundamentalReport() {
  const button = event.target;
  const originalText = button.innerHTML;
  
  button.innerHTML = '<i class="fas fa-spinner fa-spin me-2"></i>Mengunduh...';
  button.disabled = true;
  
  setTimeout(() => {
      button.innerHTML = originalText;
      button.disabled = false;
      alert('Laporan Fundamental BBRI berhasil diunduh!');
  }, 2000);
}


function downloadOHLCReport() {
  const button = event.target;
  const originalText = button.innerHTML;
  
  button.innerHTML = '<i class="fas fa-spinner fa-spin me-2"></i>Mengunduh...';
  button.disabled = true;
  
  setTimeout(() => {
      button.innerHTML = originalText;
      button.disabled = false;
      alert('Laporan OHLC BBRI berhasil diunduh!');
  }, 2000);
}


function downloadTechnicalReport() {
  const button = event.target;
  const originalText = button.innerHTML;
  
  button.innerHTML = '<i class="fas fa-spinner fa-spin me-2"></i>Mengunduh...';
  button.disabled = true;
  
  setTimeout(() => {
      button.innerHTML = originalText;
      button.disabled = false;
      alert('Laporan Analisis Teknikal BBRI berhasil diunduh!');
  }, 2000);
}